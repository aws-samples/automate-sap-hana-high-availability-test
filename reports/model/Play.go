// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT-0

package model

type Play struct {
	Duration Duration `json:"duration"`
	Id       string   `json:"id"`
	Name     string   `json:"name"`
}
