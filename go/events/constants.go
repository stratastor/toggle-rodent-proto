// Copyright 2025 The StrataSTOR Authors and Contributors
// SPDX-License-Identifier: Apache-2.0

package events

// Event type string constants organized by category
// These correspond to the enums defined in events.proto

// System Events
const (
	SystemStartup                = "system.startup"
	SystemShutdown               = "system.shutdown"
	SystemRestart                = "system.restart"
	SystemConfigChanged          = "system.config.changed"
	SystemConfigReloaded         = "system.config.reloaded"
	SystemHealthCheckPassed      = "system.health_check.passed"
	SystemHealthCheckFailed      = "system.health_check.failed"
	SystemPerformanceThreshold   = "system.performance.threshold"
	SystemLocalUserCreated       = "system.local_user.created"
	SystemLocalUserDeleted       = "system.local_user.deleted"
	SystemLocalUserModified      = "system.local_user.modified"
	SystemLocalGroupCreated      = "system.local_group.created"
	SystemLocalGroupDeleted      = "system.local_group.deleted"
	SystemPackageInstalled       = "system.package.installed"
	SystemPackageUpdated         = "system.package.updated"
	SystemPackageRemoved         = "system.package.removed"
)

// Storage Events
const (
	StoragePoolCreated           = "storage.pool.created"
	StoragePoolDestroyed         = "storage.pool.destroyed"
	StoragePoolImported          = "storage.pool.imported"
	StoragePoolExported          = "storage.pool.exported"
	StoragePoolScrubStarted      = "storage.pool.scrub.started"
	StoragePoolScrubCompleted    = "storage.pool.scrub.completed"
	StoragePoolScrubFailed       = "storage.pool.scrub.failed"
	StoragePoolStatusChanged     = "storage.pool.status.changed"
	StorageDatasetCreated        = "storage.dataset.created"
	StorageDatasetDestroyed      = "storage.dataset.destroyed"
	StorageDatasetRenamed        = "storage.dataset.renamed"
	StorageDatasetPropertyChanged = "storage.dataset.property.changed"
	StorageDatasetMounted        = "storage.dataset.mounted"
	StorageDatasetUnmounted      = "storage.dataset.unmounted"
	StorageSnapshotCreated       = "storage.snapshot.created"
	StorageSnapshotDestroyed     = "storage.snapshot.destroyed"
	StorageSnapshotRenamed       = "storage.snapshot.renamed"
	StorageSnapshotRolledBack    = "storage.snapshot.rolled_back"
	StorageSnapshotCloned        = "storage.snapshot.cloned"
	StorageTransferStarted       = "storage.transfer.started"
	StorageTransferCompleted     = "storage.transfer.completed"
	StorageTransferFailed        = "storage.transfer.failed"
	StorageTransferPaused        = "storage.transfer.paused"
	StorageTransferResumed       = "storage.transfer.resumed"
	StorageTransferCancelled     = "storage.transfer.cancelled"
	StorageVolumeFull            = "storage.volume.full"
	StorageVolumeThreshold       = "storage.volume.threshold"
	StorageQuotaExceeded         = "storage.quota.exceeded"
)

// Network Events
const (
	NetworkInterfaceUp           = "network.interface.up"
	NetworkInterfaceDown         = "network.interface.down"
	NetworkInterfaceAdded        = "network.interface.added"
	NetworkInterfaceRemoved      = "network.interface.removed"
	NetworkInterfaceConfigChanged = "network.interface.config.changed"
	NetworkConnectionEstablished = "network.connection.established"
	NetworkConnectionFailed      = "network.connection.failed"
	NetworkConnectionLost        = "network.connection.lost"
	NetworkConnectionTimeout     = "network.connection.timeout"
	NetworkDnsConfigChanged      = "network.dns.config.changed"
	NetworkDnsResolutionFailed   = "network.dns.resolution.failed"
	NetworkRouteAdded            = "network.route.added"
	NetworkRouteDeleted          = "network.route.deleted"
	NetworkRouteChanged          = "network.route.changed"
	NetworkBandwidthThreshold    = "network.bandwidth.threshold"
	NetworkPacketLossThreshold   = "network.packet_loss.threshold"
	NetworkLatencyThreshold      = "network.latency.threshold"
)

// Security Events
const (
	SecurityAuthSuccess          = "security.auth.success"
	SecurityAuthFailed           = "security.auth.failed"
	SecurityAuthTimeout          = "security.auth.timeout"
	SecurityLoginSuccess         = "security.login.success"
	SecurityLoginFailed          = "security.login.failed"
	SecurityLogout               = "security.logout"
	SecurityPermissionDenied     = "security.permission.denied"
	SecurityPrivilegeEscalation  = "security.privilege.escalation"
	SecurityUnauthorizedAccess   = "security.unauthorized_access"
	SecurityKeyGenerated         = "security.key.generated"
	SecurityKeyRotated           = "security.key.rotated"
	SecurityKeyExpired           = "security.key.expired"
	SecurityKeyCompromised       = "security.key.compromised"
	SecurityCertificateIssued    = "security.certificate.issued"
	SecurityCertificateExpired   = "security.certificate.expired"
	SecurityCertificateRevoked   = "security.certificate.revoked"
	SecuritySshKeyAdded          = "security.ssh.key.added"
	SecuritySshKeyRemoved        = "security.ssh.key.removed"
	SecuritySshConnectionEstablished = "security.ssh.connection.established"
	SecuritySshConnectionFailed  = "security.ssh.connection.failed"
	SecurityIntrusionDetected    = "security.intrusion.detected"
	SecurityMalwareDetected      = "security.malware.detected"
	SecurityPolicyViolation      = "security.policy.violation"
)

