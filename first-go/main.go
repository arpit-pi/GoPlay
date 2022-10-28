package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/alexellis/hmac"
)

func main() {
	user := os.Getenv("USER")
	fmt.Printf("Hello Mr. %s, Wassup DAWG\n", user)
	input := []byte(`input message from API`)
	secret := []byte(`so secret`)

	digest := hmac.Sign(input, secret)
	fmt.Printf("Digest: %x\n", digest)

	var inputVar string
	var inputSecret string

	flag.StringVar(&inputVar, "message", "", "message to create the digest from")
	flag.StringVar(&inputSecret, "secret", "", "secret for the digest")
	flag.Parse()

	fmt.Printf("Creating Digest For: %q\nSecret: %q", inputVar, inputSecret)

	signedVal := hmac.Sign([]byte(inputVar), []byte(inputSecret))
	fmt.Printf("\nDigest: %x\n", signedVal)
}
