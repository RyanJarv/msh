import json
import urllib.request
from pathlib import Path


def main():
    p = Path(__file__).parent / 'map.json'
    if p.exists():
        data = p.read_text()
    else:
        resp = urllib.request.urlopen('https://raw.githubusercontent.com/iann0036/iam-dataset/main/aws/map.json')
        data = resp.read().decode('utf-8')
        p.write_text(data)

    iam_map = json.loads(data)

    iam_svc_to_js_sdk = {}
    js_sdk_to_iam_svc = {}
    js_sdk_to_iam_svc_main = {}

    for full_js_op, iam_actions in iam_map['sdk_method_iam_mappings'].items():
        js_svc_name = full_js_op.split('.')[0]
        for iam_action in iam_actions:
            action_name = iam_action['action']
            iam_svc_name = action_name.split(':')[0]
            iam_svc_to_js_sdk.setdefault(iam_svc_name, set()).add(js_svc_name)
            js_sdk_to_iam_svc.setdefault(js_svc_name, set()).add(iam_svc_name)

            if iam_svc_name not in [
                'iam', 's3', 'logs', 'route53', 'acm', 'eks', 'events', 'sqs', 'ds',
                'organizations', 'sns', 'ec2', 'cloudhsm', 'kms', 'glue', 'iot',
                'amplify', 'secretsmanager', 'cloudformation', 'cloudwatch', 'kinesis',
                'ecs', 'resource-groups', 'fis', 'servicecatalog', 'codecommit',
                'lambda', 'lakeformation', 'elasticloadbalancing', 'ram',
                'cloudfront', 'autoscaling', 'sso', 'codestar-connections', 'lex',
                'sso-directory',  'app-integrations', 'ecr', 'fsx', 'mediaimport',
                'ssm', 'elasticfilesystem', 'aws-marketplace', 'inspector2', 'sms-voice',
                'cleanrooms-ml', 'outposts', 'firehose',
            ]:
                previous = js_sdk_to_iam_svc_main.get(js_svc_name, None)
                if previous and previous != iam_svc_name:
                    print(f"{js_svc_name} -> {iam_svc_name} but it's already set to {previous}")
                    continue
                js_sdk_to_iam_svc_main[js_svc_name] = iam_svc_name

    p = Path(__file__).parent / 'iam_svc_to_js_sdk.json'
    p.write_text(
        json.dumps(iam_svc_to_js_sdk, indent=2, default=list),
    )

    p = Path(__file__).parent / 'js_sdk_to_iam_svc.json'
    p.write_text(
        json.dumps(js_sdk_to_iam_svc, indent=2, default=list),
    )

    p = Path(__file__).parent / 'js_sdk_to_iam_svc_main.json'
    p.write_text(
        json.dumps(js_sdk_to_iam_svc_main, indent=2, default=list),
    )


if __name__ == '__main__':
    main()
