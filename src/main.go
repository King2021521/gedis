package main

import (
	. "client"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"
)

func main() {
	testSharding()
	//testConsistentHash()
	/*testCluster()
	time.Sleep(time.Duration(100)*time.Second)*/
	/*client:=getClient()
	multiResult,_:=client.Multi()
	fmt.Println("开启事务：",multiResult)
	client.GetShardInfo("age")
	client.GetShardInfo("name")
	client.GetShardInfo("age")
	execResult,_:=client.Discard()
	fmt.Println("终止事务结果：",execResult)*/
}

func testConsistentHash(){
	cHashRing := NewConsistent()
	//模拟5台服务器，加入到hash环中
	for i := 0; i < 5; i++ {
		si := fmt.Sprintf("%d", i)
		cHashRing.Add(NewShardInfo(i, "192.168.216.1."+si+":6379", 1))
	}
	fmt.Println("-------------", len(cHashRing.Nodes))
	//输出节点的hash值
	for k, v := range cHashRing.Nodes {
		fmt.Println("Hash:", k, " IP:", v.Url)
	}

	//模拟数据查询
	ipMap := make(map[string]int, 0)
	for i := 0; i < 10; i++ {
		si := fmt.Sprintf("key%d", i)
		k := cHashRing.GetShardInfo(si)
		if _, ok := ipMap[k.Url]; ok {
			ipMap[k.Url] += 1
		} else {
			ipMap[k.Url] = 1
		}
	}

	for k, v := range ipMap {
		fmt.Println("Node IP:", k, " count:", v)
	}
}

func testSharding(){
	var s1 = Shard{"192.168.96.232:6379", "vs959yUyx3", 50,5,200}
	var s2 = Shard{"192.168.96.4:6379", "K5re9U#mX@", 50,5,200}
	shards := []*Shard{&s1, &s2}
	var shardConfig = ShardConfig{shards, 10}
	sharding := NewSharding(shardConfig)
	rand.Seed(time.Now().UnixNano())
	for i:=0;i<20;i++{
		go func() {
			for {
				value, err := sharding.Get("teams")
				log.Printf("请求结果：%s, err: %s",value, err)
				fmt.Println("查询结果",value)
				time.Sleep(time.Duration(rand.Intn(3))*time.Second)
			}
		}()
	}
	time.Sleep(1000*10)
}

func testCluster() {
	var node7000 = Node{"127.0.0.1:7000", "123456", 50,5,200}
	var node7001 = Node{"127.0.0.1:7001", "123456", 50,5,200}
	var node7002 = Node{"127.0.0.1:7002", "123456", 50,5,200}
	var node7003 = Node{"127.0.0.1:7003", "123456", 50,5,200}
	var node7004 = Node{"127.0.0.1:7004", "123456", 50,5,200}
	var node7005 = Node{"127.0.0.1:7005", "123456", 50,5,200}

	nodes := []*Node{&node7000, &node7001, &node7002, &node7003, &node7004, &node7005}
	var clusterConfig = ClusterConfig{nodes, 10}
	cluster := NewCluster(clusterConfig)
	rand.Seed(time.Now().UnixNano())
	for i:=0;i<20;i++{
		go func() {
			for {
				value, err := cluster.Get("name")
				log.Printf("请求结果：%s, err: %s",value, err)
				time.Sleep(time.Duration(rand.Intn(3))*time.Second)
			}
		}()
	}
}

func getConn() *net.TCPConn {
	var config = ConnConfig{"127.0.0.1:6379", "root",1,1,1}
	pool := NewSingleConnPool(config)
	conn, _ := pool.GetConn()
	return conn
}

func getClient() *Client {
	var config = ConnConfig{"127.0.0.1:6379", "",1,1,1}
	pool := NewSingleConnPool(config)

	return BuildClient(pool)
}
