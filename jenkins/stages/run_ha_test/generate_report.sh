#!/bin/bash

# Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
# SPDX-License-Identifier: MIT-0

cd reports

rm -f report.html
./report_generator_linux_amd64 -log-file=log.json