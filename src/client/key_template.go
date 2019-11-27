package client

import (
	"protocol"
	"client/handler"
	"fmt"
)

func (client *Client) Keys(regex string) (interface{}, error) {
	pool := client.getConnectPool()
	conn, err := GetConn(pool)
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	result := SendCommand(conn, protocol.KEYS, protocol.SafeEncode(regex))
	return handler.HandleReply(result)
}

func (client *Client) Expire(key string, value int64) (interface{}, error) {
	pool := client.getConnectPool()
	conn, err := GetConn(pool)
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	result := SendCommand(conn, protocol.EXPIRE, protocol.SafeEncode(key), protocol.SafeEncodeInt(value))
	return handler.HandleReply(result)
}

func (client *Client) Del(key string) (interface{}, error) {
	pool := client.getConnectPool()
	conn, err := GetConn(pool)
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	result := SendCommand(conn, protocol.DEL, protocol.SafeEncode(key))
	return handler.HandleReply(result)
}

func (client *Client) Ttl(key string) (interface{}, error) {
	pool := client.getConnectPool()
	conn, err := GetConn(pool)
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	result := SendCommand(conn, protocol.TTL, protocol.SafeEncode(key))
	return handler.HandleReply(result)
}
