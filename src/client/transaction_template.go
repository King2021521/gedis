//事务处理器
//@author zxm
//@date 2020-06-02
package client

import (
	"client/handler"
	"fmt"
	"protocol"
	"strings"
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
	elements := strings.Split(result, protocol.CRLF)
	var values []string
	//解析事务批量执行的返回结果
	for i := 1; i < len(elements); {
		if elements[i] == protocol.BLANK {
			i++
			continue
		}

		if strings.HasPrefix(elements[i], protocol.DOLLARBYTE) {
			values = append(values, elements[i+1])
			i += 2
		} else {
			values = append(values, elements[i])
			i++
		}
	}
	return values, nil
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
