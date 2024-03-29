package network

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/shzy2012/common/errors"
)

func Test_PostWithNetError(t *testing.T) {

	client := NewClient()
	client.SetHTTPTimeout(5)
	_, err := client.Request("POST", "http://aaabbbxxxyyyyzzzzzz.info", []byte(""), 1)
	if err != nil {
		if serErr, ok := err.(*errors.ServerError); ok {
			t.Error("[Test_PostWithNetError]=> failed.", serErr)
		} else if _, ok := err.(*errors.ClientError); ok {

		} else {
			t.Error("[Test_PostWithNetError]=> failed.")
		}
	}
}
func Test_PostWithTransport(t *testing.T) {

	client := NewClient()
	transport := &http.Transport{
		Dial: (&net.Dialer{
			Timeout:   3 * time.Second,
			KeepAlive: 3 * time.Second,
		}).Dial,
		TLSHandshakeTimeout:   10 * time.Second,
		ResponseHeaderTimeout: 10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second, //限制client在发送包含 Expect: 100-continue的header到收到继续发送body的response之间的时间等待。
		MaxIdleConns:          30,              //连接池对所有host的最大连接数量，默认无穷大
		MaxConnsPerHost:       30,              //连接池对每个host的最大连接数量。
		IdleConnTimeout:       30 * time.Minute,
		DisableKeepAlives:     false,
	}

	client.SetTransport(transport)
	_, err := client.Request("POST", "http://aaabbbxxxyyyyzzzzzz.info", []byte(""), 1)
	if err != nil {
		if serErr, ok := err.(*errors.ServerError); ok {
			t.Error("[Test_PostWithTransport]=> failed.", serErr)
		} else if _, ok := err.(*errors.ClientError); ok {

		} else {
			t.Error("[Test_PostWithTransport]=> failed.")
		}
	}
}

func Test_Request1(t *testing.T) {
	client := NewClient()
	client.SetHTTPTimeout(3)
	_, err := client.Request("POST", "http://aaabbbxxxyyyyzzzzzz.info", []byte(`{"name":"request"}`), 3)
	if err != nil {
		if serErr, ok := err.(*errors.ServerError); ok {
			t.Error("[Test_PostWithNetError]=> failed.", serErr)
		} else if _, ok := err.(*errors.ClientError); ok {

		} else {
			t.Error("[Test_PostWithNetError]=> failed.")
		}
	}
}

func Test_Request2(t *testing.T) {

	client := NewClient()
	_, err := client.Request("POST", "http://qq.com", []byte(`错误的json`), 3)
	if err != nil {
		if serErr, ok := err.(*errors.ServerError); ok {
			t.Error("[Test_PostWithNetError]=> failed.", serErr)
		} else if _, ok := err.(*errors.ClientError); ok {

		} else {
			t.Error("[Test_PostWithNetError]=> failed.")
		}
	}
}

func Test_HttpGet(t *testing.T) {

	_, err := HTTPGet("http://qq.com")
	if err != nil {
		if serErr, ok := err.(*errors.ServerError); ok {
			t.Error("[Test_PostWithNetError]=> failed.", serErr)
		} else if _, ok := err.(*errors.ClientError); ok {

		} else {
			t.Error("[Test_PostWithNetError]=> failed.")
		}
	}
}

func Test_HTTPost(t *testing.T) {

	_, err := HTTPost("http://qq.com", []byte(`data`))
	if err != nil {
		if serErr, ok := err.(*errors.ServerError); ok {
			t.Error("[Test_PostWithNetError]=> failed.", serErr)
		} else if cErr, ok := err.(*errors.ClientError); ok {
			t.Error("[Test_Http404]=> failed.", cErr)
		} else {
			t.Error("[Test_PostWithNetError]=> failed.")
		}
	}
}

// application/x-www-form-urlencoded
// 数据被编码成以 '&' 分隔的键-值对, 同时以 '=' 分隔键和值.
// https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Methods/POST#%E7%A4%BA%E4%BE%8B
func Test_HTTPostWithXwwwFormUrlencoded(t *testing.T) {

	client := NewClient()
	client.Header["Content-Type"] = XwwwFormUrlencoded
	resp, err := client.Request(POST, "http://xxxx/login.action", []byte("username=aaa&password=bbb&os_cookie=true"), 0)
	if err != nil {
		if serErr, ok := err.(*errors.ServerError); ok {
			t.Error("[Test_PostWithNetError]=> failed.", serErr)
		} else if cErr, ok := err.(*errors.ClientError); ok {
			t.Error("[Test_Http404]=> failed.", cErr)
		} else {
			t.Error("[Test_PostWithNetError]=> failed.")
		}
	}

	fmt.Printf("%+v\n", resp.OriginHTTPResponse.Cookies())
}

func Test_PostForm(t *testing.T) {

	file := "local_or_remote.wav"
	param := fmt.Sprintf(`{
		"dialog":{
			"productId": "914010631"
		},
		"metaObject":{
			"recordId": "123457",
			"priority": 100,
			"speechSeparate": true,
			"path":"%s"
		}
	}`, file)

	formdata := map[string]io.Reader{
		"param": strings.NewReader(param),
		"file":  strings.NewReader(file),
	}

	client := NewClient()

	asrURL := "http://api.talkinggenie.com/smart/sinspection/api/v1/filePathUpload"
	resp, err := client.PostForm(asrURL, formdata)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}

	fmt.Printf("%s\n", resp.ResponseBodyBytes)
}

func Test_Http404(t *testing.T) {

	_, err := HTTP.Request("GET", "https://qq.com.cn/abc", nil, 0)
	if err != nil {
		if serErr, ok := err.(*errors.ServerError); ok {
			t.Error("[Test_Http404]=> failed.", serErr)
		} else if cErr, ok := err.(*errors.ClientError); ok {
			t.Error("[Test_Http404]=> failed.", cErr)
		} else {
			t.Error("[Test_Http404]=> failed.")
		}
	}
}
