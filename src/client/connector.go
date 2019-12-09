package client

import (
	"net"
	"os"
	"log"
)

/**
 * tcp连接
 */
func Connect(server string) *net.TCPConn{
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)

	if err != nil {
		log.Println(os.Stderr, "Fatal error: ", err)
		return nil
	}

	//建立服务器连接
	conn, err := net.DialTCP("tcp", nil, tcpAddr)

	if err != nil {
		log.Println(os.Stderr, "Fatal error:", err)
		return nil
	}

	log.Printf("server [%s] connect success\n", server)
	return conn
}