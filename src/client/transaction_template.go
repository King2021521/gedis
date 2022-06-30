//事务处理器
//@author zxm
//@date 2020-06-02
package client

import (
	"fmt"
	"gedis/src/client/handler"
	"gedis/src/protocol"
)

//开启事务
func (client *Client) Multi() (interface{}, error) {
	pool := client.getConnectPool()
	conn, err := pool.GetConn()
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	result := SendCommand(conn, protocol.MULTI)
	return handler.HandleReply(result)
}

//执行事务
func (client *Client) Exec() (interface{}, error) {
	pool := client.getConnectPool()
	conn, err := pool.GetConn()
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	result := SendCommand(conn, protocol.EXEC)
	return handler.HandleTransactionReply(result)
}

//终止事务
func (client *Client) Discard() (interface{}, error) {
	pool := client.getConnectPool()
	conn, err := pool.GetConn()
	if err != nil {
		return nil, fmt.Errorf("get conn fail")
	}
	defer pool.PutConn(conn)
	result := SendCommand(conn, protocol.DISCARD)
	return handler.HandleReply(result)
}
