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
		} else if _, ok := err.(*errors.ClientError); ok {

		} else {
			t.Error("[Test_PostWithNetError]=> failed.")
		}
	}
}

func Test_PostForm(t *testing.T) {
	// t := time.Now()
	// timestamp := fmt.Sprintf("%v", t.UnixNano()/1000000)
	file := "xxx.wav"
	param := `{
		"param":"data param",
	}`

	formdata := map[string]io.Reader{
		"param": strings.NewReader(param),
		"file":  strings.NewReader(file),
	}

	client := NewClient()
	asrURL := "http://xxx.com"
	resp, err := client.PostForm(asrURL, formdata)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}

	fmt.Printf("%+v\n", resp)
}
