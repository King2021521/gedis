package client

import (
	"fmt"
	"gedis/src/client/handler"
	"gedis/src/protocol"
)

func (cluster *Cluster) Auth(pwd string) (interface{}, error) {
	result, err := executeAuth(cluster.RandomSelect(), pwd)
	if err == nil || err.Error() != protocol.MOVED {
		return result, err
	}

	//重定向到新的节点
	return executeAuth(cluster.SelectOne(result.(string)), pwd)
}

func (sharding *Sharding) Auth(pwd string) (interface{}, error) {
	shard := sharding.cHashRing.GetShardInfo(pwd)
	return executeAuth(sharding.shardingPool[shard.Url], pwd)
}

/**
 * 认证权限
 */
func (client *Client) Auth(pwd string) (interface{}, error) {
	return executeAuth(client.getConnectPool(), pwd)
}

func executeAuth(pool *ConnPool, pwd string) (interface{}, error) {
	conn, err := pool.GetConn()
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	result := SendCommand(conn, protocol.AUTH, protocol.SafeEncode(pwd))
	return handler.HandleReply(result)
}

func (cluster *Cluster) Set(key string, value string) (interface{}, error) {
	result, err := executeSet(cluster.RandomSelect(), key, value)
	if err == nil || err.Error() != protocol.MOVED {
		return result, err
	}

	//重定向到新的节点
	return executeSet(cluster.SelectOne(result.(string)), key, value)
}

func (sharding *Sharding) Set(key string, value string) (interface{}, error) {
	return executeSet(sharding.shardingPool[sharding.cHashRing.GetShardInfo(key).Url], key, value)
}

func (client *Client) Set(key string, value string) (interface{}, error) {
	return executeSet(client.getConnectPool(), key, value)
}

func executeSet(pool *ConnPool, key string, value string) (interface{}, error) {
	conn, err := pool.GetConn()
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	result := SendCommand(conn, protocol.SET, protocol.SafeEncode(key), protocol.SafeEncode(value))
	return handler.HandleReply(result)
}

func (cluster *Cluster) Get(key string) (interface{}, error) {
	value, err := executeGet(cluster.RandomSelect(), key)
	if err == nil || err.Error() != protocol.MOVED {
		return value, err
	}

	//重定向到新的节点
	return executeGet(cluster.SelectOne(value.(string)), key)
}

func (sharding *Sharding) Get(key string) (interface{}, error) {
	return executeGet(sharding.shardingPool[sharding.cHashRing.GetShardInfo(key).Url], key)
}

func (client *Client) Get(key string) (interface{}, error) {
	pool := client.getConnectPool()
	return executeGet(pool, key)
}

func executeGet(pool *ConnPool, key string) (interface{}, error) {
	conn, err := pool.GetConn()
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	result := SendCommand(conn, protocol.GET, protocol.SafeEncode(key))
	return handler.HandleReply(result)
}

func (cluster *Cluster) Mset(keyvalues ...string) (interface{}, error) {
	value, err := executeMset(cluster.RandomSelect(), keyvalues)
	if err == nil || err.Error() != protocol.MOVED {
		return value, err
	}

	//重定向到新的节点
	return executeMset(cluster.SelectOne(value.(string)), keyvalues)
}

func (client *Client) Mset(keyvalues ...string) (interface{}, error) {
	return executeMset(client.getConnectPool(), keyvalues)
}

func executeMset(pool *ConnPool, keyvalues []string) (interface{}, error) {
	conn, err := pool.GetConn()
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	bytes := handler.HandleBulkRequest(keyvalues)
	result := SendCommand(conn, protocol.MSET, bytes...)
	return handler.HandleReply(result)
}

func (cluster *Cluster) Mget(keys ...string) (interface{}, error) {
	value, err := executeMget(cluster.RandomSelect(), keys)
	if err == nil || err.Error() != protocol.MOVED {
		return value, err
	}

	//重定向到新的节点
	return executeMget(cluster.SelectOne(value.(string)), keys)
}

func (client *Client) Mget(keys ...string) (interface{}, error) {
	return executeMget(client.getConnectPool(), keys)
}

