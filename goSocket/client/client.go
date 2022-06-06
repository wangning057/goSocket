package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	ipAndPort := "127.0.0.1:8090"
	tcpAddr, err := net.ResolveTCPAddr("tcp", ipAndPort)
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	fmt.Println("you have connected successfully !!!")

	go receiveMsg(conn)
	for {
		var data string
		fmt.Scan(&data)
		if data == "quit" {
			break
		}
		sendMsg := []byte(data + "\n")
		conn.Write(sendMsg)
	}

}
func receiveMsg(conn *net.TCPConn) {

	result := make([]byte, 256)
	for {
		readLen, err := conn.Read(result)
		checkError(err)
		fmt.Print(string(result[:readLen]))
	}
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
