from typing import Dict, Any

def load_config(name: str) -> Dict[str, Any]:
    if name == "server":
        return {"ip": "127.0.0.1", "port": 8000}
    return {}
