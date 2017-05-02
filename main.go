package main

import (
	"fmt"

	"./src"
)

func main() {
	fmt.Println("choose s|c")
	var result string
	fmt.Scanln(&result)
	switch result {
	case "s":
		src.StartServer()
	case "c":
		src.StartClient()

	}
}
