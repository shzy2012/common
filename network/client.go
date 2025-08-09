package network

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/cookiejar"
	"os"
	"strings"
	"time"

	"github.com/shzy2012/common/errors"
	"github.com/shzy2012/common/log"
)

// HTTP Client
var HTTP *Client

const (
	ContentType = "Content-Type"
)

const (
	POST    = "POST"
	GET     = "GET"
	PUT     = "PUT"
	PATCH   = "PATCH"
	DELETE  = "DELETE"
	HEAD    = "HEAD"
	OPTIONS = "OPTIONS"
	TRACE   = "TRACE"
	CONNECT = "CONNECT"
)

const (
	XwwwFormUrlencoded = "application/x-www-form-urlencoded"
	FormData           = "application/form-data"
)

func init() {
	HTTP = NewClient()
}

// 请求开始 → 从连接池获取连接 → 发送请求 → 接收响应 → 关闭响应体 → 连接返回池中 → 等待复用或超时关闭
// Client http 客户端
type Client struct {
	HttpClient          *http.Client
	Header              map[string]string
	Version             string
	MaxIdleConns        int
	MaxConnsPerHost     int
	IdleConnTimeout     time.Duration
	MaxIdleConnsPerHost int
	Debug               bool
	Auth                BasicAuth
	Cookies             []*http.Cookie
}

// BasicAuth 基础认证
type BasicAuth struct {
	Username, Password string
}

// NewClient  实例化http client
func NewClient() *Client {

	header := map[string]string{"User-Agent": "go-client 1.0"}
	jar, _ := cookiejar.New(nil)

	// 定义默认连接池配置
	maxIdleConns := 100
	maxConnsPerHost := 100
	maxIdleConnsPerHost := 10
	idleConnTimeout := 90 * time.Second

	// 创建自定义Transport
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		// 实际控制 HTTP 连接池的行为
		MaxIdleConns:          maxIdleConns,
		MaxIdleConnsPerHost:   maxIdleConnsPerHost,
		IdleConnTimeout:       idleConnTimeout,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		DisableCompression:    false,
		DisableKeepAlives:     false,
	}

	client := &Client{
		// 用于存储配置值，方便用户查询和修改
		MaxIdleConns:        maxIdleConns,
		MaxConnsPerHost:     maxConnsPerHost,
		IdleConnTimeout:     idleConnTimeout,
		MaxIdleConnsPerHost: maxIdleConnsPerHost,
		Debug:               false,
		Header:              header,
		HttpClient: &http.Client{
			Timeout:   10 * time.Second, //设置HTTP超时时间
			Transport: transport,
			Jar:       jar, //If Jar is nil, cookies are only sent if they are explicitly
		},
		Auth: BasicAuth{},
	}
	return client
}

/* Request 发起HTTP请求
* action:POST\GET\PUT\PATCH\DELETE\HEAD\OPTIONS\TRACE\CONNECT
* url:请求地址
* input:请求参数
* retry:重试次数,默认0(不重试)
 */
