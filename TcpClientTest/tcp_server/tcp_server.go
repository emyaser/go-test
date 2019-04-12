package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

//作为服务器，用来接收客户端发来的信息
func main() {
	fmt.Println("Launching server...")
	// listen on all interfaces
	ln, _ := net.Listen("tcp", "localhost:8080")
	// accept connection on port
	conn, _ := ln.Accept()
	// run loop forever (or until ctrl-c)
	for {
		// will listen for message to process ending in newline (\n)
		message, _ := bufio.NewReader(conn).ReadString('\n')
		// output message received
		fmt.Print("Message Received:", string(message))
		// sample process for string received
		newmessage := strings.ToUpper(message)
		// send new string back to client
		conn.Write([]byte(newmessage + "\n"))
	}
}
