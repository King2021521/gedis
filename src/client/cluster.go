package client

import (
	"time"
	"math/rand"
)

/**
 * 节点
 * master：主节点ip+port
 * slaves：从节点ip+port集合
 */
type Node struct {
	Url        string
	Pwd        string
	InitActive int
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

/**
 * 初始化Cluster client
 */
func NewCluster(clusterConfig ClusterConfig) *Cluster {
	nodes := clusterConfig.Nodes

	var cluster Cluster
	clusterPool := make(map[string]*ConnPool)

	for _, node := range nodes {
		var config = ConnConfig{node.Url, node.Pwd}
		pool, _ := NewConnPool(node.InitActive, config)
		clusterPool[node.Url] = pool
	}
	cluster.config = &clusterConfig
	cluster.clusterPool = clusterPool
	return &cluster
}

func (cluster *Cluster) GetClusterPool() map[string]*ConnPool {
	return cluster.clusterPool
}

func (cluster *Cluster) GetClusterNodesInfo() []*Node {
	return cluster.config.Nodes
}

func (cluster *Cluster) RandomSelect() *ConnPool{
	pools := cluster.GetClusterPool()
	nodes := cluster.GetClusterNodesInfo()
	//负载均衡，随机选择一个节点执行访问
	rand.Seed(time.Now().UnixNano())
	nodeId := rand.Intn(len(nodes))
	pool := pools[nodes[nodeId].Url]
	return pool
}

func (cluster *Cluster) SelectOne(url string) *ConnPool{
	return cluster.GetClusterPool()[url]
}
