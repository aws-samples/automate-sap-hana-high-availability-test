// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT-0

package model

type Hosts struct {
	AscsHost        HostDetails `json:"ascs-node"`
	HanaPrimHost    HostDetails `json:"hana-prim"`
	HanaSecHost     HostDetails `json:"hana-sec"`
	PasHost         HostDetails `json:"pas-node"`
	HanaPrimHostNew HostDetails `json:"hana-prim-new"`
	HanaSecHostNew  HostDetails `json:"hana-sec-new"`
}
