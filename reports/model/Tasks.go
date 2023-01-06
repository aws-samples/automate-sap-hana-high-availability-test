// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT-0

package model

type Tasks struct {
	Hosts Hosts `json:"hosts"`
	Task  Task  `json:"task"`
}
