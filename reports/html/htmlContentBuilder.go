// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT-0

package html

import (
	"reports/model"
	"strconv"
)

func buildContentForTemplates(ansibleLog model.AnsibleLog) PageData {
	summaryRows := []SummaryRow{}
	detailsSection := []DetailsSection{}
	numberSuccess := 0
	numberFailure := 0

	plays := RemovePlaysWithoutTasks(ansibleLog.Plays)

	for k, playValue := range plays {
		playFailed := false
		taskFailed := false
		currentGroupIsLowLight := true

		detailsRows := []DetailsRow{}
		lessImportantDetailsRows := []LessImportantDetailsRow{}
		importantDetailsRows := []ImportantDetailsRow{}

		tasksToIterate := UpdateReplacedHosts(playValue.Tasks)

		for _, v := range tasksToIterate {
			if !playFailed {
				playFailed = CheckIfPlayFailed(v)
			}
			taskFailed = CheckIfPlayFailed(v)

			detailMessage := parseHostsState(v)
			if CheckIsActionSkipped(detailMessage) {
				continue
			}

			isLowLight := CheckIsLowLight(v)

			if currentGroupIsLowLight != isLowLight {
				if len(lessImportantDetailsRows) > 0 {
					detailsGroup := DetailsRow{lessImportantDetailsRows, nil}
					detailsRows = append(detailsRows, detailsGroup)
					lessImportantDetailsRows = []LessImportantDetailsRow{}
				} else {
					detailsGroup := DetailsRow{nil, importantDetailsRows}
					detailsRows = append(detailsRows, detailsGroup)
					importantDetailsRows = []ImportantDetailsRow{}
				}

				currentGroupIsLowLight = isLowLight
			}

			if isLowLight {
				row := BuildLessImportant(playValue.Play.Duration.End, playValue.Play.Name, detailMessage, taskFailed)
				lessImportantDetailsRows = append(lessImportantDetailsRows, row)
			} else {
				row := BuildImportant(playValue.Play.Duration.End, playValue.Play.Name, detailMessage, taskFailed)
				importantDetailsRows = append(importantDetailsRows, row)
			}
		}
		if len(lessImportantDetailsRows) > 0 {
			detailsGroup := DetailsRow{lessImportantDetailsRows, nil}
			detailsRows = append(detailsRows, detailsGroup)
		} else {
			detailsGroup := DetailsRow{nil, importantDetailsRows}
			detailsRows = append(detailsRows, detailsGroup)
		}

		lessImportantDetailsRows = lessImportantDetailsRows[:0]
		importantDetailsRows = importantDetailsRows[:0]

		detailsSectionRow := DetailsSection{}
		detailsSectionRow = DetailsSection{
			DETAILS_ROWS: detailsRows,
			COUNT:        strconv.Itoa(k + 1),
			PLAY_NAME:    playValue.Play.Name,
			DYNAMIC_ID:   "controller" + strconv.Itoa(k),
			FAILED:       ParseFailed(playFailed),
		}

		detailsSection = append(detailsSection, detailsSectionRow)

		summaryRow := SummaryRow{
			COUNT:      strconv.Itoa(k + 1),
			NAME:       playValue.Play.Name,
			RESULT:     ParseFailed(playFailed),
			DYNAMIC_ID: "controller" + strconv.Itoa(k),
		}

		summaryRows = append(summaryRows, summaryRow)

		if playFailed {
			numberFailure += 1
		} else {
			numberSuccess += 1
		}
	}

	return PageData{
		PAGE_TITLE:       PAGE_TITLE,
		NAV_TITLE:        NAV_TITLE,
		FAILED_COUNT:     strconv.Itoa(numberFailure),
		SUCCESSFUL_COUNT: strconv.Itoa(numberSuccess),

		SUMMARY_ROWS:    summaryRows,
		DETAILS_SECTION: detailsSection,
	}
}
