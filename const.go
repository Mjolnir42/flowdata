/*-
 * Copyright (c) 2021, Jörg Pernfuß
 *
 * Use of this source code is governed by a 2-clause BSD license
 * that can be found in the LICENSE file.
 */

package flowdata // import "github.com/mjolnir42/flowdata"

const (
	flagFIN Bitmask = 1 << iota // No more data from sender
	flagSYN                     // Synchronize sequence numbers
	flagRST                     // Reset the connection
	flagPSH                     // Push Function
	flagACK                     // Acknowledgment field significant
	flagURG                     // Urgent Pointer field significant
	flagECE                     // ECN Echo
	flagCWR                     // Congestion Window Reduced
	flagNS                      // ECN Nonce Sum

	octetDeltaCount          = 1
	packetDeltaCount         = 2
	protocolIdentifier       = 4
	tcpControlBits           = 6
	sourceTransportPort      = 7
	sourceIPv4Address        = 8
	destinationTransportPort = 11
	destinationIpv4Address   = 12
	ingressInterface         = 10
	egressInterface          = 14
	sourceIPv6Address        = 27
	destinationIPv6Address   = 28
	ipVersion                = 60
	flowDirection            = 61 // 0x00: ingress, 0x01: egress
	flowStartMilliseconds    = 152
	flowEndMilliseconds      = 153

)

// vim: ts=4 sw=4 sts=4 noet fenc=utf-8 ffs=unix
