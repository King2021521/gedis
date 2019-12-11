package client

import (
	"protocol"
	"client/handler"
	"fmt"
)

func (cluster *Cluster) Sadd(set string, elements ...string) (interface{}, error) {
	result, err := executeSadd(cluster.RandomSelect(), set, elements)
	if err==nil || err.Error() != protocol.MOVED {
		return result, err
	}

	//重定向到新的节点
	return executeSadd(cluster.SelectOne(result.(string)), set, elements)
}

func (cluster *Cluster) Smembers(set string) (interface{}, error) {
	result, err := executeSmembers(cluster.RandomSelect(), set)
	if err==nil || err.Error() != protocol.MOVED {
		return result, err
	}

	//重定向到新的节点
	return executeSmembers(cluster.SelectOne(result.(string)), set)
}

func (cluster *Cluster) Srem(set string, elements ...string) (interface{}, error) {
	result, err := executeSrem(cluster.RandomSelect(), set, elements)
	if err==nil || err.Error() != protocol.MOVED {
		return result, err
	}

	//重定向到新的节点
	return executeSrem(cluster.SelectOne(result.(string)), set, elements)
}

func (cluster *Cluster) Sismember(set string, value string) (interface{}, error) {
	result, err := executeSismember(cluster.RandomSelect(), set, value)
	if err==nil || err.Error() != protocol.MOVED {
		return result, err
	}

	//重定向到新的节点
	return executeSismember(cluster.SelectOne(result.(string)), set, value)
}

func (cluster *Cluster) Scard(set string) (interface{}, error) {
	result, err := executeScard(cluster.RandomSelect(), set)
	if err==nil || err.Error() != protocol.MOVED {
		return result, err
	}

	//重定向到新的节点
	return executeScard(cluster.SelectOne(result.(string)), set)
}

func (cluster *Cluster) Srandmembers(set string, count int64) (interface{}, error) {
	result, err := executeSrandmembers(cluster.RandomSelect(), set, count)
	if err==nil || err.Error() != protocol.MOVED {
		return result, err
	}

	//重定向到新的节点
	return executeSrandmembers(cluster.SelectOne(result.(string)), set, count)
}

func (cluster *Cluster) Spop(set string) (interface{}, error) {
	result, err := executeSpop(cluster.RandomSelect(), set)
	if err==nil || err.Error() != protocol.MOVED {
		return result, err
	}

	//重定向到新的节点
	return executeSpop(cluster.SelectOne(result.(string)), set)
}

/**
 * 返回给定所有集合的差集
 */
func (cluster *Cluster) Sdiff(sets ... string) (interface{}, error) {
	result, err := executeSdiff(cluster.RandomSelect(), sets)
	if err==nil || err.Error() != protocol.MOVED {
		return result, err
	}

	//重定向到新的节点
	return executeSdiff(cluster.SelectOne(result.(string)), sets)
}

func (client *Client) Sadd(set string, elements ...string) (interface{}, error) {
	return executeSadd(client.getConnectPool(), set, elements)
}

func (client *Client) Smembers(set string) (interface{}, error) {
	return executeSmembers(client.getConnectPool(), set)
}

func (client *Client) Srem(set string, elements ...string) (interface{}, error) {
	return executeSrem(client.getConnectPool(), set, elements)
}

func (client *Client) Sismember(set string, value string) (interface{}, error) {
	return executeSismember(client.getConnectPool(), set, value)
}

func (client *Client) Scard(set string) (interface{}, error) {
	return executeScard(client.getConnectPool(), set)
}

func (client *Client) Srandmembers(set string, count int64) (interface{}, error) {
	return executeSrandmembers(client.getConnectPool(), set, count)
}

func (client *Client) Spop(set string) (interface{}, error) {
	return executeSpop(client.getConnectPool(), set)
}

/**
 * 返回给定所有集合的差集
 */
func (client *Client) Sdiff(sets ... string) (interface{}, error) {
	return executeSdiff(client.getConnectPool(), sets)
}

func executeSadd(pool *ConnPool, set string, elements []string) (interface{}, error) {
	conn, err := pool.GetConn()
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	bytes := handler.HandleMultiBulkRequest(set, elements)
	result := SendCommand(conn, protocol.SADD, bytes...)
	return handler.HandleReply(result)
}

func executeSmembers(pool *ConnPool, set string) (interface{}, error) {
	conn, err := pool.GetConn()
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	result := SendCommand(conn, protocol.SMEMBERS, protocol.SafeEncode(set))
	return handler.HandleReply(result)
}

func executeSrem(pool *ConnPool, set string, elements []string) (interface{}, error) {
	conn, err := pool.GetConn()
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	bytes := handler.HandleMultiBulkRequest(set, elements)
	result := SendCommand(conn, protocol.SREM, bytes...)
	return handler.HandleReply(result)
}

func executeSismember(pool *ConnPool, set string, value string) (interface{}, error) {
	conn, err := pool.GetConn()
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	result := SendCommand(conn, protocol.SISMEMBER, protocol.SafeEncode(set), protocol.SafeEncode(value))
	return handler.HandleReply(result)
}

func executeScard(pool *ConnPool, set string) (interface{}, error) {
	conn, err := pool.GetConn()
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	result := SendCommand(conn, protocol.SCARD, protocol.SafeEncode(set))
	return handler.HandleReply(result)
}

func executeSrandmembers(pool *ConnPool, set string, count int64) (interface{}, error) {
	conn, err := pool.GetConn()
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	result := SendCommand(conn, protocol.SRANDMEMBER, protocol.SafeEncode(set), protocol.SafeEncodeInt(count))
	return handler.HandleReply(result)
}

func executeSpop(pool *ConnPool, set string) (interface{}, error) {
	conn, err := pool.GetConn()
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
func executeSdiff(pool *ConnPool, sets []string) (interface{}, error) {
	conn, err := pool.GetConn()
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	bytes := handler.HandleBulkRequest(sets)
	result := SendCommand(conn, protocol.SDIFF, bytes...)
	return handler.HandleReply(result)
}
