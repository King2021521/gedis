package client

import (
	"errors"
	"net"
	"sync"
	"protocol"
)

const defaultMaxActive = 100
const defaultMinActive = 5
const defaultIntActive = 10

type ConnConfig struct {
	ConnString string
	Pwd        string
	InitActive int
	MinActive  int
	MaxActive  int
}

type ConnPool struct {
	connPool   chan *net.TCPConn
	initActive int
	minActive  int
	maxActive  int
}

var pool *ConnPool
var oSingle sync.Once

/**
 * 单例的连接池（线程安全）
 */
func NewSingleConnPool(config ConnConfig) *ConnPool {
	oSingle.Do(func() {
		pool, _ = NewConnPool(config)
	})
	return pool
}

/**
 * 初始化连接池
 */
func NewConnPool(config ConnConfig) (*ConnPool, error) {
	config.validate()

	var pool ConnPool
	pool.initActive = config.InitActive
	pool.minActive = config.MinActive
	pool.maxActive = config.MaxActive

	channel := make(chan *net.TCPConn, config.InitActive)

	for index := 0; index < config.InitActive; index++ {
		//初始化一个连接
		conn := Connect(config.ConnString)
		//设置keepalive
		conn.SetKeepAlive(true)
		//为当前连接授权
		auth(conn, config.Pwd)
		//将连接加入连接池
		channel <- conn
	}

	pool.connPool = channel
	return &pool, nil
}

func (config *ConnConfig) validate() {
	if config.InitActive < 0 {
		config.InitActive = defaultIntActive
	}
	if config.MinActive < 0 {
		config.MinActive = defaultMinActive
	}
	if config.MaxActive < 0 {
		config.MaxActive = defaultMaxActive
	}
	if config.MinActive > config.InitActive {
		config.MinActive = config.InitActive
	}
	if config.MaxActive < config.InitActive {
		config.MaxActive = config.InitActive
	}
}

func auth(conn *net.TCPConn, pwd string) {
	SendCommand(conn, protocol.AUTH, protocol.SafeEncode(pwd))
}

/**
 * 从连接池中获取连接
 */
func (pool *ConnPool) GetConn() (*net.TCPConn, error) {
	if pool.PoolSize() == 0 {
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
func (pool *ConnPool) PoolSize() int {
	return len(pool.connPool)
}

func (pool *ConnPool) setMaxActive(maxActive int) {
	pool.maxActive = maxActive
}

func (pool *ConnPool) setMinActive(minActive int) {
	pool.minActive = minActive
}

func (pool *ConnPool) setInitActive(initActive int) {
	pool.initActive = initActive
}
