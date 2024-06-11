import json
from pathlib import Path
from typing import Dict

import urllib.request

from botocore.model import ServiceModel
import botocore.loaders
import botocore.model
import boto3

p = Path(__file__).parent/'map.json'
if p.exists():
    data = p.read_text()
else:
    resp = urllib.request.urlopen('https://raw.githubusercontent.com/iann0036/iam-dataset/main/aws/map.json')
    data = resp.read().decode('utf-8')
    p.write_text(data)

iam_map = json.loads(data)
# iam_map['sdk_method_iam_mappings']

def main():
    iam_action_to_boto3_map = {}
    p = Path(__file__).parent/'boto3_iam_map.json'
    boto3_iam_map = json.loads(p.read_text())
    for boto3_svc_name, op_names in boto3_iam_map.items():
        for boto3_op_name, actions in op_names.items():
            for action in actions:
                full_iam_action = action.get('action')
                if not full_iam_action:
                    continue
                iam_action_to_boto3_map.setdefault(full_iam_action, []).append(f"{boto3_svc_name}:{boto3_op_name}")

    p = Path(__file__).parent/'iam_to_boto3_map.json'
    p.write_text(json.dumps(iam_action_to_boto3_map, indent=2))


if __name__ == '__main__':
    main()
