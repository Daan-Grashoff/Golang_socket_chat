package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	service := "145.24.240.41:1201"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	defer conn.Close()
	checkError(err)

	reader := bufio.NewReader(os.Stdin)

	go func() {
		for {
			var buf [512]byte
			n, err := conn.Read(buf[0:])
			checkError(err)
			fmt.Print(string(buf[0:n]))
		}
	}()

	for {
		fmt.Print("Enter text: ")
		text, _ := reader.ReadString('\n')
		_, err = conn.Write([]byte(text))
		checkError(err)
	}

	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
