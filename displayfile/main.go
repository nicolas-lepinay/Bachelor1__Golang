package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	args := os.Args[1:]

	if len(args) > 1 {
		fmt.Print("Too many arguments")
		fmt.Print("\n")
	}

	if len(args) < 1 {
		fmt.Print("File name missing")
		fmt.Print("\n")
	}

	if len(args) == 1 {

		content, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
		}

		fmt.Printf("%s", content)
	}
}
