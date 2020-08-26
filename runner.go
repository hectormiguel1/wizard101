package main

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"os"
	"time"
	"wizard101/parser"
)

const (
	NumOfArgsExp                     = 3
	SourceIndex                      = 2
	ModeIndex                        = 1
	LiveCapture                      = "live"
	OfflineCapture                   = "file"
	SnapShotLength                   = 1024
	Wizard101BPFFilter               = "src net 165.193.0.0/16 or dst net 165.193.0.0/16"
	TimeOut            time.Duration = 30 * time.Second
)

func main() {
	args := os.Args
	if len(args) < NumOfArgsExp {
		panic("Failed to pass file to parse packets from! (File must be in pcap format)")
	}
	analyze(args[ModeIndex], args[SourceIndex])
}

func analyze(analysisType string, source string) {
	handle := &pcap.Handle{}
	err := error(nil)
	if analysisType == LiveCapture {
		handle, err = pcap.OpenLive(source, SnapShotLength, true, TimeOut)
	} else if analysisType == OfflineCapture {
		handle, err = pcap.OpenOffline(source)
	}
	if err == nil {
		_ = handle.SetBPFFilter(Wizard101BPFFilter)
		packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
		for packet := range packetSource.Packets() {
			packet, isKINetworkProtocol := parser.TestPacket(packet.Data())
			if isKINetworkProtocol {
				message := parser.BuildMessage(packet)
				if message.ProtocolMessage != nil {
					fmt.Printf("%#v \n\n", message.ProtocolMessage)
				}

			}
		}
	}
}
