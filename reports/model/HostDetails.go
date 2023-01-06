// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT-0

package model

type HostDetails struct {
	Action       string       `json:"action"`
	Changed      bool         `json:"changed"`
	Skipped      bool         `json:"skipped"`
	Failed       bool         `json:"failed"`
	Msg          interface{}  `json:"msg"`
	AnsibleFacts AnsibleFacts `json:"ansible_facts"`
	Cmd          string       `json:"cmd"`
	Attempts     int          `json:"attempts"`
	StartDate    string       `json:"start"`
	EndDate      string       `json:"end"`
	Stderr       string       `json:"STDERR"`
	Stderr_lines []string     `json:"STDERR_LINES"`
	Stdout       string       `json:"STDOUT"`
	Stdout_lines []string     `json:"STDOUT_LINES"`
}
