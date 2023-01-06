// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT-0

package files

import (
	"fmt"
	"os"
)

func Read(path string) (fileContents []byte) {
	checkIfFileExists(path)

	logFile, err := os.ReadFile(path)
	check(err)

	return logFile
}

func checkIfFileExists(fileName string) {
	if _, err := os.Stat(fileName); err != nil {
		fmt.Println("File name doesn't exist. Please check your \"-log-file\" param")
		os.Exit(1)
	}
}

func check(e error) {
	if e != nil {
		fmt.Println(ERROR_READING_LOG_FILE)
	}
}
