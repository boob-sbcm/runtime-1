// Copyright (c) 2018 IBM
//
// SPDX-License-Identifier: Apache-2.0
//

package main

import "testing"

func getExpectedHostDetails(tmpdir string) (HostInfo, error) {
	expectedVendor := ""
	expectedModel := "POWER8"
	return genericGetExpectedHostDetails(tmpdir, expectedVendor, expectedModel)
}

func TestEnvGetEnvInfoSetsCPUType(t *testing.T) {
	testEnvGetEnvInfoSetsCPUTypeGeneric(t)
}
