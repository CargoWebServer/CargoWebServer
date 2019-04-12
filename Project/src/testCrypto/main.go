package main

import (
	"log"

	"golang.org/x/crypto/chacha20poly1305"
)

func main() {
	c, err := chacha20poly1305.New(make([]byte, chacha20poly1305.KeySize))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("cipher: %+v", c)
}
