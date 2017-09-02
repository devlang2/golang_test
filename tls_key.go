package main

import (
	"crypto/rand"
	"crypto/rsa"
	"log"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	privKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalf("generating random key: %v", err)
	}

	spew.Dump(privKey)
}
