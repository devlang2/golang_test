package main

import (
	"bufio"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net"
)

func main() {
	log.SetFlags(log.Lshortfile)

	cer, err := tls.LoadX509KeyPair("cert.pem", "key.pem")
	if err != nil {
		log.Println(err)
		return
	}

	// Load our CA certificate
	clientCACert, err := ioutil.ReadFile("cert.pem")
	if err != nil {
		log.Fatal("Unable to open cert", err)
	}
	clientCertPool := x509.NewCertPool()
	clientCertPool.AppendCertsFromPEM(clientCACert)

	config := &tls.Config{
		ClientAuth:   tls.RequireAndVerifyClientCert,
		Certificates: []tls.Certificate{cer},
		RootCAs:      clientCertPool,
		ClientCAs:    clientCertPool,
		CipherSuites: []uint16{
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
		},
		PreferServerCipherSuites: true,
		MinVersion:               tls.VersionTLS12,
	}
	ln, err := tls.Listen("tcp", ":8080", config)
	if err != nil {
		log.Println(err)
		return
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(conn)
	for {
		msg, err := r.ReadString('\n')
		if err != nil {
			log.Println(err)
			return
		}

		println(msg)

		n, err := conn.Write([]byte("world\n"))
		if err != nil {
			log.Println(n, err)
			return
		}
	}
}
