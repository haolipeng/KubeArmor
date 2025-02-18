// Copyright 2021 Authors of KubeArmor
// SPDX-License-Identifier: Apache-2.0

package feeder

import (
	"testing"

	tp "github.com/kubearmor/KubeArmor/KubeArmor/types"
)

func TestFeeder(t *testing.T) {
	// node
	node := tp.Node{}
	node.NodeName = "nodeName"
	node.NodeIP = "nodeIP"

	// create logger
	logger := NewFeeder("Default", node, "32767", "none")
	if logger == nil {
		t.Log("[FAIL] Failed to create logger")
		return
	}

	t.Log("[PASS] Created logger")

	// destroy logger
	if err := logger.DestroyFeeder(); err != nil {
		t.Log("[FAIL] Failed to destroy logger")
		return
	}

	t.Log("[PASS] Destroyed logger")
}
