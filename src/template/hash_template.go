package template

import (
	"net"
	"protocol"
	"template/handler"
)

func Hset(conn *net.TCPConn, hash string, key string, value string) (interface{}, error) {
	result := SendCommand(conn, protocol.HSET, protocol.SafeEncode(hash), protocol.SafeEncode(key), protocol.SafeEncode(value))
	return handler.HandleReply(result)
}

func Hget(conn *net.TCPConn, hash string, key string) (interface{}, error) {
	result := SendCommand(conn, protocol.HGET, protocol.SafeEncode(hash), protocol.SafeEncode(key))
	return handler.HandleReply(result)
}

func Hmset(conn *net.TCPConn, hash string, keyvalues ...string) (interface{}, error) {
	bytes := handler.HandleMultiBulkRequest(hash, keyvalues)
	result := SendCommand(conn, protocol.HMSET, bytes...)
	return handler.HandleReply(result)
}

func Hmget(conn *net.TCPConn, hash string, keys ...string) (interface{}, error) {
	bytes := handler.HandleMultiBulkRequest(hash, keys)
	result := SendCommand(conn, protocol.HMGET, bytes...)
	return handler.HandleReply(result)
}

func Hgetall(conn *net.TCPConn, hash string) (interface{}, error) {
	result := SendCommand(conn, protocol.HGETALL, protocol.SafeEncode(hash))
	return handler.HandleReply(result)
}

func Hexists(conn *net.TCPConn, hash string, key string) (interface{}, error) {
	result := SendCommand(conn, protocol.HEXISTS, protocol.SafeEncode(hash), protocol.SafeEncode(key))
	return handler.HandleReply(result)
}

func Hdel(conn *net.TCPConn, hash string, key string) (interface{}, error) {
	result := SendCommand(conn, protocol.HDEL, protocol.SafeEncode(hash), protocol.SafeEncode(key))
	return handler.HandleReply(result)
}

func Hkeys(conn *net.TCPConn, hash string) (interface{}, error) {
	result := SendCommand(conn, protocol.HKEYS, protocol.SafeEncode(hash))
	return handler.HandleReply(result)
}

func Hvalues(conn *net.TCPConn, hash string) (interface{}, error) {
	result := SendCommand(conn, protocol.HVALS, protocol.SafeEncode(hash))
	return handler.HandleReply(result)
}

func Hlen(conn *net.TCPConn, hash string) (interface{}, error) {
	result := SendCommand(conn, protocol.HLEN, protocol.SafeEncode(hash))
	return handler.HandleReply(result)
}
