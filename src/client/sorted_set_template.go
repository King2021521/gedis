package client

import (
	"protocol"
	"client/handler"
	"fmt"
)

func (client *Client) Zadd(zset string, scoresvalues ...string) (interface{}, error) {
	pool := client.getConnectPool()
	conn, err := GetConn(pool)
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	bytes := handler.HandleMultiBulkRequest(zset, scoresvalues)
	result := SendCommand(conn, protocol.ZADD, bytes...)
	return handler.HandleReply(result)
}

func (client *Client) Zscore(zset string, value string) (interface{}, error) {
	pool := client.getConnectPool()
	conn, err := GetConn(pool)
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	result := SendCommand(conn, protocol.ZSCORE, protocol.SafeEncode(zset), protocol.SafeEncode(value))
	return handler.HandleReply(result)
}

func (client *Client) Zrange(zset string, start int64, end int64) (interface{}, error) {
	pool := client.getConnectPool()
	conn, err := GetConn(pool)
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	result := SendCommand(conn, protocol.ZRANGE, protocol.SafeEncode(zset), protocol.SafeEncodeInt(start), protocol.SafeEncodeInt(end))
	return handler.HandleReply(result)
}

func (client *Client) ZrangeWithScores(zset string, start int64, end int64) (interface{}, error) {
	pool := client.getConnectPool()
	conn, err := GetConn(pool)
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	result := SendCommand(conn, protocol.ZRANGE, protocol.SafeEncode(zset), protocol.SafeEncodeInt(start), protocol.SafeEncodeInt(end), protocol.SafeEncode("WITHSCORES"))
	return handler.HandleReply(result)
}

func (client *Client) Zrevrange(zset string, start int64, end int64) (interface{}, error) {
	pool := client.getConnectPool()
	conn, err := GetConn(pool)
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	result := SendCommand(conn, protocol.ZREVRANGE, protocol.SafeEncode(zset), protocol.SafeEncodeInt(start), protocol.SafeEncodeInt(end))
	return handler.HandleReply(result)
}

func (client *Client) ZrevrangeWithScores(zset string, start int64, end int64) (interface{}, error) {
	pool := client.getConnectPool()
	conn, err := GetConn(pool)
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	result := SendCommand(conn, protocol.ZREVRANGE, protocol.SafeEncode(zset), protocol.SafeEncodeInt(start), protocol.SafeEncodeInt(end), protocol.SafeEncode("WITHSCORES"))
	return handler.HandleReply(result)
}

func (client *Client) Zcard(zset string) (interface{}, error) {
	pool := client.getConnectPool()
	conn, err := GetConn(pool)
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	result := SendCommand(conn, protocol.ZCARD, protocol.SafeEncode(zset))
	return handler.HandleReply(result)
}

func (client *Client) Zrem(zset string, elements ...string) (interface{}, error) {
	pool := client.getConnectPool()
	conn, err := GetConn(pool)
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	bytes := handler.HandleMultiBulkRequest(zset, elements)
	result := SendCommand(conn, protocol.ZREM, bytes...)
	return handler.HandleReply(result)
}

func (client *Client) Zrank(zset string, value string) (interface{}, error) {
	pool := client.getConnectPool()
	conn, err := GetConn(pool)
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	result := SendCommand(conn, protocol.ZRANK, protocol.SafeEncode(zset))
	return handler.HandleReply(result)
}

func (client *Client) Zrevrank(zset string, value string) (interface{}, error) {
	pool := client.getConnectPool()
	conn, err := GetConn(pool)
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	result := SendCommand(conn, protocol.ZREVRANK, protocol.SafeEncode(zset))
	return handler.HandleReply(result)
}
