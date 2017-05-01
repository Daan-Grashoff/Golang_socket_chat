package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	service := ":1201"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	sendC := make(chan []byte)
	var receiverList [](chan []byte)

	go func(){
		for message := range sendC {
			fmt.Print(string(message))
			for _, receiver := range receiverList{
				receiver <- message
			}
		}
	}()



	for {
		conn, err := listener.Accept()
		receiveC := make(chan []byte)
		receiverList = append(receiverList, receiveC)
		if err != nil {
			continue
		}
		// run as a goroutine
		go handleClient(conn, sendC, receiveC)
	}
}
func handleClient(conn net.Conn, sendC chan []byte, receiver chan []byte) {
	// close connection on exit
	defer conn.Close()
	var buf [512]byte

	go func(){
		for message := range receiver{
			_, err2 := conn.Write(message)
			if err2 != nil {
				return
			}
		}
	}()
	
	for {
		n, err := conn.Read(buf[0:])
		if err != nil{
			return
		}
		// write the n bytes read
		sendC <- buf[0:n]
		
	}
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
