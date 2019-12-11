package client

import (
	"protocol"
	"client/handler"
	"fmt"
)

func (cluster *Cluster) Keys(regex string) (interface{}, error) {
	result, err := executeKeys(cluster.RandomSelect(), regex)
	if err==nil || err.Error() != protocol.MOVED {
		return result, err
	}

	//重定向到新的节点
	return executeKeys(cluster.SelectOne(result.(string)), regex)
}

func (client *Client) Keys(regex string) (interface{}, error) {
	return executeKeys(client.getConnectPool(), regex)
}

func executeKeys(pool *ConnPool, regex string) (interface{}, error) {
	conn, err := pool.GetConn()
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	result := SendCommand(conn, protocol.KEYS, protocol.SafeEncode(regex))
	return handler.HandleReply(result)
}

func (cluster *Cluster) Expire(key string, value int64) (interface{}, error) {
	result, err := executeExpire(cluster.RandomSelect(), key, value)
	if err==nil || err.Error() != protocol.MOVED {
		return result, err
	}

	//重定向到新的节点
	return executeExpire(cluster.SelectOne(result.(string)), key, value)
}

func (client *Client) Expire(key string, value int64) (interface{}, error) {
	return executeExpire(client.getConnectPool(), key, value)
}

func executeExpire(pool *ConnPool, key string, value int64) (interface{}, error) {
	conn, err := pool.GetConn()
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	result := SendCommand(conn, protocol.EXPIRE, protocol.SafeEncode(key), protocol.SafeEncodeInt(value))
	return handler.HandleReply(result)
}

func (cluster *Cluster) Del(key string) (interface{}, error) {
	result, err := executeDel(cluster.RandomSelect(), key)
	if err==nil || err.Error() != protocol.MOVED {
		return result, err
	}

	//重定向到新的节点
	return executeDel(cluster.SelectOne(result.(string)), key)
}

func (client *Client) Del(key string) (interface{}, error) {
	return executeDel(client.getConnectPool(), key)
}

func executeDel(pool *ConnPool, key string) (interface{}, error) {
	conn, err := pool.GetConn()
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	result := SendCommand(conn, protocol.DEL, protocol.SafeEncode(key))
	return handler.HandleReply(result)
}

func (cluster *Cluster) Ttl(key string) (interface{}, error) {
	result, err := executeTtl(cluster.RandomSelect(), key)
	if err==nil || err.Error() != protocol.MOVED {
		return result, err
	}

	//重定向到新的节点
	return executeTtl(cluster.SelectOne(result.(string)), key)
}

func (client *Client) Ttl(key string) (interface{}, error) {
	return executeTtl(client.getConnectPool(), key)
}

func executeTtl(pool *ConnPool, key string) (interface{}, error) {
	conn, err := pool.GetConn()
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	result := SendCommand(conn, protocol.TTL, protocol.SafeEncode(key))
	return handler.HandleReply(result)
}

func (client *Client) Ping() (interface{}, error) {
	return executePing(client.getConnectPool())
}

func executePing(pool *ConnPool)(interface{}, error){
	conn, err := pool.GetConn()
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	result := SendCommand(conn, protocol.PING)
	return handler.HandleReply(result)
}
