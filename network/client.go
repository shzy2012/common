package network

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/cookiejar"
	"os"
	"strings"
	"time"

	"github.com/shzy2012/common/errors"
	"github.com/shzy2012/common/log"
)

//HTTP Client
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

//Client http 客户端
type Client struct {
	HttpClient      *http.Client
	Header          map[string]string
	Version         string
	MaxIdleConns    int
	MaxConnsPerHost int
	Debug           bool
	Auth            BasicAuth
	Cookies         []*http.Cookie
}

//BasicAuth 基础认证
type BasicAuth struct {
	Username, Password string
}

//NewClient  实例化http client
func NewClient() *Client {

	header := map[string]string{"User-Agent": "Go-Client 1.0"}
	jar, _ := cookiejar.New(nil)
	client := &Client{
		MaxIdleConns:    100,
		MaxConnsPerHost: 100,
		Debug:           false,
		Header:          header, //make(map[string]string),
		HttpClient: &http.Client{
			Timeout: 6 * time.Second, //设置HTTP超时时间
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // disable security checks globally for all requests of the default client
			},
			Jar: jar, //If Jar is nil, cookies are only sent if they are explicitly
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
		log.Debugf("[http req]=>%s to %s \n%s\n", action, url, input)
	}

	//处理 http action
	if strings.ToUpper(action) == "POST" {
		action = "POST"
	} else if strings.ToUpper(action) == "GET" {
		action = "GET"
	} else {
		//默认支持其他method
		action = strings.ToUpper(action)
	}

	//构建HTTP请求
	req, err := http.NewRequest(action, url, bytes.NewReader(input))
	if err != nil {
		errMsg := fmt.Sprintf(errors.NetWorkErrorMessage, err.Error())
		return response, errors.NewClientError(errors.NetWorkErrorCode, errMsg, err)
	}

	//增加 BasicAuth
	if strings.TrimSpace(c.Auth.Username) != "" {
		req.SetBasicAuth(c.Auth.Username, c.Auth.Password)
	}

	//设置默认
	if req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/json;charset=utf-8")
	}

	//设置header
	for k, v := range c.Header {
		req.Header.Set(k, v)
		if c.Debug {
			log.Debugf("[header]=>%s:%s \n", k, v)
		}
	}

	//设置cookies
	for _, c := range c.Cookies {
		req.AddCookie(c)
	}

	//默认 retry
	if retry < 0 {
		retry = 0
	}

	//发起HTTP请求
	var resp *http.Response
	for i := 0; i <= retry; i++ {
		resp, err = c.HttpClient.Do(req)
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
	case 200, 201, 202:
		return response, nil
	default:
		response.Message = string(response.ResponseBodyBytes)
		return response, errors.NewServerError(resp.StatusCode, response.Message, err)
	}
}

//SetTransport 设置Transport
func (c *Client) SetTransport(transport http.RoundTripper) {
	c.HttpClient.Transport = transport
}

//SetHTTPTimeout 设置http 超时时间
func (c *Client) SetHTTPTimeout(timeout time.Duration) {
	c.HttpClient.Timeout = timeout
}

//SetDebug 设置debug
func (c *Client) SetDebug(d bool) {
	c.Debug = true
}

//SetCookie 添加cookie
func (c *Client) SetCookie(cookie *http.Cookie) {
	c.Cookies = append(c.Cookies, cookie)
}

//ClearCookies 清除cookies
func (c *Client) ClearCookies() {
	c.Cookies = c.Cookies[0:0]
}

//HTTPGet 发起HTTP Get请求
func HTTPGet(url string) ([]byte, error) {
	response, err := HTTP.Request("GET", url, nil, 0)
	return response.ResponseBodyBytes, err
}

//HTTPost 发起HTTP Post请求
func HTTPost(url string, input []byte) ([]byte, error) {
	response, err := HTTP.Request("POST", url, input, 0)
	return response.ResponseBodyBytes, err
}

//PostForm 发起PostForm请求
//from-data方式
func (c *Client) PostForm(url string, values map[string]io.Reader) (*HTTPResponse, error) {

	var err error
	response := &HTTPResponse{}
	// Prepare a form that you will submit to that URL.
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	for key, r := range values {
		var fw io.Writer
		if x, ok := r.(io.Closer); ok {
			defer x.Close()
		}
		// file
		if x, ok := r.(*os.File); ok {
			if fw, err = w.CreateFormFile(key, x.Name()); err != nil {
				return response, nil
			}
		} else {
			// Add other fields
			if fw, err = w.CreateFormField(key); err != nil {
				return response, nil
			}
		}
		if _, err = io.Copy(fw, r); err != nil {
			return response, err
		}

	}
	// Don't forget to close the multipart writer.
	// If you don't close it, your request will be missing the terminating boundary.
	w.Close()

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
	case 200, 201, 202:
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

	// x-www-form-urlencoded
	// if condition {

	// }
	return response, err
}
