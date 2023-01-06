#!/bin/bash

# Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
# SPDX-License-Identifier: MIT-0

FOLDER_PATH="./jenkins-pipelines/sap-pipeline-hana-pas-ascs/stages/prepare/check_aws_resources"

# Check KeyPair
$FOLDER_PATH/check_key_pair.sh
if [ $? -ne 0 ]; then
    exit 104
fi

exit 0