package tcp

import (
	"errors"
	"net"
	"fmt"
	"protocol"
)

type ConnConfig struct {
	ConnString string
	Pwd        string
}

type ConnPool struct {
	connPool   chan *net.TCPConn
	maxActive  int
	initActive int
}

/**
 * 初始化连接池
 */
func NewConnPool(maxActive int, initActive int, config ConnConfig) (*ConnPool, error) {
	if maxActive <= 0 {
		return nil, errors.New("maxActive must gt than 0")
	}
	if initActive <= 0 {
		return nil, errors.New("maxActive must gt than 0")
	}

	var pool ConnPool
	channel := make(chan *net.TCPConn, initActive)
	pool.maxActive = maxActive
	pool.initActive = initActive

	for index := 0; index < initActive; index++ {
		//初始化一个连接
		conn := Connect(config.ConnString)
		//设置keepalive
		conn.SetKeepAlive(true)
		//为当前连接授权
		auth(config.Pwd, conn)
		//将连接加入连接池
		channel <- conn
	}

	pool.connPool = channel
	return &pool, nil
}

/**
 * 授权当前连接
 */
func auth(pwd string, conn *net.TCPConn) string {
	var content = fmt.Sprintf(protocol.AuthCmdFormat, len(pwd), pwd)
	result := Sender(conn, content)
	return result
}

/**
 * 从连接池中获取连接
 */
func GetConn(pool *ConnPool) (*net.TCPConn, error) {
	if PoolSize(pool) == 0 {
		return nil, errors.New("连接数已不足")
	}

	conn := <-pool.connPool
	if conn == nil {
		return nil, errors.New("连接数已不足")
	}

	return conn, nil
}

/**
 * 将连接归还到连接池
 */
func (pool *ConnPool) PutConn(conn *net.TCPConn) error {
	if conn == nil {
		return errors.New("连接为空")
	}
	pool.connPool <- conn
	return nil
}

/**
 * 返回连接池当前连接数
 */
func PoolSize(pool *ConnPool) int {
	return len(pool.connPool)
}
