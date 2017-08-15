package main

import (
	//	"bytes"
	"encoding/gob"
	"flag"
	"fmt"
	"math/rand"
	"net"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/davecgh/go-spew/spew"
	"github.com/icrowley/fake"
)

const (
	DefaultServerName = "tcpsender"
	DefaultEventCount = 1
	DefaultEventLoop  = false
	Sleep             = 5000 // ms
	MacChars          = "abcdef0123456789"
)

var (
	fs       *flag.FlagSet
	count    *int
	loop     *bool
	random   *rand.Rand // Rand for this package.
	osBit    = [2]int{32, 64}
	osVer    = [10]float32{10, 10.0, 5.0, 5.1, 5.2, 6, 6.0, 6.1, 6.2, 6.3}
	osServer = [2]int{0, 1}
)

type Event struct {
	Time               time.Time
	Guid               uuid.UUID // AD2BDBE0-BB14-4CBA-A1A4-F9CFD096774F
	IP                 net.IP    // IP
	Mac                string    // MAC
	ComputerName       string    // WSAHN-PC
	OsVersionNumber    float32   // 10.0
	OsIsServer         int       // 0
	OsBit              int       // 64
	FullPolicyVersion  string    // 1026
	TodayPolicyVersion string    // 1028
}

func init() {
	// Set CPU
	runtime.GOMAXPROCS(1)
	rand.Seed(time.Now().Unix())

	// Check flags
	fs = flag.NewFlagSet("", flag.ExitOnError)
	count = fs.Int("count", DefaultEventCount, "Event count")
	loop = fs.Bool("loop", DefaultEventLoop, "Event loop")
	fs.Usage = printHelp
	fs.Parse(os.Args[1:])

	// Get random value
	random = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8808")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer conn.Close()

	//	var network bytes.Buffer        // Stand-in for a network connection
	//	enc := gob.NewEncoder(&network) // Will write to network.
	encoder := gob.NewEncoder(conn)
	for i := 0; i < *count; i++ {
		agent := NewEvent()
		encoder.Encode(*agent)
		spew.Dump(agent)
	}

	//	a := NewEvent()
	//	spew.Dump(a)
	// Initialize the encoder and decoder.  Normally enc and dec would be
	// bound to network connections and the encoder and decoder would
	// run in different processes.
	//var network bytes.Buffer        // Stand-in for a network connection
	//enc := gob.NewEncoder(&network) // Will write to network.
	//dec := gob.NewDecoder(&network) // Will read from network.
	//// Encode (send) the value.
	//err := enc.Encode(P{3, 4, 5, "Pythagoras"})
	//if err != nil {
	//    log.Fatal("encode error:", err)
	//}
	//// Decode (receive) the value.
	//var q Q
	//err = dec.Decode(&q)
	//if err != nil {
	//    log.Fatal("decode error:", err)
	//}
	//fmt.Printf("%q: {%d,%d}\n", q.Name, *q.X, *q.Y)
}

func printHelp() {
	fmt.Println(DefaultServerName + " [options]")
	fs.PrintDefaults()
}

func NewEvent() *Event {
	return &Event{
		Time:               time.Now(),
		Guid:               uuid.New(),
		IP:                 net.ParseIP(fake.IPv4()),
		Mac:                getVirtualMac(),
		ComputerName:       strings.ToUpper(fake.FirstName()) + "-PC",
		OsVersionNumber:    osVer[rand.Intn(len(osVer))],
		OsIsServer:         osServer[rand.Intn(len(osServer))],
		OsBit:              osBit[rand.Intn(len(osBit))],
		FullPolicyVersion:  fake.DigitsN(2),
		TodayPolicyVersion: fake.DigitsN(2),
	}
}

func getVirtualMac() string {
	result := make([]byte, 17)
	for i := range result {
		if i > 0 && i%3 == 2 {
			result[i] = ':'
		} else {
			result[i] = MacChars[random.Intn(len(MacChars))]
		}
	}

	return string(result)
}
