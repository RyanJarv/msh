#!/usr/bin/env python3

import datetime


def lambda_handler(event, context):
    input_date_str = event['date']
    input_date = datetime.datetime.fromisoformat(input_date_str)

    current_date = datetime.datetime.utcnow()
    delta = current_date - input_date

    return {
        'result': delta.total_seconds() < event['seconds']
    }
