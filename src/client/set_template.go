package client

import (
	"protocol"
	"client/handler"
	"fmt"
)

func (client *Client) Sadd(set string, elements ...string) (interface{}, error) {
	pool := client.getConnectPool()
	conn, err := GetConn(pool)
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	bytes := handler.HandleMultiBulkRequest(set, elements)
	result := SendCommand(conn, protocol.SADD, bytes...)
	return handler.HandleReply(result)
}

func (client *Client) Smembers(set string) (interface{}, error) {
	pool := client.getConnectPool()
	conn, err := GetConn(pool)
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	result := SendCommand(conn, protocol.SMEMBERS, protocol.SafeEncode(set))
	return handler.HandleReply(result)
}

func (client *Client) Srem(set string, elements ...string) (interface{}, error) {
	pool := client.getConnectPool()
	conn, err := GetConn(pool)
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	bytes := handler.HandleMultiBulkRequest(set, elements)
	result := SendCommand(conn, protocol.SREM, bytes...)
	return handler.HandleReply(result)
}

func (client *Client) Sismember(set string, value string) (interface{}, error) {
	pool := client.getConnectPool()
	conn, err := GetConn(pool)
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	result := SendCommand(conn, protocol.SISMEMBER, protocol.SafeEncode(set), protocol.SafeEncode(value))
	return handler.HandleReply(result)
}

func (client *Client) Scard(set string) (interface{}, error) {
	pool := client.getConnectPool()
	conn, err := GetConn(pool)
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	result := SendCommand(conn, protocol.SCARD, protocol.SafeEncode(set))
	return handler.HandleReply(result)
}

func (client *Client) Srandmembers(set string, count int64) (interface{}, error) {
	pool := client.getConnectPool()
	conn, err := GetConn(pool)
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	result := SendCommand(conn, protocol.SRANDMEMBER, protocol.SafeEncode(set), protocol.SafeEncodeInt(count))
	return handler.HandleReply(result)
}

func (client *Client) Spop(set string) (interface{}, error) {
	pool := client.getConnectPool()
	conn, err := GetConn(pool)
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	result := SendCommand(conn, protocol.SPOP, protocol.SafeEncode(set))
	return handler.HandleReply(result)
}

/**
 * 返回给定所有集合的差集
 */
func (client *Client) Sdiff(sets ... string) (interface{}, error) {
	pool := client.getConnectPool()
	conn, err := GetConn(pool)
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	bytes := handler.HandleBulkRequest(sets)
	result := SendCommand(conn, protocol.SDIFF, bytes...)
	return handler.HandleReply(result)
}
