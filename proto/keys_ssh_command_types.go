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

	// Peer authorization operations (manages authorized_keys on the destination/accepter)
	CmdSSHPeerAuthorize   = "keys.ssh.peer.authorize"
	CmdSSHPeerDeauthorize = "keys.ssh.peer.deauthorize"
	CmdSSHPeerList        = "keys.ssh.peer.list"

	// Host key operations (manages known_hosts on the source/requester)
	// GetHostKey retrieves this machine's SSH server host public key
	CmdSSHHostKeyGet = "keys.ssh.hostkey.get"
	// AddKnownHost adds a remote host's key to this machine's known_hosts
	CmdSSHKnownHostAdd = "keys.ssh.knownhost.add"
	// RemoveKnownHost removes a host entry from known_hosts by peering ID
	CmdSSHKnownHostRemove = "keys.ssh.knownhost.remove"
)
