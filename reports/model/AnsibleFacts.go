// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT-0

package model

type AnsibleFacts struct {
	AnsibleAllIPv4Addresses []string   `json:"ansible_all_ipv4_addresses"`
	AnsibleAllIPv6Addresses []string   `json:"ansible_all_ipv6_addresses"`
	AnsibleDistribution     string     `json:"ansible_distribution"`
	AnsibleDomain           string     `json:"ansible_domain"`
	AnsibleEnv              AnsibleEnv `json:"ansible_env"`
	AnsibleFQDN             string     `json:"ansible_fqdn"`
	AnsibleHostname         string     `json:"ansible_hostname"`
	AnsibleNodeName         string     `json:"ansible_nodename"`
	AnsibleOSFamily         string     `json:"ansible_os_family"`
}