func (c *Client) Request(action, url string, input []byte, retry int) (*HTTPResponse, error) {

	var err error
	response := &HTTPResponse{}

	if c.Debug {
		log.Debugf("[http_request]=>%s to %s \n%s\n", action, url, input)
	}

	// 简化HTTP方法处理
	action = strings.ToUpper(action)

	// 构建HTTP请求
	req, err := http.NewRequest(action, url, bytes.NewReader(input))
	if err != nil {
		errMsg := fmt.Sprintf(errors.NetWorkErrorMessage, err.Error())
		return response, errors.NewClientError(errors.NetWorkErrorCode, errMsg, err)
	}

	// 增加 BasicAuth
	if strings.TrimSpace(c.Auth.Username) != "" {
		req.SetBasicAuth(c.Auth.Username, c.Auth.Password)
	}

	// 设置默认Content-Type
	if req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/json;charset=utf-8")
	}

	// 设置header
	for k, v := range c.Header {
		req.Header.Set(k, v)
		if c.Debug {
			log.Debugf("[http_request_header]=>%s:%s \n", k, v)
		}
	}

	// 设置cookies
	for _, cookie := range c.Cookies {
		req.AddCookie(cookie)
	}

	// 确保retry非负
	if retry < 0 {
		retry = 0
	}

	// 改进重试逻辑
	var resp *http.Response
	for i := 0; i <= retry; i++ {
		resp, err = c.HttpClient.Do(req)
		if err == nil {
			break
		}

		// 如果不是最后一次重试，添加退避时间
		if i < retry {
			backoffTime := time.Duration(i+1) * time.Second
			time.Sleep(backoffTime)
		}
	}

	if err != nil {
		errMsg := fmt.Sprintf(errors.NetWorkErrorMessage, err.Error())
		return response, errors.NewClientError(errors.NetWorkErrorCode, errMsg, err)
	}

	// 确保响应体被正确关闭
	if resp != nil && resp.Body != nil {
		defer func() {
			// 关闭响应体
			closeErr := resp.Body.Close()
			if err == nil && closeErr != nil {
				err = closeErr
			}
		}()
	}

	response.StatusCode = resp.StatusCode
	response.Status = resp.Status
	response.OriginHTTPResponse = resp // 原始的Http Response

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		errMsg := fmt.Sprintf(errors.NetWorkErrorMessage, err.Error())
		return response, errors.NewClientError(errors.NetWorkErrorCode, errMsg, err)
	}

	if c.Debug {
		log.Debugf("[http_response]=>%s \n", bytes)
	}

	response.ResponseBodyBytes = bytes // http 响应体

	// 处理HTTP状态码
	switch response.StatusCode {
	case 200, 201, 202, 203, 204, 205, 206:
		return response, nil
	default:
		response.Message = string(response.ResponseBodyBytes)
		return response, errors.NewServerError(resp.StatusCode, response.Message, err)
	}
}

// SetTransport 设置Transport
func (c *Client) SetTransport(transport http.RoundTripper) {
	c.HttpClient.Transport = transport
}

// SetHTTPTimeout 设置http 超时时间
func (c *Client) SetHTTPTimeout(timeout time.Duration) {
	c.HttpClient.Timeout = timeout
}

// SetConnectionPool 设置连接池参数
func (c *Client) SetConnectionPool(maxIdleConns, maxConnsPerHost, maxIdleConnsPerHost int, idleConnTimeout time.Duration) {
	c.MaxIdleConns = maxIdleConns
	c.MaxConnsPerHost = maxConnsPerHost
	c.MaxIdleConnsPerHost = maxIdleConnsPerHost
	c.IdleConnTimeout = idleConnTimeout

	// 更新Transport配置
	if transport, ok := c.HttpClient.Transport.(*http.Transport); ok {
		transport.MaxIdleConns = maxIdleConns
		transport.MaxConnsPerHost = maxConnsPerHost
		transport.MaxIdleConnsPerHost = maxIdleConnsPerHost
		transport.IdleConnTimeout = idleConnTimeout
	}
}

// SetMaxIdleConns 设置最大空闲连接数
func (c *Client) SetMaxIdleConns(maxIdleConns int) {
	c.MaxIdleConns = maxIdleConns
	if transport, ok := c.HttpClient.Transport.(*http.Transport); ok {
		transport.MaxIdleConns = maxIdleConns
	}
}

// SetMaxConnsPerHost 设置每个主机的最大连接数
func (c *Client) SetMaxConnsPerHost(maxConnsPerHost int) {
	c.MaxConnsPerHost = maxConnsPerHost
	if transport, ok := c.HttpClient.Transport.(*http.Transport); ok {
		transport.MaxConnsPerHost = maxConnsPerHost
	}
}

// SetIdleConnTimeout 设置空闲连接超时时间
func (c *Client) SetIdleConnTimeout(timeout time.Duration) {
	c.IdleConnTimeout = timeout
	if transport, ok := c.HttpClient.Transport.(*http.Transport); ok {
		transport.IdleConnTimeout = timeout
	}
}

// SetMaxIdleConnsPerHost 设置每个主机的最大空闲连接数
func (c *Client) SetMaxIdleConnsPerHost(maxIdleConnsPerHost int) {
	c.MaxIdleConnsPerHost = maxIdleConnsPerHost
	if transport, ok := c.HttpClient.Transport.(*http.Transport); ok {
		transport.MaxIdleConnsPerHost = maxIdleConnsPerHost
	}
}

