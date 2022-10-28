package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/alexellis/hmac"
)

func main() {
	var mode string
	var inputVar string
	var secret string
	var digest string

	flag.StringVar(&inputVar, "input", "", "Input Message")
	flag.StringVar(&secret, "secret", "", "Secret for Digest")
	flag.StringVar(&digest, "digest", "", "Secret for Digest")
	flag.StringVar(&mode, "mode", "", "GoApp Mode")

	flag.Parse()

	if len(strings.TrimSpace(secret)) == 0 {
		panic("--secret is required")
	}

	if mode == "sign" {
		digest := hmac.Sign([]byte(inputVar), []byte(secret))
		fmt.Printf("DIGEST: %x", digest)
	} else if mode == "validate" {
		if len(strings.TrimSpace(digest)) == 0 {
			panic("--digest is required")
		}
		err := hmac.Validate([]byte(inputVar), fmt.Sprintf("sha1=%s", digest), string(secret))
		if err != nil {
			panic(err)
		}
		fmt.Printf("Validated the hmac.")
	} else {
		panic("Unsupported Mode Selected.")
	}
}
