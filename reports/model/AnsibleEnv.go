// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT-0

package model

type AnsibleEnv struct {
	HOME         string `json:"HOME"`
	PATH         string `json:"PATH"`
	SUDO_GID     string `json:"SUDO_GID"`
	SUDO_USER    string `json:"SUDO_USER"`
	TERM         string `json:"TERM"`
	USER         string `json:"USER"`
	SUDO_UID     string `json:"SUDO_UID"`
	LOGNAME      string `json:"LOGNAME"`
	PWD          string `json:"PWD"`
	SUDO_COMMAND string `json:"SUDO_COMMAND"`
	SHELL        string `json:"SHELL"`
}
