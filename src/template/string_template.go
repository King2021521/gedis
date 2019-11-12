package template

import (
	"net"
	"protocol"
	"template/handler"
)

/**
 * 认证权限
 */
func Auth(conn *net.TCPConn, pwd string) (interface{}, error) {
	result := SendCommand(conn, protocol.AUTH, protocol.SafeEncode(pwd))
	return handler.HandleReply(result)
}

func Set(conn *net.TCPConn, key string, value string) (interface{}, error) {
	result := SendCommand(conn, protocol.SET, protocol.SafeEncode(key), protocol.SafeEncode(value))
	return handler.HandleReply(result)
}

func Get(conn *net.TCPConn, key string) (interface{}, error) {
	result := SendCommand(conn, protocol.GET, protocol.SafeEncode(key))
	return handler.HandleReply(result)
}

func Mset(conn *net.TCPConn, keyvalues ...string) (interface{}, error) {
	bytes := handler.HandleBulkRequest(keyvalues)
	result := SendCommand(conn, protocol.MSET, bytes...)
	return handler.HandleReply(result)
}

func Mget(conn *net.TCPConn, keys ...string) (interface{}, error) {
	bytes := handler.HandleBulkRequest(keys)
	result := SendCommand(conn, protocol.MGET, bytes...)
	return handler.HandleReply(result)
}

func Setnx(conn *net.TCPConn, key string, value string) (interface{}, error) {
	result := SendCommand(conn, protocol.SETNX, protocol.SafeEncode(key), protocol.SafeEncode(value))
	return handler.HandleReply(result)
}

func Incr(conn *net.TCPConn, key string) (interface{}, error) {
	result := SendCommand(conn, protocol.INCR, protocol.SafeEncode(key))
	return handler.HandleReply(result)
}

func Decr(conn *net.TCPConn, key string) (interface{}, error) {
	result := SendCommand(conn, protocol.DECR, protocol.SafeEncode(key))
	return handler.HandleReply(result)
}

func Setex(conn *net.TCPConn, key string, time int64, value string) (interface{}, error) {
	result := SendCommand(conn, protocol.SETEX, protocol.SafeEncode(key), protocol.SafeEncodeInt(time), protocol.SafeEncode(value))
	return handler.HandleReply(result)
}
