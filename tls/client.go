package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// Flag set
var fs *flag.FlagSet

func main() {
	// Set flags
	fs = flag.NewFlagSet("", flag.ExitOnError)
	var (
		addr = fs.String("addr", "", "Server address")
	)
	fs.Usage = printHelp
	fs.Parse(os.Args[1:])

	log.SetFlags(log.Lshortfile)

	// Load our TLS key pair to use for authentication
	cert, err := tls.LoadX509KeyPair("cert.pem", "key.pem")
	if err != nil {
		log.Fatalln("Unable to load cert", err)
	}

	// Load our CA certificate
	clientCACert, err := ioutil.ReadFile("cert.pem")
	if err != nil {
		log.Fatal("Unable to open cert", err)
	}

	clientCertPool := x509.NewCertPool()
	clientCertPool.AppendCertsFromPEM(clientCACert)

	conf := &tls.Config{
		CipherSuites: []uint16{
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
		},
		InsecureSkipVerify:       true,
		Certificates:             []tls.Certificate{cert},
		RootCAs:                  clientCertPool,
		MinVersion:               tls.VersionTLS12,
		MaxVersion:               tls.VersionTLS12,
		PreferServerCipherSuites: true,
	}
	conf.BuildNameToCertificate()

	conn, err := tls.Dial("tcp", *addr, conf)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	n, err := conn.Write([]byte("hello\n"))
	if err != nil {
		log.Println(n, err)
		return
	}

	buf := make([]byte, 100)
	n, err = conn.Read(buf)
	if err != nil {
		log.Println(n, err)
		return
	}

	println(string(buf[:n]))
}

func printHelp() {
	fmt.Println("[options]")
	fs.PrintDefaults()
}
