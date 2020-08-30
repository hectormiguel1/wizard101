package main

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	log "github.com/sirupsen/logrus"
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
	TimeOut            time.Duration = 100 * time.Millisecond
)

var (
	// Log is the package logger
	Log *log.Logger
)

func init() {
	// set and build package logger before start
	Log = log.New()
	configureLogger(Log)
}

func main() {
	args := os.Args
	if len(args) < NumOfArgsExp {
    usage()
    os.Exit(1)
	}

	// timestamp isn't important reading from capture file
	if args[ModeIndex] == OfflineCapture {
		Log.Formatter.(*log.TextFormatter).DisableTimestamp = true
	}
  
	analyze(args[ModeIndex], args[SourceIndex])
}

func usage() {
	fmt.Println("To run the program use the following instructions: \n" +
		"wizard101 [file|live] [source] \n" +
		"file : this is used to read from source as a file in pcapng format, source must be file (this can be run without admin privileges.) \n" +
		"live: this is the option for live packet analysis, requires source to be network device where packages are captured from (requires admin privileges.) \n" +
		"source: this is the file or the network interface where the packets will be read from. (file must be in pcapng format)")
}

func analyze(analysisType string, source string) {
	handle := &pcap.Handle{}
	err := error(nil)
	switch analysisType {
	case LiveCapture:
		Log.Infof("Listening via traffic filter '%v'", Wizard101BPFFilter)
		handle, err = pcap.OpenLive(source, SnapShotLength, true, TimeOut)
	case OfflineCapture:
		Log.Infof("Parsing pcap via traffic filter '%v'", Wizard101BPFFilter)
		handle, err = pcap.OpenOffline(source)
	default:
		panic("unknown capture type")
	}

	if err != nil {
		Log.WithField("err", err).Fatal("Failed to open traffic source")
		panic("failed to open source")
	}
	_ = handle.SetBPFFilter(Wizard101BPFFilter)
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		ipLayer := packet.Layer(layers.LayerTypeIPv4)
		ip, _ := ipLayer.(*layers.IPv4)
		packet, isKINetworkProtocol := parser.TestPacket(packet.Data())
		if !isKINetworkProtocol {
			continue
		}

		message := parser.BuildMessage(packet, ip.SrcIP.String(), ip.DstIP.String())
		if message.ProtocolMessage != nil {
			Log.Infof("%#v", message)
		} else {
			Log.WithField("msg", message).Error("Failed to parse message")
		}
	}
}

// configures the internal logger to its defaults
func configureLogger(l *log.Logger) {
	formatter := new(log.TextFormatter)
	formatter.FullTimestamp = true
	formatter.DisableColors = false
	formatter.TimestampFormat = "06-01-02 15:04:05.00"
	l.Formatter = formatter
}
