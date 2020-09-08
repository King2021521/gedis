package client

import (
	"hash/crc32"
	"sort"
	"strconv"
	"sync"
)
//默认每个节点的虚拟节点数量
const DEFAULT_REPLICAS = 160
//hash环
type HashRing []uint32

func (c HashRing) Len() int {
	return len(c)
}

func (c HashRing) Less(i, j int) bool {
	return c[i] < c[j]
}

func (c HashRing) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
//分片信息
type ShardInfo struct {
	Id       int
	Url       string
	Weight   int
}

func NewShardInfo(id int, url string, weight int) *ShardInfo {
	return &ShardInfo{
		Id:       id,
		Url:       url,
		Weight:   weight,
	}
}

type Consistent struct {
	Nodes     map[uint32]ShardInfo
	numReps   int
	Resources map[int]bool
	ring      HashRing
	sync.RWMutex
}

func NewConsistent() *Consistent {
	return &Consistent{
		Nodes:     make(map[uint32]ShardInfo),
		numReps:   DEFAULT_REPLICAS,
		Resources: make(map[int]bool),
		ring:      HashRing{},
	}
}

//将节点加入到ring上
func (c *Consistent) Add(node *ShardInfo) bool {
	c.Lock()
	defer c.Unlock()

	if _, ok := c.Resources[node.Id]; ok {
		return false
	}
	//散列在环上的虚拟节点数量（基础值*权重）
	count := c.numReps * node.Weight
	for i := 0; i < count; i++ {
		str := c.joinStr(i, node)
		c.Nodes[c.hashStr(str)] = *(node)
	}
	c.Resources[node.Id] = true
	c.sortHashRing()
	return true
}

func (c *Consistent) sortHashRing() {
	c.ring = HashRing{}
	for k := range c.Nodes {
		c.ring = append(c.ring, k)
	}
	sort.Sort(c.ring)
}

func (c *Consistent) joinStr(i int, node *ShardInfo) string {
	return node.Url + "*" + strconv.Itoa(node.Weight) +
		"-" + strconv.Itoa(i) +
		"-" + strconv.Itoa(node.Id)
}

// MurMurHash算法 :https://github.com/spaolacci/murmur3
func (c *Consistent) hashStr(key string) uint32 {
	return crc32.ChecksumIEEE([]byte(key))
}

//根据key获取分片信息
func (c *Consistent) GetShardInfo(key string) ShardInfo {
	c.RLock()
	defer c.RUnlock()

	hash := c.hashStr(key)
	i := c.search(hash)

	return c.Nodes[c.ring[i]]
}

func (c *Consistent) search(hash uint32) int {

	i := sort.Search(len(c.ring), func(i int) bool { return c.ring[i] >= hash })
	if i < len(c.ring) {
		if i == len(c.ring)-1 {
			return 0
		} else {
			return i
		}
	} else {
		return len(c.ring) - 1
	}
}

func (c *Consistent) Remove(node *ShardInfo) {
	c.Lock()
	defer c.Unlock()

	if _, ok := c.Resources[node.Id]; !ok {
		return
	}

	delete(c.Resources, node.Id)

	count := c.numReps * node.Weight
	for i := 0; i < count; i++ {
		str := c.joinStr(i, node)
		delete(c.Nodes, c.hashStr(str))
	}
	c.sortHashRing()
}
