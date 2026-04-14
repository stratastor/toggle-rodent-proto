// Copyright 2025 Raamsri Kumar <raam@tinkershack.in>
// Copyright 2025 The StrataSTOR Authors and Contributors
// SPDX-License-Identifier: Apache-2.0

package proto

// Command type constants for TLS certificate management operations.
// Toggle issues certs via ACME (DNS-01/Route53) and delivers them to nodes
// via gRPC. Nodes never touch Route53 or ACME directly.
const (
	// CmdCertDeliver delivers a TLS certificate to the node.
	// Payload: {"cert_pem": "...", "key_pem": "...", "domain": "...", "dest_dir": "/etc/dremio/tls"}
	CmdCertDeliver = "certs.deliver"

	// CmdCertStatus requests the current cert status from the node.
	// Payload: {"dest_dir": "/etc/dremio/tls"}
	// Response: {"domain": "...", "expires_at": "...", "exists": true}
	CmdCertStatus = "certs.status"
)
