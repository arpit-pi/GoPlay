package cmd

import (
	"fmt"
	"os"
)

//starting a variable, function or struct with a capital letter means that it will be available outside the package to other consumers.
func ExecuteLs(path string) (string, error) {
	enteries, err := os.ReadDir(path)
	if err != nil {
		return "", err
	}

	output := fmt.Sprintf("Files in %s:\n", path)
	output += "Name\t Directory\t\n"

	for _, e := range enteries {
		output += fmt.Sprintf("%s\t%v\n", e.Name(), e.IsDir())
	}

	return output, nil
}
