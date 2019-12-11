package main

import (
	. "client"
	"net"
	"fmt"
	"time"
)

func main() {
	testCluster()
	time.Sleep(time.Duration(100)*time.Second)
}

func testCluster() {
	var node7000 = Node{"127.0.0.1:7000", "123456", 10,5,20}
	var node7001 = Node{"127.0.0.1:7001", "123456", 10,5,20}
	var node7002 = Node{"127.0.0.1:7002", "123456", 10,5,20}
	var node7003 = Node{"127.0.0.1:7003", "123456", 10,5,20}
	var node7004 = Node{"127.0.0.1:7004", "123456", 10,5,20}
	var node7005 = Node{"127.0.0.1:7005", "123456", 10,5,20}

	nodes := []*Node{&node7000, &node7001, &node7002, &node7003, &node7004, &node7005}
	var clusterConfig = ClusterConfig{nodes, 10}
	cluster := NewCluster(clusterConfig)
	value, err := cluster.Get("name")
	fmt.Println(value, err)
	for i := 0; i < 100; i++ {
		value, err := cluster.Get("name")
		fmt.Println(value, err)
		time.Sleep(time.Duration(2) * time.Second)
	}
}

func getConn() *net.TCPConn {
	var config = ConnConfig{"127.0.0.1:6379", "root",1,1,1}
	pool := NewSingleConnPool(config)
	conn, _ := pool.GetConn()
	return conn
}

func getClient() *Client {
	var config = ConnConfig{"127.0.0.1:7002", "123456",1,1,1}
	pool := NewSingleConnPool(config)

	return BuildClient(pool)
}
