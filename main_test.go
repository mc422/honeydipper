// Copyright 2019 Honey Science Corporation
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, you can obtain one at http://mozilla.org/MPL/2.0/.

package main

import (
	"github.com/honeydipper/honeydipper/pkg/dipper"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	if dipper.Logger == nil {
		logFile, err := os.Create("test.log")
		if err != nil {
			panic(err)
		}
		dipper.GetLogger("test", "INFO", logFile, logFile)
	}
	os.Exit(m.Run())
}
