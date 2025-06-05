// Copyright 2025 Raamsri Kumar <raam@tinkershack.in>
// Copyright 2025 The StrataSTOR Authors and Contributors
// SPDX-License-Identifier: Apache-2.0

package proto

// Command type constants for service management operations
const (
	// Service listing and status operations
	CmdServicesList    = "services.list"
	CmdServicesStatuses = "services.statuses"
	CmdServiceStatus   = "services.status"

	// Service lifecycle operations
	CmdServiceStart   = "services.start"
	CmdServiceStop    = "services.stop"
	CmdServiceRestart = "services.restart"

	// Service startup management operations
	CmdServiceIsEnabled = "services.is_enabled"
	CmdServiceEnable    = "services.enable"
	CmdServiceDisable   = "services.disable"
)