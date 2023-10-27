#!/usr/bin/env python3
# Borrowed from https://gist.github.com/dagrz/9b006d116302ba9dc4a1234b867c30c7, thanks dagrz!

import json
import sys
import awscli.clidriver
from io import StringIO

driver = awscli.clidriver.create_clidriver()


def lambda_handler(event, context):
    old_stdout = sys.stdout
    sys.stdout = mystdout = StringIO()

    driver.main(args=event['command'])

    sys.stdout = old_stdout

    mystdout.seek(0)

    return json.loads(mystdout.read())


if __name__ == '__main__':
    lambda_handler({'command': sys.argv[1:]}, {})
    sys.exit(0)
