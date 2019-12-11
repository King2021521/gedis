package client

import (
	"net"
	"protocol"
	"client/handler"
	"fmt"
)

func (cluster *Cluster) Lpush(list string, elements ...string) (interface{}, error) {
	result, err := executeLpush(cluster.RandomSelect(), list, elements)
	if err==nil || err.Error() != protocol.MOVED {
		return result, err
	}

	//重定向到新的节点
	return executeLpush(cluster.SelectOne(result.(string)), list, elements)
}

func (cluster *Cluster) Rpush(list string, elements ...string) (interface{}, error) {
	result, err := executeRpush(cluster.RandomSelect(), list, elements)
	if err==nil || err.Error() != protocol.MOVED {
		return result, err
	}

	//重定向到新的节点
	return executeRpush(cluster.SelectOne(result.(string)), list, elements)
}

func (cluster *Cluster) Lrange(list string, start int64, end int64) (interface{}, error) {
	result, err := executeLrange(cluster.RandomSelect(), list, start, end)
	if err==nil || err.Error() != protocol.MOVED {
		return result, err
	}

	//重定向到新的节点
	return executeLrange(cluster.SelectOne(result.(string)), list, start, end)
}

func (cluster *Cluster) Lpop(list string) (interface{}, error) {
	result, err := executeLpop(cluster.RandomSelect(), list)
	if err==nil || err.Error() != protocol.MOVED {
		return result, err
	}

	//重定向到新的节点
	return executeLpop(cluster.SelectOne(result.(string)), list)
}

func (cluster *Cluster) Rpop(list string) (interface{}, error) {
	result, err := executeRpop(cluster.RandomSelect(), list)
	if err==nil || err.Error() != protocol.MOVED {
		return result, err
	}

	//重定向到新的节点
	return executeRpop(cluster.SelectOne(result.(string)), list)
}

func (cluster *Cluster) Llen(list string) (interface{}, error) {
	result, err := executeLlen(cluster.RandomSelect(), list)
	if err==nil || err.Error() != protocol.MOVED {
		return result, err
	}

	//重定向到新的节点
	return executeLlen(cluster.SelectOne(result.(string)), list)
}

/**
 *  count > 0 : 从表头开始向表尾搜索，移除与 VALUE 相等的元素，数量为 COUNT 。
 *  count < 0 : 从表尾开始向表头搜索，移除与 VALUE 相等的元素，数量为 COUNT 的绝对值。
 *  count = 0 : 移除表中所有与 VALUE 相等的值。
 */
func (cluster *Cluster) Lrem(list string, count int64, value string) (interface{}, error) {
	result, err := executeLrem(cluster.RandomSelect(), list, count, value)
	if err==nil || err.Error() != protocol.MOVED {
		return result, err
	}

	//重定向到新的节点
	return executeLrem(cluster.SelectOne(result.(string)), list, count, value)
}

func (cluster *Cluster) Lindex(list string, pos int64) (interface{}, error) {
	result, err := executeLindex(cluster.RandomSelect(), list, pos)
	if err==nil || err.Error() != protocol.MOVED {
		return result, err
	}

	//重定向到新的节点
	return executeLindex(cluster.SelectOne(result.(string)), list, pos)
}

func (cluster *Cluster) Lset(list string, pos int64, value string) (interface{}, error) {
	result, err := executeLset(cluster.RandomSelect(), list, pos, value)
	if err==nil || err.Error() != protocol.MOVED {
		return result, err
	}

	//重定向到新的节点
	return executeLset(cluster.SelectOne(result.(string)), list, pos, value)
}

/**
 * target 目标元素
 */
func (cluster *Cluster) LinsertBefore(list string, target string, value string) (interface{}, error) {
	result, err := executeLinsertBefore(cluster.RandomSelect(), list, target, value)
	if err==nil || err.Error() != protocol.MOVED {
		return result, err
	}

	//重定向到新的节点
	return executeLinsertBefore(cluster.SelectOne(result.(string)), list, target, value)
}

/**
 * target 目标元素
 */
func (cluster *Cluster) LinsertAfter(list string, target string, value string) (interface{}, error) {
	result, err := executeLinsertAfter(cluster.RandomSelect(), list, target, value)
	if err==nil || err.Error() != protocol.MOVED {
		return result, err
	}

	//重定向到新的节点
	return executeLinsertAfter(cluster.SelectOne(result.(string)), list, target, value)
}

func (client *Client) Lpush(list string, elements ...string) (interface{}, error) {
	return executeLpush(client.getConnectPool(), list, elements)
}

func (client *Client) Rpush(list string, elements ...string) (interface{}, error) {
	return executeRpush(client.getConnectPool(), list, elements)
}

func (client *Client) Lrange(list string, start int64, end int64) (interface{}, error) {
	return executeLrange(client.getConnectPool(), list, start, end)
}

func (client *Client) Lpop(list string) (interface{}, error) {
	return executeLpop(client.getConnectPool(), list)
}

func (client *Client) Rpop(list string) (interface{}, error) {
	return executeRpop(client.getConnectPool(), list)
}

