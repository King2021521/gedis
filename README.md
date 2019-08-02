# gedis
golang redis client  
##demo  
```
func main(){
    var config = tcp.ConnConfig{"127.0.0.1:6379","123456"}
	pool,err:=tcp.NewConnPool(1,config)
	if err!=nil{
		fmt.Println(err)
	}

	conn,_:=tcp.GetConn(pool)
	fmt.Println(conn.RemoteAddr())
	sendResult := template.Set("name", "james", conn)
	fmt.Println("send result:" + sendResult)
	result := template.Get("name", conn)
	fmt.Println("get result:" + result)

	pool.PutConn(conn)

	conn1,_:=tcp.GetConn(pool)
	fmt.Println(conn1.RemoteAddr())
	sendResult1 := template.Set("name", "james", conn)
	fmt.Println("send result:" + sendResult1)
	result1 := template.Get("name", conn)
	fmt.Println("get result:" + result1)
}
```  

特性：  
基于原生golang开发  
连接池管理  
keepalive支持  
redisTemplate提供多种命令支持  
更多内容持续更新中  

author: tony  
热衷开源，欢迎大家加入
