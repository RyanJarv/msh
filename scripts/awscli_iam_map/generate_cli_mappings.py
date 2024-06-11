import json
from pathlib import Path

import awscli.customizations.cloudfront
import awscli.customizations.cloudtrail
import awscli.customizations.cloudformation
import awscli.customizations.waiters
import awscli.clidriver
import urllib.request

import botocore.model
import boto3

def main():
    driver = awscli.clidriver.create_clidriver()
    cmd_table = driver._get_command_table()

    p = Path(__file__).parent/'map.json'
    if p.exists():
        data = p.read_text()
    else:
        resp = urllib.request.urlopen('https://raw.githubusercontent.com/iann0036/iam-dataset/main/aws/map.json')
        data = resp.read().decode('utf-8')
        p.write_text(data)

    iam_map = json.loads(data)
    # iam_map['sdk_method_iam_mappings']


    svc_id_to_js_names = json.loads(
        Path('svc_id_to_js_names.json').read_text(),
    )


    def boto3_service(svc_name: str) -> botocore.model.ServiceModel:
        client: boto3.client = boto3.client(svc_name)
        svc_model: botocore.model.ServiceModel = client.meta.service_model
        return svc_model


    cli_svc_info = {}

    for key, value in cmd_table.items():
        if key in ['configure', 'history']:
            continue

        if hasattr(value, 'service_model'):
            cli_svc_info[key] = {
                "cmd_table": value,
                "model": value.service_model,
                "js_sdk_name": svc_id_to_js_names[value.service_model.service_id],
            }
        else:
            svc_model = boto3_service(key)
            cli_svc_info[key] = {
                "cmd_table": value,
                "model": svc_model,
                "js_sdk_name": svc_id_to_js_names[svc_model.service_id],
            }

    awscli_iam_mappings = {}

    for cli_svc_name, v in cli_svc_info.items():
        try:
            svc_cmd_table = v['cmd_table']._get_command_table()
        except AttributeError:
            print(f"Could not find command table: {cli_svc_name}")
            continue


        awscli_iam_mappings[cli_svc_name] = {}

        for cli_cmd_name, cli_cmd in svc_cmd_table.items():
            cli_cmd: awscli.clidriver.ServiceOperation

            try:
                op_model: botocore.model.OperationModel = cli_cmd._operation_model
            except AttributeError:
                if type(cli_cmd) in [
                    awscli.customizations.cloudformation.package.PackageCommand,
                    awscli.customizations.cloudfront.SignCommand,
                    awscli.customizations.cloudtrail.subscribe.CloudTrailSubscribe,
                    awscli.customizations.cloudtrail.subscribe.CloudTrailUpdate,
                    awscli.customizations.cloudtrail.validation.CloudTrailValidateLogs,
                ]:
                    continue


                op_models = []

                for name, subcmd in cli_cmd.subcommand_table.items():
                    try:
                        op_models.append(subcmd._operation_model)
                    except AttributeError:
                        continue

                if op_models:
                    # TODO: probably need to handle all of these
                    op_model = op_models[0]
                else:
                    print(f"Could not find operation_model: {cli_svc_name}.{cli_cmd_name}: {type(cli_cmd)}")
                    continue

            try:
                map = iam_map['sdk_method_iam_mappings'][f"{v['js_sdk_name']}.{op_model.name}"]
                awscli_iam_mappings[cli_svc_name][cli_cmd_name] = map
            except KeyError:
                print(f"NotFound: {v['js_sdk_name']}.{op_model.name}")
                awscli_iam_mappings[cli_svc_name][cli_cmd_name] = [{
                    "note": "NotFound",
                    "action": f"{v['model'].service_id}:{op_model.name}",
                    "resource_mappings": {},
                }]
                continue


    Path('../../cmd/aws/cli_iam_map.json').write_text(
        json.dumps(awscli_iam_mappings, indent=2),
    )


    # for name, mapping in iam_map['sdk_method_iam_mappings'].items():
    #     js_name, action_name = name.split('.')
    #     svc_id = js_names_to_svc_ids[js_name]
    #     cmd_name = svc_ids_cmd_names[svc_id]
    #
    #     xform_name
    #

if __name__ == '__main__':
    main()
