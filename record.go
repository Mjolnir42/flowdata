/*-
 * Copyright (c) 2021, Jörg Pernfuß
 *
 * Use of this source code is governed by a 2-clause BSD license
 * that can be found in the LICENSE file.
 */

package flowdata // import "github.com/mjolnir42/flowdata"

import "time"

type Record struct {
	OctetCount     uint64    `json:"OctetCount"`
	PacketCount    uint64    `json:"PacketCount"`
	ProtocolID     uint8     `json:"ProtocolID"`
	Protocol       string    `json:"Protocol,omitempty"`
	IPVersion      uint8     `json:"IPVersion"`
	SrcAddress     string    `json:"SrcAddress"`
	SrcPort        uint16    `json:"SrcPort"`
	DstAddress     string    `json:"DstAddress"`
	DstPort        uint16    `json:"DstPort"`
	TcpControlBits Bitmap    `json:"TcpControlBits,string"`
	TcpFlags       Flags     `json:"TcpFlags"`
	IngressIf      uint32    `json:"IngressIf"`
	EgressIf       uint32    `json:"EgressIf"`
	FlowDirection  uint8     `json:"FlowDirection"`
	StartMilli     time.Time `json:"StartDateTimeMilli"`
	EndMilli       time.Time `json:"EndDateTimeMilli"`
	AgentID        string    `json:"AgentID"`
}

type Flags struct {
	NS  bool `json:"ns,string"`
	CWR bool `json:"cwr,string"`
	ECE bool `json:"ece,string"`
	URG bool `json:"urg,string"`
	ACK bool `json:"ack,string"`
	PSH bool `json:"psh,string"`
	RST bool `json:"rst,string"`
	SYN bool `json:"syn,string"`
	FIN bool `json:"fin,string"`
}

// vim: ts=4 sw=4 sts=4 noet fenc=utf-8 ffs=unix
