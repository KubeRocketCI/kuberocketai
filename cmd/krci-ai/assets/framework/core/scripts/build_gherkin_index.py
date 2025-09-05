#!/usr/bin/env python3
"""
Build a local Gherkin index (lexical/structured), not semantic RAG.

Outputs:
- JSON index (entries per scenario) with structured fields
- Optional SQLite database with FTS5 for fast local search

Usage:
  python ./.krci-ai/scripts/build_gherkin_index.py \
    --root . \
    --out-json ./.krci-ai/indexes/gherkin-lex.json \
    --out-sqlite ./.krci-ai/indexes/gherkin-lex.sqlite

The script is dependency-free (std lib only). FTS will be used if available.
"""

from __future__ import annotations

import argparse
import datetime
import json
import os
import re
import sqlite3
from dataclasses import dataclass, asdict
from pathlib import Path
from typing import Iterable, List, Optional, Tuple


GHERKIN_STEP_RE = re.compile(r"^\s*(Given|When|Then|And|But)\b\s*(.*)")
FEATURE_TITLE_RE = re.compile(r"^\s*Feature:\s*(.+?)\s*$")
SCENARIO_TITLE_RE = re.compile(r"^\s*Scenario(?: Outline)?:\s*(.+?))\s*$")
TAGS_LINE_RE = re.compile(r"^\s*(@\S+(?:\s+@\S+)*)\s*$")
EXAMPLES_RE = re.compile(r"^\s*Examples?\s*:\s*(.*)$")
URL_RE = re.compile(r"https?://\S+")
CONSTANT_RE = re.compile(r"\b[A-Z][A-Z0-9_]{2,}\b")


@dataclass
class RagEntry:
    filePath: str
    featureTitle: Optional[str]
    scenarioTitle: Optional[str]
    tags: List[str]
    stepKeywords: List[str]
    stepTextSample: str
    artifacts: List[str]
    examplesHeaders: List[str]
    provider: Optional[str]
    layer: str  # UI | API | utility
    topic: Optional[str]


def detect_provider_layer_topic(file_path: Path) -> Tuple[Optional[str], str, Optional[str]]:
    path_parts = file_path.parts
    provider: Optional[str] = None
    layer: str = "UI"
    topic: Optional[str] = None

    if "api_tests" in path_parts:
        layer = "API"
    if "utility" in path_parts:
        layer = "utility"

    try:
        idx = path_parts.index("tekton")
        if idx + 1 < len(path_parts):
            nxt = path_parts[idx + 1]
            if nxt in {"bitbucket", "github", "gerrit", "gitlab"}:
                provider = nxt
    except ValueError:
        pass

    topic = file_path.stem
    return provider, layer, topic


def parse_feature_file(file_path: Path) -> List[RagEntry]:
    entries: List[RagEntry] = []

    feature_title: Optional[str] = None
    current_tags: List[str] = []
    scenario_title: Optional[str] = None
    step_keywords: List[str] = []
    step_texts: List[str] = []
    artifacts: List[str] = []
    examples_headers: List[str] = []

    def flush_scenario():
        if scenario_title is None and not step_texts and not current_tags:
            return
        p, l, t = detect_provider_layer_topic(file_path)
        entries.append(
            RagEntry(
                filePath=str(file_path).replace("\\", "/"),
                featureTitle=feature_title,
                scenarioTitle=scenario_title,
                tags=list(dict.fromkeys(current_tags)),
                stepKeywords=list(dict.fromkeys(step_keywords)),
                stepTextSample=" \n".join(step_texts[:10]),
                artifacts=sorted(list({*artifacts})),
                examplesHeaders=list(dict.fromkeys(examples_headers)),
                provider=p,
                layer=l,
                topic=t,
            )
        )

    with file_path.open("r", encoding="utf-8", errors="ignore") as f:
        for raw_line in f:
            line = raw_line.rstrip("\n")

            m_feat = FEATURE_TITLE_RE.match(line)
            if m_feat:
                feature_title = m_feat.group(1).strip()
                continue

            m_tags = TAGS_LINE_RE.match(line)
            if m_tags:
                tag_line = m_tags.group(1)
                current_tags.extend(tag_line.split())
                continue

            m_scn = SCENARIO_TITLE_RE.match(line)
            if m_scn:
                if scenario_title is not None:
                    flush_scenario()
                    current_tags_local = list(current_tags)
                    current_tags.clear()
                    current_tags.extend(current_tags_local)
                scenario_title = m_scn.group(1).strip()
                step_keywords.clear()
                step_texts.clear()
                artifacts.clear()
                examples_headers.clear()
                continue

            m_step = GHERKIN_STEP_RE.match(line)
            if m_step:
                step_keywords.append(m_step.group(1))
                step_text = m_step.group(2).strip()
                step_texts.append(step_text)
                artifacts.extend(URL_RE.findall(line))
                artifacts.extend(CONSTANT_RE.findall(line))
                continue

            m_ex = EXAMPLES_RE.match(line)
            if m_ex:
                continue

            if "|" in line and scenario_title is not None:
                parts = [col.strip() for col in line.split("|") if col.strip()]
                if parts and len(parts) >= 2:
                    examples_headers = parts

        if scenario_title is not None:
            flush_scenario()

    return entries


