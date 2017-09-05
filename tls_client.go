package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/levigross/grequests"
)

func main() {
	cert, err := tls.LoadX509KeyPair("cert.pem", "key.pem")
	if err != nil {
		log.Fatalln("Unable to load cert", err)
	}

	clientCACert, err := ioutil.ReadFile("cert.pem")
	if err != nil {
		log.Fatal("Unable to open cert", err)
	}

	clientCertPool := x509.NewCertPool()
	clientCertPool.AppendCertsFromPEM(clientCACert)

	tlsConfig := &tls.Config{
		Certificates:             []tls.Certificate{cert},
		RootCAs:                  clientCertPool,
		InsecureSkipVerify:       true,
		MinVersion:               tls.VersionTLS12,
		MaxVersion:               tls.VersionTLS12,
		PreferServerCipherSuites: true,
	}

	tlsConfig.BuildNameToCertificate()
	ro := &grequests.RequestOptions{
		HTTPClient: &http.Client{
			Transport: &http.Transport{TLSClientConfig: tlsConfig},
		},
	}
	resp, err := grequests.Get("https://192.168.72.128:8080", ro)
	if err != nil {
		log.Println("Unable to speak to our server", err)
	}

	log.Println(resp.String())

	//	conn, err := tls.Dial("tcp", "192.168.72.128:8080", tlsConfig)
	//	if err != nil {
	//		log.Println(err)
	//		return
	//	}
	//	defer conn.Close()

	//	n, err := conn.Write([]byte("hello\n"))
	//	if err != nil {
	//		log.Println(n, err)
	//		return
	//	}

	//	buf := make([]byte, 100)
	//	n, err = conn.Read(buf)
	//	if err != nil {
	//		log.Println(n, err)
	//		return
	//	}

	//	println(string(buf[:n]))
}

//state := conn.ConnectionState()
//for _, v := range state.PeerCertificates {
//	fmt.Println(x509.MarshalPKIXPublicKey(v.PublicKey))
//	fmt.Println(v.Subject)
//}
//log.Println("client: handshake: ", state.HandshakeComplete)
//log.Println("client: mutual: ", state.NegotiatedProtocolIsMutual)

//message := "Hello\n"
//n, err := io.WriteString(conn, message)
//if err != nil {
//	log.Fatalf("client: write: %s", err)
//}
//log.Printf("client: wrote %q (%d bytes)", message, n)

//reply := make([]byte, 256)
//n, err = conn.Read(reply)
//log.Printf("client: read %q (%d bytes)", string(reply[:n]), n)
//log.Print("client: exiting")
