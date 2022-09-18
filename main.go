package main

import (
	"crypto/rand"
	"io"
	"log"
)

func main() {
	k := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, k); err != nil {
		log.Println(err)
	}

	log.Printf("%x", k)
}
