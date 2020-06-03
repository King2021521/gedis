# gedis
golang redis client  
##demo  
```
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
//集群版
func testCluster(){
	var node7000 = Node{"127.0.0.1:7000", "123456", 10}
	var node7001 = Node{"127.0.0.1:7001", "123456", 10}
	var node7002 = Node{"127.0.0.1:7002", "123456", 10}
	var node7003 = Node{"127.0.0.1:7003", "123456", 10}
	var node7004 = Node{"127.0.0.1:7004", "123456", 10}
	var node7005 = Node{"127.0.0.1:7005", "123456", 10}

	nodes := []*Node{&node7000, &node7001, &node7002, &node7003, &node7004, &node7005}
	var clusterConfig = ClusterConfig{nodes,10}
	cluster := NewCluster(clusterConfig)
	value,err:=cluster.Get("name")
	fmt.Println(value, err)
}

func getConn() *net.TCPConn {
	var config = ConnConfig{"127.0.0.1:6379", "root"}
	pool := NewSingleConnPool(1, config)
	conn, _ := GetConn(pool)
	return conn
}
//单机版
func getClient() *Client {
	var config = ConnConfig{"127.0.0.1:7002", "123456"}
	pool := NewSingleConnPool(1, config)

	return BuildClient(pool)
}
```  

1.0 特性：  
基于原生golang开发  
连接池管理  
keepalive  
redisTemplate提供多种命令支持  
2.0 特性   
cluster支持  
loadbalance支持  
heartBeat支持  
连接池的监控及动态扩容  
更多内容持续更新中  

author: tony  
博客地址：https://blog.csdn.net/u012737673  
热衷开源，欢迎大家加入
