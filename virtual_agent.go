package main

import (
	//	"crypto/aes"
	//	"crypto/cipher"
	//	"crypto/rand"
	"os/signal"
	"syscall"
	//	"encoding/hex"
	"flag"
	"fmt"
	//	"io"
	"log"
	mathrand "math/rand"
	"net"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/iwondory/encryption"

	//	"github.com/davecgh/go-spew/spew"
	"github.com/icrowley/fake"
	"github.com/satori/go.uuid"
)

const (
	DefaultAgentCount = 1
	Sleep             = 5000
	MacChars          = "abcdef0123456789"
)

var (
	fs    *flag.FlagSet
	count *int

	agents []*Agent
	random *mathrand.Rand // Rand for this package.
	//	iv     = []byte("2981eeca66b5c3cd")                 // internal vector
	key = []byte("c43ac86d84469030f28c0a9656b1c533") // key
)

type Agent struct {
	Code               string  // 00
	Guid               string  // AD2BDBE0-BB14-4CBA-A1A4-F9CFD096774F
	Eth                string  // 10.0.7.194:18-67-b0-47-c0-cc,192.168.1.70:18-67-b0-47-c0-cc,192.168.184.1:00-50-56-c0-00-01,192.168.239.1:00-50-56-c0-00-08,169.254.179.182:00-ff-b7-ad-74-1c
	ComputerName       string  // WSAHN-PC
	OsVersionNumber    float64 // 10.0
	OsIsServer         int64   // 0
	OsBit              int64   // 64
	FullPolicyVersion  string  // 1026
	TodayPolicyVersion string  // 1028
	Data               []byte
}

func init() {
	// Set CPU
	runtime.GOMAXPROCS(1)

	// Check flags
	fs = flag.NewFlagSet("", flag.ExitOnError)
	count = fs.Int("count", DefaultAgentCount, "Agent count")
	fs.Usage = printHelp
	fs.Parse(os.Args[1:])

	// Get random value
	random = mathrand.New(mathrand.NewSource(time.Now().UnixNano()))

	// Create virtual agent
	log.Printf("Creating virtual agent.. %d", *count)
	for i := 0; i < *count; i++ {
		agents = append(agents, NewAgent())
	}

}
func main() {

	// Set network
	log.Printf("Setting network..")
	ServerAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:19902")
	CheckError(err)
	LocalAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
	CheckError(err)

	// Connect to server
	Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
	CheckError(err)
	defer Conn.Close()

	// Send virtual agent data
	go func() {
		i := 0
		for i < 20 {
			t1 := time.Now()
			for _, a := range agents {
				//			spew.Dump(a.Data)
				_, err = Conn.Write(a.Data)
			}

			log.Printf("Complete. Count: %d, Took: %s", len(agents), time.Since(t1))
			time.Sleep(Sleep * time.Millisecond)

		}

	}()
	//	spew.Dump(count)
	waitForSignals()

}

func NewAgent() *Agent {
	guid := uuid.NewV4()
	mac := getIP() + ":" + getVirtualMac()

	agent := &Agent{
		Code:               "00",
		Guid:               strings.ToUpper(guid.String()),
		Eth:                mac,
		ComputerName:       fake.FirstName() + "-PC",
		OsVersionNumber:    10.1,
		OsIsServer:         0,
		OsBit:              64,
		FullPolicyVersion:  "10",
		TodayPolicyVersion: "11",
	}

	text := fmt.Sprintf("%s|%s|%s|%s|%.1f|%d|%d|%s|%s", agent.Code, agent.Guid, agent.Eth, agent.ComputerName, agent.OsVersionNumber, agent.OsIsServer, agent.OsBit, agent.FullPolicyVersion, agent.TodayPolicyVersion)
	b, _ := encryption.Encrypt(key, []byte(text))
	agent.Data = b
	return agent
}

//func Encrypt(str string) []byte {

//	b := []byte(str)
//	//	spew.Dump(b)
//	plaintext := pad(b)
//	spew.Dump(plaintext)

//	if len(plaintext)%aes.BlockSize != 0 {
//		panic("plaintext is not a multiple of the block size")
//	}

//	block, err := aes.NewCipher(key)
//	if err != nil {
//		panic(err)
//	}

//	// The IV needs to be unique, but not secure. Therefore it's common to
//	// include it at the beginning of the ciphertext.
//	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
//	iv := ciphertext[:aes.BlockSize]
//	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
//		panic(err)
//	}

//	mode := cipher.NewCBCEncrypter(block, iv)
//	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

//	// It's important to remember that ciphertexts must be authenticated
//	// (i.e. by using crypto/hmac) as well as being encrypted in order to
//	// be secure.

//	//	fmt.Printf("%x\n", ciphertext)
//	return ciphertext
//}

//func Decrypt(data []byte) ([]byte, error) {
//	block, err := aes.NewCipher(key)
//	if err != nil {
//		return []byte(""), err
//	}

//	if len(data) < aes.BlockSize {
//		return []byte(""), fmt.Errorf("ciphertext too short")
//	}

//	if len(data)%aes.BlockSize != 0 {
//		return []byte(""), fmt.Errorf("ciphertext is not a multiple of the block size")
//	}

//	data_temp := make([]byte, len(data))
//	copy(data_temp, data)

//	mode := cipher.NewCBCDecrypter(block, iv)
//	mode.CryptBlocks(data_temp, data)

//	return data_temp, nil

//}

//func pad(in []byte) []byte {
//	remains := aes.BlockSize - (len(in) % aes.BlockSize)
//	if remains > 0 {
//		for i := 0; i < remains; i++ {
//			in = append(in, byte(0))
//		}
//	}
//	return in
//}

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

func getIP() string {
	size := 4
	ip := make([]byte, size)
	for i := 0; i < size; i++ {
		ip[i] = byte(random.Intn(256))
	}
	return net.IP(ip).To4().String()
}

func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func waitForSignals() {
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)
	select {
	case <-signalCh:
		log.Println("signal received, shutting down...")
	}
}

func printHelp() {
	fmt.Println("virtual_agent [options]")
	fs.PrintDefaults()
}
