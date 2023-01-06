// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT-0

package html

import (
	"fmt"
	"reports/model"
	"strconv"
)

func parseHostsState(task model.Tasks) MessageCell {
	return MessageCell{
		NAME: task.Task.Name,
		HOSTS_STATE: []HostDetail{
			{
				HOST_NAME:     "HANA Pri",
				ACTION_NAME:   task.Hosts.HanaPrimHost.Action,
				CHANGED:       strconv.FormatBool(task.Hosts.HanaPrimHost.Changed),
				SKIPPED:       strconv.FormatBool(task.Hosts.HanaPrimHost.Skipped),
				FAILED:        strconv.FormatBool(CheckIsActionFailed(task.Task.Name, task.Hosts.HanaPrimHost)),
				MSG:           parseInterfaceMessage(task.Hosts.HanaPrimHost.Msg),
				ANSIBLE_FACTS: buildAnsibleFacts(task.Hosts.HanaPrimHost.AnsibleFacts),
				CMD:           task.Hosts.HanaPrimHost.Cmd,
				ATTEMPTS:      parseWaitTimeToString(task.Hosts.HanaPrimHost.Attempts),
				STARTDATE:     task.Hosts.HanaPrimHost.StartDate,
				ENDDATE:       task.Hosts.HanaPrimHost.EndDate,
				STDERR_LINES:  task.Hosts.HanaPrimHost.Stderr_lines,
				STDOUT_LINES:  task.Hosts.HanaPrimHost.Stdout_lines,
			},
			{
				HOST_NAME:     "HANA Sec",
				ACTION_NAME:   task.Hosts.HanaSecHost.Action,
				CHANGED:       strconv.FormatBool(task.Hosts.HanaSecHost.Changed),
				SKIPPED:       strconv.FormatBool(task.Hosts.HanaSecHost.Skipped),
				FAILED:        strconv.FormatBool(CheckIsActionFailed(task.Task.Name, task.Hosts.HanaSecHost)),
				MSG:           parseInterfaceMessage(task.Hosts.HanaSecHost.Msg),
				ANSIBLE_FACTS: buildAnsibleFacts(task.Hosts.HanaSecHost.AnsibleFacts),
				CMD:           task.Hosts.HanaSecHost.Cmd,
				ATTEMPTS:      parseWaitTimeToString(task.Hosts.HanaSecHost.Attempts),
				STARTDATE:     task.Hosts.HanaSecHost.StartDate,
				ENDDATE:       task.Hosts.HanaSecHost.EndDate,
				STDERR_LINES:  task.Hosts.HanaSecHost.Stderr_lines,
				STDOUT_LINES:  task.Hosts.HanaSecHost.Stdout_lines,
			},
			{
				HOST_NAME:     "ASCS",
				ACTION_NAME:   task.Hosts.AscsHost.Action,
				CHANGED:       strconv.FormatBool(task.Hosts.AscsHost.Changed),
				SKIPPED:       strconv.FormatBool(task.Hosts.AscsHost.Skipped),
				FAILED:        strconv.FormatBool(CheckIsActionFailed(task.Task.Name, task.Hosts.AscsHost)),
				MSG:           parseInterfaceMessage(task.Hosts.AscsHost.Msg),
				ANSIBLE_FACTS: buildAnsibleFacts(task.Hosts.AscsHost.AnsibleFacts),
				CMD:           task.Hosts.AscsHost.Cmd,
				ATTEMPTS:      parseWaitTimeToString(task.Hosts.AscsHost.Attempts),
				STARTDATE:     task.Hosts.AscsHost.StartDate,
				ENDDATE:       task.Hosts.AscsHost.EndDate,
				STDERR_LINES:  task.Hosts.AscsHost.Stderr_lines,
				STDOUT_LINES:  task.Hosts.AscsHost.Stdout_lines,
			},
			{
				HOST_NAME:     "PAS",
				ACTION_NAME:   task.Hosts.PasHost.Action,
				CHANGED:       strconv.FormatBool(task.Hosts.PasHost.Changed),
				SKIPPED:       strconv.FormatBool(task.Hosts.PasHost.Skipped),
				FAILED:        strconv.FormatBool(CheckIsActionFailed(task.Task.Name, task.Hosts.PasHost)),
				MSG:           parseInterfaceMessage(task.Hosts.PasHost.Msg),
				ANSIBLE_FACTS: buildAnsibleFacts(task.Hosts.PasHost.AnsibleFacts),
				CMD:           task.Hosts.PasHost.Cmd,
				ATTEMPTS:      parseWaitTimeToString(task.Hosts.PasHost.Attempts),
				STARTDATE:     task.Hosts.PasHost.StartDate,
				ENDDATE:       task.Hosts.PasHost.EndDate,
				STDERR_LINES:  task.Hosts.PasHost.Stderr_lines,
				STDOUT_LINES:  task.Hosts.PasHost.Stdout_lines,
			},
		},
	}
}

func parseInterfaceMessage(msg interface{}) []string {
	stringList := []string{}

	switch msg.(type) {
	case string:
		if msg != "" {
			stringList = append(stringList, fmt.Sprintf("%v", msg))
		}
		return stringList
	case []interface{}:
		for _, v := range msg.([]interface{}) {
			stringList = append(stringList, v.(string))
		}
		return stringList

	default:
		return stringList
	}
}

func buildAnsibleFacts(ansibleFacts model.AnsibleFacts) AnsibleFacts {
	return AnsibleFacts{
		ANSIBLE_ALL_IPV4_ADDRESSES: ansibleFacts.AnsibleAllIPv4Addresses,
		ANSIBLE_ALL_IPV6_ADDRESSES: ansibleFacts.AnsibleAllIPv6Addresses,
		ANSIBLE_DISTRIBUTION:       ansibleFacts.AnsibleDistribution,
		ANSIBLE_DOMAIN:             ansibleFacts.AnsibleDomain,
		ANSIBLE_ENV:                buildAnsibleEnv(ansibleFacts.AnsibleEnv),
		ANSIBLE_FQDN:               ansibleFacts.AnsibleFQDN,
		ANSIBLE_HOSTNAME:           ansibleFacts.AnsibleHostname,
		ANSIBLE_NODE_NAME:          ansibleFacts.AnsibleNodeName,
		ANSIBLE_OS_FAMILY:          ansibleFacts.AnsibleOSFamily,
	}
}

func buildAnsibleEnv(ansibleEnv model.AnsibleEnv) AnsibleEnv {
	return AnsibleEnv{
		HOME:         ansibleEnv.HOME,
		PATH:         ansibleEnv.PATH,
		SUDO_GID:     ansibleEnv.SUDO_GID,
		SUDO_USER:    ansibleEnv.SUDO_USER,
		TERM:         ansibleEnv.TERM,
		USER:         ansibleEnv.USER,
		SUDO_UID:     ansibleEnv.SUDO_UID,
		LOGNAME:      ansibleEnv.LOGNAME,
		PWD:          ansibleEnv.PWD,
		SUDO_COMMAND: ansibleEnv.SUDO_COMMAND,
		SHELL:        ansibleEnv.SHELL,
	}
}

func parseWaitTimeToString(attempts int) string {
	attemptsInSeconds := calculateAttemptsToSeconds(attempts)

	if attemptsInSeconds < 5 {
		return "< 5s"
	} else if attemptsInSeconds < 60 {
		return "< 1min"
	} else {
		return strconv.Itoa(attemptsInSeconds/60) + "min"
	}
}

func calculateAttemptsToSeconds(attempts int) int {
	return attempts * 5
}