func (client *Client) Llen(list string) (interface{}, error) {
	return executeLlen(client.getConnectPool(), list)
}

/**
 *  count > 0 : 从表头开始向表尾搜索，移除与 VALUE 相等的元素，数量为 COUNT 。
 *  count < 0 : 从表尾开始向表头搜索，移除与 VALUE 相等的元素，数量为 COUNT 的绝对值。
 *  count = 0 : 移除表中所有与 VALUE 相等的值。
 */
func (client *Client) Lrem(list string, count int64, value string) (interface{}, error) {
	return executeLrem(client.getConnectPool(), list, count, value)
}

func (client *Client) Lindex(list string, pos int64) (interface{}, error) {
	return executeLindex(client.getConnectPool(), list, pos)
}

func (client *Client) Lset(list string, pos int64, value string) (interface{}, error) {
	return executeLset(client.getConnectPool(), list, pos, value)
}

/**
 * target 目标元素
 */
func (client *Client) LinsertBefore(list string, target string, value string) (interface{}, error) {
	return executeLinsertBefore(client.getConnectPool(), list, target, value)
}

/**
 * target 目标元素
 */
func (client *Client) LinsertAfter(list string, target string, value string) (interface{}, error) {
	return executeLinsertAfter(client.getConnectPool(), list, target, value)
}

func executeLpush(pool *ConnPool, list string, elements []string) (interface{}, error) {
	conn, err := pool.GetConn()
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	return push(conn, list, protocol.LPUSH, elements)
}

func executeRpush(pool *ConnPool, list string, elements []string) (interface{}, error) {
	conn, err := pool.GetConn()
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	return push(conn, list, protocol.RPUSH, elements)
}

func executeLrange(pool *ConnPool, list string, start int64, end int64) (interface{}, error) {
	conn, err := pool.GetConn()
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	result := SendCommand(conn, protocol.LRANGE, protocol.SafeEncode(list), protocol.SafeEncodeInt(start), protocol.SafeEncodeInt(end))
	return handler.HandleReply(result)
}

func executeLpop(pool *ConnPool, list string) (interface{}, error) {
	conn, err := pool.GetConn()
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	return pop(conn, list, protocol.LPOP)
}

func executeRpop(pool *ConnPool, list string) (interface{}, error) {
	conn, err := pool.GetConn()
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	return pop(conn, list, protocol.RPOP)
}

func executeLlen(pool *ConnPool, list string) (interface{}, error) {
	conn, err := pool.GetConn()
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	result := SendCommand(conn, protocol.LLEN, protocol.SafeEncode(list))
	return handler.HandleReply(result)
}

/**
 *  count > 0 : 从表头开始向表尾搜索，移除与 VALUE 相等的元素，数量为 COUNT 。
 *  count < 0 : 从表尾开始向表头搜索，移除与 VALUE 相等的元素，数量为 COUNT 的绝对值。
 *  count = 0 : 移除表中所有与 VALUE 相等的值。
 */
func executeLrem(pool *ConnPool, list string, count int64, value string) (interface{}, error) {
	conn, err := pool.GetConn()
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	result := SendCommand(conn, protocol.LREM, protocol.SafeEncode(list), protocol.SafeEncodeInt(count), protocol.SafeEncode(value))
	return handler.HandleReply(result)
}

func executeLindex(pool *ConnPool, list string, pos int64) (interface{}, error) {
	conn, err := pool.GetConn()
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	result := SendCommand(conn, protocol.LINDEX, protocol.SafeEncode(list), protocol.SafeEncodeInt(pos))
	return handler.HandleReply(result)
}

func executeLset(pool *ConnPool, list string, pos int64, value string) (interface{}, error) {
	conn, err := pool.GetConn()
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	result := SendCommand(conn, protocol.LSET, protocol.SafeEncode(list), protocol.SafeEncodeInt(pos), protocol.SafeEncode(value))
	return handler.HandleReply(result)
}

/**
 * target 目标元素
 */
func executeLinsertBefore(pool *ConnPool, list string, target string, value string) (interface{}, error) {
	conn, err := pool.GetConn()
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	result := SendCommand(conn, protocol.LINSERT, protocol.SafeEncode(list), protocol.SafeEncode("BEFORE"), protocol.SafeEncode(target), protocol.SafeEncode(value))
	return handler.HandleReply(result)
}

/**
 * target 目标元素
 */
func executeLinsertAfter(pool *ConnPool, list string, target string, value string) (interface{}, error) {
	conn, err := pool.GetConn()
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	result := SendCommand(conn, protocol.LINSERT, protocol.SafeEncode(list), protocol.SafeEncode("AFTER"), protocol.SafeEncode(target), protocol.SafeEncode(value))
	return handler.HandleReply(result)
}

func push(conn *net.TCPConn, list string, cmd string, elements []string) (interface{}, error) {
	bytes := handler.HandleMultiBulkRequest(list, elements)
	result := SendCommand(conn, cmd, bytes...)
	return handler.HandleReply(result)
}

func pop(conn *net.TCPConn, list string, cmd string) (interface{}, error) {
	result := SendCommand(conn, cmd, protocol.SafeEncode(list))
	return handler.HandleReply(result)
}
