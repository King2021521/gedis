package main

import (
	. "client"
	"fmt"
	"net"
	"time"
	"math/rand"
	"log"
)

func main() {
	/*testCluster()
	time.Sleep(time.Duration(100)*time.Second)*/
	client:=getClient()
	multiResult,_:=client.Multi()
	fmt.Println("开启事务：",multiResult)
	client.Get("age")
	client.Get("name")
	client.Get("age")
	execResult,_:=client.Discard()
	fmt.Println("终止事务结果：",execResult)
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
