/*-
 * Copyright (c) 2021, Jörg Pernfuß
 *
 * Use of this source code is governed by a 2-clause BSD license
 * that can be found in the LICENSE file.
 */

package flowdata // import "github.com/mjolnir42/flowdata"

import "github.com/davecgh/go-spew/spew"

type Message struct {
	AgentID  string `json:"AgentID"`
	Header   Header `json:"Header"`
	DataSets []Data `json:"DataSets"`
}

type Header struct {
	Version    int `json:"Version"`
	Length     int `json:"Length"`
	ExportTime int `json:"ExportTime"`
	SequenceNo int `json:"SequenceNo"`
	DomainID   int `json:"DomainID"`
}

type Data []kvpair

type kvpair map[string]interface{}

func (m *Message) Convert() <-chan Record {
	spew.Dump(`flowdata.message.M:`, m)
	ret := make(chan Record)
	go func() {
		for i := range m.DataSets {
			res := Record{AgentID: m.AgentID}
			spew.Dump(`flowdata.message.res/before`, res)
			for _, pair := range m.DataSets[i] {
				if key, ok := pair[`I`].(uint64); ok {
					switch key {
					case octetDeltaCount:
						if val, ok := pair[`V`].(uint64); ok {
							res.OctetCount = val
						}
					case packetDeltaCount:
						if val, ok := pair[`V`].(uint64); ok {
							res.PacketCount = val
						}
					case protocolIdentifier:
						if val, ok := pair[`V`].(uint8); ok {
							res.ProtocolID = val
							switch res.ProtocolID {
							case 1:
								res.Protocol = `ICMP`
							case 2:
								res.Protocol = `IGMP`
							case 3:
								res.Protocol = `IPv4`
							case 6:
								res.Protocol = `TCP`
							case 17:
								res.Protocol = `UDP`
							case 41:
								res.Protocol = `IPv6`
							case 47:
								res.Protocol = `GRE`
							case 50:
								res.Protocol = `ESP`
							case 51:
								res.Protocol = `AH`
							case 58:
								res.Protocol = `IPv6-ICMP`
							case 115:
								res.Protocol = `L2TP`
							case 132:
								res.Protocol = `SCTP`
							case 136:
								res.Protocol = `UDPLite`
							case 137:
								res.Protocol = `MPLS-in-IP`
							}
						}
					case tcpControlBits:
						if val, ok := pair[`V`].(string); ok {
							res.TcpControlBits = ParseBitmap(val)
							res.TcpFlags.FIN = res.TcpControlBits.Check(flagFIN)
							res.TcpFlags.SYN = res.TcpControlBits.Check(flagSYN)
							res.TcpFlags.RST = res.TcpControlBits.Check(flagRST)
							res.TcpFlags.PSH = res.TcpControlBits.Check(flagPSH)
							res.TcpFlags.ACK = res.TcpControlBits.Check(flagACK)
							res.TcpFlags.URG = res.TcpControlBits.Check(flagURG)
							res.TcpFlags.ECE = res.TcpControlBits.Check(flagECE)
							res.TcpFlags.CWR = res.TcpControlBits.Check(flagCWR)
							res.TcpFlags.NS = res.TcpControlBits.Check(flagNS)
						}
					case sourceTransportPort:
						if val, ok := pair[`V`].(uint16); ok {
							res.SrcPort = val
						}
					case sourceIPv4Address:
						if val, ok := pair[`V`].(string); ok {
							res.SrcAddress = FormatIP(val)
						}
					case destinationTransportPort:
						if val, ok := pair[`V`].(uint16); ok {
							res.DstPort = val
						}
					case destinationIpv4Address:
						if val, ok := pair[`V`].(string); ok {
							res.SrcAddress = FormatIP(val)
						}
					case ingressInterface:
						if val, ok := pair[`V`].(uint32); ok {
							res.IngressIf = val
						}
					case egressInterface:
						if val, ok := pair[`V`].(uint32); ok {
							res.EgressIf = val
						}
					case sourceIPv6Address:
						if val, ok := pair[`V`].(string); ok {
							res.SrcAddress = FormatIP(val)
						}
					case destinationIPv6Address:
						if val, ok := pair[`V`].(string); ok {
							res.SrcAddress = FormatIP(val)
						}
					case ipVersion:
						if val, ok := pair[`V`].(uint8); ok {
							res.IPVersion = val
						}
					case flowDirection:
						if val, ok := pair[`V`].(uint8); ok {
							res.FlowDirection = val
						}
					case flowStartMilliseconds:
						if val, ok := pair[`V`].(int64); ok {
							res.StartMilli = unix2time(val)
						}
					case flowEndMilliseconds:
						if val, ok := pair[`V`].(int64); ok {
							res.EndMilli = unix2time(val)
						}
					}
				}
			}
			spew.Dump(`flowdata.message.res/after`, res)
			ret <- res
		}
	}()
	return ret
}

// vim: ts=4 sw=4 sts=4 noet fenc=utf-8 ffs=unix
