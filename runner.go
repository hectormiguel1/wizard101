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
	if args[ModeIndex] == OfflineCapture {
		if _, err := os.Stat(args[SourceIndex]); err != nil {
			errorString := fmt.Sprintf("Coult not open %s!\n", args[SourceIndex])
			panic(errorString)
		} else {
			offlineCapture(args[SourceIndex])
		}
	} else if args[ModeIndex] == LiveCapture {
		liveCapture(args[SourceIndex])
	}
}

func liveCapture(device string) {
	handle, err := pcap.OpenLive(device, SnapShotLength, false, TimeOut)
	if err != nil {
		panic("Failure to Open Device for live Capture!")
	} else {
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

func offlineCapture(fileToOpen string) {
	if handle, err := pcap.OpenOffline(fileToOpen); err == nil {
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
