package http

type Request struct {
	Method string
	Header Header
	URL    string
	Body   []byte
	Host   string
}

func NewRequest(method, url string, body []byte) *Request {
	return &Request{
		Method: method,
		URL:    url,
		Body:   body,
	}
}
