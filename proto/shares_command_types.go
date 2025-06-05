// Copyright 2025 Raamsri Kumar <raam@tinkershack.in>
// Copyright 2025 The StrataSTOR Authors and Contributors
// SPDX-License-Identifier: Apache-2.0

package proto

// Shares command constants for Rodent gRPC API
const (
	// SMB Share operations
	CmdSharesSMBList       = "shares.smb.list"
	CmdSharesSMBGet        = "shares.smb.get"
	CmdSharesSMBCreate     = "shares.smb.create"
	CmdSharesSMBUpdate     = "shares.smb.update"
	CmdSharesSMBDelete     = "shares.smb.delete"
	CmdSharesSMBStats      = "shares.smb.stats"
	CmdSharesSMBBulkUpdate = "shares.smb.bulk_update"

	// SMB Global config operations
	CmdSharesSMBGlobalGet    = "shares.smb.global.get"
	CmdSharesSMBGlobalUpdate = "shares.smb.global.update"

	CmdSharesSMBRegenerateConfig = "shares.smb.regenerate"

	// SMB Service operations
	CmdSharesSMBServiceStatus  = "shares.smb.service.status"
	CmdSharesSMBServiceStart   = "shares.smb.service.start"
	CmdSharesSMBServiceStop    = "shares.smb.service.stop"
	CmdSharesSMBServiceRestart = "shares.smb.service.restart"
	CmdSharesSMBServiceReload  = "shares.smb.service.reload"
)
