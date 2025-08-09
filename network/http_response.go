package network

import (
	"net/http"
)

// HTTPResponse http响应
type HTTPResponse struct {
	StatusCode         int
	Status             string
	Message            string
	ResponseBodyBytes  []byte
	OriginHTTPResponse *http.Response
}

// ToString 将http body转化为字符串
func (r *HTTPResponse) ToString() string {
	return string(r.ResponseBodyBytes)
}
