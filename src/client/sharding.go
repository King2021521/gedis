package client

import (
	"client/handler"
	"log"
	"protocol"
	"sync"
	"time"
)

//默认心跳检测轮询时间间隔，单位s
var _defaultHeartBeatInterval = 10

/**
 * 分片信息
 *
 */
type Shard struct {
	Url        string
	Pwd        string
	InitActive int
	MinActive  int
	MaxActive  int
}
//分片配置
type ShardConfig struct {
	Shards             []*Shard
	HeartBeatInterval int
}

/**
 * 客户端分片
 * heartBeatInterval 心跳检测时间间隔，单位s
 * shardingPool key：连接串 value:连接池
 */
type Sharding struct {
	cHashRing   *Consistent
	config      *ShardConfig
	shardingPool map[string]*ConnPool
	m *sync.RWMutex
}

//初始化日志
func init() {
	LoggerInit()
}

//初始化分片客户端
func NewSharding(shardConfig ShardConfig) *Sharding {
	shards := shardConfig.Shards
	//初始化一致性hash环
	cHashRing := NewConsistent()

	var sharding Sharding
	shardingPool := make(map[string]*ConnPool)

	for index, shard := range shards {
		var config = ConnConfig{shard.Url, shard.Pwd, shard.InitActive, shard.MinActive, shard.MaxActive}
		pool, _ := NewConnPool(config)
		shardingPool[shard.Url] = pool
		//将分片散列到hash环中
		cHashRing.Add(NewShardInfo(index, shard.Url, 1))
	}
	sharding.cHashRing = cHashRing
	sharding.config = &shardConfig
	sharding.shardingPool = shardingPool

	if sharding.m == nil {
		sharding.m = new(sync.RWMutex)
	}
	return &sharding
}

//获取分片连接池
func (sharding *Sharding) GetShardingPool() map[string]*ConnPool {
	return sharding.shardingPool
}

//获取分片信息
func (sharding *Sharding) GetShardInfo() []*Shard {
	return sharding.config.Shards
}

/**
 * 连接池心跳检测，定时ping各个节点，ping失败的，从连接池退出，并将节点加入失败队列
 * 定时轮询失败节点队列，检测节点是否已恢复连接，若恢复，则重新创建连接池，并从失败队列中移除
 */
func (sharding *Sharding) heartBeat() {
	shardingPool := sharding.GetShardingPool()
	interval := sharding.config.HeartBeatInterval
	if interval <= 0 {
		interval = _defaultHeartBeatInterval
	}
	var shards = make(map[string]*Shard)

	for i := 0; i < len(sharding.GetShardInfo()); i++ {
		shard := sharding.GetShardInfo()[i]
		shards[shard.Url] = shard
	}

	var failNodes = make(map[string]*Shard)
	for {
		for url := range shardingPool {
			result, err := _ping(shards[url])
			if err != nil {
				log.Printf("节点[%s] 健康检查异常，原因[%s], 节点将被移除\n", url, err)
				//加锁
				sharding.m.Lock()
				failNodes[url] = shards[url]
				delete(shardingPool, url)
				sharding.m.Unlock()
			} else {
				log.Printf("节点[%s] 健康检查结果[%s]\n", url, result)
			}
		}
		//恢复检测
		_recover(failNodes, shardingPool)

		time.Sleep(time.Duration(interval) * time.Second)
	}
}

//ping redis server
func _ping(node *Shard) (interface{}, error) {
	conn := Connect(node.Url)
	//设置keepalive
	conn.SetKeepAlive(true)
	//为当前连接授权
	auth(conn, node.Pwd)
	defer conn.Close()
	result := SendCommand(conn, protocol.PING)
	return handler.HandleReply(result)
}

/**
 * 检测fail节点是否已恢复正常
 */
func _recover(failNodes map[string]*Shard, shardingPool map[string]*ConnPool) {
	for url, node := range failNodes {
		conn := Connect(url)
		if conn != nil {
			//节点重连,恢复连接
			var config = ConnConfig{url, node.Pwd, node.InitActive, node.MinActive, node.MaxActive}
			pool, _ := NewConnPool(config)
			//加锁
			m.Lock()
			shardingPool[node.Url] = pool
			delete(failNodes, url)
			m.Unlock()
			log.Printf("节点[%s] 已重连\n", url)
		}
	}
}