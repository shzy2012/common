package tool

import "net/url"

//GetParam 获取Url的值
func GetParam(url *url.URL, name string) string {
	keys, ok := url.Query()[name]
	if !ok || len(keys[0]) < 1 {
		return ""
	}
	return keys[0]
}

//GetParams 获取Url的值
func GetParams(url *url.URL, name string) []string {
	keys, _ := url.Query()[name]
	return keys
}
