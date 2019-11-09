package network

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/shzy2012/common/errors"
	"github.com/shzy2012/common/log"
)

//Client http 客户端
type Client struct {
	httpClient      *http.Client
	ContentType     string
	Version         string
	MaxIdleConns    int
	MaxConnsPerHost int
	Debug           bool
}

//NewClient  实例化http client
func NewClient() *Client {
	client := &Client{
		MaxIdleConns:    30,
		MaxConnsPerHost: 30,
		Debug:           false,
		httpClient: &http.Client{
			Timeout: 8 * time.Second, //设置HTTP超时时间
		},
	}
	return client
}

//Request 发起HTTP请求
func (c *Client) Request(action, url string, input []byte, retry int) (*HTTPResponse, error) {

	var err error
	response := &HTTPResponse{}

	if c.Debug {
		log.Debugf("[http req]=>%s to %s \n%s\n", action, url, input)
	}

	//处理 http action
	if strings.ToUpper(action) == "POST" {
		action = "POST"
	} else if strings.ToUpper(action) == "GET" {
		action = "GET"
	} else {
		errorMsg := fmt.Sprintf(errors.UnsupportedTypeErrorMessage, action, "POST,GET")
		return nil, errors.NewClientError(errors.UnsupportedTypeErrorCode, errorMsg, err)
	}

	req, err := http.NewRequest(action, url, bytes.NewReader(input))
	if err != nil {
		errMsg := fmt.Sprintf(errors.NetWorkErrorMessage, err.Error())
		return response, errors.NewClientError(errors.NetWorkErrorCode, errMsg, err)
	}

	//设置Content-Type
	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	if c.ContentType != "" {
		req.Header.Set("Content-Type", c.ContentType)
	}

	//默认 retry
	if retry < 0 {
		retry = 0
	}

	var resp *http.Response
	for i := 0; i <= retry; i++ {
		resp, err = c.httpClient.Do(req)
		if err == nil {
			break
		}

		time.Sleep(time.Second * 1)
	}

	if err != nil {
		errMsg := fmt.Sprintf(errors.NetWorkErrorMessage, err.Error())
		return response, errors.NewClientError(errors.NetWorkErrorCode, errMsg, err)
	}
	if resp != nil {
		defer func() {
			err = resp.Body.Close()
		}()
	}

	response.StatusCode = resp.StatusCode
	response.Status = resp.Status
	response.OriginHTTPResponse = resp //原始的Http Response
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		errMsg := fmt.Sprintf(errors.NetWorkErrorMessage, err.Error())
		return response, errors.NewClientError(errors.NetWorkErrorCode, errMsg, err)
	}

	response.ResponseBodyBytes = bytes //http 响应体
	//如果StatusCode不等于200,则错误
	if response.StatusCode != 200 {
		response.Message = string(response.ResponseBodyBytes)
		return response, errors.NewServerError(resp.StatusCode, response.Message, err)
	}

	if c.Debug {
		log.Debugf("[http resp]=>%s \n", bytes)
	}

	return response, nil
}

//SetTransport 设置Transport
func (c *Client) SetTransport(transport http.RoundTripper) {
	c.httpClient.Transport = transport
}

//SetHTTPTimeout 设置http 超时时间
func (c *Client) SetHTTPTimeout(timeout time.Duration) {
	c.httpClient.Timeout = timeout
}

//HTTPGet 发起HTTP Get请求
func (c *Client) HTTPGet(url string) (*HTTPResponse, error) {
	return c.Request("GET", url, nil, 0)
}

//HTTPost 发起HTTP Post请求
func (c *Client) HTTPost(url string, input []byte) (*HTTPResponse, error) {
	return c.Request("POST", url, input, 0)
}
