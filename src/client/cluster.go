package client

import (
	"time"
	"sync"
	"log"
	"fmt"
	"protocol"
	"client/handler"
)

//默认心跳检测轮询时间间隔，单位s
var defaultHeartBeatInterval = 10

var m *sync.RWMutex

/**
 * 节点
 * master：主节点ip+port
 * slaves：从节点ip+port集合
 */
type Node struct {
	Url        string
	Pwd        string
	InitActive int
	MinActive  int
	MaxActive  int
}

type ClusterConfig struct {
	Nodes             []*Node
	HeartBeatInterval int
}

/**
 * 集群客户端
 * heartBeatInterval 心跳检测时间间隔，单位s
 * clusterPool key：连接串 value:连接池
 */
type Cluster struct {
	config      *ClusterConfig
	clusterPool map[string]*ConnPool
}

func init() {
	LoggerInit()
}

/**
 * 初始化Cluster client
 */
func NewCluster(clusterConfig ClusterConfig) *Cluster {
	nodes := clusterConfig.Nodes

	var cluster Cluster
	clusterPool := make(map[string]*ConnPool)

	for _, node := range nodes {
		var config = ConnConfig{node.Url, node.Pwd, node.InitActive, node.MinActive, node.MaxActive}
		pool, _ := NewConnPool(config)
		clusterPool[node.Url] = pool
	}
	cluster.config = &clusterConfig
	cluster.clusterPool = clusterPool
	//初始化节点健康检测线程
	defer func() {
		go cluster.heartBeat()
	}()
	if m == nil {
		m = new(sync.RWMutex)
	}
	return &cluster
}

func (cluster *Cluster) GetClusterPool() map[string]*ConnPool {
	return cluster.clusterPool
}

func (cluster *Cluster) GetClusterNodesInfo() []*Node {
	return cluster.config.Nodes
}

func (cluster *Cluster) RandomSelect() *ConnPool {
	m.RLock()
	defer m.RUnlock()
	pools := cluster.GetClusterPool()
	for _, pool := range pools {
		if pool != nil {
			return pool
		}
	}
	fmt.Errorf("none pool can be used")
	return nil
}

func (cluster *Cluster) SelectOne(url string) *ConnPool {
	m.RLock()
	defer m.RUnlock()
	return cluster.GetClusterPool()[url]
}

/**
 * 连接池心跳检测，定时ping各个节点，ping失败的，从连接池退出，并将节点加入失败队列
 * 定时轮询失败节点队列，检测节点是否已恢复连接，若恢复，则重新创建连接池，并从失败队列中移除
 */
func (cluster *Cluster) heartBeat() {
	clusterPool := cluster.GetClusterPool()
	interval := cluster.config.HeartBeatInterval
	if interval <= 0 {
		interval = defaultHeartBeatInterval
	}
	var nodes = make(map[string]*Node)

	for i := 0; i < len(cluster.GetClusterNodesInfo()); i++ {
		node := cluster.GetClusterNodesInfo()[i]
		nodes[node.Url] = node
	}

	var failNodes = make(map[string]*Node)
	for {
		for url := range clusterPool {
			result, err := ping(nodes[url])
			if err != nil {
				log.Printf("节点[%s] 健康检查异常，原因[%s], 节点将被移除\n", url, err)
				//加锁
				m.Lock()
				failNodes[url] = nodes[url]
				delete(clusterPool, url)
				m.Unlock()
			} else {
				log.Printf("节点[%s] 健康检查结果[%s]\n", url, result)
			}
		}
		//恢复检测
		recover(failNodes, clusterPool)

		time.Sleep(time.Duration(interval) * time.Second)
	}
}

//ping redis server
func ping(node *Node) (interface{}, error) {
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
func recover(failNodes map[string]*Node, clusterPool map[string]*ConnPool) {
	for url, node := range failNodes {
		conn := Connect(url)
		if conn != nil {
			//节点重连,恢复连接
			var config = ConnConfig{url, node.Pwd, node.InitActive, node.MinActive, node.MaxActive}
			pool, _ := NewConnPool(config)
			//加锁
			m.Lock()
			clusterPool[node.Url] = pool
			delete(failNodes, url)
			m.Unlock()
			log.Printf("节点[%s] 已重连\n", url)
		}
	}
}
