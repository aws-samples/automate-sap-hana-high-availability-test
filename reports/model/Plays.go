// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT-0

package model

type Plays struct {
	Play  Play    `json:"play"`
	Tasks []Tasks `json:"tasks"`
}
