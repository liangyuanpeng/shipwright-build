// Copyright The Shipwright Contributors
//
// SPDX-License-Identifier: Apache-2.0

package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestImageProcessingCmd(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Image Processing Command Suite")
}