// GetConnectionPoolStats 获取连接池统计信息
func (c *Client) GetConnectionPoolStats() map[string]interface{} {
	stats := map[string]interface{}{
		"MaxIdleConns":        c.MaxIdleConns,
		"MaxConnsPerHost":     c.MaxConnsPerHost,
		"MaxIdleConnsPerHost": c.MaxIdleConnsPerHost,
		"IdleConnTimeout":     c.IdleConnTimeout,
	}

	if transport, ok := c.HttpClient.Transport.(*http.Transport); ok {
		stats["CurrentMaxIdleConns"] = transport.MaxIdleConns
		stats["CurrentMaxConnsPerHost"] = transport.MaxConnsPerHost
		stats["CurrentMaxIdleConnsPerHost"] = transport.MaxIdleConnsPerHost
		stats["CurrentIdleConnTimeout"] = transport.IdleConnTimeout
	}

	return stats
}

// SetDebug 设置debug
func (c *Client) SetDebug(d bool) {
	c.Debug = d
}

// SetCookie 添加cookie
func (c *Client) SetCookie(cookie *http.Cookie) {
	c.Cookies = append(c.Cookies, cookie)
}

// ClearCookies 清除cookies
func (c *Client) ClearCookies() {
	c.Cookies = c.Cookies[0:0]
}

// Close 释放客户端资源
func (c *Client) Close() error {
	// 关闭 HTTP 客户端的传输层
	if c.HttpClient != nil && c.HttpClient.Transport != nil {
		if transport, ok := c.HttpClient.Transport.(*http.Transport); ok {
			// 关闭所有空闲连接
			transport.CloseIdleConnections()
		}
	}

	// 清空 cookies
	c.ClearCookies()

	// 清空 headers
	c.Header = make(map[string]string)

	return nil
}

// HTTPGet 发起HTTP Get请求
func HTTPGet(url string) ([]byte, error) {
	response, err := HTTP.Request("GET", url, nil, 0)
	return response.ResponseBodyBytes, err
}

// HTTPost 发起HTTP Post请求
func HTTPost(url string, input []byte) ([]byte, error) {
	response, err := HTTP.Request("POST", url, input, 0)
	return response.ResponseBodyBytes, err
}

// HTTPPostForm 发起HTTP Post x-www-form-urlencoded 请求
func HTTPPostForm(url string, values map[string]string) ([]byte, error) {
	response, err := HTTP.PostForm2(url, values)
	return response.ResponseBodyBytes, err
}

//PostForm 发起PostForm请求
//from-data方式
/* example
file, _ := os.Open("file.png")             //读取文件
defer file.Close()

form := map[string]io.Reader{}             //定义form
form["source"] = strings.NewReader("post") //字符串
form["file"] = file                        //文件类型
*/
func (c *Client) PostForm(url string, form map[string]io.Reader) (*HTTPResponse, error) {

	var err error
	response := &HTTPResponse{}

	// Prepare a form that you will submit to that URL.
	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	// 确保 multipart writer 被正确关闭
	defer func() {
		if closeErr := w.Close(); err == nil && closeErr != nil {
			err = closeErr
		}
	}()

	for key, r := range form {
		var fw io.Writer
		if x, ok := r.(io.Closer); ok {
			defer func(closer io.Closer) {
				if closeErr := closer.Close(); err == nil && closeErr != nil {
					err = closeErr
				}
			}(x)
		}
		// file
		if x, ok := r.(*os.File); ok {
			if fw, err = w.CreateFormFile(key, x.Name()); err != nil {
				return response, errors.NewClientError(errors.NetWorkErrorCode, "Failed to create form file", err)
			}
		} else {
			// Add other fields
			if fw, err = w.CreateFormField(key); err != nil {
				return response, errors.NewClientError(errors.NetWorkErrorCode, "Failed to create form field", err)
			}
		}
		if _, err = io.Copy(fw, r); err != nil {
			return response, errors.NewClientError(errors.NetWorkErrorCode, "Failed to copy form data", err)
		}
	}

	req, err := http.NewRequest("POST", url, &b)
	if err != nil {
		errMsg := fmt.Sprintf(errors.NetWorkErrorMessage, err.Error())
		return response, errors.NewClientError(errors.NetWorkErrorCode, errMsg, err)
	}

	//设置header
	for k, v := range c.Header {
		req.Header.Set(k, v)
		if c.Debug {
			log.Debugf("[header]=>%s:%s \n", k, v)
		}
	}

	// Don't forget to set the content type, this will contain the boundary.
	req.Header.Set("Content-Type", w.FormDataContentType())

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		errMsg := fmt.Sprintf(errors.NetWorkErrorMessage, err.Error())
		return response, errors.NewClientError(errors.NetWorkErrorCode, errMsg, err)
	}

	// 确保响应体被正确关闭
	if resp != nil && resp.Body != nil {
		defer func() {
			closeErr := resp.Body.Close()
			if err == nil && closeErr != nil {
				err = closeErr
			}
		}()
	}

	response.StatusCode = resp.StatusCode
	response.Status = resp.Status
	response.OriginHTTPResponse = resp //原始的Http Response

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		errMsg := fmt.Sprintf(errors.NetWorkErrorMessage, err.Error())
		return response, errors.NewClientError(errors.NetWorkErrorCode, errMsg, err)
	}

	if c.Debug {
		log.Debugf("[http resp]=>%s \n", bytes)
	}

	response.ResponseBodyBytes = bytes //http 响应体
	/*
		https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Status/201
		200 OK
		201 Created
		202 Accepted
	*/
	switch response.StatusCode {
	case 200, 201, 202, 203, 204, 205, 206:
		return response, nil
	default:
		response.Message = string(response.ResponseBodyBytes)
		return response, errors.NewServerError(resp.StatusCode, response.Message, err)
	}
}

