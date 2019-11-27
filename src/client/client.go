package client

/**
 * gedis 客户端
 */
type Client struct {
	pool *ConnPool
}

/**
 * 初始化客户端
 * @param pool 连接池
 */
func BuildClient(pool *ConnPool) *Client {
	var client Client
	client.pool = pool
	return &client
}

/**
 * 获取连接池
 */
func (client *Client) getConnectPool() *ConnPool {
	return client.pool
}
