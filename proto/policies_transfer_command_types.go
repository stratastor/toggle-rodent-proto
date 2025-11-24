// Copyright 2025 Raamsri Kumar <raam@tinkershack.in>
// Copyright 2025 The StrataSTOR Authors and Contributors
// SPDX-License-Identifier: Apache-2.0

package proto

// Command type definitions for transfer policy operations
const (
	// Transfer policy operations
	CmdPoliciesTransferList    = "policies.transfer.list"
	CmdPoliciesTransferGet     = "policies.transfer.get"
	CmdPoliciesTransferCreate  = "policies.transfer.create"
	CmdPoliciesTransferUpdate  = "policies.transfer.update"
	CmdPoliciesTransferDelete  = "policies.transfer.delete"
	CmdPoliciesTransferRun     = "policies.transfer.run"
	CmdPoliciesTransferEnable  = "policies.transfer.enable"
	CmdPoliciesTransferDisable = "policies.transfer.disable"
)
