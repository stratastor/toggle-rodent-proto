// Copyright 2025 Raamsri Kumar <raam@tinkershack.in>
// Copyright 2025 The StrataSTOR Authors and Contributors
// SPDX-License-Identifier: Apache-2.0

package proto

// FACL command constants for Rodent gRPC API
const (
	// Path ACL operations
	CmdFACLGet    = "facl.get"
	CmdFACLSet    = "facl.set"
	CmdFACLModify = "facl.modify"
	CmdFACLRemove = "facl.remove"

	// These are the same operations but with recursive=true parameter
	// Adding them as separate commands makes the API easier to use
	CmdFACLGetRecursive    = "facl.recursive.get"
	CmdFACLSetRecursive    = "facl.recursive.set"
	CmdFACLModifyRecursive = "facl.recursive.modify"
	CmdFACLRemoveRecursive = "facl.recursive.remove"
)
