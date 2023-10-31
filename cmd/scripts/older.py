#!/usr/bin/env python3

from datetime import datetime, timezone
import os


def lambda_handler(event, context):
    input_date_str = event['UsageOperationUpdateTime']
    print(input_date_str)
    input_date = datetime.fromisoformat(input_date_str)

    current_date = datetime.now(timezone.utc)
    delta = current_date - input_date

    return delta.total_seconds() > int(os.environ['ARG1'])