// PostForm2
// x-www-form-urlencoded 方式
func (c *Client) PostForm2(url string, values map[string]string) (*HTTPResponse, error) {
	var err error
	response := &HTTPResponse{}

	// 构建 x-www-form-urlencoded 数据
	data := make([]string, 0, len(values))
	for key, value := range values {
		// URL编码参数
		encodedKey := strings.ReplaceAll(key, " ", "+")
		encodedValue := strings.ReplaceAll(value, " ", "+")
		data = append(data, fmt.Sprintf("%s=%s", encodedKey, encodedValue))
	}

	// 将参数用 & 连接
	encodedData := strings.Join(data, "&")

	if c.Debug {
		log.Debugf("[form-data]=>%s \n", encodedData)
	}

	// 创建请求
	req, err := http.NewRequest("POST", url, strings.NewReader(encodedData))
	if err != nil {
		errMsg := fmt.Sprintf(errors.NetWorkErrorMessage, err.Error())
		return response, errors.NewClientError(errors.NetWorkErrorCode, errMsg, err)
	}

	// 设置Content-Type为x-www-form-urlencoded
	req.Header.Set("Content-Type", XwwwFormUrlencoded)

	// 设置header
	for k, v := range c.Header {
		req.Header.Set(k, v)
		if c.Debug {
			log.Debugf("[header]=>%s:%s \n", k, v)
		}
	}

	// 增加 BasicAuth
	if strings.TrimSpace(c.Auth.Username) != "" {
		req.SetBasicAuth(c.Auth.Username, c.Auth.Password)
	}

	// 设置cookies
	for _, cookie := range c.Cookies {
		req.AddCookie(cookie)
	}

	// 发送请求
	resp, err := c.HttpClient.Do(req)
	if err != nil {
		errMsg := fmt.Sprintf(errors.NetWorkErrorMessage, err.Error())
		return response, errors.NewClientError(errors.NetWorkErrorCode, errMsg, err)
	}

	// 确保响应体被正确关闭
	if resp != nil && resp.Body != nil {
		defer func() {
			closeErr := resp.Body.Close()
			if err == nil && closeErr != nil {
				err = closeErr
			}
		}()
	}

	response.StatusCode = resp.StatusCode
	response.Status = resp.Status
	response.OriginHTTPResponse = resp // 原始的Http Response

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		errMsg := fmt.Sprintf(errors.NetWorkErrorMessage, err.Error())
		return response, errors.NewClientError(errors.NetWorkErrorCode, errMsg, err)
	}

	if c.Debug {
		log.Debugf("[http resp]=>%s \n", bytes)
	}

	response.ResponseBodyBytes = bytes // http 响应体

	// 处理HTTP状态码
	switch response.StatusCode {
	case 200, 201, 202, 203, 204, 205, 206:
		return response, nil
	default:
		response.Message = string(response.ResponseBodyBytes)
		return response, errors.NewServerError(resp.StatusCode, response.Message, err)
	}
}
