package client

import (
	"protocol"
	"client/handler"
	"fmt"
)

func (cluster *Cluster) Zadd(zset string, scoresvalues ...string) (interface{}, error) {
	result, err := executeZadd(cluster.RandomSelect(), zset, scoresvalues)
	if err.Error() != protocol.MOVED {
		return result, err
	}

	//重定向到新的节点
	return executeZadd(cluster.SelectOne(result.(string)), zset, scoresvalues)
}

func (cluster *Cluster) Zscore(zset string, value string) (interface{}, error) {
	result, err := executeZscore(cluster.RandomSelect(), zset, value)
	if err.Error() != protocol.MOVED {
		return result, err
	}

	//重定向到新的节点
	return executeZscore(cluster.SelectOne(result.(string)), zset, value)
}

func (cluster *Cluster) Zrange(zset string, start int64, end int64) (interface{}, error) {
	result, err := executeZrange(cluster.RandomSelect(), zset, start, end)
	if err.Error() != protocol.MOVED {
		return result, err
	}

	//重定向到新的节点
	return executeZrange(cluster.SelectOne(result.(string)), zset, start, end)
}

func (cluster *Cluster) ZrangeWithScores(zset string, start int64, end int64) (interface{}, error) {
	result, err := executeZrangeWithScores(cluster.RandomSelect(), zset, start, end)
	if err.Error() != protocol.MOVED {
		return result, err
	}

	//重定向到新的节点
	return executeZrangeWithScores(cluster.SelectOne(result.(string)), zset, start, end)
}

func (cluster *Cluster) Zrevrange(zset string, start int64, end int64) (interface{}, error) {
	result, err := executeZrevrange(cluster.RandomSelect(), zset, start, end)
	if err.Error() != protocol.MOVED {
		return result, err
	}

	//重定向到新的节点
	return executeZrevrange(cluster.SelectOne(result.(string)), zset, start, end)
}

func (cluster *Cluster) ZrevrangeWithScores(zset string, start int64, end int64) (interface{}, error) {
	result, err := executeZrevrangeWithScores(cluster.RandomSelect(), zset, start, end)
	if err.Error() != protocol.MOVED {
		return result, err
	}

	//重定向到新的节点
	return executeZrevrangeWithScores(cluster.SelectOne(result.(string)), zset, start, end)
}

func (cluster *Cluster) Zcard(zset string) (interface{}, error) {
	result, err := executeZcard(cluster.RandomSelect(), zset)
	if err.Error() != protocol.MOVED {
		return result, err
	}

	//重定向到新的节点
	return executeZcard(cluster.SelectOne(result.(string)), zset)
}

func (cluster *Cluster) Zrem(zset string, elements ...string) (interface{}, error) {
	result, err := executeZrem(cluster.RandomSelect(), zset, elements)
	if err.Error() != protocol.MOVED {
		return result, err
	}

	//重定向到新的节点
	return executeZrem(cluster.SelectOne(result.(string)), zset, elements)
}

func (cluster *Cluster) Zrank(zset string, value string) (interface{}, error) {
	result, err := executeZrank(cluster.RandomSelect(), zset, value)
	if err.Error() != protocol.MOVED {
		return result, err
	}

	//重定向到新的节点
	return executeZrank(cluster.SelectOne(result.(string)), zset, value)
}

func (cluster *Cluster) Zrevrank(zset string, value string) (interface{}, error) {
	result, err := executeZrevrank(cluster.RandomSelect(), zset, value)
	if err.Error() != protocol.MOVED {
		return result, err
	}

	//重定向到新的节点
	return executeZrevrank(cluster.SelectOne(result.(string)), zset, value)
}

func (client *Client) Zadd(zset string, scoresvalues ...string) (interface{}, error) {
	return executeZadd(client.getConnectPool(), zset, scoresvalues)
}

func (client *Client) Zscore(zset string, value string) (interface{}, error) {
	return executeZscore(client.getConnectPool(), zset, value)
}

func (client *Client) Zrange(zset string, start int64, end int64) (interface{}, error) {
	return executeZrange(client.getConnectPool(), zset, start, end)
}

func (client *Client) ZrangeWithScores(zset string, start int64, end int64) (interface{}, error) {
	return executeZrangeWithScores(client.getConnectPool(), zset, start, end)
}

