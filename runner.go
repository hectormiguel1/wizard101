package main

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"os"
	"wizard101/parser"
)

const NumOfArgsExp = 2
const FilePathLocation = 1

func main() {
	parser.InitializeLoginMap()
	args := os.Args
	if len(args) < NumOfArgsExp {
		panic("Failed to pass file to parse packets from! (File must be in pcap format)")
	}
	if _, err := os.Stat(args[FilePathLocation]); err != nil {
		errorString := fmt.Sprintf("Coult not open %s!\n", args[FilePathLocation])
		panic(errorString)
	} else {
		if handle, err := pcap.OpenOffline(args[FilePathLocation]); err == nil {
			_ = handle.SetBPFFilter("src net 165.193.0.0/16 or dst net 165.193.0.0/16")
			packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
			for packet := range packetSource.Packets() {
				packet, isKINetworkProtocol := parser.TestPacket(packet.Data())
				if isKINetworkProtocol {
					message := parser.BuildMessage(packet)
					fmt.Printf("%#v \n", message)
				}
			}
		}
	}
}
