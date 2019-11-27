package client

import (
	"protocol"
	"client/handler"
	"fmt"
)

/**
 * 认证权限
 */
func (client *Client) Auth(pwd string) (interface{}, error) {
	pool := client.getConnectPool()
	conn, err := GetConn(pool)
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	result := SendCommand(conn, protocol.AUTH, protocol.SafeEncode(pwd))
	return handler.HandleReply(result)
}

func (client *Client) Set(key string, value string) (interface{}, error) {
	pool := client.getConnectPool()
	conn, err := GetConn(pool)
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	result := SendCommand(conn, protocol.SET, protocol.SafeEncode(key), protocol.SafeEncode(value))
	return handler.HandleReply(result)
}

func (client *Client) Get(key string) (interface{}, error) {
	pool := client.getConnectPool()
	conn, err := GetConn(pool)
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	result := SendCommand(conn, protocol.GET, protocol.SafeEncode(key))
	return handler.HandleReply(result)
}

func (client *Client) Mset(keyvalues ...string) (interface{}, error) {
	pool := client.getConnectPool()
	conn, err := GetConn(pool)
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	bytes := handler.HandleBulkRequest(keyvalues)
	result := SendCommand(conn, protocol.MSET, bytes...)
	return handler.HandleReply(result)
}

func (client *Client) Mget(keys ...string) (interface{}, error) {
	pool := client.getConnectPool()
	conn, err := GetConn(pool)
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	bytes := handler.HandleBulkRequest(keys)
	result := SendCommand(conn, protocol.MGET, bytes...)
	return handler.HandleReply(result)
}

func (client *Client) Setnx(key string, value string) (interface{}, error) {
	pool := client.getConnectPool()
	conn, err := GetConn(pool)
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	result := SendCommand(conn, protocol.SETNX, protocol.SafeEncode(key), protocol.SafeEncode(value))
	return handler.HandleReply(result)
}

func (client *Client) Incr(key string) (interface{}, error) {
	pool := client.getConnectPool()
	conn, err := GetConn(pool)
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	result := SendCommand(conn, protocol.INCR, protocol.SafeEncode(key))
	return handler.HandleReply(result)
}

func (client *Client) Decr(key string) (interface{}, error) {
	pool := client.getConnectPool()
	conn, err := GetConn(pool)
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	result := SendCommand(conn, protocol.DECR, protocol.SafeEncode(key))
	return handler.HandleReply(result)
}

func (client *Client) Setex(key string, time int64, value string) (interface{}, error) {
	pool := client.getConnectPool()
	conn, err := GetConn(pool)
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	result := SendCommand(conn, protocol.SETEX, protocol.SafeEncode(key), protocol.SafeEncodeInt(time), protocol.SafeEncode(value))
	return handler.HandleReply(result)
}