def iter_feature_files(root: Path) -> Iterable[Path]:
    features_root = root / "docs" / "testing" / "features"
    if not features_root.exists():
        return []
    return features_root.rglob("*.feature")


def write_json(entries: List[RagEntry], out_path: Path) -> None:
    out_path.parent.mkdir(parents=True, exist_ok=True)
    payload = {
        "generatedAt": datetime.datetime.now(datetime.UTC).isoformat(),
        "entryCount": len(entries),
        "entries": [asdict(e) for e in entries],
    }
    with out_path.open("w", encoding="utf-8") as f:
        json.dump(payload, f, ensure_ascii=False, indent=2)


def write_sqlite(entries: List[RagEntry], out_path: Path) -> None:
    out_path.parent.mkdir(parents=True, exist_ok=True)
    conn = sqlite3.connect(str(out_path))
    try:
        cur = conn.cursor()
        cur.execute("PRAGMA journal_mode=WAL;")
        cur.execute("DROP TABLE IF EXISTS entries;")
        cur.execute(
            """
            CREATE TABLE entries (
              filePath TEXT,
              featureTitle TEXT,
              scenarioTitle TEXT,
              tags TEXT,
              stepKeywords TEXT,
              stepTextSample TEXT,
              artifacts TEXT,
              examplesHeaders TEXT,
              provider TEXT,
              layer TEXT,
              topic TEXT
            );
            """
        )
        cur.executemany(
            """
            INSERT INTO entries (
              filePath, featureTitle, scenarioTitle, tags, stepKeywords, stepTextSample,
              artifacts, examplesHeaders, provider, layer, topic
            ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);
            """,
            [
                (
                    e.filePath,
                    e.featureTitle,
                    e.scenarioTitle,
                    " ".join(e.tags),
                    " ".join(e.stepKeywords),
                    e.stepTextSample,
                    " ".join(e.artifacts),
                    " ".join(e.examplesHeaders),
                    e.provider,
                    e.layer,
                    e.topic,
                )
                for e in entries
            ],
        )
        try:
            cur.execute("DROP TABLE IF EXISTS entries_fts;")
            cur.execute(
                """
                CREATE VIRTUAL TABLE entries_fts USING fts5(
                  filePath, featureTitle, scenarioTitle, tags, stepKeywords, stepTextSample,
                  artifacts, examplesHeaders, provider, layer, topic
                );
                """
            )
            cur.execute(
                "INSERT INTO entries_fts(rowid, filePath, featureTitle, scenarioTitle, tags, stepKeywords, stepTextSample, artifacts, examplesHeaders, provider, layer, topic) "
                "SELECT rowid, filePath, featureTitle, scenarioTitle, tags, stepKeywords, stepTextSample, artifacts, examplesHeaders, provider, layer, topic FROM entries;"
            )
        except sqlite3.OperationalError:
            pass
        conn.commit()
    finally:
        conn.close()


def main() -> None:
    parser = argparse.ArgumentParser(description="Build local Gherkin index (lexical)")
    parser.add_argument("--root", type=str, default=".", help="Repository root (default: current dir)")
    parser.add_argument("--out-json", type=str, default="./.krci-ai/indexes/gherkin-lex.json", help="Output JSON path")
    parser.add_argument("--out-sqlite", type=str, default="./.krci-ai/indexes/gherkin-lex.sqlite", help="Output SQLite path")
    args = parser.parse_args()

    root = Path(args.root).resolve()
    files = list(iter_feature_files(root))
    entries: List[RagEntry] = []
    for fp in files:
        try:
            entries.extend(parse_feature_file(fp))
        except Exception as exc:
            print(f"[WARN] Failed parsing {fp}: {exc}")

    write_json(entries, Path(args.out_json))
    write_sqlite(entries, Path(args.out_sqlite))
    print(f"Wrote JSON index with {len(entries)} entries to {args.out_json}")
    print(f"Wrote SQLite index to {args.out_sqlite}")


if __name__ == "__main__":
    main()
