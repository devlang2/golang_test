package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

// Flag set
var fs *flag.FlagSet

const (
	DefaultServerName = "Syslog Generator"
	DefaultCount      = 10
	DefaultDstIp      = "127.0.0.1"
	DefaultDstPort    = "514"
)

func main() {
	start := time.Now()

	// Set flags
	fs = flag.NewFlagSet("", flag.ExitOnError)
	var (
		count    = fs.Int("count", DefaultCount, "Count")
		udpIface = fs.String("destination", DefaultDstIp+":"+DefaultDstPort, "Destination IP:Port")
	)
	fs.Usage = printHelp
	fs.Parse(os.Args[1:])

	log.SetPrefix("[SyslogGen] ")

	ServerAddr, err := net.ResolveUDPAddr("udp", *udpIface)
	CheckError(err)

	LocalAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
	CheckError(err)

	Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
	CheckError(err)
	defer Conn.Close()

	i := 1
	success := 0
	for i <= *count {
		t := time.Now().Format("2006-01-02T15:04:05Z")
		str := fmt.Sprintf("<17>1 %s mymachine.example1.com evntslog - ID41 #%d BOMA1 application #", t, i)
		_, err = Conn.Write([]byte(str))
		i++
		if err == nil {
			success++
		}
	}
	elapsed := time.Since(start)
	log.Printf("UDP data are transmitted. (Execution time: %s, Destination: %s, Success: %d, Failure: %d\n", elapsed, *udpIface, success, *count-success)

}

func printHelp() {
	fmt.Println(DefaultServerName + " [options]")
	fs.PrintDefaults()
}

//func waitForSignals() {
//	signalCh := make(chan os.Signal, 1)
//	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)

//	// Block until one of the signals above is received
//	select {
//	case <-signalCh:
//		log.Println("signal received, shutting down...")
//	}
//}

func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
