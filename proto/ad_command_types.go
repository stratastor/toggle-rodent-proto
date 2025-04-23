// Copyright 2025 Raamsri Kumar <raam@tinkershack.in>
// Copyright 2025 The StrataSTOR Authors and Contributors
// SPDX-License-Identifier: Apache-2.0

package proto

// AD command constants for Rodent gRPC API
const (
	// User operations
	CmdADUserList   = "ad.user.list"
	CmdADUserGet    = "ad.user.get"
	CmdADUserCreate = "ad.user.create"
	CmdADUserUpdate = "ad.user.update"
	CmdADUserDelete = "ad.user.delete"
	CmdADUserGroups = "ad.user.groups"

	// Group operations
	CmdADGroupList          = "ad.group.list"
	CmdADGroupGet           = "ad.group.get"
	CmdADGroupCreate        = "ad.group.create"
	CmdADGroupUpdate        = "ad.group.update"
	CmdADGroupDelete        = "ad.group.delete"
	CmdADGroupMembers       = "ad.group.members"
	CmdADGroupAddMembers    = "ad.group.members.add"
	CmdADGroupRemoveMembers = "ad.group.members.remove"

	// Computer operations
	CmdADComputerList   = "ad.computer.list"
	CmdADComputerGet    = "ad.computer.get"
	CmdADComputerCreate = "ad.computer.create"
	CmdADComputerUpdate = "ad.computer.update"
	CmdADComputerDelete = "ad.computer.delete"
)