package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"math/rand"
	"net"
	"os"

	//	"github.com/davecgh/go-spew/spew"
)

var fs *flag.FlagSet

func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func main() {

	// Flag
	fs = flag.NewFlagSet("", flag.ExitOnError)
	var (
		count = fs.Int("count", 1, "Count to send")
	)
	fs.Usage = printHelp
	fs.Parse(os.Args[1:])

	// Set network
	ServerAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:19902")
	CheckError(err)
	LocalAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
	CheckError(err)

	// Connect to server
	Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
	CheckError(err)
	defer Conn.Close()

	rand.Seed(100)

	i := 1
	str := "6ce89c7ea611a74d7b01129585a877927cf4c534cf42ca5b05cc7ac972cd976f1f8cede5628c7e47ae6ee63ca06b4d9fc0fd41910d6ee81341faeb19a2876fc8beb6b548e5e47726e53cb5de15d0147fb878bbfbfa6d26879858a34ff89d0d6db5d5b7d0dbefc646be21a0e3edd15f6b076097257f39d6a779b42fe0776feadd646b4a05eb263abce3e8d087a1daac1beae7c5cd8cf6773db524a00971594d9f2714c671c9b5c32b5c38fe3cf264d3fcb4462e62e394755824d1668d45a1e7a9cb8d036cc82cf647b28f256139100a9176d9797ba7f3d8fbf9cd3d022b9c6c94779bb2c95f1dae4158c5fe144e4dfd6d"
	for i <= *count {

		b, _ := hex.DecodeString(str)
		_, err = Conn.Write(b)
		i++
	}
}

func printHelp() {
	fmt.Println("agent_client [options]")
	fs.PrintDefaults()
}
