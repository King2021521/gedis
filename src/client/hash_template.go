package client

import (
	"fmt"
	"gedis/src/client/handler"
	"gedis/src/protocol"
)

func (cluster *Cluster) Hset(hash string, key string, value string) (interface{}, error) {
	result, err := executeHset(cluster.RandomSelect(), hash, key, value)
	if err == nil || err.Error() != protocol.MOVED {
		return result, err
	}

	//重定向到新的节点
	return executeHset(cluster.SelectOne(result.(string)), hash, key, value)
}

func (sharding *Sharding) Hset(hash string, key string, value string) (interface{}, error) {
	return executeHset(sharding.shardingPool[sharding.cHashRing.GetShardInfo(key).Url], hash, key, value)
}

func (client *Client) Hset(hash string, key string, value string) (interface{}, error) {
	return executeHset(client.getConnectPool(), hash, key, value)
}

func executeHset(pool *ConnPool, hash string, key string, value string) (interface{}, error) {
	conn, err := pool.GetConn()
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	result := SendCommand(conn, protocol.HSET, protocol.SafeEncode(hash), protocol.SafeEncode(key), protocol.SafeEncode(value))
	return handler.HandleReply(result)
}

func (cluster *Cluster) Hget(hash string, key string) (interface{}, error) {
	result, err := executeHget(cluster.RandomSelect(), hash, key)
	if err == nil || err.Error() != protocol.MOVED {
		return result, err
	}

	//重定向到新的节点
	return executeHget(cluster.SelectOne(result.(string)), hash, key)
}

func (sharding *Sharding) Hget(hash string, key string) (interface{}, error) {
	return executeHget(sharding.shardingPool[sharding.cHashRing.GetShardInfo(key).Url], hash, key)
}

func (client *Client) Hget(hash string, key string) (interface{}, error) {
	return executeHget(client.getConnectPool(), hash, key)
}

func executeHget(pool *ConnPool, hash string, key string) (interface{}, error) {
	conn, err := pool.GetConn()
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	result := SendCommand(conn, protocol.HGET, protocol.SafeEncode(hash), protocol.SafeEncode(key))
	return handler.HandleReply(result)
}

func (cluster *Cluster) Hmset(hash string, keyvalues ...string) (interface{}, error) {
	result, err := executeHmset(cluster.RandomSelect(), hash, keyvalues)
	if err == nil || err.Error() != protocol.MOVED {
		return result, err
	}

	//重定向到新的节点
	return executeHmset(cluster.SelectOne(result.(string)), hash, keyvalues)
}

func (client *Client) Hmset(hash string, keyvalues ...string) (interface{}, error) {
	return executeHmset(client.getConnectPool(), hash, keyvalues)
}

func executeHmset(pool *ConnPool, hash string, keyvalues []string) (interface{}, error) {
	conn, err := pool.GetConn()
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	bytes := handler.HandleMultiBulkRequest(hash, keyvalues)
	result := SendCommand(conn, protocol.HMSET, bytes...)
	return handler.HandleReply(result)
}

func (cluster *Cluster) Hmget(hash string, keys ...string) (interface{}, error) {
	result, err := executeHmget(cluster.RandomSelect(), hash, keys)
	if err == nil || err.Error() != protocol.MOVED {
		return result, err
	}

	//重定向到新的节点
	return executeHmget(cluster.SelectOne(result.(string)), hash, keys)
}

func (client *Client) Hmget(hash string, keys ...string) (interface{}, error) {
	return executeHmget(client.getConnectPool(), hash, keys)
}

func executeHmget(pool *ConnPool, hash string, keys []string) (interface{}, error) {
	conn, err := pool.GetConn()
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	bytes := handler.HandleMultiBulkRequest(hash, keys)
	result := SendCommand(conn, protocol.HMGET, bytes...)
	return handler.HandleReply(result)
}

