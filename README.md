#### golang tool


#### log example
```go
package main

import "github.com/shzy2012/common/log"

func main() {
	log.Info("测试 log")
	log.Debug("测试 log")
	log.Error("测试 log")
}


[INFO]2019/10/17 16:32:33 main.go:6: 测试 log
[DEBG]2019/10/17 16:32:33 main.go:7: 测试 log
[ERRO]2019/10/17 16:32:33 main.go:8: 测试 log
````

#### RandomString example
```go
package main

import (
	"github.com/shzy2012/common/log"
	"github.com/shzy2012/common/tool"
)

func main() {
	t := tool.NewTool()
	log.Info(t.GetRandomString(64))
}

[INFO]2019/10/17 16:45:34 main.go:10: 5CDIiY3fJ1X501Ri5jsbsuomCUPjjLR2tLXYzQ5p5N0kZRnRHqGhDWrC7Hnw7YMx
````

#### StringBuilder example
```go
package main

import (
	"github.com/shzy2012/common/log"
	"github.com/shzy2012/common/tool"
)

func main() {
	origin := "apple,iphone,apple"
	result := tool.StringBuilder(origin).Replace("apple", "fruit").Replace("iphone", "phone").Build()
	log.Info(result)
}

[INFO]2019/10/17 16:43:17 main.go:11: fruit,phone,fruit
```
