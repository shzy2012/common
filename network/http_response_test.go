package network

import (
	"net/http"
	"testing"
)

func TestHTTPResponse_ToString(t *testing.T) {
	tests := []struct {
		name     string
		response *HTTPResponse
		want     string
	}{
		{
			name: "normal string",
			response: &HTTPResponse{
				ResponseBodyBytes: []byte("hello world"),
			},
			want: "hello world",
		},
		{
			name: "empty bytes",
			response: &HTTPResponse{
				ResponseBodyBytes: []byte{},
			},
			want: "",
		},
		{
			name: "nil bytes",
			response: &HTTPResponse{
				ResponseBodyBytes: nil,
			},
			want: "",
		},
		{
			name: "chinese characters",
			response: &HTTPResponse{
				ResponseBodyBytes: []byte("你好世界"),
			},
			want: "你好世界",
		},
		{
			name: "special characters",
			response: &HTTPResponse{
				ResponseBodyBytes: []byte("!@#$%^&*()"),
			},
			want: "!@#$%^&*()",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.response.ToString(); got != tt.want {
				t.Errorf("HTTPResponse.ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 测试空响应体
func TestHTTPResponse_ToString_Empty(t *testing.T) {
	response := &HTTPResponse{
		ResponseBodyBytes: []byte{},
	}

	result := response.ToString()
	expected := ""

	if result != expected {
		t.Errorf("Expected empty string, got '%s'", result)
	}
}

func TestHTTPResponse_Fields(t *testing.T) {
	// 创建一个模拟的 http.Response
	mockHTTPResp := &http.Response{
		StatusCode: 200,
		Status:     "200 OK",
	}

	response := &HTTPResponse{
		StatusCode:         200,
		Status:             "200 OK",
		Message:            "Success",
		ResponseBodyBytes:  []byte("test response"),
		OriginHTTPResponse: mockHTTPResp,
	}

	// 测试所有字段是否正确设置
	if response.StatusCode != 200 {
		t.Errorf("Expected StatusCode 200, got %d", response.StatusCode)
	}

	if response.Status != "200 OK" {
		t.Errorf("Expected Status '200 OK', got '%s'", response.Status)
	}

	if response.Message != "Success" {
		t.Errorf("Expected Message 'Success', got '%s'", response.Message)
	}

	if string(response.ResponseBodyBytes) != "test response" {
		t.Errorf("Expected ResponseBodyBytes 'test response', got '%s'", string(response.ResponseBodyBytes))
	}

	if response.OriginHTTPResponse != mockHTTPResp {
		t.Errorf("Expected OriginHTTPResponse to match mockHTTPResp")
	}
}

func TestHTTPResponse_EmptyResponse(t *testing.T) {
	response := &HTTPResponse{}

	// 测试空响应的默认值
	if response.StatusCode != 0 {
		t.Errorf("Expected default StatusCode 0, got %d", response.StatusCode)
	}

	if response.Status != "" {
		t.Errorf("Expected default Status empty string, got '%s'", response.Status)
	}

	if response.Message != "" {
		t.Errorf("Expected default Message empty string, got '%s'", response.Message)
	}

	if response.ResponseBodyBytes != nil {
		t.Errorf("Expected default ResponseBodyBytes nil, got %v", response.ResponseBodyBytes)
	}

	if response.OriginHTTPResponse != nil {
		t.Errorf("Expected default OriginHTTPResponse nil, got %v", response.OriginHTTPResponse)
	}

	// 测试空响应的 ToString 方法
	if got := response.ToString(); got != "" {
		t.Errorf("Expected empty string from ToString(), got '%s'", got)
	}
}

func TestHTTPResponse_ToStringWithLargeData(t *testing.T) {
	// 测试大数据量的情况
	largeData := make([]byte, 10000)
	for i := range largeData {
		largeData[i] = byte(i % 256)
	}

	response := &HTTPResponse{
		ResponseBodyBytes: largeData,
	}

	result := response.ToString()
	if len(result) != 10000 {
		t.Errorf("Expected result length 10000, got %d", len(result))
	}
}

func BenchmarkHTTPResponse_ToString(b *testing.B) {
	response := &HTTPResponse{
		ResponseBodyBytes: []byte("benchmark test data"),
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = response.ToString()
	}
}
