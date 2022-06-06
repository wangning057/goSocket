package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

var connMap = make(map[string]net.Conn)

func main() {
	ipAndPort := "127.0.0.1:8090"
	tcpAddr, err := net.ResolveTCPAddr("tcp", ipAndPort)
	checkError(err)
	fmt.Printf("server start at %v\n", tcpAddr)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		fmt.Println(conn.RemoteAddr().String() + " has connected successfully !!!")
		connMap[conn.RemoteAddr().String()] = conn

		go handleClient(conn)
	}
}
func handleClient(conn net.Conn) {
	conn.SetReadDeadline(time.Now().Add(2 * time.Minute)) //两分钟内没有信息传输，自动关闭conn
	defer conn.Close()

	for {
		request := make([]byte, 128)
		readLen, err := conn.Read(request)
		if err != nil {
			fmt.Println(conn.RemoteAddr().String() + " has quit !!!")
			delete(connMap, conn.RemoteAddr().String()) //将退出的conn移除
			break
		}

		fmt.Print(conn.RemoteAddr().String() + ": " + string(request[:readLen]))
		for _, value := range connMap {
			if value != conn {
				_, err := value.Write([]byte(conn.RemoteAddr().String() + ": " + string(request[:readLen])))
				checkError(err)
			}
		}

	}

}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
