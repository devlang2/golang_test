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

	//	"github.com/davecgh/go-spew/spew"
	"github.com/icrowley/fake"
)

const (
	DefaultServerName = "tcpsender"

	DefaultEventSize = 1200
	DefaultCount     = 200
	DefaultInterval  = 100 // milliseconds
	//	DefaultDuration   = 5   // seconds
	MacChars          = "abcdef0123456789"
	DefaultServerAddr = "127.0.0.1:8808"
)

var (
	fs       *flag.FlagSet
	size     *int
	interval *int
	count    *int
	addr     *string
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
	Sequence           int64
}

func init() {
	// Set CPU
	runtime.GOMAXPROCS(1)
	rand.Seed(time.Now().Unix())

	// Check flags
	fs = flag.NewFlagSet("", flag.ExitOnError)
	interval = fs.Int("interval", DefaultInterval, "Interval (ms)")
	count = fs.Int("count", DefaultCount, "Count")
	size = fs.Int("size", DefaultEventSize, "Event size")
	addr = fs.String("addr", DefaultServerAddr, "Server address")

	fs.Usage = printHelp
	fs.Parse(os.Args[1:])

	// Get random value
	random = rand.New(rand.NewSource(time.Now().UnixNano()))
}
func main() {
	fmt.Printf("Size: %d, Count: %d, Interval: %d(ms)\n", *size, *count, *interval)
	conn, err := net.Dial("tcp", *addr)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//	defer conn.Close()
	defer func() { conn.Close(); fmt.Println("exit") }()

	t0 := time.Now()
	seq := int64(0)
	c := 0
	for c < *count {
		events := make([]*Event, 0, *size)
		for i := 0; i < *size; i++ {
			events = append(events, NewEvent(seq))
			seq++
		}
		encoder := gob.NewEncoder(conn)
		err := encoder.Encode(events)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		time.Sleep(time.Duration(*interval) * time.Millisecond)
		c++
	}
	fmt.Printf("Count: %d, EPS: %5.1f\n", seq, float64(seq)/time.Since(t0).Seconds())
}

//func main2() {
//	conn, err := net.Dial("tcp", "127.0.0.1:8808")
//	if err != nil {
//		fmt.Println(err.Error())
//		return
//	}
//	defer conn.Close()

//	events := make([]*Event, 0, *count)
//	for i := 0; i < *count; i++ {
//		events = append(events, NewEvent())
//	}
//	encoder := gob.NewEncoder(conn)
//	encoder.Encode(events)
//	spew.Dump(events)
//}

//func main1() {
//	conn, err := net.Dial("tcp", "127.0.0.1:8808")
//	if err != nil {
//		fmt.Println(err.Error())
//		return
//	}
//	defer conn.Close()

//	var events [2]Event
//	events[0] = NewEvent()
//	events[1] = NewEvent()

//	encoder := gob.NewEncoder(conn)
//	encoder.Encode(events)
//	spew.Dump(events)
//}

func printHelp() {
	fmt.Println(DefaultServerName + " [options]")
	fs.PrintDefaults()
}

func NewEvent(seq int64) *Event {
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
		Sequence:           seq,
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
