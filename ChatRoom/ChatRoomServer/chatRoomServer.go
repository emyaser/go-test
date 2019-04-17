package main

import (
	"fmt"
	"net"
	"os"
)

var clients []net.Conn

func main() {
	var (
		host   = "localhost"
		port   = "8000"
		remote = host + ":" + port
		data   = make([]byte, 1024)
	)
	fmt.Println("Initiating server...")
	lis, err := net.Listen("tcp", remote)
	defer lis.Close()

	if err != nil {
		fmt.Println("Error when listen:", remote, err)
		os.Exit(-1)
	}

	for {
		var res string
		conn, err := lis.Accept()
		if err != nil {
			fmt.Println("Error accept client:", err.Error())
			os.Exit(0)
		}
		clients = append(clients, conn)
		go func(con net.Conn) {
			fmt.Println("New Connection:", con.RemoteAddr())
			length, err := con.Read(data)
			if err != nil {
				fmt.Println("Client quit:", con.RemoteAddr())
				con.Close()
				disconnect(con, con.RemoteAddr().String())
				return
			}
			name := string(data[:length])
			comeStr := name + "entered the room"
			notify(con, comeStr)

			// Begin receive message from client
			for {
				length, err := con.Read(data)
				if err != nil {
					fmt.Printf("Client %s quit.\n", name)
					con.Close()
					disconnect(con, name)
					return
				}
				res = string(data[:length])
				sprdMsg := name + " said: " + res
				fmt.Println(sprdMsg)
				res = "You said:" + res
				con.Write([]byte(res))
				notify(con, sprdMsg)
			}
		}(conn)
	}
}

func notify(conn net.Conn, msg string) {
	for _, con := range clients {
		if con.RemoteAddr() != conn.RemoteAddr() {
			con.Write([]byte(msg))
		}
	}
}

func disconnect(conn net.Conn, name string) {
	for index, con := range clients {
		if con.RemoteAddr() == conn.RemoteAddr() {
			disMsg := name + " has left the room."
			fmt.Println(disMsg)
			clients = append(clients[:index], clients[index+1:]...)
			notify(conn, disMsg)
		}
	}
}
