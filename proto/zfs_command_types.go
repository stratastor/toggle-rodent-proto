// Copyright 2025 Raamsri Kumar <raam@tinkershack.in>
// Copyright 2025 The StrataSTOR Authors and Contributors
// SPDX-License-Identifier: Apache-2.0

package proto

// ZFS command type constants
// These should be used as the command_type field in CommandRequest
// when interacting with ZFS domain services over gRPC
const (
	// Pool operations
	ZFSCmdPoolList          = "zfs.pool.list"
	ZFSCmdPoolStatus        = "zfs.pool.status"
	ZFSCmdPoolCreate        = "zfs.pool.create"
	ZFSCmdPoolDestroy       = "zfs.pool.destroy"
	ZFSCmdPoolImport        = "zfs.pool.import"
	ZFSCmdPoolExport        = "zfs.pool.export"
	ZFSCmdPoolPropertyList  = "zfs.pool.property.list"
	ZFSCmdPoolPropertyGet   = "zfs.pool.property.get"
	ZFSCmdPoolPropertySet   = "zfs.pool.property.set"
	ZFSCmdPoolScrub         = "zfs.pool.scrub"
	ZFSCmdPoolResilver      = "zfs.pool.resilver"
	ZFSCmdPoolDeviceAttach  = "zfs.pool.device.attach"
	ZFSCmdPoolDeviceDetach  = "zfs.pool.device.detach"
	ZFSCmdPoolDeviceReplace = "zfs.pool.device.replace"

	// Dataset operations
	ZFSCmdDatasetList   = "zfs.dataset.list"
	ZFSCmdDatasetDelete = "zfs.dataset.delete"
	ZFSCmdDatasetRename = "zfs.dataset.rename"
	ZFSCmdDatasetDiff   = "zfs.dataset.diff"

	// Property operations
	ZFSCmdDatasetPropertyList    = "zfs.dataset.property.list"
	ZFSCmdDatasetPropertyGet     = "zfs.dataset.property.get"
	ZFSCmdDatasetPropertySet     = "zfs.dataset.property.set"
	ZFSCmdDatasetPropertyInherit = "zfs.dataset.property.inherit"

	// Filesystem operations
	ZFSCmdFilesystemList    = "zfs.filesystem.list"
	ZFSCmdFilesystemCreate  = "zfs.filesystem.create"
	ZFSCmdFilesystemMount   = "zfs.filesystem.mount"
	ZFSCmdFilesystemUnmount = "zfs.filesystem.unmount"

	// Volume operations
	ZFSCmdVolumeList   = "zfs.volume.list"
	ZFSCmdVolumeCreate = "zfs.volume.create"

	// Snapshot operations
	ZFSCmdSnapshotList     = "zfs.snapshot.list"
	ZFSCmdSnapshotCreate   = "zfs.snapshot.create"
	ZFSCmdSnapshotRollback = "zfs.snapshot.rollback"

	// Clone operations
	ZFSCmdCloneCreate  = "zfs.clone.create"
	ZFSCmdClonePromote = "zfs.clone.promote"

	// Bookmark operations
	ZFSCmdBookmarkList   = "zfs.bookmark.list"
	ZFSCmdBookmarkCreate = "zfs.bookmark.create"

	// Permission operations
	ZFSCmdPermissionList    = "zfs.permission.list"
	ZFSCmdPermissionAllow   = "zfs.permission.allow"
	ZFSCmdPermissionUnallow = "zfs.permission.unallow"

	// Share operations
	ZFSCmdShareDataset   = "zfs.share.dataset"
	ZFSCmdUnshareDataset = "zfs.unshare.dataset"

	// Data transfer operations
	ZFSCmdTransferSend        = "zfs.transfer.send"
	ZFSCmdTransferResumeToken = "zfs.transfer.resume_token"

	// Managed transfer operations
	ZFSCmdTransferStart  = "zfs.transfer.start"
	ZFSCmdTransferList   = "zfs.transfer.list"
	ZFSCmdTransferGet    = "zfs.transfer.get"
	ZFSCmdTransferPause  = "zfs.transfer.pause"
	ZFSCmdTransferResume = "zfs.transfer.resume"
	ZFSCmdTransferStop   = "zfs.transfer.stop"
	ZFSCmdTransferDelete = "zfs.transfer.delete"
)