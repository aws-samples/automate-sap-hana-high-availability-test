#!/bin/bash

# Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
# SPDX-License-Identifier: MIT-0

terraform -chdir="$PWD/$TERRAFORM_FOLDER_NAME" \
                init \
                -backend-config "bucket=$BUCKET_NAME_CHKD" \
                -backend-config "key=environment.tfstate" \
                -backend-config "region=$AWS_REGION_CHKD" \
                -backend-config "access_key=$AWS_ACCOUNT_CREDENTIALS_USR" \
                -backend-config "secret_key=$AWS_ACCOUNT_CREDENTIALS_PSW" \
                -backend-config "dynamodb_table=$DYNAMO_TABLE_TF_LOCK" \
                > /dev/null

if [ $? -ne 0 ]; then
    echo "Error with Terraform Init. Please check again"
    exit 100
fi