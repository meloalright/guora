![ui](https://user-images.githubusercontent.com/11075892/90159118-80a65600-ddc2-11ea-91f4-b1afa0fe7818.png)

# Guora

[![license](https://img.shields.io/github/license/meloalright/guora)](https://opensource.org/licenses/MIT)
[![go-mod](https://img.shields.io/github/go-mod/go-version/meloalright/guora)](https://github.com/meloalright/guora)

🖖🏻 A self-hosted Quora like web application written in Go

基于 Golang 类似知乎的私有部署问答应用 包含问答、评论、点赞、管理后台等功能

## Quick Start (Docker Deploy)

```sh
$ docker run -d --name guora -p 8080:8080 meloalright/guora:beta3
```

## Development (Non-Dockerized Deploy)

### 1.Clone Source Code

```shell
$ git clone https://github.com/meloalright/guora

$ cd guora
```

### 2.Download Requirements

```shell
$ go mod download
```

### 3.Create Configuration

```shell
$ touch /etc/guora/configuration.yaml
```

```yaml
# configuration example
sql:
  sqlite3: true
  addr: ./guora.db
redis:
  addr: localhost:6379
  password:
  db: 0
admin:
  name: Guora Robot (管理员)
  mail: admin@localhost
  password: mypassword
secretkey: JustWriteSomethingWhatYouLike
lang: en
address: :8080
```

| Param     | Description                       | 备注                           |
| --------- | --------------------------------- | ------------------------------ |
| sql       | Database configure                | 数据库配置                     |
| redis     | Redis configure                   | Redis 配置                     |
| admin     | Administrator info                | 管理员信息                     |
| secretkey | Secret string for token signature | Token 密钥                     |
| lang      | languages, such as en, zh         | 语言: en 为英文; zh 为简体中文 |
| address   | Listening port                    | 监听端口                       |

### 4.Init and Run

```shell
$ (sudo) go run ./cmd/guora -init
```

### 5. Visit Website

visit [localhost:8080](http://localhost:8080) and log in as admin

|                     |                 |
| ------------------- | --------------- |
| mail (默认邮箱)     | admin@localhost |
| password (默认密码) | mypassword      |

## Run Test

```shell
$ (sudo) go test ./cmd/guora
```

## Source

Repository: [guora](https://github.com/meloalright/guora)

Author: [meloalright](https://github.com/meloalright)

Contributors: [contributors](https://github.com/meloalright/guora/graphs/contributors)

## ChangeLog

Documented in [Releases](https://github.com/meloalright/guora/releases)

## License

[MIT](https://opensource.org/licenses/MIT)
