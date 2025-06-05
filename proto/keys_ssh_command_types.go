// Copyright 2025 Raamsri Kumar <raam@tinkershack.in>
// Copyright 2025 The StrataSTOR Authors and Contributors
// SPDX-License-Identifier: Apache-2.0

package proto

// Command type constants for SSH keys operations
const (
	// Key pair operations
	CmdSSHKeyPairGenerate = "keys.ssh.generate"
	CmdSSHKeyPairGet      = "keys.ssh.get"
	CmdSSHKeyPairList     = "keys.ssh.list"
	CmdSSHKeyPairRemove   = "keys.ssh.remove"

	// Peer authorization operations (manages both authorized_keys and known_hosts)
	CmdSSHPeerAuthorize   = "keys.ssh.peer.authorize"
	CmdSSHPeerDeauthorize = "keys.ssh.peer.deauthorize"
	CmdSSHPeerList        = "keys.ssh.peer.list"
)
