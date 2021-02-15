/*-
 * Copyright (c) 2021, Jörg Pernfuß
 *
 * Use of this source code is governed by a 2-clause BSD license
 * that can be found in the LICENSE file.
 */

package flowdata // import "github.com/mjolnir42/flowdata"

const (
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

	flagFIN = "0x0001"
	flagSYN = "0x0002"
	flagRST = "0x0004"
	flagPSH = "0x0008"
	flagACK = "0x0010"
	flagURG = "0x0020"
	flagECE = "0x0040"
	flagCWR = "0x0080"
	flagNS  = "0x0100"
)

// vim: ts=4 sw=4 sts=4 noet fenc=utf-8 ffs=unix
