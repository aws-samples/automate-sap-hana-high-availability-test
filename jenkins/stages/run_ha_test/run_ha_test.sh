#!/bin/bash

# Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
# SPDX-License-Identifier: MIT-0

ansibleHanaDir="$PWD"
export ANSIBLE_DIR=$ansibleHanaDir

# ------------------------------------------------------------------
# Grab data from Terraform
# ------------------------------------------------------------------
hana_public_ips=$(terraform -chdir="$PWD/$TERRAFORM_FOLDER_NAME" output -json hana_instance_public_ips)
if [ -z "$hana_public_ips" ]; then
    echo "No Hana instance IPs were found. Please check Terraform step"
    exit 100
fi
export HOSTS_IPS=$hana_public_ips

ascs_private_ip=$(terraform -chdir="$PWD/$TERRAFORM_FOLDER_NAME" output -json ascs_instance_public_ips | jq -r '.[0]')
if [ -z "$ascs_private_ip" ]; then
    echo "No ASCS instance IP was found. Please check Terraform step"
    exit 101
fi

pas_private_ip=$(terraform -chdir="$PWD/$TERRAFORM_FOLDER_NAME" output -json app_instance_public_ips | jq -r '.[0]')
if [ -z "$pas_private_ip" ]; then
    echo "No PAS instance IP was found. Please check Terraform step"
    exit 102
fi

# ------------------------------------------------------------------
# Create hosts file
# ------------------------------------------------------------------
# Create hosts_runtime.yml file
FOLDER_PATH="./jenkins/stages/run_ha_test"
$FOLDER_PATH/create_hosts_file.sh
if [ $? -ne 0 ]; then
    echo "There was an error creating the hosts file. Please check again"
    exit 104
fi

hostsFile="$ansibleHanaDir/hosts_runtime.yml"

export VAR_FILE_FULL_PATH="$ansibleHanaDir/var_file.yaml"
rm $VAR_FILE_FULL_PATH 2> /dev/null
touch $VAR_FILE_FULL_PATH

# ------------------------------------------------------------------
# Add variables to VAR_FILE_FULL_PATH
# ------------------------------------------------------------------
echo "---" >> $VAR_FILE_FULL_PATH
echo "INPUT_HANA_SID: $HANA_SID_CHKD" >> $VAR_FILE_FULL_PATH
echo "INPUT_HANA_INSTANCE_NUMBER: $HANA_INSTANCE_NUMBER_CHKD" >> $VAR_FILE_FULL_PATH
echo "INPUT_HANA_SYNC_MODE: SYNC" >> $VAR_FILE_FULL_PATH
echo "INPUT_ASCS_SID: $SAP_SID_CHKD" >> $VAR_FILE_FULL_PATH
echo "INPUT_ASCS_INSTANCE_NUMBER: $ASCS_INSTANCE_NUMBER_CHKD" >> $VAR_FILE_FULL_PATH
echo "INPUT_PAS_SID: $SAP_SID_CHKD" >> $VAR_FILE_FULL_PATH
echo "INPUT_PAS_INSTANCE_NUMBER: $PAS_INSTANCE_NUMBER_CHKD" >> $VAR_FILE_FULL_PATH
echo "INPUT_CHECK_R3_TRANS: true" >> $VAR_FILE_FULL_PATH

echo "INPUT_SYSTEM_USER: SYSTEM" >> $VAR_FILE_FULL_PATH
echo "INPUT_SYSTEM_PASSWORD: $MASTER_PASSWORD_CHKD" >> $VAR_FILE_FULL_PATH

echo "INPUT_AWS_REGION: $AWS_REGION_CHKD" >> $VAR_FILE_FULL_PATH
echo "INPUT_AWS_CLI_PROFILE: $CLI_PROFILE_CHKD" >> $VAR_FILE_FULL_PATH

echo "INPUT_PRIVATE_SSH_KEY: $SSH_KEYPAIR_FILE_CHKD" >> $VAR_FILE_FULL_PATH

# ------------------------------------------------------------------
# Run playbook
# ------------------------------------------------------------------
ANSIBLE_HOST_KEY_CHECKING=False
ANSIBLE_BECOME_EXE="sudo su -"

rm -f log.json

ansible-playbook $ansibleHanaDir/main.yml \
                    --inventory-file "$hostsFile" \
                    --extra-vars "@$VAR_FILE_FULL_PATH"

result_value=$?

if [[ $result_value == 4 ]]; then
    echo "ERROR! It looks like at least one of the hosts were not reachable. Double check your SSH_KEYPAIR_NAME, SSH_KEYPAIR_FILE variables and if your IP is allowed to go through port 22 on the security groups used by your instances"
    exit 110
elif [[ $result_value != 0 ]]; then
    echo "ERROR! There was an error during installation. Check the logs and try again"
    exit 111
fi

exit 0