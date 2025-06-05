// Copyright 2025 Raamsri Kumar <raam@tinkershack.in>
// Copyright 2025 The StrataSTOR Authors and Contributors
// SPDX-License-Identifier: Apache-2.0

package proto

// Command type definitions for autosnap policy operations
const (
	// Snapshot policy operations
	CmdPoliciesAutosnapList   = "policies.autosnap.list"
	CmdPoliciesAutosnapGet    = "policies.autosnap.get"
	CmdPoliciesAutosnapCreate = "policies.autosnap.create"
	CmdPoliciesAutosnapUpdate = "policies.autosnap.update"
	CmdPoliciesAutosnapDelete = "policies.autosnap.delete"
	CmdPoliciesAutosnapRun    = "policies.autosnap.run"
)
