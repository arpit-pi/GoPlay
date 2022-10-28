package main

import (
	"fmt"
	"multiple_packages/cmd"
	"os"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		// Writes to the Provided pipe Stderr/ Stdout etc. returns no. of bytes and error
		fmt.Fprintf(os.Stderr, "Unable to Get Working Directory: %s",
			err.Error())
	}

	res, err := cmd.ExecuteLs(wd)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to Fetch the files in %s, error: %s",
			wd, err.Error())
	}

	fmt.Printf("%s\n", res)
}
