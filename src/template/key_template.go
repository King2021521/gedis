package template

import (
	"net"
	"protocol"
	"template/handler"
)

func Keys(conn *net.TCPConn, regex string) (interface{}, error) {
	result := SendCommand(conn, protocol.KEYS, protocol.SafeEncode(regex))
	return handler.HandleReply(result)
}

func Expire(conn *net.TCPConn, key string, value int64) (interface{}, error) {
	result := SendCommand(conn, protocol.EXPIRE, protocol.SafeEncode(key), protocol.SafeEncodeInt(value))
	return handler.HandleReply(result)
}

func Del(conn *net.TCPConn, key string) (interface{}, error) {
	result := SendCommand(conn, protocol.DEL, protocol.SafeEncode(key))
	return handler.HandleReply(result)
}

func Ttl(conn *net.TCPConn, key string) (interface{}, error) {
	result := SendCommand(conn, protocol.TTL, protocol.SafeEncode(key))
	return handler.HandleReply(result)
}
