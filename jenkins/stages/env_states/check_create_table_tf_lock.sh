#!/bin/bash

# Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
# SPDX-License-Identifier: MIT-0

# ------------------------------------------------------------------
# Check if the variable is present. If not, send back default value
# ------------------------------------------------------------------
table_exists=$(aws dynamodb list-tables --profile "$CLI_PROFILE_CHKD" | grep -c "$DYNAMO_TABLE_NAME_CHKD")

if [ "$table_exists" -eq 0 ]; then
    echo "The DynamoDB table $DYNAMO_TABLE_NAME_CHKD doesn't exist and therefore no Terraform state file could be found"
    exit 1
fi

exit 0