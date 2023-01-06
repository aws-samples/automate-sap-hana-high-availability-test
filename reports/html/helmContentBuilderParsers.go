// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT-0

package html

func ParseFailed(playFailed bool) string {
	if playFailed {
		return "FAIL"
	}
	return "PASS"
}
