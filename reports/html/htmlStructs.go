// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT-0

package html

type SummaryRow struct {
	COUNT      string
	NAME       string
	RESULT     string
	DYNAMIC_ID string
}

type DetailsSection struct {
	DYNAMIC_ID   string
	COUNT        string
	PLAY_NAME    string
	DETAILS_ROWS []DetailsRow
	FAILED       string
}

type LessImportantDetailsRow struct {
	TIMESTAMP string
	SCENARIO  string
	FAILED    string
	DETAILS   MessageCell
}

type ImportantDetailsRow struct {
	TIMESTAMP string
	SCENARIO  string
	FAILED    string
	DETAILS   MessageCell
}

type DetailsRow struct {
	LESS_IMPORTANT_DETAILS_ROWS []LessImportantDetailsRow
	IMPORTANT_DETAILS_ROWS      []ImportantDetailsRow
}

type MessageCell struct {
	NAME        string
	HOSTS_STATE []HostDetail
}

type HostDetail struct {
	HOST_NAME     string
	ACTION_NAME   string
	CHANGED       string
	SKIPPED       string
	FAILED        string
	MSG           []string
	ANSIBLE_FACTS AnsibleFacts
	CMD           string
	ATTEMPTS      string
	STARTDATE     string
	ENDDATE       string
	STDERR_LINES  []string
	STDOUT_LINES  []string
}

type AnsibleFacts struct {
	ANSIBLE_ALL_IPV4_ADDRESSES []string
	ANSIBLE_ALL_IPV6_ADDRESSES []string
	ANSIBLE_DISTRIBUTION       string
	ANSIBLE_DOMAIN             string
	ANSIBLE_ENV                AnsibleEnv
	ANSIBLE_FQDN               string
	ANSIBLE_HOSTNAME           string
	ANSIBLE_NODE_NAME          string
	ANSIBLE_OS_FAMILY          string
}

type AnsibleEnv struct {
	HOME         string
	PATH         string
	SUDO_GID     string
	SUDO_USER    string
	TERM         string
	USER         string
	SUDO_UID     string
	LOGNAME      string
	PWD          string
	SUDO_COMMAND string
	SHELL        string
}

type PageData struct {
	PAGE_TITLE       string
	NAV_TITLE        string
	FAILED_COUNT     string
	SUCCESSFUL_COUNT string

	SUMMARY_ROWS    []SummaryRow
	DETAILS_SECTION []DetailsSection
}
