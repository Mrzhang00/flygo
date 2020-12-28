# flygo

## Introduction

A simple and lightweight web framework, pure native and no third dependencies.

## Features

- Pure native
- No third dependencies
- Middleware supports
- Session supports
- RESTful controllers

## Middlewares

- `not_found` Not Found resource handler 
- `logger` Built in logger implemention
- `recovery` Recover catch handler
- `method_not_allowed` Method not allowed handler
- `static` Static resource handler
- `cors` Cors handler
- `uploadfile` Upload files
- `downfile` Download files
- `redisauth` Redis simple authentication
- `redistoken` Redis simple authorization
- `session` Session implemention(providers of memory or redis)

## Install

1. GOPATH

```
mkdir -p $GOPATH/src/github.com/billcoding/flygo

cd $GOPATH/src/github.com/billcoding

git clone https://github.com/billcoding/flygo.git flygo
```

2. Go mod

```
require github.com/billcoding/flygo TAG
```