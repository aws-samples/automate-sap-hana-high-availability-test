#!/bin/bash

# Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
# SPDX-License-Identifier: MIT-0

location=$(PWD)
osFamily=`uname`
processorArchitecture=`uname -p`

cd reports

if [[ "$osFamily" == "Darwin" ]]; then
    if [[ "$processorArchitecture" == "arm" ]]; then
        ./bin/report_generator_darwin_arm64 -log-file=$location/log.json
    else
        ./bin/report_generator_darwin_amd64 -log-file=$location/log.json
    fi

else
    if [[ "$processorArchitecture" == "arm" ]]; then
        ./bin/report_generator_linux_arm64 -log-file=$location/log.json
    else
        ./bin/report_generator_linux_amd64 -log-file=$location/log.json
    fi
fi

fileName=$(ls -t | awk 'NR==1{print $1}')
fileFullPath="$location/reports/$fileName"
echo "Test finished! Your report is stored in $fileFullPath"

cd ..