# flygo

## Introduction

A simple and lightweight web framework, pure native and no third dependencies.

## Features

- Pure native
- No third dependencies
- Middleware supports
- Session supports
- RESTful controllers

## Embedded Middlewares

- [x] `not_found` Not Found resource handler
- [x] `logger` Built in logger implemention
- [x] `recovery` Recover catch handler
- [x] `method_not_allowed` Method not allowed handler
- [x] `static` Static resource handler
- [x] `cors` Cors handler
- [x] `uploadfile` Upload files
- [x] `downfile` Download files
- [x] `redisauth` Redis simple authentication
- [x] `redistoken` Redis simple authorization
- [x] `session` Session implemention(providers of memory or redis)

## Implements Middlewares
- [x] [Captcha middleware](https://github.com/flygotm/captcha)
- [x] [GZIP compression](https://github.com/flygotm/gzip)
- [x] [Deflate compression](https://github.com/flygotm/deflate)
- [x] [Brotli compression](https://github.com/flygotm/brotli)

## Install

1. GOPATH

```
mkdir -p $GOPATH/src/github.com/billcoding/flygo

cd $GOPATH/src/github.com/billcoding

git clone https://github.com/billcoding/flygo.git flygo
```

2. Go mod

```
require github.com/billcoding/flygo latest
```

## Docs
[Wiki](https://flygotm.github.com/)



