package http

type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

type Response struct {
	StatusCode int
	Body       []byte
}

func (c *Client) Do(req *Request) (*Response, error) {
	return nil, nil
}
