package template

import (
	"net"
	"protocol"
	"template/handler"
)

func Sadd(conn *net.TCPConn, set string, elements ...string) (interface{}, error) {
	bytes := handler.HandleMultiBulkRequest(set, elements)
	result := SendCommand(conn, protocol.SADD, bytes...)
	return handler.HandleReply(result)
}

func Smembers(conn *net.TCPConn, set string) (interface{}, error) {
	result := SendCommand(conn, protocol.SMEMBERS, protocol.SafeEncode(set))
	return handler.HandleReply(result)
}

func Srem(conn *net.TCPConn, set string, elements ...string) (interface{}, error) {
	bytes := handler.HandleMultiBulkRequest(set, elements)
	result := SendCommand(conn, protocol.SREM, bytes...)
	return handler.HandleReply(result)
}

func Sismember(conn *net.TCPConn, set string, value string) (interface{}, error) {
	result := SendCommand(conn, protocol.SISMEMBER, protocol.SafeEncode(set), protocol.SafeEncode(value))
	return handler.HandleReply(result)
}

func Scard(conn *net.TCPConn, set string) (interface{}, error) {
	result := SendCommand(conn, protocol.SCARD, protocol.SafeEncode(set))
	return handler.HandleReply(result)
}

func Srandmembers(conn *net.TCPConn, set string, count int64) (interface{}, error) {
	result := SendCommand(conn, protocol.SRANDMEMBER, protocol.SafeEncode(set), protocol.SafeEncodeInt(count))
	return handler.HandleReply(result)
}

func Spop(conn *net.TCPConn, set string) (interface{}, error) {
	result := SendCommand(conn, protocol.SPOP, protocol.SafeEncode(set))
	return handler.HandleReply(result)
}

/**
 * 返回给定所有集合的差集
 */
func Sdiff(conn *net.TCPConn, sets ... string) (interface{}, error) {
	bytes := handler.HandleBulkRequest(sets)
	result := SendCommand(conn, protocol.SDIFF, bytes...)
	return handler.HandleReply(result)
}
