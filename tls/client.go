package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
)

func main() {
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
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
		},
		InsecureSkipVerify: true,
		Certificates:       []tls.Certificate{cert},
		RootCAs:            clientCertPool,
	}
	conf.BuildNameToCertificate()

	conn, err := tls.Dial("tcp", "127.0.0.1:8080", conf)
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

//tlsConfig := &tls.Config{
//	Certificates: []tls.Certificate{cert},
//	RootCAs:      clientCertPool,
//}

//tlsConfig.BuildNameToCertificate()
