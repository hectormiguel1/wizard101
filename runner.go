package main

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"os"
	"syscall"
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
	TimeOut            time.Duration = 100 * time.Millisecond
)

func main() {
	args := os.Args
	if len(args) < NumOfArgsExp {
		usage()
		os.Exit(1)
	} else {
		analyze(args[ModeIndex], args[SourceIndex])
	}
}

func usage() {
	fmt.Println("To run the program use the fallowing instructions: \n" +
		"wizard101 [file|live] [source] \n" +
		"file : this is used to read from source as a file in pcapng format, source must be file (this can be run without admin privileges.) \n" +
		"live: this is the option for live packet analysis, requires source to be network device where packages are captured from (requires admin privileges.) \n" +
		"source: this is the file or the network interface where the packets will be read from. (file must be in pcapng format)")
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
			ipLayer := packet.Layer(layers.LayerTypeIPv4)
			ip, _ := ipLayer.(*layers.IPv4)
			packet, isKINetworkProtocol := parser.TestPacket(packet.Data())
			if isKINetworkProtocol {
				message := parser.BuildMessage(packet, ip.SrcIP.String(), ip.DstIP.String())
				if message.ProtocolMessage != nil {
					_, _ = syscall.Write(syscall.Stdout, []byte(fmt.Sprintf("%#v \n\n", message)))
				}

			}
		}
	}
}
