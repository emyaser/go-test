package main

import "net"
import "fmt"
import "bufio"
import "os"

//作为客户端，用来连接服务器，并发送信息到服务器
func main() {

	// connect to this socket
	conn, _ := net.Dial("tcp", "localhost:8080")
	for {
		// read in input from stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		// send to socket
		fmt.Fprintf(conn, text+"\n")
		// listen for reply
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: " + message)
	}
}
