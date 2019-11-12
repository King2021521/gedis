package tcp

import (
	"errors"
	"net"
	"template"
)

type ConnConfig struct {
	ConnString string
	Pwd        string
}

type ConnPool struct {
	connPool   chan *net.TCPConn
	initActive int
}

/**
 * 初始化连接池
 */
func NewConnPool(initActive int, config ConnConfig) (*ConnPool, error) {
	if initActive <= 0 {
		return nil, errors.New("maxActive must gt than 0")
	}

	var pool ConnPool
	channel := make(chan *net.TCPConn, initActive)
	pool.initActive = initActive

	for index := 0; index < initActive; index++ {
		//初始化一个连接
		conn := Connect(config.ConnString)
		//设置keepalive
		conn.SetKeepAlive(true)
		//为当前连接授权
		template.Auth(conn, config.Pwd)
		//将连接加入连接池
		channel <- conn
	}

	pool.connPool = channel
	return &pool, nil
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
