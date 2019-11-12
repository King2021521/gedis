package template

import (
	"net"
	"protocol"
	"template/handler"
)

func Zadd(conn *net.TCPConn, zset string, scoresvalues ...string) (interface{}, error) {
	bytes := handler.HandleMultiBulkRequest(zset, scoresvalues)
	result := SendCommand(conn, protocol.ZADD, bytes...)
	return handler.HandleReply(result)
}

func Zscore(conn *net.TCPConn, zset string, value string) (interface{}, error) {
	result := SendCommand(conn, protocol.ZSCORE, protocol.SafeEncode(zset), protocol.SafeEncode(value))
	return handler.HandleReply(result)
}

func Zrange(conn *net.TCPConn, zset string, start int64, end int64) (interface{}, error) {
	result := SendCommand(conn, protocol.ZRANGE, protocol.SafeEncode(zset), protocol.SafeEncodeInt(start), protocol.SafeEncodeInt(end))
	return handler.HandleReply(result)
}

func ZrangeWithScores(conn *net.TCPConn, zset string, start int64, end int64) (interface{}, error) {
	result := SendCommand(conn, protocol.ZRANGE, protocol.SafeEncode(zset), protocol.SafeEncodeInt(start), protocol.SafeEncodeInt(end), protocol.SafeEncode("WITHSCORES"))
	return handler.HandleReply(result)
}

func Zrevrange(conn *net.TCPConn, zset string, start int64, end int64) (interface{}, error) {
	result := SendCommand(conn, protocol.ZREVRANGE, protocol.SafeEncode(zset), protocol.SafeEncodeInt(start), protocol.SafeEncodeInt(end))
	return handler.HandleReply(result)
}

func ZrevrangeWithScores(conn *net.TCPConn, zset string, start int64, end int64) (interface{}, error) {
	result := SendCommand(conn, protocol.ZREVRANGE, protocol.SafeEncode(zset), protocol.SafeEncodeInt(start), protocol.SafeEncodeInt(end), protocol.SafeEncode("WITHSCORES"))
	return handler.HandleReply(result)
}

func Zcard(conn *net.TCPConn, zset string) (interface{}, error) {
	result := SendCommand(conn, protocol.ZCARD, protocol.SafeEncode(zset))
	return handler.HandleReply(result)
}

func Zrem(conn *net.TCPConn, zset string, elements ...string) (interface{}, error) {
	bytes := handler.HandleMultiBulkRequest(zset, elements)
	result := SendCommand(conn, protocol.ZREM, bytes...)
	return handler.HandleReply(result)
}

func Zrank(conn *net.TCPConn, zset string, value string) (interface{}, error) {
	result := SendCommand(conn, protocol.ZRANK, protocol.SafeEncode(zset))
	return handler.HandleReply(result)
}

func Zrevrank(conn *net.TCPConn, zset string, value string) (interface{}, error) {
	result := SendCommand(conn, protocol.ZREVRANK, protocol.SafeEncode(zset))
	return handler.HandleReply(result)
}
