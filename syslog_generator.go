package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"time"

	"github.com/icrowley/fake"
)

// Flag set
var fs *flag.FlagSet

const (
	DefaultServerName = "Syslog Generator"
	DefaultCount      = 10
	DefaultDst        = "127.0.0.1"
	DefaultDstUDPPort = "514"
	DefaultDstTCPPort = "5514"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	start := time.Now()

	// Set flags
	fs = flag.NewFlagSet("", flag.ExitOnError)
	var (
		count   = fs.Int("count", DefaultCount, "Count")
		dst     = fs.String("destination", DefaultDst, "Destination IP:Port")
		udpPort = fs.String("UDP port", DefaultDstUDPPort, "Destination UDP port")
		//		tcpPort = fs.String("TCP port", DefaultDstTCPPort, "Destination TCP port")
	)
	fs.Usage = printHelp
	fs.Parse(os.Args[1:])

	log.SetPrefix("[SyslogGen] ")

	ServerAddr, err := net.ResolveUDPAddr("udp", *dst+":"+*udpPort)
	CheckError(err)
	LocalAddr, err := net.ResolveUDPAddr("udp", "127.0.0.2:0")
	CheckError(err)
	Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
	CheckError(err)
	defer Conn.Close()

	for i := 0; i < *count; i++ {
		continue

		msg := fmt.Sprintf("<%s>%s %s %s %s %s %s %s %s",
			fake.DigitsN(2), // facility
			fake.DigitsN(1), // version
			time.Now().Add(time.Duration(rand.Int31n(1000))*time.Second).Format(time.RFC3339), // timestamp
			fake.DomainName(),  // hostname
			fake.WordsN(1),     // app-name
			fake.DigitsN(4),    // process id
			"ID"+fake.Digits(), // message id
			fmt.Sprintf(`[%s@%s iut="%s" eventSource="%s" eventID="%s"]`, fake.Word(), fake.DigitsN(5), fake.DigitsN(1), fake.Word(), fake.DigitsN(4)), //  STRUCTURED-DATA
			fake.Words(), // message
		)
		_, err = Conn.Write([]byte(msg))
	}

	elapsed := time.Since(start)
	log.Printf("Data transmitted.(Count: %d, Execution time: %s, Destination: %s)\n", *count, elapsed, *dst+":"+*udpPort)

	conn2, _ := net.Dial("tcp", "127.0.0.1:5514")
	defer conn2.Close()
	for i := 0; i < *count; i++ {

		msg := fmt.Sprintf("<%s>%s %s %s %s %s %s %s %s",
			fake.DigitsN(2), // facility
			fake.DigitsN(1), // version
			time.Now().Add(time.Duration(rand.Int31n(1000))*time.Second).Format(time.RFC3339), // timestamp
			fake.DomainName(),  // hostname
			fake.WordsN(1),     // app-name
			fake.DigitsN(4),    // process id
			"ID"+fake.Digits(), // message id
			fmt.Sprintf(`[%s@%s iut="%s" eventSource="%s" eventID="%s"]`, fake.Word(), fake.DigitsN(5), fake.DigitsN(1), fake.Word(), fake.DigitsN(4)), //  STRUCTURED-DATA
			fake.Words(), // message
		)
		_, err = conn2.Write([]byte(msg))
		fmt.Println(msg)
		time.Sleep(1 * time.Second)
	}

	
}

func printHelp() {
	fmt.Println(DefaultServerName + " [options]")
	fs.PrintDefaults()
}
func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
