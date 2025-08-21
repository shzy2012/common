package tools

import (
	"regexp"
	"strings"
)

func CleanLLMOutput(output string) string {
	output = strings.TrimSpace(output)
	output = strings.Replace(output, "```", "", -1) //将```替换为空
	output = strings.Replace(output, `""`, `"`, -1) //将两个"替换为1个"
	return output
}

// find mongo id in content
func FindMongoID(content string) []string {
	if IsEmpty(content) {
		return []string{}
	}

	results := make([]string, 0)

	// 定义正则表达式模式
	// MongoDB的ID是24个字符的十六进制字符串
	pattern := `\b[a-f0-9]{24}\b`
	re := regexp.MustCompile(pattern) // 编译正则表达式

	for _, ele := range strings.Split(content, "\n") {
		// 查找所有匹配的字符串
		matches := re.FindAllString(ele, -1)
		if len(matches) > 0 {
			results = append(results, matches...)
		}
	}

	return results
}
