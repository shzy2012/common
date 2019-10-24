

<p align="center">
	<img src="https://github.com/shzy2012/static/blob/master/toolbox.png?raw=true" width="120" height="120">
</p>

<h1 align="center">Toolbox</h1>

<p align="center">

[![build status][travis-image]][travis-url]   [![codecov][cov-image]][cov-url] [![GitHub license](https://img.shields.io/github/license/laiye-ai/wulai-openapi-sdk-golang?style=social)](https://travis-ci.org/shzy2012/common/blob/master/LICENSE)


[travis-image]: https://travis-ci.org/shzy2012/common.svg?branch=master

[travis-url]: https://travis-ci.org/shzy2012/common

[cov-image]: https://codecov.io/gh/shzy2012/common/branch/master/graph/badge.svg

[cov-url]: https://codecov.io/gh/shzy2012/common

</p>

## 安装
使用 `go get` 下载安装 SDK

```sh
$ go get -u github.com/shzy2012/common
```

如果您使用了 glide 管理依赖，您也可以使用 glide 来安装 SDK

```sh
$ glide get github.com/shzy2012/common
```

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

#### random string example
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

#### string builder example
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
