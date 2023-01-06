// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT-0

package html

import (
	"reports/model"
	"strings"
)

func CheckIfPlayFailed(task model.Tasks) bool {
	if checkIsFalsePositiveFailure(task.Task.Name) {
		return false

	} else if checkStderrIsErrored(task.Hosts.AscsHost) ||
		checkStderrIsErrored(task.Hosts.PasHost) ||
		checkStderrIsErrored(task.Hosts.HanaPrimHost) ||
		checkStderrIsErrored(task.Hosts.HanaSecHost) {

		return true

	} else if CheckIsActionFailed(task.Task.Name, task.Hosts.HanaPrimHost) ||
		CheckIsActionFailed(task.Task.Name, task.Hosts.HanaSecHost) ||
		CheckIsActionFailed(task.Task.Name, task.Hosts.AscsHost) ||
		CheckIsActionFailed(task.Task.Name, task.Hosts.PasHost) {

		return true
	}

	return false
}

func checkStderrIsErrored(host model.HostDetails) bool {
	return host.Stderr != "" && !checkCommandContainsIgnoreErrors(host.Cmd)
}

func checkCommandContainsIgnoreErrors(cmd string) bool {
	return strings.Contains(cmd, IGNORE_ERRORS)
}

func CheckIsActionFailed(taskName string, host model.HostDetails) bool {
	if checkIsFalsePositiveFailure(taskName) {
		return false
	}

	return (host.Action == "fail" && host.Failed) ||
		(host.Failed && !checkCommandContainsIgnoreErrors(host.Cmd))
}

func CheckIsLowLight(task model.Tasks) bool {
	if task.Hosts.HanaPrimHost.Action == "shell" ||
		task.Hosts.HanaSecHost.Action == "shell" ||
		task.Hosts.AscsHost.Action == "shell" ||
		task.Hosts.PasHost.Action == "shell" {

		return false
	}
	return !CheckIfPlayFailed(task)
}

func CheckIsActionSkipped(messageCell MessageCell) bool {
	skipped := true
	for _, v := range messageCell.HOSTS_STATE {
		if v.SKIPPED == "false" {
			skipped = false
		}
	}
	return skipped
}

func checkIsFalsePositiveFailure(taskName string) bool {
	specialTaskToIgnore := map[string]string{
		"ansible-role-check-inputs : Ensure non mandatory input parameters were sent as required": "true",
		"ansible-role-bridge-task : Check HA configuration status after self-healing actions":     "true",
	}

	if _, ok := specialTaskToIgnore[taskName]; ok {
		return true
	}

	return false
}
