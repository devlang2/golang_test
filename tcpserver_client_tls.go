package main

import (
	"crypto/tls"
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
	"github.com/icrowley/fake"
)

const (
	DefaultServerName = "tcpsender"

	DefaultEventSize = 2
	DefaultCount     = 1
	DefaultInterval  = 100 // milliseconds
	//	DefaultDuration   = 5   // seconds
	MacChars          = "abcdef0123456789"
	DefaultServerAddr = "192.168.0.3:8808"
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
	Dummy              string
}

type Result struct {
	Code    int
	Message string
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

	config, err := newTLSConfig("server.crt", "server.key")
	//	config := &tls.Config{
	//		InsecureSkipVerify: true,
	//	}

	//	if err != nil {
	//		fmt.Println(err.Error())
	//		return
	//	}

	//	conn, err := net.Dial("tcp", *addr)

	conn, err := tls.Dial("tcp", *addr, config)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//	defer conn.Close()
	defer func() { conn.Close(); fmt.Println("exit") }()

	//	t0 := time.Now()
	//	seq := int64(0)
	//	c := 0
	//	for c < *count {
	//		t0 := time.Now()
	//		events := make([]*Event, 0, *size)
	//		for i := 0; i < *size; i++ {
	//			e := NewEvent(seq)
	//			events = append(events, e)
	//			seq++
	//		}
	//		t1 := time.Now()
	//		fmt.Printf("Generating: %4.1f, ", time.Since(t0).Seconds())

	//		encoder := gob.NewEncoder(conn)
	//		err := encoder.Encode(events)
	//		if err != nil {
	//			fmt.Println(err.Error())
	//			return
	//		}

	//		fmt.Printf("Sending: %4.1f\n", time.Since(t1).Seconds())
	//		time.Sleep(time.Duration(*interval) * time.Millisecond)
	//		c++
	//	}
	//	fmt.Printf("Count: %d, EPS: %5.1f\n", seq, float64(seq)/time.Since(t0).Seconds())

	// Create event
	t0 := time.Now()
	seq := int64(0)
	events := make([]*Event, 0, *size)
	for i := 0; i < *size; i++ {
		events = append(events, NewEvent(seq))
		seq++
	}
	fmt.Printf("Generating: %4.1f\n", time.Since(t0).Seconds())

	// Send events
	c := 0
	for c < *count {
		t1 := time.Now()
		encoder := gob.NewEncoder(conn)
		err := encoder.Encode(events)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		//		decoder := gob.NewDecoder(conn)
		//		err = decoder.Decode(result)
		//		if err != nil {
		//			fmt.Println(err.Error())
		//			return
		//		}

		time.Sleep(time.Duration(*interval) * time.Millisecond)
		fmt.Printf("Sending: %4.1f\n", time.Since(t1).Seconds())
		c++
	}

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
		Dummy:              fake.CharactersN(1800),
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

func newTLSConfig(caPemPath, caKeyPath string) (*tls.Config, error) {
	var config *tls.Config

	//	caPem, err := ioutil.ReadFile(caPemPath)
	//	if err != nil {
	//		return nil, err
	//	}
	//	ca, err := x509.ParseCertificate(caPem)
	//	if err != nil {
	//		return nil, err
	//	}

	//	caKey, err := ioutil.ReadFile(caKeyPath)
	//	if err != nil {
	//		return nil, err
	//	}
	//	key, err := x509.ParsePKCS1PrivateKey(caKey)
	//	if err != nil {
	//		return nil, err
	//	}
	//	pool := x509.NewCertPool()
	//	pool.AddCert(ca)

	//	cert := tls.Certificate{
	//		Certificate: [][]byte{caPem},
	//		PrivateKey:  key,
	//	}

	//	config = &tls.Config{
	//		ClientAuth:   tls.RequireAndVerifyClientCert,
	//		Certificates: []tls.Certificate{cert},
	//		ClientCAs:    pool,
	//	}

	//	config.Rand = rand.Reader

	//	return config, nil

	cer, err := tls.LoadX509KeyPair(caPemPath, caKeyPath)
	if err != nil {
		fmt.Println(err)

		return nil, err
	}

	config = &tls.Config{
		//		Certificates:       []tls.Certificate{cer},
		//		MinVersion:         tls.VersionTLS12,
		//		InsecureSkipVerify: true,
		Certificates: []tls.Certificate{cer},
		MinVersion:   tls.VersionTLS12,
		//		InsecureSkipVerify: true,
		CurvePreferences: []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		},
	}

	return config, nil
}
