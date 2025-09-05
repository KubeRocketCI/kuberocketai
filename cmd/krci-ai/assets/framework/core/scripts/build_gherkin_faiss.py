#!/usr/bin/env python3
"""
Build a FAISS vector index for Gherkin scenarios (RAG-ready).

Requires:
  pip install sentence-transformers faiss-cpu

Usage:
  python ./.krci-ai/scripts/build_gherkin_faiss.py \
    --root . \
    --model sentence-transformers/all-MiniLM-L6-v2 \
    --out-index ./.krci-ai/indexes/gherkin-faiss.index \
    --out-meta ./.krci-ai/indexes/gherkin-faiss.meta.json
"""

from __future__ import annotations

import argparse
import datetime
import json
from pathlib import Path
from typing import List

import faiss  # type: ignore
from sentence_transformers import SentenceTransformer  # type: ignore

try:
    from .build_gherkin_index import iter_feature_files, parse_feature_file  # type: ignore
except Exception:
    from build_gherkin_index import iter_feature_files, parse_feature_file  # type: ignore


def scenario_text_for_embedding(feature_title: str | None, scenario_title: str | None, tags: List[str], steps: str) -> str:
    parts = []
    if feature_title:
        parts.append(f"Feature: {feature_title}")
    if scenario_title:
        parts.append(f"Scenario: {scenario_title}")
    if tags:
        parts.append("Tags: " + " ".join(tags))
    if steps:
        parts.append("Steps:\n" + steps)
    return "\n".join(parts)


def main() -> None:
    parser = argparse.ArgumentParser(description="Build FAISS index for Gherkin scenarios")
    parser.add_argument("--root", type=str, default=".")
    parser.add_argument("--model", type=str, default="sentence-transformers/all-MiniLM-L6-v2")
    parser.add_argument("--out-index", type=str, default="./.krci-ai/indexes/gherkin-faiss.index")
    parser.add_argument("--out-meta", type=str, default="./.krci-ai/indexes/gherkin-faiss.meta.json")
    args = parser.parse_args()

    root = Path(args.root).resolve()
    files = list(iter_feature_files(root))

    # Collect entries per scenario
    meta: List[dict] = []
    corpus: List[str] = []
    for fp in files:
        for e in parse_feature_file(fp):
            doc = scenario_text_for_embedding(e.featureTitle, e.scenarioTitle, e.tags, e.stepTextSample)
            corpus.append(doc)
            meta.append({
                "filePath": e.filePath,
                "featureTitle": e.featureTitle,
                "scenarioTitle": e.scenarioTitle,
                "tags": e.tags,
                "provider": e.provider,
                "layer": e.layer,
                "topic": e.topic,
            })

    # Embed
    model = SentenceTransformer(args.model)
    embeddings = model.encode(corpus, show_progress_bar=True, normalize_embeddings=True)

    # Build FAISS index (cosine via inner-product with normalized vectors)
    dim = embeddings.shape[1]
    index = faiss.IndexFlatIP(dim)
    index.add(embeddings)

    # Persist index and metadata
    out_index = Path(args.out_index)
    out_index.parent.mkdir(parents=True, exist_ok=True)
    faiss.write_index(index, str(out_index))

    out_meta = Path(args.out_meta)
    out_meta.parent.mkdir(parents=True, exist_ok=True)
    with out_meta.open("w", encoding="utf-8") as f:
        json.dump({
            "generatedAt": datetime.datetime.now(datetime.UTC).isoformat(),
            "model": args.model,
            "dimension": dim,
            "count": len(meta),
            "entries": meta,
        }, f, ensure_ascii=False, indent=2)

    print(f"FAISS index written to {out_index}")
    print(f"Metadata written to {out_meta}")


if __name__ == "__main__":
    main()
