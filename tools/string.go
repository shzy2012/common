package tools

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/rand"
	"net/url"
	"path"
	"regexp"
	"strings"
	"time"

	"github.com/shzy2012/common/log"
)

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// StringBuilderBody 对string进行链式处理
type StringBuilderBody struct {
	origin string
}

// StringBuilder new stringbuilder
func StringBuilder(origin string) *StringBuilderBody {
	return &StringBuilderBody{
		origin: origin,
	}
}

// Replace 替换函数
func (s *StringBuilderBody) Replace(old, new string) *StringBuilderBody {
	s.origin = strings.Replace(s.origin, old, new, -1)
	return s
}

// Build 返回字符串
func (s *StringBuilderBody) Build() string {
	return s.origin
}

// 获取随机字符串
func GetRandomString(length uint64) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	for i := uint64(0); i < length; i++ {
		result = append(result, bytes[r.Intn(int(len(bytes)))])
	}
	return string(result)
}

// 获取随机数字
func GetRandomNumber(length uint64) string {
	str := "0123456789"
	bytes := []byte(str)
	result := make([]byte, length)
	for i := uint64(0); i < length; i++ {
		result = append(result, bytes[r.Intn(int(len(bytes)))])
	}
	return string(result)
}

// Join 返回包含引号("")的字符串
func Join(a []string, sep string) string {
	switch len(a) {
	case 0:
		return ""
	case 1:
		return "\"" + a[0] + "\""
	}
	n := len(sep) * (len(a) - 1)
	for i := 0; i < len(a); i++ {
		n += len(a[i]) + 2
	}

	var b strings.Builder
	b.Grow(n)
	b.WriteString("\"" + a[0] + "\"")
	for _, s := range a[1:] {
		b.WriteString(sep)
		b.WriteString("\"" + s + "\"")
	}

	return b.String()
}

func ToHex(src []byte) string {
	return hex.EncodeToString(src)
}

func ToHex2(src string) string {
	return fmt.Sprintf("%x", src)
}

func Base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func Base64Decode(str string) string {
	dist, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		log.Infof("%s\n", err.Error())
	}
	return string(dist)
}

// 使用多种符号分割字符串
// text:字符串
// symbols:分隔符.比如匹配逗号、分号: ,;
func Splits(text, symbols string) []string {
	re := regexp.MustCompile(fmt.Sprintf(`[%s]+`, symbols))
	// 使用正则表达式进行分割
	return re.Split(text, -1)
}

/*
s:字符串
fillChar:填充符号
width:目标宽度
*/
func PaddingFromLeft(s string, fillChar string, width int) string {
	if len(s) < width {
		return strings.Repeat(fillChar, width-len(s)) + s
	}

	return s
}

/*
s:字符串
fillChar:填充符号
width:目标宽度
*/
func PaddingFromRight(s string, fillChar string, width int) string {
	// 自定义填充函数
	if len(s) < width {
		return s + strings.Repeat(fillChar, width-len(s))
	}
	return s
}

// 将时间转化为  x min ago
func TimeAgo(t time.Time) string {
	duration := time.Since(t)

	switch {
	case duration < time.Minute:
		return "just now"
	case duration < time.Hour:
		return fmt.Sprintf("%d min ago", int(duration.Minutes()))
	case duration < 24*time.Hour:
		return fmt.Sprintf("%d hours ago", int(duration.Hours()))
	case duration < 30*24*time.Hour:
		return fmt.Sprintf("%d days ago", int(duration.Hours()/24))
	case duration < 12*30*24*time.Hour:
		return fmt.Sprintf("%d months ago", int(duration.Hours()/(24*30)))
	default:
		return fmt.Sprintf("%d years ago", int(duration.Hours()/(24*365)))
	}
}

func IsEmpty(str string) bool {
	return strings.TrimSpace(str) == ""
}

// URL Join
func JoinURL(base string, paths ...string) string {
	u, _ := url.Parse(base)
	u.Path = path.Join(u.Path, path.Join(paths...))
	return u.String()
}

// Trim  trim string
func Trim(str string) string {
	return strings.TrimSpace(str)
}

// 截取包含中文的字符串
func SubString(str string, f, e int) string {
	return string([]rune(str)[f:e])
}

// slice 去重
func Unique(slice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, ele := range slice {
		if _, value := keys[ele]; !value {
			keys[ele] = true
			list = append(list, ele)
		}
	}
	return list
}

// MergeAndUnique 合并多个字符串数组，并确保内容唯一
func MergeAndUnique(arrays ...[]string) []string {
	uniqueMap := make(map[string]struct{})

	// 遍历每个数组并添加到 map 中
	for _, arr := range arrays {
		for _, item := range arr {
			uniqueMap[item] = struct{}{}
		}
	}

	// 将唯一元素添加到切片中
	var merged []string
	for key := range uniqueMap {
		merged = append(merged, key)
	}

	return merged
}

// 字符串首字母大写
func UpperA(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func FindVariables(text string) []string {

	result := make([]string, 0)
	re := regexp.MustCompile(`<([^>]+)>`)
	matches := re.FindAllStringSubmatch(text, -1)
	if len(matches) > 0 {
		for _, match := range matches {
			// match[0] 是整个匹配的字符串，包括尖括号。
			// match[1] 是捕获组匹配的内容，即尖括号之间的字符串。
			result = append(result, match[1])
		}
	}
	return result
}

// 将Map key转化为数组
func MapKey2Array[T comparable](m map[T]interface{}) []T {
	a := make([]T, 0)
	for k := range m {
		a = append(a, k)
	}
	return a
}

// 避免显示乱码
func CheckNil(i interface{}) string {
	switch v := i.(type) {
	case nil:
		return ""
	case int, float64, bool:
		return fmt.Sprintf("%v", v)
	case string:
		return v
	default:
		return ""
	}
}

// 添加随机延迟 0-n 毫秒
func Sleep(t int) {
	time.Sleep(time.Duration(r.Intn(t)) * time.Millisecond)
}
