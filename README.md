![ui](https://user-images.githubusercontent.com/11075892/90159118-80a65600-ddc2-11ea-91f4-b1afa0fe7818.png)

# Guora

[![license](https://img.shields.io/github/license/meloalright/guora)](https://opensource.org/licenses/MIT)
[![go-mod](https://img.shields.io/github/go-mod/go-version/meloalright/guora)](https://github.com/meloalright/guora)

🖖🏻 A self-hosted Quora like web application written in Go

基于 Golang 类似知乎的私有部署问答应用 包含问答、评论、点赞、管理后台等功能

## Quick Start

`1.Clone Source Code`

```shell
$ git clone https://github.com/meloalright/guora

$ cd guora
```

`2.Download Requirements`

```shell
$ go mod download
```

`3.Edit Configuration`

Open the `configuration.yaml` and edit your redis config.

打开 `configuration.yaml` 编辑你的 redis 环境配置。

`4.Init and Run`

```shell
$ (sudo) go run init/init.go

$ (sudo) go run main.go
```

`5.Visit localhost:8080 and Log in as Admin`

|                 |                      |
| --------------- | -------------------- |
| mail (邮箱)     | admin@guora.mydomain |
| password (密码) | mypassword           |

## Run Test

```shell
$ (sudo) go run init/init.go

$ (sudo) go test
```

## Source

Repository: [Guora](https://github.com/meloalright/guora)

Author: [Meloalright](https://github.com/meloalright)

## License

[MIT](https://opensource.org/licenses/MIT)
