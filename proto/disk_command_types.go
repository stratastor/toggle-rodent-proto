// Copyright 2025 Raamsri Kumar <raam@tinkershack.in>
// Copyright 2025 The StrataSTOR Authors and Contributors
// SPDX-License-Identifier: Apache-2.0

package proto

// Disk management command constants for Toggle gRPC communication
const (
	// Disk inventory and discovery commands
	CmdDiskList          = "disk.list"           // List all disks
	CmdDiskListAvailable = "disk.list.available" // List available disks
	CmdDiskGet           = "disk.get"            // Get disk details
	CmdDiskDiscover      = "disk.discover"       // Trigger disk discovery

	// Disk health and SMART commands
	CmdDiskHealthCheck   = "disk.health.check"   // Trigger health check for all disks
	CmdDiskHealthGet     = "disk.health.get"     // Get disk health status
	CmdDiskSMARTGet      = "disk.smart.get"      // Get SMART data
	CmdDiskSMARTRefresh  = "disk.smart.refresh"  // Refresh SMART data

	// SMART probe (self-test) commands
	CmdDiskProbeStart    = "disk.probe.start"    // Start a SMART probe (manual)
	CmdDiskProbeCancel   = "disk.probe.cancel"   // Cancel a running probe
	CmdDiskProbeGet      = "disk.probe.get"      // Get probe details
	CmdDiskProbeList     = "disk.probe.list"     // List probe executions
	CmdDiskProbeHistory  = "disk.probe.history"  // Get probe history for a disk

	// Probe schedule management commands
	CmdDiskProbeScheduleList   = "disk.probe.schedule.list"   // List probe schedules
	CmdDiskProbeScheduleGet    = "disk.probe.schedule.get"    // Get schedule details
	CmdDiskProbeScheduleCreate = "disk.probe.schedule.create" // Create probe schedule
	CmdDiskProbeScheduleUpdate = "disk.probe.schedule.update" // Update probe schedule
	CmdDiskProbeScheduleDelete = "disk.probe.schedule.delete" // Delete probe schedule
	CmdDiskProbeScheduleEnable = "disk.probe.schedule.enable" // Enable schedule
	CmdDiskProbeScheduleDisable = "disk.probe.schedule.disable" // Disable schedule

	// Physical topology commands
	CmdDiskTopologyGet       = "disk.topology.get"        // Get complete topology
	CmdDiskTopologyRefresh   = "disk.topology.refresh"    // Refresh topology
	CmdDiskTopologyControllers = "disk.topology.controllers" // List controllers
	CmdDiskTopologyEnclosures  = "disk.topology.enclosures"  // List enclosures

	// Fault domain commands
	CmdDiskFaultDomainsAnalyze = "disk.fault-domains.analyze" // Analyze fault domains
	CmdDiskFaultDomainsGet     = "disk.fault-domains.get"     // Get fault domain info

	// Disk state management commands
	CmdDiskStateGet     = "disk.state.get"      // Get disk state
	CmdDiskStateSet     = "disk.state.set"      // Set disk state (e.g., quarantine)
	CmdDiskValidate     = "disk.validate"       // Validate a disk
	CmdDiskQuarantine   = "disk.quarantine"     // Quarantine a disk

	// Disk naming and configuration commands
	CmdDiskNamingStrategyGet = "disk.naming-strategy.get" // Get naming strategy
	CmdDiskNamingStrategySet = "disk.naming-strategy.set" // Set naming strategy
	CmdDiskVdevConfGenerate  = "disk.vdev-conf.generate"  // Generate vdev_id.conf

	// Disk metadata commands
	CmdDiskTagsSet    = "disk.tags.set"     // Set disk tags
	CmdDiskTagsDelete = "disk.tags.delete"  // Delete disk tags
	CmdDiskNotesSet   = "disk.notes.set"    // Set disk notes

	// Statistics and monitoring commands
	CmdDiskStatsGet       = "disk.stats.get"        // Get disk statistics
	CmdDiskStatsGlobal    = "disk.stats.global"     // Get global statistics
	CmdDiskMonitoringGet  = "disk.monitoring.get"   // Get monitoring config
	CmdDiskMonitoringSet  = "disk.monitoring.set"   // Set monitoring config

	// Configuration management commands
	CmdDiskConfigGet     = "disk.config.get"      // Get disk manager config
	CmdDiskConfigUpdate  = "disk.config.update"   // Update disk manager config
	CmdDiskConfigReload  = "disk.config.reload"   // Reload configuration
)
