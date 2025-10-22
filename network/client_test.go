package network

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/shzy2012/common/errors"
)

// 测试服务器响应函数类型
type testHandlerFunc func(w http.ResponseWriter, r *http.Request)

// 创建测试服务器
func createTestServer(handler testHandlerFunc) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(handler))
}

// 测试 NewClient 函数
func TestNewClient(t *testing.T) {
	client := NewClient()

	if client == nil {
		t.Fatal("NewClient() returned nil")
	}

	// 检查默认配置
	if client.HttpClient == nil {
		t.Error("HttpClient should not be nil")
	}

	if client.Header == nil {
		t.Error("Header should not be nil")
	}

	if client.Header["User-Agent"] != "go-client 1.0" {
		t.Errorf("Expected User-Agent 'go-client 1.0', got '%s'", client.Header["User-Agent"])
	}

	if client.maxIdleConns != 100 {
		t.Errorf("Expected MaxIdleConns 100, got %d", client.maxIdleConns)
	}

	if client.maxConnsPerHost != 100 {
		t.Errorf("Expected MaxConnsPerHost 100, got %d", client.maxConnsPerHost)
	}

	if client.maxIdleConnsPerHost != 10 {
		t.Errorf("Expected MaxIdleConnsPerHost 10, got %d", client.maxIdleConnsPerHost)
	}

	if client.idleConnTimeout != 90*time.Second {
		t.Errorf("Expected IdleConnTimeout 90s, got %v", client.idleConnTimeout)
	}

	if client.Debug {
		t.Error("Debug should be false by default")
	}
}

// 测试 GET 请求
func TestClient_Request_GET(t *testing.T) {
	// 创建测试服务器
	server := createTestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "success"}`))
	})
	defer server.Close()

	client := NewClient()
	response, err := client.Request("GET", server.URL, nil, 0)

	if err != nil {
		t.Fatalf("Request failed: %v", err)
	}

	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", response.StatusCode)
	}

	expectedBody := `{"message": "success"}`
	if string(response.ResponseBodyBytes) != expectedBody {
		t.Errorf("Expected body '%s', got '%s'", expectedBody, string(response.ResponseBodyBytes))
	}
}

// 测试 POST 请求
func TestClient_Request_POST(t *testing.T) {
	// 创建测试服务器
	server := createTestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Expected POST request, got %s", r.Method)
		}

		// 读取请求体
		body, _ := io.ReadAll(r.Body)
		if string(body) != `{"test": "data"}` {
			t.Errorf("Expected body '{\"test\": \"data\"}', got '%s'", string(body))
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "success"}`))
	})
	defer server.Close()

	client := NewClient()
	requestBody := []byte(`{"test": "data"}`)
	response, err := client.Request("POST", server.URL, requestBody, 0)

	if err != nil {
		t.Fatalf("Request failed: %v", err)
	}

	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", response.StatusCode)
	}
}

// 测试重试机制
func TestClient_Request_Retry(t *testing.T) {
	attempts := 0
	// 创建测试服务器，前两次请求失败，第三次成功
	server := createTestServer(func(w http.ResponseWriter, r *http.Request) {
		attempts++
		if attempts < 3 {
			// 模拟网络错误 - 关闭连接来模拟网络错误
			hj, ok := w.(http.Hijacker)
			if ok {
				conn, _, _ := hj.Hijack()
				conn.Close()
			}
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "success"}`))
	})
	defer server.Close()

	client := NewClient()
	client.SetHTTPTimeout(5 * time.Second) // 设置较短的超时时间

	response, err := client.Request("GET", server.URL, nil, 2) // 重试2次

	if err != nil {
		t.Fatalf("Request failed after retries: %v", err)
	}

	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", response.StatusCode)
	}

	if attempts != 3 {
		t.Errorf("Expected 3 attempts, got %d", attempts)
	}
}

// 测试 BasicAuth
func TestClient_Request_BasicAuth(t *testing.T) {
	// 创建测试服务器
	server := createTestServer(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			t.Error("BasicAuth not found in request")
		}
		if username != "testuser" || password != "testpass" {
			t.Errorf("Expected username 'testuser' and password 'testpass', got '%s' and '%s'", username, password)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "success"}`))
	})
	defer server.Close()

	client := NewClient()
	client.Auth = BasicAuth{Username: "testuser", Password: "testpass"}

	response, err := client.Request("GET", server.URL, nil, 0)

	if err != nil {
		t.Fatalf("Request failed: %v", err)
	}

	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", response.StatusCode)
	}
}

