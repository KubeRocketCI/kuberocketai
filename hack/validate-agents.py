#!/usr/bin/env python3
"""
Agent Validation Tool for KubeRocketAI Framework

Usage: ./hack/validate-agents.py [agent-file.yaml|--all]
"""

import json
import os
import sys
import subprocess
from pathlib import Path

# Project paths
SCRIPT_DIR = Path(__file__).parent
PROJECT_ROOT = SCRIPT_DIR.parent
VENV_PATH = PROJECT_ROOT / ".venv"
REQUIREMENTS_PATH = SCRIPT_DIR / "requirements.txt"

def setup_venv():
    """Create venv and install requirements if needed"""
    if not VENV_PATH.exists():
        print("📦 Creating virtual environment...")
        subprocess.run([sys.executable, "-m", "venv", str(VENV_PATH)], check=True)
    
    # Get venv python path
    venv_python = VENV_PATH / "bin" / "python"
    if not venv_python.exists():
        venv_python = VENV_PATH / "Scripts" / "python.exe"  # Windows
    
    # Check if requirements are installed
    try:
        subprocess.run([str(venv_python), "-c", "import yaml, jsonschema"], 
                      check=True, capture_output=True)
    except subprocess.CalledProcessError:
        print("📥 Installing requirements...")
        subprocess.run([str(venv_python), "-m", "pip", "install", "-r", str(REQUIREMENTS_PATH)], 
                      check=True)
    
    return venv_python

def validate_agent(venv_python, agent_file):
    """Validate single agent file"""
    validation_code = f'''
import yaml
import json
import jsonschema
from pathlib import Path

# Load schema
schema_path = Path("{PROJECT_ROOT}/assets/schemas/agent-schema.json")
with open(schema_path) as f:
    schema = json.load(f)

# Load agent file
agent_path = Path("{agent_file}")
if not agent_path.exists():
    print(f"❌ File not found: {{agent_path}}")
    exit(1)

with open(agent_path) as f:
    agent_data = yaml.safe_load(f)

# Validate
try:
    jsonschema.validate(agent_data, schema)
    print(f"✅ VALID: {{agent_path.name}} passes schema validation")
except jsonschema.ValidationError as e:
    print(f"❌ INVALID: {{agent_path.name}} - {{e.message}}")
    exit(1)
except Exception as e:
    print(f"❌ ERROR: {{e}}")
    exit(1)
'''
    
    result = subprocess.run([str(venv_python), "-c", validation_code], 
                           capture_output=True, text=True)
    print(result.stdout)
    if result.stderr:
        print(result.stderr)
    return result.returncode == 0

def validate_all_agents(venv_python):
    """Validate all agent files"""
    agents_dir = PROJECT_ROOT / "assets/framework/core/agents"
    if not agents_dir.exists():
        print(f"❌ Agents directory not found: {agents_dir}")
        return False
    
    agent_files = list(agents_dir.glob("*.yaml"))
    if not agent_files:
        print("❌ No agent files found")
        return False
    
    print(f"🔍 Validating {len(agent_files)} agent files...")
    all_valid = True
    
    for agent_file in agent_files:
        if not validate_agent(venv_python, agent_file):
            all_valid = False
    
    if all_valid:
        print(f"✅ All {len(agent_files)} agents are valid!")
    else:
        print("❌ Some agents failed validation")
    
    return all_valid

def main():
    if len(sys.argv) < 2:
        print("Usage: ./hack/validate-agents.py [agent-file.yaml|--all]")
        sys.exit(1)
    
    # Setup environment
    try:
        venv_python = setup_venv()
    except Exception as e:
        print(f"❌ Failed to setup environment: {e}")
        sys.exit(1)
    
    # Run validation
    if sys.argv[1] == "--all":
        success = validate_all_agents(venv_python)
    else:
        agent_file = sys.argv[1]
        success = validate_agent(venv_python, agent_file)
    
    sys.exit(0 if success else 1)

if __name__ == "__main__":
    main()