package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net"
	"os"

	"github.com/davecgh/go-spew/spew"
)

const (
	CONN_HOST = ""
	CONN_PORT = "8080"
	CONN_TYPE = "tcp"
)

func main() {
	tlsConfig, _ := getTLSConfig("cert.pem", "key.pem")
	l, err := tls.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT, tlsConfig)

	//	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	conn.Write([]byte("Hello Client."))
	conn.Close()
}

func getTLSConfig(caPemPath, caKeyPath string) (*tls.Config, error) {
	var config *tls.Config

	caPem, err := ioutil.ReadFile(caPemPath)
	if err != nil {
		return nil, err
	}
	spew.Println("===")
	ca, err := x509.ParseCertificate(caPem)
	spew.Dump(ca)
	if err != nil {
		return nil, err
	}

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

	return config, nil
	////////////////////////////////////////////////////////////////////

	//	cer, err := tls.LoadX509KeyPair(caPemPath, caKeyPath)
	//	if err != nil {
	//		log.Println(err)
	//		return nil, err
	//	}

	//	//	config = &tls.Config{
	//	//		Certificates: []tls.Certificate{cer},
	//	//		MinVersion:   tls.VersionTLS12,
	//	//		MaxVersion:   tls.VersionTLS12,
	//	//	}
	//	config := &tls.Config{
	//		Certificates: []tls.Certificate{cer},
	//		MinVersion:   tls.VersionTLS12,
	//		//		InsecureSkipVerify: true,
	//		CurvePreferences: []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
	//		CipherSuites: []uint16{
	//			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
	//			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
	//			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
	//			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
	//		},
	//	}
	//	return config, nil
	///////////////////////////////////////
	//	cert, err := tls.LoadX509KeyPair(caPemPath, caKeyPath)
	//	if err != nil {
	//		fmt.Errorf(err.Error())
	//		return nil, err
	//	}

	//	// Load CA cert
	//	caCert, err := ioutil.ReadFile(*caFile)
	//	if err != nil {
	//		fmt.Errorf(err.Error())
	//	}
	//	caCertPool := x509.NewCertPool()
	//	caCertPool.AppendCertsFromPEM(caCert)

}
