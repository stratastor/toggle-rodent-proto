// Copyright 2025 Raamsri Kumar <raam@tinkershack.in>
// Copyright 2025 The StrataSTOR Authors and Contributors
// SPDX-License-Identifier: Apache-2.0

package proto

// System management command constants for Toggle gRPC communication
const (
	// System information commands
	CmdSystemInfoGet         = "system.info.get"
	CmdSystemInfoCPUGet      = "system.info.cpu.get"
	CmdSystemInfoMemoryGet   = "system.info.memory.get"
	CmdSystemInfoOSGet       = "system.info.os.get"
	CmdSystemInfoPerformanceGet = "system.info.performance.get"
	CmdSystemHealthGet       = "system.health.get"

	// Hostname management commands
	CmdSystemHostnameGet = "system.hostname.get"
	CmdSystemHostnameSet = "system.hostname.set"

	// User management commands
	CmdSystemUsersList           = "system.users.list"
	CmdSystemUsersCreate         = "system.users.create"
	CmdSystemUsersDelete         = "system.users.delete"
	CmdSystemUsersGet            = "system.users.get"
	CmdSystemUsersUpdate         = "system.users.update"
	CmdSystemUsersPasswordSet    = "system.users.password.set"
	CmdSystemUsersLock           = "system.users.lock"
	CmdSystemUsersUnlock         = "system.users.unlock"
	CmdSystemUsersGroupsList     = "system.users.groups.list"
	CmdSystemUsersGroupsAdd      = "system.users.groups.add"
	CmdSystemUsersGroupsRemove   = "system.users.groups.remove"
	CmdSystemUsersGroupsPrimary  = "system.users.groups.primary.set"

	// Group management commands
	CmdSystemGroupsList   = "system.groups.list"
	CmdSystemGroupsGet    = "system.groups.get"
	CmdSystemGroupsCreate = "system.groups.create"
	CmdSystemGroupsDelete = "system.groups.delete"

	// Power management commands
	CmdSystemPowerStatusGet        = "system.power.status.get"
	CmdSystemPowerScheduledGet     = "system.power.scheduled.get"
	CmdSystemPowerShutdown         = "system.power.shutdown"
	CmdSystemPowerReboot           = "system.power.reboot"
	CmdSystemPowerScheduleShutdown = "system.power.schedule-shutdown.create"
	CmdSystemPowerScheduledDelete  = "system.power.scheduled.delete"

	// System configuration commands
	CmdSystemTimezoneGet = "system.config.timezone.get"
	CmdSystemTimezoneSet = "system.config.timezone.set"
	CmdSystemLocaleGet   = "system.config.locale.get"
	CmdSystemLocaleSet   = "system.config.locale.set"
)