// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT-0

package html

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"reports/model"
)

func buildReport(ansibleLog model.AnsibleLog) string {

	files := getTemplateFiles()

	dataForTemplates := buildContentForTemplates(ansibleLog)

	var templatedData bytes.Buffer
	ts, err := template.ParseFiles(files...)
	check(err)

	err = ts.ExecuteTemplate(&templatedData, "html", dataForTemplates)
	check(err)

	templatedHTML := templatedData.String()

	return templatedHTML
}

func getTemplateFiles() []string {
	return []string{
		"templates/html.html",
		"templates/body.html",
		"templates/head.html",
		"templates/nav.html",
		"templates/summary_div.html",
		"templates/summary_table.html",
		"templates/summary_table_row.html",
		"templates/details.html",
		"templates/details_row.html",
		"templates/details_row_important.html",
		"templates/details_row_less_important.html",
		"templates/message_cell.html",
	}
}

func ExportToHTML(ansibleLog model.AnsibleLog) {
	htmlData := buildReport(ansibleLog)

	os.WriteFile("report.html", []byte(htmlData), 0644)
}

func check(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