// 测试自定义 Header
func TestClient_Request_CustomHeader(t *testing.T) {
	// 创建测试服务器
	server := createTestServer(func(w http.ResponseWriter, r *http.Request) {
		customHeader := r.Header.Get("X-Custom-Header")
		if customHeader != "custom-value" {
			t.Errorf("Expected X-Custom-Header 'custom-value', got '%s'", customHeader)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "success"}`))
	})
	defer server.Close()

	client := NewClient()
	client.Header["X-Custom-Header"] = "custom-value"

	response, err := client.Request("GET", server.URL, nil, 0)

	if err != nil {
		t.Fatalf("Request failed: %v", err)
	}

	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", response.StatusCode)
	}
}

// 测试 Cookie
func TestClient_Request_Cookie(t *testing.T) {
	// 创建测试服务器
	server := createTestServer(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session")
		if err != nil {
			t.Errorf("Cookie 'session' not found: %v", err)
		}
		if cookie.Value != "abc123" {
			t.Errorf("Expected cookie value 'abc123', got '%s'", cookie.Value)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "success"}`))
	})
	defer server.Close()

	client := NewClient()
	cookie := &http.Cookie{Name: "session", Value: "abc123"}
	client.SetCookie(cookie)

	response, err := client.Request("GET", server.URL, nil, 0)

	if err != nil {
		t.Fatalf("Request failed: %v", err)
	}

	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", response.StatusCode)
	}
}

// 测试错误状态码
func TestClient_Request_ErrorStatusCode(t *testing.T) {
	// 创建测试服务器
	server := createTestServer(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not Found", http.StatusNotFound)
	})
	defer server.Close()

	client := NewClient()
	response, err := client.Request("GET", server.URL, nil, 0)

	if err == nil {
		t.Error("Expected error for 404 status code")
	}

	if response.StatusCode != http.StatusNotFound {
		t.Errorf("Expected status 404, got %d", response.StatusCode)
	}

	// 检查错误类型
	if _, ok := err.(*errors.ServerError); !ok {
		t.Errorf("Expected ServerError, got %T", err)
	}
}

// 测试 SetHTTPTimeout
func TestClient_SetHTTPTimeout(t *testing.T) {
	client := NewClient()
	timeout := 30 * time.Second
	client.SetHTTPTimeout(timeout)

	if client.HttpClient.Timeout != timeout {
		t.Errorf("Expected timeout %v, got %v", timeout, client.HttpClient.Timeout)
	}
}

// 测试 SetConnectionPool
func TestClient_SetConnectionPool(t *testing.T) {
	client := NewClient()

	maxIdleConns := 50
	maxConnsPerHost := 50
	maxIdleConnsPerHost := 5
	idleConnTimeout := 60 * time.Second

	client.SetConnectionPool(maxIdleConns, maxConnsPerHost, maxIdleConnsPerHost, idleConnTimeout)

	if client.maxIdleConns != maxIdleConns {
		t.Errorf("Expected MaxIdleConns %d, got %d", maxIdleConns, client.maxIdleConns)
	}

	if client.maxConnsPerHost != maxConnsPerHost {
		t.Errorf("Expected MaxConnsPerHost %d, got %d", maxConnsPerHost, client.maxConnsPerHost)
	}

	if client.maxIdleConnsPerHost != maxIdleConnsPerHost {
		t.Errorf("Expected MaxIdleConnsPerHost %d, got %d", maxIdleConnsPerHost, client.maxIdleConnsPerHost)
	}

	if client.idleConnTimeout != idleConnTimeout {
		t.Errorf("Expected IdleConnTimeout %v, got %v", idleConnTimeout, client.idleConnTimeout)
	}
}

// 测试 SetDebug
func TestClient_SetDebug(t *testing.T) {
	client := NewClient()

	if client.Debug {
		t.Error("Debug should be false by default")
	}

	client.SetDebug(true)
	if !client.Debug {
		t.Error("Debug should be true after SetDebug(true)")
	}

	client.SetDebug(false)
	if client.Debug {
		t.Error("Debug should be false after SetDebug(false)")
	}
}

// 测试 Cookie 操作
func TestClient_CookieOperations(t *testing.T) {
	client := NewClient()

	// 测试 SetCookie
	cookie1 := &http.Cookie{Name: "session", Value: "abc123"}
	cookie2 := &http.Cookie{Name: "user", Value: "john"}

	client.SetCookie(cookie1)
	client.SetCookie(cookie2)

	if len(client.Cookies) != 2 {
		t.Errorf("Expected 2 cookies, got %d", len(client.Cookies))
	}

	// 测试 ClearCookies
	client.ClearCookies()
	if len(client.Cookies) != 0 {
		t.Errorf("Expected 0 cookies after clear, got %d", len(client.Cookies))
	}
}

// 测试 Close
func TestClient_Close(t *testing.T) {
	client := NewClient()

	// 添加一些数据
	client.Header["Test-Header"] = "test-value"
	client.SetCookie(&http.Cookie{Name: "test", Value: "value"})

	err := client.Close()
	if err != nil {
		t.Errorf("Close() returned error: %v", err)
	}

	// 检查是否清空了数据
	if len(client.Header) != 0 {
		t.Error("Header should be empty after Close()")
	}

	if len(client.Cookies) != 0 {
		t.Error("Cookies should be empty after Close()")
	}
}
