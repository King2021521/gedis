package client

import (
	"net"
	"protocol"
	"client/handler"
	"fmt"
)

func (client *Client) Hset(hash string, key string, value string) (interface{}, error) {
	pool := client.getConnectPool()
	conn, err := GetConn(pool)
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	result := SendCommand(conn, protocol.HSET, protocol.SafeEncode(hash), protocol.SafeEncode(key), protocol.SafeEncode(value))
	return handler.HandleReply(result)
}

func (client *Client) Hget(hash string, key string) (interface{}, error) {
	pool := client.getConnectPool()
	conn, err := GetConn(pool)
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	result := SendCommand(conn, protocol.HGET, protocol.SafeEncode(hash), protocol.SafeEncode(key))
	return handler.HandleReply(result)
}

func (client *Client) Hmset(hash string, keyvalues ...string) (interface{}, error) {
	pool := client.getConnectPool()
	conn, err := GetConn(pool)
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	bytes := handler.HandleMultiBulkRequest(hash, keyvalues)
	result := SendCommand(conn, protocol.HMSET, bytes...)
	return handler.HandleReply(result)
}

func (client *Client) Hmget(hash string, keys ...string) (interface{}, error) {
	pool := client.getConnectPool()
	conn, err := GetConn(pool)
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	bytes := handler.HandleMultiBulkRequest(hash, keys)
	result := SendCommand(conn, protocol.HMGET, bytes...)
	return handler.HandleReply(result)
}

func (client *Client) Hgetall(hash string) (interface{}, error) {
	pool := client.getConnectPool()
	conn, err := GetConn(pool)
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	result := SendCommand(conn, protocol.HGETALL, protocol.SafeEncode(hash))
	return handler.HandleReply(result)
}

func (client *Client) Hexists(hash string, key string) (interface{}, error) {
	pool := client.getConnectPool()
	conn, err := GetConn(pool)
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	result := SendCommand(conn, protocol.HEXISTS, protocol.SafeEncode(hash), protocol.SafeEncode(key))
	return handler.HandleReply(result)
}

func (client *Client) Hdel(hash string, key string) (interface{}, error) {
	pool := client.getConnectPool()
	conn, err := GetConn(pool)
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	result := SendCommand(conn, protocol.HDEL, protocol.SafeEncode(hash), protocol.SafeEncode(key))
	return handler.HandleReply(result)
}

func (client *Client) Hkeys(hash string) (interface{}, error) {
	pool := client.getConnectPool()
	conn, err := GetConn(pool)
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	result := SendCommand(conn, protocol.HKEYS, protocol.SafeEncode(hash))
	return handler.HandleReply(result)
}

func (client *Client) Hvalues(hash string) (interface{}, error) {
	pool := client.getConnectPool()
	conn, err := GetConn(pool)
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	result := SendCommand(conn, protocol.HVALS, protocol.SafeEncode(hash))
	return handler.HandleReply(result)
}

func (client *Client) Hlen(hash string) (interface{}, error) {
	pool := client.getConnectPool()
	conn, err := GetConn(pool)
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	result := SendCommand(conn, protocol.HLEN, protocol.SafeEncode(hash))
	return handler.HandleReply(result)
}