func (cluster *Cluster) Hgetall(hash string) (interface{}, error) {
	result, err := executeHgetall(cluster.RandomSelect(), hash)
	if err == nil || err.Error() != protocol.MOVED {
		return result, err
	}

	//重定向到新的节点
	return executeHgetall(cluster.SelectOne(result.(string)), hash)
}

func (client *Client) Hgetall(hash string) (interface{}, error) {
	return executeHgetall(client.getConnectPool(), hash)
}

func executeHgetall(pool *ConnPool, hash string) (interface{}, error) {
	conn, err := pool.GetConn()
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	result := SendCommand(conn, protocol.HGETALL, protocol.SafeEncode(hash))
	return handler.HandleReply(result)
}

func (cluster *Cluster) Hexists(hash string, key string) (interface{}, error) {
	result, err := executeHexists(cluster.RandomSelect(), hash, key)
	if err == nil || err.Error() != protocol.MOVED {
		return result, err
	}

	//重定向到新的节点
	return executeHexists(cluster.SelectOne(result.(string)), hash, key)
}

func (client *Client) Hexists(hash string, key string) (interface{}, error) {
	return executeHexists(client.getConnectPool(), hash, key)
}

func executeHexists(pool *ConnPool, hash string, key string) (interface{}, error) {
	conn, err := pool.GetConn()
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	result := SendCommand(conn, protocol.HEXISTS, protocol.SafeEncode(hash), protocol.SafeEncode(key))
	return handler.HandleReply(result)
}

func (cluster *Cluster) Hdel(hash string, key string) (interface{}, error) {
	result, err := executeHdel(cluster.RandomSelect(), hash, key)
	if err == nil || err.Error() != protocol.MOVED {
		return result, err
	}

	//重定向到新的节点
	return executeHdel(cluster.SelectOne(result.(string)), hash, key)
}

func (client *Client) Hdel(hash string, key string) (interface{}, error) {
	return executeHdel(client.getConnectPool(), hash, key)
}

func executeHdel(pool *ConnPool, hash string, key string) (interface{}, error) {
	conn, err := pool.GetConn()
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	result := SendCommand(conn, protocol.HDEL, protocol.SafeEncode(hash), protocol.SafeEncode(key))
	return handler.HandleReply(result)
}

func (cluster *Cluster) Hkeys(hash string) (interface{}, error) {
	result, err := executeHkeys(cluster.RandomSelect(), hash)
	if err == nil || err.Error() != protocol.MOVED {
		return result, err
	}

	//重定向到新的节点
	return executeHkeys(cluster.SelectOne(result.(string)), hash)
}

func (client *Client) Hkeys(hash string) (interface{}, error) {
	return executeHkeys(client.getConnectPool(), hash)
}

func executeHkeys(pool *ConnPool, hash string) (interface{}, error) {
	conn, err := pool.GetConn()
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	result := SendCommand(conn, protocol.HKEYS, protocol.SafeEncode(hash))
	return handler.HandleReply(result)
}

func (cluster *Cluster) Hvalues(hash string) (interface{}, error) {
	result, err := executeHvalues(cluster.RandomSelect(), hash)
	if err == nil || err.Error() != protocol.MOVED {
		return result, err
	}

	//重定向到新的节点
	return executeHvalues(cluster.SelectOne(result.(string)), hash)
}

func (client *Client) Hvalues(hash string) (interface{}, error) {
	return executeHvalues(client.getConnectPool(), hash)
}

func executeHvalues(pool *ConnPool, hash string) (interface{}, error) {
	conn, err := pool.GetConn()
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	result := SendCommand(conn, protocol.HVALS, protocol.SafeEncode(hash))
	return handler.HandleReply(result)
}

func (client *Client) Hlen(hash string) (interface{}, error) {
	return executeHlen(client.getConnectPool(), hash)
}

func executeHlen(pool *ConnPool, hash string) (interface{}, error) {
	conn, err := pool.GetConn()
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	result := SendCommand(conn, protocol.HLEN, protocol.SafeEncode(hash))
	return handler.HandleReply(result)
}
