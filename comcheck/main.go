package main

import (
	"fmt"
	"os"
)

func main() {

	args := os.Args[1:]

	for _, element := range args {
		if element == "01" || element == "galaxy" || element == "galaxy 01" {
			fmt.Println("Alert!!!")
			break
		}
	}
	return
}