// Service Events
const (
	ServiceStarted               = "service.started"
	ServiceStopped               = "service.stopped"
	ServiceRestarted             = "service.restarted"
	ServiceFailed                = "service.failed"
	ServiceCrashed               = "service.crashed"
	ServiceKilled                = "service.killed"
	ServiceConfigChanged         = "service.config.changed"
	ServiceConfigReloaded        = "service.config.reloaded"
	ServiceConfigInvalid         = "service.config.invalid"
	ServiceEnabledAtStartup      = "service.enabled.at_startup"
	ServiceDisabledAtStartup     = "service.disabled.at_startup"
	ServiceHealthCheckPassed     = "service.health_check.passed"
	ServiceHealthCheckFailed     = "service.health_check.failed"
	ServiceDependencyFailed      = "service.dependency.failed"
)

// Identity Events (AD/LDAP)
const (
	IdentityUserCreated          = "identity.user.created"
	IdentityUserUpdated          = "identity.user.updated"
	IdentityUserDeleted          = "identity.user.deleted"
	IdentityUserEnabled          = "identity.user.enabled"
	IdentityUserDisabled         = "identity.user.disabled"
	IdentityUserPasswordChanged  = "identity.user.password.changed"
	IdentityUserPasswordExpired  = "identity.user.password.expired"
	IdentityUserLocked           = "identity.user.locked"
	IdentityUserUnlocked         = "identity.user.unlocked"
	IdentityGroupCreated         = "identity.group.created"
	IdentityGroupUpdated         = "identity.group.updated"
	IdentityGroupDeleted         = "identity.group.deleted"
	IdentityGroupMemberAdded     = "identity.group.member.added"
	IdentityGroupMemberRemoved   = "identity.group.member.removed"
	IdentityComputerJoined       = "identity.computer.joined"
	IdentityComputerLeft         = "identity.computer.left"
	IdentityComputerUpdated      = "identity.computer.updated"
	IdentityComputerDisabled     = "identity.computer.disabled"
	IdentityDomainSyncStarted    = "identity.domain.sync.started"
	IdentityDomainSyncCompleted  = "identity.domain.sync.completed"
	IdentityDomainSyncFailed     = "identity.domain.sync.failed"
	IdentityDomainControllerConnected = "identity.domain_controller.connected"
	IdentityDomainControllerDisconnected = "identity.domain_controller.disconnected"
)

// Access Control Events
const (
	AccessAclCreated             = "access.acl.created"
	AccessAclUpdated             = "access.acl.updated"
	AccessAclDeleted             = "access.acl.deleted"
	AccessAclPermissionGranted   = "access.acl.permission.granted"
	AccessAclPermissionRevoked   = "access.acl.permission.revoked"
	AccessFileAccessGranted      = "access.file.access.granted"
	AccessFileAccessDenied       = "access.file.access.denied"
	AccessDirectoryAccessGranted = "access.directory.access.granted"
	AccessDirectoryAccessDenied  = "access.directory.access.denied"
	AccessOwnershipChanged       = "access.ownership.changed"
	AccessPermissionsChanged     = "access.permissions.changed"
	AccessInheritanceChanged     = "access.inheritance.changed"
	AccessSecurityDescriptorChanged = "access.security_descriptor.changed"
	AccessDefaultAclChanged      = "access.default_acl.changed"
)

// File Sharing Events
const (
	SharingShareCreated          = "sharing.share.created"
	SharingShareUpdated          = "sharing.share.updated"
	SharingShareDeleted          = "sharing.share.deleted"
	SharingShareEnabled          = "sharing.share.enabled"
	SharingShareDisabled         = "sharing.share.disabled"
	SharingSmbShareConnected     = "sharing.smb.share.connected"
	SharingSmbShareDisconnected  = "sharing.smb.share.disconnected"
	SharingSmbSessionEstablished = "sharing.smb.session.established"
	SharingSmbSessionTerminated  = "sharing.smb.session.terminated"
	SharingSmbFileOpened         = "sharing.smb.file.opened"
	SharingSmbFileClosed         = "sharing.smb.file.closed"
	SharingNfsExportCreated      = "sharing.nfs.export.created"
	SharingNfsExportDeleted      = "sharing.nfs.export.deleted"
	SharingNfsMount              = "sharing.nfs.mount"
	SharingNfsUnmount            = "sharing.nfs.unmount"
	SharingShareAccessGranted    = "sharing.share.access.granted"
	SharingShareAccessDenied     = "sharing.share.access.denied"
	SharingQuotaExceeded         = "sharing.quota.exceeded"
	SharingConnectionLimitReached = "sharing.connection.limit_reached"
	SharingBandwidthThreshold    = "sharing.bandwidth.threshold"
)

// Standard metadata keys
const (
	MetaComponent      = "component"
	MetaAction         = "action"
	MetaModule         = "module"
	MetaResourceType   = "resource_type"
	MetaResourceID     = "resource_id"
	MetaResourceName   = "resource_name"
	MetaOperation      = "operation"
	MetaStatus         = "status"
	MetaError          = "error"
	MetaUser           = "user"
	MetaHost           = "host"
	MetaService        = "service"
	MetaInterface      = "interface"
	MetaPool           = "pool"
	MetaDataset        = "dataset"
	MetaSnapshot       = "snapshot"
	MetaTransferID     = "transfer_id"
	MetaShareName      = "share_name"
	MetaDomain         = "domain"
	MetaGroup          = "group"
)