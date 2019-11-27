package client

import (
	"net"
	"fmt"
	"os"
)

/**
 * tcp连接
 */
func Connect(server string) *net.TCPConn{
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)

	if err != nil {
		fmt.Println(os.Stderr, "Fatal error: ", err)
		os.Exit(1)
	}

	//建立服务器连接
	conn, err := net.DialTCP("tcp", nil, tcpAddr)

	if err != nil {
		fmt.Println(conn.RemoteAddr().String(), os.Stderr, "Fatal error:", err)
		os.Exit(1)
	}

	fmt.Println("connection success")
	return conn
}