func (client *Client) Zrevrange(zset string, start int64, end int64) (interface{}, error) {
	return executeZrevrange(client.getConnectPool(), zset, start, end)
}

func (client *Client) ZrevrangeWithScores(zset string, start int64, end int64) (interface{}, error) {
	return executeZrevrangeWithScores(client.getConnectPool(), zset, start, end)
}

func (client *Client) Zcard(zset string) (interface{}, error) {
	return executeZcard(client.getConnectPool(), zset)
}

func (client *Client) Zrem(zset string, elements ...string) (interface{}, error) {
	return executeZrem(client.getConnectPool(), zset, elements)
}

func (client *Client) Zrank(zset string, value string) (interface{}, error) {
	return executeZrank(client.getConnectPool(), zset, value)
}

func (client *Client) Zrevrank(zset string, value string) (interface{}, error) {
	return executeZrevrank(client.getConnectPool(), zset, value)
}

func executeZadd(pool *ConnPool, zset string, scoresvalues []string) (interface{}, error) {
	conn, err := GetConn(pool)
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	bytes := handler.HandleMultiBulkRequest(zset, scoresvalues)
	result := SendCommand(conn, protocol.ZADD, bytes...)
	return handler.HandleReply(result)
}

func executeZscore(pool *ConnPool, zset string, value string) (interface{}, error) {
	conn, err := GetConn(pool)
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	result := SendCommand(conn, protocol.ZSCORE, protocol.SafeEncode(zset), protocol.SafeEncode(value))
	return handler.HandleReply(result)
}

func executeZrange(pool *ConnPool, zset string, start int64, end int64) (interface{}, error) {
	conn, err := GetConn(pool)
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	result := SendCommand(conn, protocol.ZRANGE, protocol.SafeEncode(zset), protocol.SafeEncodeInt(start), protocol.SafeEncodeInt(end))
	return handler.HandleReply(result)
}

func executeZrangeWithScores(pool *ConnPool, zset string, start int64, end int64) (interface{}, error) {
	conn, err := GetConn(pool)
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	result := SendCommand(conn, protocol.ZRANGE, protocol.SafeEncode(zset), protocol.SafeEncodeInt(start), protocol.SafeEncodeInt(end), protocol.SafeEncode("WITHSCORES"))
	return handler.HandleReply(result)
}

func executeZrevrange(pool *ConnPool, zset string, start int64, end int64) (interface{}, error) {
	conn, err := GetConn(pool)
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	result := SendCommand(conn, protocol.ZREVRANGE, protocol.SafeEncode(zset), protocol.SafeEncodeInt(start), protocol.SafeEncodeInt(end))
	return handler.HandleReply(result)
}

func executeZrevrangeWithScores(pool *ConnPool, zset string, start int64, end int64) (interface{}, error) {
	conn, err := GetConn(pool)
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	result := SendCommand(conn, protocol.ZREVRANGE, protocol.SafeEncode(zset), protocol.SafeEncodeInt(start), protocol.SafeEncodeInt(end), protocol.SafeEncode("WITHSCORES"))
	return handler.HandleReply(result)
}

func executeZcard(pool *ConnPool, zset string) (interface{}, error) {
	conn, err := GetConn(pool)
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	result := SendCommand(conn, protocol.ZCARD, protocol.SafeEncode(zset))
	return handler.HandleReply(result)
}

func executeZrem(pool *ConnPool, zset string, elements []string) (interface{}, error) {
	conn, err := GetConn(pool)
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	bytes := handler.HandleMultiBulkRequest(zset, elements)
	result := SendCommand(conn, protocol.ZREM, bytes...)
	return handler.HandleReply(result)
}

func executeZrank(pool *ConnPool, zset string, value string) (interface{}, error) {
	conn, err := GetConn(pool)
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	result := SendCommand(conn, protocol.ZRANK, protocol.SafeEncode(zset))
	return handler.HandleReply(result)
}

func executeZrevrank(pool *ConnPool, zset string, value string) (interface{}, error) {
	conn, err := GetConn(pool)
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	result := SendCommand(conn, protocol.ZREVRANK, protocol.SafeEncode(zset))
	return handler.HandleReply(result)
}
