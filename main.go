package main

import (
	"fmt"

	"github.com/Daan-Grashoff/Golang_socket_chat/src"
)

func main() {
	fmt.Println("choose")
	var result string
	fmt.Scanln(&result)
	switch {
	case result == "s":
		src.StartServer()
	case result == "c":
		src.StartClient()

	}
}
