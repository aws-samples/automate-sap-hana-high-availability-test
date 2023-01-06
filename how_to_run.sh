#!/bin/bash

# Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
# SPDX-License-Identifier: MIT-0 

rm -f log.json

sudo ansible-playbook main.yaml \
            --inventory-file hosts.yaml \
            --extra-vars "@var_file.yaml"

location=$(PWD)
cd reports
./report_generator -log-file=$location/log.json
open report.html
cd ..