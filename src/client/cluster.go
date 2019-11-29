package client

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
	heartBeatInterval int
	clusterPool       map[string]*ConnPool
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
		pool,_ := NewConnPool(node.InitActive, config)
		clusterPool[node.Url] = pool
	}
	cluster.heartBeatInterval = clusterConfig.HeartBeatInterval
	cluster.clusterPool = clusterPool
	return &cluster
}

func (cluster *Cluster) GetClusterPool()map[string]*ConnPool{
	return cluster.clusterPool
}
