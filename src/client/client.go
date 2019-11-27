package client

type Client struct {
	pool *ConnPool
}

func BuildClient(pool *ConnPool) *Client {
	var client Client
	client.pool = pool
	return &client
}

func (client *Client) getConnectPool() *ConnPool {
	return client.pool
}
