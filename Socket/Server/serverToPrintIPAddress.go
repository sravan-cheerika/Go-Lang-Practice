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
        fmt.Println("IP address : ", string(message), conn.RemoteAddr())
		
		fmt.Println("time stamp : ", string(message), time.Now().Format("2006-01-02 15:04:05"))
		
    }
}

var totalNumOfConnections int = 0
func main() {
    fmt.Println("Start the server ")
    ln, err := net.Listen("tcp", port)
    for {
        conn, _ := ln.Accept()

        fmt.Println(" Connected Client : ", conn.RemoteAddr())
		// count 
		totalNumOfConnections++
        fmt.Println("Total Connected Clients:", totalNumOfConnections)
		// count 
        if err != nil {
            fmt.Println("Error ", err)
            continue
        }
        go echoMessage(conn)
    }
}