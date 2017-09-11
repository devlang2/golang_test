package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	//	"io/ioutil"
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
	const rootPEM = `
-----BEGIN CERTIFICATE-----
MIIEBDCCAuygAwIBAgIDAjppMA0GCSqGSIb3DQEBBQUAMEIxCzAJBgNVBAYTAlVT
MRYwFAYDVQQKEw1HZW9UcnVzdCBJbmMuMRswGQYDVQQDExJHZW9UcnVzdCBHbG9i
YWwgQ0EwHhcNMTMwNDA1MTUxNTU1WhcNMTUwNDA0MTUxNTU1WjBJMQswCQYDVQQG
EwJVUzETMBEGA1UEChMKR29vZ2xlIEluYzElMCMGA1UEAxMcR29vZ2xlIEludGVy
bmV0IEF1dGhvcml0eSBHMjCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEB
AJwqBHdc2FCROgajguDYUEi8iT/xGXAaiEZ+4I/F8YnOIe5a/mENtzJEiaB0C1NP
VaTOgmKV7utZX8bhBYASxF6UP7xbSDj0U/ck5vuR6RXEz/RTDfRK/J9U3n2+oGtv
h8DQUB8oMANA2ghzUWx//zo8pzcGjr1LEQTrfSTe5vn8MXH7lNVg8y5Kr0LSy+rE
ahqyzFPdFUuLH8gZYR/Nnag+YyuENWllhMgZxUYi+FOVvuOAShDGKuy6lyARxzmZ
EASg8GF6lSWMTlJ14rbtCMoU/M4iarNOz0YDl5cDfsCx3nuvRTPPuj5xt970JSXC
DTWJnZ37DhF5iR43xa+OcmkCAwEAAaOB+zCB+DAfBgNVHSMEGDAWgBTAephojYn7
qwVkDBF9qn1luMrMTjAdBgNVHQ4EFgQUSt0GFhu89mi1dvWBtrtiGrpagS8wEgYD
VR0TAQH/BAgwBgEB/wIBADAOBgNVHQ8BAf8EBAMCAQYwOgYDVR0fBDMwMTAvoC2g
K4YpaHR0cDovL2NybC5nZW90cnVzdC5jb20vY3Jscy9ndGdsb2JhbC5jcmwwPQYI
KwYBBQUHAQEEMTAvMC0GCCsGAQUFBzABhiFodHRwOi8vZ3RnbG9iYWwtb2NzcC5n
ZW90cnVzdC5jb20wFwYDVR0gBBAwDjAMBgorBgEEAdZ5AgUBMA0GCSqGSIb3DQEB
BQUAA4IBAQA21waAESetKhSbOHezI6B1WLuxfoNCunLaHtiONgaX4PCVOzf9G0JY
/iLIa704XtE7JW4S615ndkZAkNoUyHgN7ZVm2o6Gb4ChulYylYbc3GrKBIxbf/a/
zG+FA1jDaFETzf3I93k9mTXwVqO94FntT0QJo544evZG0R0SnU++0ED8Vf4GXjza
HFa9llF7b1cq26KqltyMdMKVvvBulRP/F/A8rLIQjcxz++iPAsbw+zOzlTvjwsto
WHPbqCRiOwY1nQ2pM714A5AuTHhdUDqB1O6gyHA43LL5Z/qHQF1hwFGPa4NrzQU6
yuGnBXj8ytqU0CwIPX4WecigUCAkVDNx
-----END CERTIFICATE-----`

	// First, create the set of root certificates. For this example we only
	// have one. It's also possible to omit this in order to use the
	// default root set of the current operating system.
	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM([]byte(rootPEM))
	if !ok {
		panic("failed to parse root certificate")
	}
	//	spew.Println(roots)
	spew.Println("###")
	//	const rootPEM = `
	//-----BEGIN CERTIFICATE-----
	//MIIEBDCCAuygAwIBAgIDAjppMA0GCSqGSIb3DQEBBQUAMEIxCzAJBgNVBAYTAlVT
	//MRYwFAYDVQQKEw1HZW9UcnVzdCBJbmMuMRswGQYDVQQDExJHZW9UcnVzdCBHbG9i
	//YWwgQ0EwHhcNMTMwNDA1MTUxNTU1WhcNMTUwNDA0MTUxNTU1WjBJMQswCQYDVQQG
	//EwJVUzETMBEGA1UEChMKR29vZ2xlIEluYzElMCMGA1UEAxMcR29vZ2xlIEludGVy
	//bmV0IEF1dGhvcml0eSBHMjCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEB
	//AJwqBHdc2FCROgajguDYUEi8iT/xGXAaiEZ+4I/F8YnOIe5a/mENtzJEiaB0C1NP
	//VaTOgmKV7utZX8bhBYASxF6UP7xbSDj0U/ck5vuR6RXEz/RTDfRK/J9U3n2+oGtv
	//h8DQUB8oMANA2ghzUWx//zo8pzcGjr1LEQTrfSTe5vn8MXH7lNVg8y5Kr0LSy+rE
	//ahqyzFPdFUuLH8gZYR/Nnag+YyuENWllhMgZxUYi+FOVvuOAShDGKuy6lyARxzmZ
	//EASg8GF6lSWMTlJ14rbtCMoU/M4iarNOz0YDl5cDfsCx3nuvRTPPuj5xt970JSXC
	//DTWJnZ37DhF5iR43xa+OcmkCAwEAAaOB+zCB+DAfBgNVHSMEGDAWgBTAephojYn7
	//qwVkDBF9qn1luMrMTjAdBgNVHQ4EFgQUSt0GFhu89mi1dvWBtrtiGrpagS8wEgYD
	//VR0TAQH/BAgwBgEB/wIBADAOBgNVHQ8BAf8EBAMCAQYwOgYDVR0fBDMwMTAvoC2g
	//K4YpaHR0cDovL2NybC5nZW90cnVzdC5jb20vY3Jscy9ndGdsb2JhbC5jcmwwPQYI
	//KwYBBQUHAQEEMTAvMC0GCCsGAQUFBzABhiFodHRwOi8vZ3RnbG9iYWwtb2NzcC5n
	//ZW90cnVzdC5jb20wFwYDVR0gBBAwDjAMBgorBgEEAdZ5AgUBMA0GCSqGSIb3DQEB
	//BQUAA4IBAQA21waAESetKhSbOHezI6B1WLuxfoNCunLaHtiONgaX4PCVOzf9G0JY
	///iLIa704XtE7JW4S615ndkZAkNoUyHgN7ZVm2o6Gb4ChulYylYbc3GrKBIxbf/a/
	//zG+FA1jDaFETzf3I93k9mTXwVqO94FntT0QJo544evZG0R0SnU++0ED8Vf4GXjza
	//HFa9llF7b1cq26KqltyMdMKVvvBulRP/F/A8rLIQjcxz++iPAsbw+zOzlTvjwsto
	//WHPbqCRiOwY1nQ2pM714A5AuTHhdUDqB1O6gyHA43LL5Z/qHQF1hwFGPa4NrzQU6
	//yuGnBXj8ytqU0CwIPX4WecigUCAkVDNx
	//-----END CERTIFICATE-----`

	//	caPem, err := ioutil.ReadFile(caPemPath)
	//	if err != nil {
	//		return nil, err
	//	}
	//	spew.Dump(caPem)
	//	spew.Println("===")
	//	var testRSACertificateIssuer = fromHex("3082021930820182a003020102020900ca5e4e811a965964300d06092a864886f70d01010b0500301f310b3009060355040a1302476f3110300e06035504031307476f20526f6f74301e170d3136303130313030303030305a170d3235303130313030303030305a301f310b3009060355040a1302476f3110300e06035504031307476f20526f6f7430819f300d06092a864886f70d010101050003818d0030818902818100d667b378bb22f34143b6cd2008236abefaf2852adf3ab05e01329e2c14834f5105df3f3073f99dab5442d45ee5f8f57b0111c8cb682fbb719a86944eebfffef3406206d898b8c1b1887797c9c5006547bb8f00e694b7a063f10839f269f2c34fff7a1f4b21fbcd6bfdfb13ac792d1d11f277b5c5b48600992203059f2a8f8cc50203010001a35d305b300e0603551d0f0101ff040403020204301d0603551d250416301406082b0601050507030106082b06010505070302300f0603551d130101ff040530030101ff30190603551d0e041204104813494d137e1631bba301d5acab6e7b300d06092a864886f70d01010b050003818100c1154b4bab5266221f293766ae4138899bd4c5e36b13cee670ceeaa4cbdf4f6679017e2fe649765af545749fe4249418a56bd38a04b81e261f5ce86b8d5c65413156a50d12449554748c59a30c515bc36a59d38bddf51173e899820b282e40aa78c806526fd184fb6b4cf186ec728edffa585440d2b3225325f7ab580e87dd76")
	//	spew.Dump(testRSACertificateIssuer)

	//	_, err := x509.ParseCertificate([]byte(rootPEM))
	//	//	//	spew.Dump(ca)
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

func fromHex(s string) []byte {
	b, _ := hex.DecodeString(s)
	return b
}
