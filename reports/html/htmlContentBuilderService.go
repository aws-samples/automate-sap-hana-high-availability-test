// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT-0

package html

import (
	"reports/model"
	"strconv"
)

func BuildLessImportant(durationEnd string, scenarioName string, detailMessage MessageCell, failed bool) LessImportantDetailsRow {
	return LessImportantDetailsRow{
		TIMESTAMP: durationEnd,
		SCENARIO:  scenarioName,
		FAILED:    strconv.FormatBool(failed),
		DETAILS:   detailMessage,
	}
}

func BuildImportant(durationEnd string, scenarioName string, detailMessage MessageCell, failed bool) ImportantDetailsRow {
	return ImportantDetailsRow{
		TIMESTAMP: durationEnd,
		SCENARIO:  scenarioName,
		FAILED:    strconv.FormatBool(failed),
		DETAILS:   detailMessage,
	}
}

func RemovePlaysWithoutTasks(plays []model.Plays) []model.Plays {
	playsToReturn := []model.Plays{}

	for _, playValue := range plays {
		if len(playValue.Tasks) > 0 {
			playsToReturn = append(playsToReturn, playValue)
		}
	}

	return playsToReturn
}

func UpdateReplacedHosts(tasks []model.Tasks) []model.Tasks {
	tasksToReturn := []model.Tasks{}

	for _, v := range tasks {

		if v.Hosts.HanaPrimHost.Action == "" {
			v.Hosts.HanaPrimHost = v.Hosts.HanaPrimHostNew
		}
		if v.Hosts.HanaSecHost.Action == "" {
			v.Hosts.HanaSecHost = v.Hosts.HanaSecHostNew
		}

		tasksToReturn = append(tasksToReturn, v)
	}

	return tasksToReturn
}
