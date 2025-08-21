package tools

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"
	"sync/atomic"
	"time"
)

const Long = "2006-01-02 15:04:05"
const Short = "2006-01-02"
const Time = "15:04:05"

var counter int64
var iset bool

// AutoIncrementing 自动自增(base设置自增起始点)
func AutoIncrementing(base int64) int64 {
	if !iset {
		counter = base
		iset = true
	}
	return atomic.AddInt64(&counter, 1)
}

// Sum Get sum of numbers
func Sum(numbers []int) int {
	sum := 0
	for _, v := range numbers {
		sum = sum + v
	}

	return sum
}

// Max Get max value from numbers
func Max(numbers []int) int {
	max := math.MinInt64
	for _, n := range numbers {
		if n > max {
			max = n
		}
	}

	return max
}

// Min Get min value from numbers
func Min(numbers []int) int {
	min := math.MaxInt64
	for _, n := range numbers {
		if n < min {
			min = n
		}
	}

	return min
}

// F64SliceToStringSlice convert float64 slice to string slice
func F64SliceToStringSlice(list []float64) []string {
	tmp := make([]string, 0)
	for _, ele := range list {
		tmp = append(tmp, fmt.Sprintf("%v", ele))
	}
	return tmp
}

// IsNum 是否是数字
func IsNum(b byte) bool {
	if b >= 48 && b <= 57 {
		return true
	}
	return false
}

// IsChar 是否是字母 A-Z and a-z
func IsChar(b byte) bool {
	if (b >= 65 && b <= 90) || (b >= 97 && b <= 122) {
		return true
	}
	return false
}

func ToBytes(t interface{}) []byte {
	bytes, err := json.Marshal(t)
	if err != nil {
		log.Printf("[ToBytes]:%s\n", err.Error())
	}
	return bytes
}

// https://juejin.cn/post/6844903648045039624
func IsChinesePhone(phone string) bool {
	reg1 := regexp.MustCompile(`^1(?:3[0-9]|4[5-9]|5[0-9]|6[12456]|7[0-8]|8[0-9]|9[0-9])[0-9]{8}$`)
	if reg1 == nil {
		return false
	}
	//根据规则提取关键信息
	if len(reg1.FindAllStringSubmatch(phone, 1)) > 0 {
		return true
	}

	return false
}

// 手机手机号 136****1389
func HideMiddlePhone(phoneNumber string) string {
	if len(phoneNumber) != 11 {
		return "Invalid phone number"
	}
	// 隐藏中间6位
	return phoneNumber[:3] + "****" + phoneNumber[7:]
}

// Convert string to int
func GetInt(v string) int {
	if strings.TrimSpace(v) == "" {
		return 0
	}

	i, err := strconv.Atoi(v)
	if err != nil {
		fmt.Printf("%s", err.Error())
		return 0
	}

	return i
}

// Convert string to uint
func GetuInt(v string) uint {
	if strings.TrimSpace(v) == "" {
		return 0
	}

	i, err := strconv.Atoi(v)
	if err != nil {
		fmt.Printf("%s", err.Error())
		return 0
	}

	return uint(i)
}

// Convert string to int
func GetInt64(v string) int64 {
	if strings.TrimSpace(v) == "" {
		return 0
	}

	i, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		fmt.Printf("%s", err.Error())
		return 0
	}

	return i
}

// Convert string to float64
func GetFloat64(v string) float64 {
	if strings.TrimSpace(v) == "" {
		return float64(0)
	}

	i, err := strconv.ParseFloat(v, 64)
	if err != nil {
		fmt.Printf("%s", err.Error())
		return float64(0)
	}

	return i
}

// 将普通时间格式转化为 RFC3339
func GetRFC3339(v string) string {
	if v == "" {
		return ""
	}
	t, err := time.Parse("2006-01-02 15:04:05", v)
	if err != nil {
		fmt.Printf("[GetRFC3339]=>%s\n", err.Error())
		return ""
	}

	return t.Format("2006-01-02T15:04:05Z07:00")
}

// return local date
func GetDate(v string) (time.Time, error) {
	if v == "" {
		return time.Time{}, fmt.Errorf("empty date string")
	}

	// Try common date formats
	formats := []string{
		"2006-01-02",
		"2006/01/02",
		"02-01-2006",
		"02/01/2006",
		"2006.01.02",
		"02.01.2006",
	}

	for _, format := range formats {
		t, err := time.Parse(format, v)
		if err == nil {
			// Set time to 00:00:00
			return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local), nil
		}
	}

	return time.Time{}, fmt.Errorf("invalid date format: %s", v)
}

// return local datetime
func GetDatetime(v string) (time.Time, error) {
	if v == "" {
		return time.Time{}, fmt.Errorf("empty date string")
	}

	// Try parsing with time component
	formats := []string{
		"2006-01-02 15:04:05",
		"2006/01/02 15:04:05",
		"02-01-2006 15:04:05",
		"02/01/2006 15:04:05",
	}

	for _, format := range formats {
		t, err := time.Parse(format, v)
		if err == nil {
			return t.Local(), nil
		}
	}

	return time.Time{}, fmt.Errorf("invalid date format: %s", v)
}

// 禁用Unicode转义
func ToBytes2(t interface{}) []byte {
	buf := &bytes.Buffer{}
	encoder := json.NewEncoder(buf)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(t); err != nil {
		log.Printf("%s\n", err.Error())
	}
	return buf.Bytes()
}

// 将列索引转换为 Excel 中的列字母（例如：1 -> A, 2 -> B, 27 -> AA）
func ColumnIndexToLetters(index int) string {
	var result string
	for {
		if index > 0 {
			index--
			result = string(rune('A'+index%26)) + result
			index /= 26
		} else {
			break
		}
	}
	return result
}

func Weekday(t time.Time) string {
	weekdays := map[time.Weekday]string{
		time.Sunday:    "星期日",
		time.Monday:    "星期一",
		time.Tuesday:   "星期二",
		time.Wednesday: "星期三",
		time.Thursday:  "星期四",
		time.Friday:    "星期五",
		time.Saturday:  "星期六",
	}
	return weekdays[t.Weekday()]
}

// getContentType 根据文件扩展名返回正确的MIME类型
func GetContentType(ext string) string {
	switch strings.ToLower(ext) {
	case ".pdf":
		return "application/pdf"
	case ".doc":
		return "application/msword"
	case ".docx":
		return "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
	case ".xls":
		return "application/vnd.ms-excel"
	case ".xlsx":
		return "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
	case ".ppt":
		return "application/vnd.ms-powerpoint"
	case ".pptx":
		return "application/vnd.openxmlformats-officedocument.presentationml.presentation"
	case ".txt":
		return "text/plain"
	case ".jpg", ".jpeg":
		return "image/jpeg"
	case ".png":
		return "image/png"
	case ".gif":
		return "image/gif"
	case ".bmp":
		return "image/bmp"
	case ".svg":
		return "image/svg+xml"
	case ".mp4":
		return "video/mp4"
	case ".avi":
		return "video/x-msvideo"
	case ".mov":
		return "video/quicktime"
	case ".mp3":
		return "audio/mpeg"
	case ".wav":
		return "audio/wav"
	case ".zip":
		return "application/zip"
	case ".rar":
		return "application/x-rar-compressed"
	case ".7z":
		return "application/x-7z-compressed"
	default:
		return "application/octet-stream"
	}
}
