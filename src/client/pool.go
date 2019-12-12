package client

import (
	"errors"
	"net"
	"sync"
	"protocol"
	"time"
	"log"
)

const defaultMaxActive = 100
const defaultMinActive = 5
const defaultIntActive = 10
const timeCheckPoolSeconds = 30

type ConnConfig struct {
	ConnString string
	Pwd        string
	InitActive int
	MinActive  int
	MaxActive  int
}

type ConnPool struct {
	connPool   chan *net.TCPConn
	connConfig ConnConfig
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
	pool.connConfig = config

	channel := make(chan *net.TCPConn, config.InitActive)
	create(channel, pool.connConfig, config.InitActive)
	defer func() { go pool.startupMonitor() }()
	pool.connPool = channel
	return &pool, nil
}

func create(channel chan *net.TCPConn, config ConnConfig, size int) {
	for index := 0; index < size; index++ {
		//初始化一个连接
		conn := Connect(config.ConnString)
		//设置keepalive
		conn.SetKeepAlive(true)
		//为当前连接授权
		auth(conn, config.Pwd)
		//将连接加入连接池
		channel <- conn
	}
}

//连接池监控任务，当连接池的连接数量不足时，进行扩充；连接数过多时进行回收
func (pool *ConnPool) startupMonitor() {
	for {
		time.Sleep(time.Duration(timeCheckPoolSeconds) * time.Second)
		log.Printf("执行连接池连接数监控，当前节点：{%s}", pool.connConfig.ConnString)
		size := pool.PoolSize()
		log.Printf("节点{%s}当前连接数{%d}：", pool.connConfig.ConnString, size)
		if size < pool.connConfig.MinActive {
			//连接数不足
			create(pool.connPool, pool.connConfig, pool.connConfig.MinActive-size)
			continue
		}
		if size > pool.connConfig.MaxActive {
			//回收过多的连接
			for i := 0; i < (size - pool.connConfig.MaxActive); i++ {
				conn := <-pool.connPool
				conn.Close()
			}
		}
	}
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
	pool.connConfig.MaxActive = maxActive
}

func (pool *ConnPool) setMinActive(minActive int) {
	pool.connConfig.MinActive = minActive
}

func (pool *ConnPool) setInitActive(initActive int) {
	pool.connConfig.InitActive = initActive
}
