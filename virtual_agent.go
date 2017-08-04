package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	//	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"log"
	mathrand "math/rand"
	"net"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/icrowley/fake"
	"github.com/satori/go.uuid"
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

var (
	fs     *flag.FlagSet
	count  *int
	agents []*Agent
	random *mathrand.Rand                               // Rand for this package.
	iv     = []byte("2981eeca66b5c3cd")                 // internal vector
	key    = []byte("c43ac86d84469030f28c0a9656b1c533") // key
)

const (
	chars = "abcdef0123456789"
)

func NewAgent() *Agent {
	guid := uuid.NewV4()
	mac := getIP() + ":" + getMac()

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
	b := encrypt(text)
	agent.Data = b
	return agent
}

func encrypt(str string) []byte {

	plaintext := pad([]byte(str))

	if len(plaintext)%aes.BlockSize != 0 {
		panic("plaintext is not a multiple of the block size")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	// It's important to remember that ciphertexts must be authenticated
	// (i.e. by using crypto/hmac) as well as being encrypted in order to
	// be secure.

	//	fmt.Printf("%x\n", ciphertext)
	return ciphertext
}

func pad(in []byte) []byte {
	padding := aes.BlockSize - (len(in) % aes.BlockSize)
	if padding == 0 {
		padding = aes.BlockSize
	}
	for i := 0; i < padding; i++ {
		in = append(in, byte(padding))
	}
	return in
}

func getMac() string {
	result := make([]byte, 17)
	for i := range result {
		if i > 0 && i%3 == 2 {
			result[i] = ':'
		} else {
			result[i] = chars[random.Intn(len(chars))]
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

func init() {
	// Flag
	fs = flag.NewFlagSet("", flag.ExitOnError)
	count = fs.Int("count", 2, "Agent count")

	// Random
	random = mathrand.New(mathrand.NewSource(time.Now().UnixNano()))

	// Create virtual agent
	for i := 0; i < *count; i++ {
		agents = append(agents, NewAgent())
		//		spew.Dump()
	}

}

func main() {
	log.Println("main")
	//	for _, a := range agents {
	//		//		fmt.Printf("%s|%s|%s|%s|%.1f|%d|%d|%s|%s\n", a.Code, a.Guid, a.Eth, a.ComputerName, a.OsVersionNumber, a.OsIsServer, a.OsBit, a.FullPolicyVersion, a.TodayPolicyVersion)
	//	}

	// Set network
	ServerAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:19902")
	CheckError(err)
	LocalAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
	CheckError(err)

	// Connect to server
	Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
	CheckError(err)
	defer Conn.Close()

	spew.Dump()
	//	rand.Seed(100)

	for _, a := range agents {
		_, err = Conn.Write(a.Data)
	}

	//	i := 1
	//	str := "6ce89c7ea611a74d7b01129585a877927cf4c534cf42ca5b05cc7ac972cd976f1f8cede5628c7e47ae6ee63ca06b4d9fc0fd41910d6ee81341faeb19a2876fc8beb6b548e5e47726e53cb5de15d0147fb878bbfbfa6d26879858a34ff89d0d6db5d5b7d0dbefc646be21a0e3edd15f6b076097257f39d6a779b42fe0776feadd646b4a05eb263abce3e8d087a1daac1beae7c5cd8cf6773db524a00971594d9f2714c671c9b5c32b5c38fe3cf264d3fcb4462e62e394755824d1668d45a1e7a9cb8d036cc82cf647b28f256139100a9176d9797ba7f3d8fbf9cd3d022b9c6c94779bb2c95f1dae4158c5fe144e4dfd6d"
	//	for i <= *count {

	//		b, _ := hex.DecodeString(str)
	//		_, err = Conn.Write(b)
	//		i++
	//	}
}
