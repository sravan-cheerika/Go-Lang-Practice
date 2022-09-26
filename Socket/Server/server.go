package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"time"
)

var port = "0.0.0.0:9001"

func echoMessage(conn net.Conn) {
	r := bufio.NewReader(conn)
	for {
		message, err := r.ReadBytes(byte('\n'))
		switch err {
		case nil:
			break
		case io.EOF:
		default:
			fmt.Println("Error ", err)
		}
		conn.Write([]byte("From Server :"))
		conn.Write(message)
		fmt.Println("Message Received From Client : ", string(message)) 
		fmt.Println("IP address : ", string(message), conn.RemoteAddr())
		
		fmt.Println("time stamp : ", string(message), time.Now())

	}
}
func main() {
	fmt.Println("Start the server ")
	ln, err := net.Listen("tcp", port)
	for {
		fmt.Println("1st msg :")
		conn, _ := ln.Accept() // pbm check for firewall connection 
		fmt.Println("2nd msg :")
		if err != nil {
			fmt.Println("Error ", err)
			continue
		}
		go echoMessage(conn)
	}
}