func executeMget(pool *ConnPool, keys []string) (interface{}, error) {
	conn, err := pool.GetConn()
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	bytes := handler.HandleBulkRequest(keys)
	result := SendCommand(conn, protocol.MGET, bytes...)
	return handler.HandleReply(result)
}

func (cluster *Cluster) Setnx(key string, value string) (interface{}, error) {
	result, err := executeSetnx(cluster.RandomSelect(), key, value)
	if err == nil || err.Error() != protocol.MOVED {
		return result, err
	}

	//重定向到新的节点
	return executeSetnx(cluster.SelectOne(result.(string)), key, value)
}

func (sharding *Sharding) Setnx(key string, value string) (interface{}, error) {
	return executeSetnx(sharding.shardingPool[sharding.cHashRing.GetShardInfo(key).Url], key, value)
}

func (client *Client) Setnx(key string, value string) (interface{}, error) {
	return executeSetnx(client.getConnectPool(), key, value)
}

func executeSetnx(pool *ConnPool, key string, value string) (interface{}, error) {
	conn, err := pool.GetConn()
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	result := SendCommand(conn, protocol.SETNX, protocol.SafeEncode(key), protocol.SafeEncode(value))
	return handler.HandleReply(result)
}

func (sharding *Sharding) Incr(key string) (interface{}, error) {
	return executeIncr(sharding.shardingPool[sharding.cHashRing.GetShardInfo(key).Url], key)
}

func (client *Client) Incr(key string) (interface{}, error) {
	return executeIncr(client.getConnectPool(), key)
}

func (cluster *Cluster) Incr(key string) (interface{}, error) {
	result, err := executeIncr(cluster.RandomSelect(), key)
	if err == nil || err.Error() != protocol.MOVED {
		return result, err
	}

	//重定向到新的节点
	return executeIncr(cluster.SelectOne(result.(string)), key)
}

func executeIncr(pool *ConnPool, key string) (interface{}, error) {
	conn, err := pool.GetConn()
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	result := SendCommand(conn, protocol.INCR, protocol.SafeEncode(key))
	return handler.HandleReply(result)
}

func (cluster *Cluster) Decr(key string) (interface{}, error) {
	result, err := executeDecr(cluster.RandomSelect(), key)
	if err == nil || err.Error() != protocol.MOVED {
		return result, err
	}

	//重定向到新的节点
	return executeDecr(cluster.SelectOne(result.(string)), key)
}

func (sharding *Sharding) Decr(key string) (interface{}, error) {
	return executeDecr(sharding.shardingPool[sharding.cHashRing.GetShardInfo(key).Url], key)
}

func (client *Client) Decr(key string) (interface{}, error) {
	return executeDecr(client.getConnectPool(), key)
}

func executeDecr(pool *ConnPool, key string) (interface{}, error) {
	conn, err := pool.GetConn()
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	result := SendCommand(conn, protocol.DECR, protocol.SafeEncode(key))
	return handler.HandleReply(result)
}

func (cluster *Cluster) Setex(key string, time int64, value string) (interface{}, error) {
	result, err := executeSetex(cluster.RandomSelect(), key, time, value)
	if err == nil || err.Error() != protocol.MOVED {
		return result, err
	}

	//重定向到新的节点
	return executeSetex(cluster.SelectOne(result.(string)), key, time, value)
}

func (sharding *Sharding) Setex(key string, time int64, value string) (interface{}, error) {
	return executeSetex(sharding.shardingPool[sharding.cHashRing.GetShardInfo(key).Url], key, time, value)
}

func (client *Client) Setex(key string, time int64, value string) (interface{}, error) {
	return executeSetex(client.getConnectPool(), key, time, value)
}

func executeSetex(pool *ConnPool, key string, time int64, value string) (interface{}, error) {
	conn, err := pool.GetConn()
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	result := SendCommand(conn, protocol.SETEX, protocol.SafeEncode(key), protocol.SafeEncodeInt(time), protocol.SafeEncode(value))
	return handler.HandleReply(result)
}
