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
iam_svc_to_js_sdk = json.loads(
    (Path(__file__).parent/'iam_svc_to_js_sdk.json').read_text()
)

iam_map['sdk_service_mappings'].update({
    'APIGateway': 'apigateway',
    'Bedrock': 'bedrock',
    'B2bi': 'b2bi',
    'ARCZonalShift': 'arc-zonal-shift',
    'BackupStorage': 'backup-storage',
    'BCMDataExports': 'bcm-data-exports',
})

def main():
    loader = botocore.loaders.create_loader()

    svc_id_to_js_names = json.loads(
        Path('svc_id_to_js_names.json').read_text(),
    )

    svcs: Dict[str, ServiceModel] = {
        svc_name: ServiceModel(loader.load_service_model(svc_name, 'service-2'))
        for svc_name in loader.list_available_services('service-2')
    }

    svc_id_to_boto3_names = {
        ServiceModel(loader.load_service_model(svc_name, 'service-2')).service_id: svc_name
        for svc_name in loader.list_available_services('service-2')
    }

    def boto3_service(svc_name: str) -> botocore.model.ServiceModel:
        client: boto3.client = boto3.client(svc_name)
        svc_model: botocore.model.ServiceModel = client.meta.service_model
        return svc_model


    boto3_svc_info = {}

    # for key, value in cmd_table.items():
    for service_id, boto3_svc_name in svc_id_to_boto3_names.items():
        boto3_svc_info[boto3_svc_name] = {
            "operation_names": svcs[boto3_svc_name].operation_names,
            "model": svcs[boto3_svc_name],
            "js_sdk_name": svc_id_to_js_names.get(service_id, None),
        }

    boto3_iam_mappings = {}

    for boto3_svc_name, info in boto3_svc_info.items():
        boto3_iam_mappings[boto3_svc_name] = {}

        for boto3_op_name in info['operation_names']:
            try:
                map = iam_map['sdk_method_iam_mappings'][f"{info['js_sdk_name']}.{boto3_op_name}"]
                boto3_iam_mappings[boto3_svc_name][boto3_op_name] = map
            except KeyError:
                print(f"NotFound: {info['js_sdk_name']}.{boto3_op_name} -- {boto3_svc_name}.{boto3_op_name}")
                boto3_iam_mappings[boto3_svc_name][boto3_op_name] = [{
                    "note": "NotFound",
                    "action": f"{iam_map['sdk_service_mappings'][info['js_sdk_name']]}:{boto3_op_name}",
                    "js_op": f"{info['js_sdk_name']}.{boto3_op_name}",
                    "resource_mappings": {},
                }]
                continue

    p = Path(__file__).parent/'boto3_iam_map.json'
    p.write_text(
        json.dumps(boto3_iam_mappings, indent=2),
    )

if __name__ == '__main__':
    main()
