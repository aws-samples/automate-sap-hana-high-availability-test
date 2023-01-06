// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT-0

package model

type Stat struct {
	AscsHost     PlayStats `json:"ascs-node"`
	HanaPrimHost PlayStats `json:"hana-prim"`
	HanaSecHost  PlayStats `json:"hana-sec"`
	PasHost      PlayStats `json:"pas-node"`
}
