// Copyright 2025 Raamsri Kumar <raam@tinkershack.in>
// Copyright 2025 The StrataSTOR Authors and Contributors
// SPDX-License-Identifier: Apache-2.0

package proto

// Network management command constants for Toggle gRPC communication
const (
	// Interface management commands
	CmdNetworkInterfacesList     = "network.interfaces.list"
	CmdNetworkInterfacesGet      = "network.interfaces.get"
	CmdNetworkInterfacesSetState = "network.interfaces.state.set"
	CmdNetworkInterfacesGetStats = "network.interfaces.stats.get"

	// IP address management commands
	CmdNetworkAddressesAdd    = "network.addresses.add"
	CmdNetworkAddressesRemove = "network.addresses.remove"
	CmdNetworkAddressesGet    = "network.addresses.get"

	// Route management commands
	CmdNetworkRoutesList   = "network.routes.list"
	CmdNetworkRoutesAdd    = "network.routes.add"
	CmdNetworkRoutesRemove = "network.routes.remove"

	// Netplan configuration commands
	CmdNetworkNetplanGetConfig = "network.netplan.config.get"
	CmdNetworkNetplanSetConfig = "network.netplan.config.set"
	CmdNetworkNetplanApply     = "network.netplan.apply"
	CmdNetworkNetplanTry       = "network.netplan.try"
	CmdNetworkNetplanSafeApply = "network.netplan.safe_apply"
	CmdNetworkNetplanStatus    = "network.netplan.status"
	CmdNetworkNetplanDiff      = "network.netplan.diff"

	// Backup and restore commands
	CmdNetworkBackupsList      = "network.backups.list"
	CmdNetworkBackupsCreate    = "network.backups.create"
	CmdNetworkBackupsGet       = "network.backups.get"
	CmdNetworkBackupsDelete    = "network.backups.delete"
	CmdNetworkBackupsDeleteAll = "network.backups.purge"
	CmdNetworkBackupsRestore   = "network.backups.restore"

	// System information commands
	CmdNetworkSystemInfo = "network.system.info"

	// Global DNS management commands
	CmdNetworkDNSGetGlobal = "network.dns.global.get"
	CmdNetworkDNSSetGlobal = "network.dns.global.set"

	// Validation commands
	CmdNetworkValidateIP            = "network.validate.ip"
	CmdNetworkValidateInterfaceName = "network.validate.interface_name"
	CmdNetworkValidateNetplanConfig = "network.validate.netplan_config"
)
