// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT-0

package files

import (
	"strings"
)

func TrimToJson(logFile string) (trimmedLogFile string) {
	logFile = removeLeadingChars(logFile)
	logFile = removeTrailingChars(logFile)

	return logFile
}

func removeLeadingChars(logFile string) string {
	strLength := len(logFile)

	for pos, char := range logFile {
		if char == []rune("{")[0] {
			logFile = substr(logFile, pos, strLength)
			break
		}
	}

	return logFile
}

func removeTrailingChars(logFile string) string {
	for i := len(logFile) - 1; i >= 0; i-- {
		if rune(logFile[i]) == []rune("}")[0] {
			logFile = substr(logFile, 0, i+1)
			break
		}
	}

	return logFile
}

func IsEmptyFile(logFile string) bool {
	if strings.TrimSpace(logFile) == "" {
		return true
	}
	return false
}

func substr(input string, start int, length int) string {
	asRunes := []rune(input)

	if start >= len(asRunes) {
		return ""
	}

	if start+length > len(asRunes) {
		length = len(asRunes) - start
	}

	return string(asRunes[start : start+length])
}
