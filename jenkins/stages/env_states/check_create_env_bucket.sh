#!/bin/bash

# Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
# SPDX-License-Identifier: MIT-0

# ------------------------------------------------------------------
# Check if the variable is present. If not, send back default value
# ------------------------------------------------------------------
aws s3 ls $BUCKET_NAME_CHKD --profile "$CLI_PROFILE_CHKD" > /dev/null

if [ $? -ne 0 ]; then
    echo "The bucket $BUCKET_NAME_CHKD doesn't exist and therefore no Terraform state file could be found"
    exit 1
fi

exit 0