package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"reports/files"
	"reports/html"
	"reports/model"
)

func main() {

	validateParams(os.Args)
	fileName := getMandatoryParam(os.Args, "-log-file")

	logFileBytes := files.Read(fileName)
	logFile := string(logFileBytes)

	validateFile(logFile)

	logFile = files.TrimToJson(logFile)

	ansibleLog := model.AnsibleLog{}
	json.Unmarshal([]byte(logFile), &ansibleLog)

	html.ExportToHTML(ansibleLog)
}

func validateFile(logFile string) {
	if files.IsEmptyFile(logFile) {
		fmt.Println("File is empty. Check your -log-file param")
		os.Exit(1)
	}
}

func dismissCLIName(params []string) []string {
	realParams := []string{}
	for k, v := range params {
		if k != 0 {
			realParams = append(realParams, v)
		}
	}
	return realParams
}

func validateParams(params []string) {
	params = dismissCLIName(params)

	if params == nil || len(params) < 1 {
		fmt.Println("Invalid parameter options. You must enter mandatory parameters: -log-file")
		os.Exit(1)
	}

	validateParamLogFile(params)
}

func validateParamLogFile(params []string) {
	for _, v := range params {
		if len(strings.Split(v, "=")) != 2 {
			fmt.Println("Parameter \"" + v + "\" sent in an invalid format. It should be key=value")
			os.Exit(1)
		}
	}
}

func getMandatoryParam(params []string, paramName string) string {
	for _, v := range params {
		keyValue := strings.Split(v, "=")

		if keyValue[0] == paramName {
			return keyValue[1]
		}
	}

	fmt.Println("Mandatory parameter \"" + paramName + "\" not found")
	os.Exit(1)

	return ""
}
