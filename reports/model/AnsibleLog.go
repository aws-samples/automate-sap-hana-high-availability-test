// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT-0

package model

type AnsibleLog struct {
	Custom_stats        map[string]interface{} `json:"custom_stats"`
	Global_custom_stats map[string]interface{} `json:"global_custom_stats"`
	Plays               []Plays                `json:"plays"`
	Stats               Stat                   `json:"stats"`
}
