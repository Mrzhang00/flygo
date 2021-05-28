# flygo

[![Go Report Card](https://goreportcard.com/badge/github.com/billcoding/flygo)](https://goreportcard.com/report/github.com/billcoding/flygo)
[![GoDoc](https://pkg.go.dev/badge/github.com/billcoding/flygo?status.svg)](https://pkg.go.dev/github.com/billcoding/flygo?tab=doc)

## Overview

- [Introduction](#Introduction)
- [Features](#Features)
- [Install](#Install)
    - [Install by Go PATH](#install-by-go-path)
    - [Install by Go Module](#install-by-go-module)
- [Quickstart](#Quickstart)
    - [Build helloworld App](#build-helloworld-app)
    - [Build simple route App](#build-simple-route-app)
    - [Build dynamic route App](#build-dynamic-route-app)
    - [Build Router App](#build-router-app)
    - [Build Router Group App](#build-router-group-app)
- [Route](#Route)
    - [Simple Route](#simple-route)
    - [Dynamic Route](#dynamic-route)
    - [Router Route](#router-route)
    - [Router Group Route](#router-group-route)
    - [RESTful Controller Route](#restful-controller-route)
- [RESTful Controller](#restful-controller)
    - [How to define a RESTful Controller](#how-to-define-a-restful-controller)
- [Middleware](#Middleware)
    - [How to implement a Middleware](#how-to-implement-a-middleware)
    - [Embedded implemented Middlewares](#embedded-implemented-middlewares)
    - [Extra implemented Middlewares](#extra-implemented-middlewares)
- [Session Support](#session-support)
    - [How to use Session Middleware](#how-to-use-session-middleware)
    - [How to register Session Listener](#how-to-register-session-listener)
- [Binding and Validator](#binding-and-validator)
- [Configuration](#Configuration)
    - [HTTP Server Configuration](#http-server-configuration)
        - [HTTP Server Max Header Size Configuration](#http-server-max-header-size-configuration)
        - [HTTP Server Read Timeout Configuration](#http-server-read-timeout-configuration)
        - [HTTP Server Read Header Timeout Configuration](#http-server-read-header-timeout-configuration)
        - [HTTP Server Write Timeout Configuration](#http-server-write-timeout-configuration)
        - [HTTP Server Idle Timeout Configuration](#http-server-idle-timeout-configuration)
    - [Application Configuration](#application-configuration)
        - [Application Config File Configuration](#application-config-file-configuration)
        - [Application Serve Host Configuration](#application-serve-host-configuration)
        - [Application Serve Port Configuration](#application-serve-port-configuration)
        - [Application Serve TLS Enable Configuration](#application-serve-tls-enable-configuration)
        - [Application Serve TLS Cert File Configuration](#application-serve-tls-cert-file-configuration)
        - [Application Serve TLS Key File Configuration](#application-serve-tls-key-file-configuration)
        - [Application Banner Enable Configuration](#application-banner-enable-configuration)
        - [Application Banner Type Configuration](#application-banner-type-configuration)
        - [Application Banner Text Configuration](#application-banner-text-configuration)
        - [Application Banner File Configuration](#application-banner-file-configuration)
        - [Application Template Enable Configuration](#application-template-enable-configuration)
        - [Application Template Cache Enable Configuration](#application-template-cache-enable-configuration)
        - [Application Template Root Path Configuration](#application-template-root-configuration)
        - [Application Template Suffix Configuration](#application-template-suffix-configuration)
    - [Default YAML Configuration File](#default-yaml-configuration-file)
    - [All Environment Variables Table](#all-environment-variables-table)
- [go-web-framework-benchmark](https://github.com/smallnest/go-web-framework-benchmark#flygo)
- [Pull Request](https://github.com/billcoding/flygo/pulls)
- [Issues](https://github.com/billcoding/flygo/issues)

## Introduction

A simple and lightweight web framework, pure native and no third dependencies.

It can help you to build good web apps.

**That's freedom.**

**We are flying in Go.**

## Features

- Pure native
- No third dependencies
- Middleware supports
- Session supports
- RESTful controllers
- Binding & Validator
- Session supports
- Basic & Variable & Group router
- Multiple supports
- Rich render supports

## Install

### Install by Go PATH

```
mkdir -p $GOPATH/src/github.com/billcoding/flygo

cd $GOPATH/src/github.com/billcoding

git clone https://github.com/billcoding/flygo.git flygo
```

### Install by Go Module

```
require github.com/billcoding/flygo latest
```

## Quickstart

### Build helloworld App

```go
package main

import (
	"fmt"
	"github.com/billcoding/flygo"
	. "github.com/billcoding/flygo/context"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	go func() {
		time.AfterFunc(time.Second, func() {
			response, _ := http.Get("http://localhost")
			resp, _ := ioutil.ReadAll(response.Body)
			fmt.Println(string(resp))
		})
	}()

	flygo.GetApp().GET("/", func(c *Context) {
		c.Text("Helloworld, flygo!")
	}).Run()
}
```

You can see following outputs:

```
......
Helloworld, flygo!
```

### Build simple route App

```go
package main

import (
	"fmt"
	"github.com/billcoding/flygo"
	. "github.com/billcoding/flygo/context"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {

	go func() {
		time.AfterFunc(time.Second, func() {
			response, _ := http.Get("http://localhost/index")
			resp, _ := ioutil.ReadAll(response.Body)
			fmt.Println(string(resp))
		})
	}()

	flygo.GetApp().GET("/index", func(c *Context) {
		c.Text("index")
	}).Run()
}
```

You can see following outputs:

```
......
index
```

### Build dynamic route App

```go
package main

import (
	"fmt"
	"github.com/billcoding/flygo"
	. "github.com/billcoding/flygo/context"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	go func() {
		time.AfterFunc(time.Second, func() {
			response, _ := http.Get("http://localhost/helloworld")
			resp, _ := ioutil.ReadAll(response.Body)
			fmt.Println(string(resp))
		})
	}()

	flygo.GetApp().GET("/{message}", func(c *Context) {
		c.Text(c.Param("message"))
	}).Run()
}
```

You can see following outputs:

```
......
helloworld
```

### Build Router App

```go
package main

import (
	"fmt"
	"github.com/billcoding/flygo"
	. "github.com/billcoding/flygo/context"
	. "github.com/billcoding/flygo/router"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	go func() {
		time.AfterFunc(time.Second, func() {
			response, _ := http.Get("http://localhost/route")
			resp, _ := ioutil.ReadAll(response.Body)
			fmt.Println(string(resp))
		})
	}()

	flygo.GetApp().AddRouter(NewRouter().GET("/route", func(c *Context) {
		c.Text("Routed")
	})).Run()
}
```

You can see following outputs:

```
......
Routed
```

### Build Router Group App

```go
package main

import (
	"fmt"
	"github.com/billcoding/flygo"
	. "github.com/billcoding/flygo/context"
	. "github.com/billcoding/flygo/router"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	go func() {
		time.AfterFunc(time.Second, func() {
			response, _ := http.Get("http://localhost/r/route")
			resp, _ := ioutil.ReadAll(response.Body)
			fmt.Println(string(resp))
		})
	}()

	flygo.GetApp().AddRouterGroup(NewGroupWithPrefix("/r").Add(NewRouter().GET("/route", func(c *Context) {
		c.Text("Group Routed")
	}))).Run()
}
```

You can see following outputs:

```
......
Group Routed
```

## Route

### Simple Route

```go
package main

import (
	"fmt"
	"github.com/billcoding/flygo"
	. "github.com/billcoding/flygo/context"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {

	go func() {
		time.AfterFunc(time.Second, func() {
			response, _ := http.Get("http://localhost/req")
			resp, _ := ioutil.ReadAll(response.Body)
			fmt.Println(string(resp))
			response, _ = http.Post("http://localhost/req", "", nil)
			resp, _ = ioutil.ReadAll(response.Body)
			fmt.Println(string(resp))
			response, _ = http.Get("http://localhost/index")
			resp, _ = ioutil.ReadAll(response.Body)
			fmt.Println(string(resp))
			response, _ = http.Post("http://localhost/index", "", nil)
			resp, _ = ioutil.ReadAll(response.Body)
			fmt.Println(string(resp))
			req, _ := http.NewRequest(http.MethodPut, "http://localhost/index", nil)
			resp2, _ := http.DefaultClient.Do(req)
			resp, _ = ioutil.ReadAll(resp2.Body)
			fmt.Println(string(resp))
			req, _ = http.NewRequest(http.MethodDelete, "http://localhost/index", nil)
			resp2, _ = http.DefaultClient.Do(req)
			resp, _ = ioutil.ReadAll(resp2.Body)
			fmt.Println(string(resp))
			req, _ = http.NewRequest(http.MethodPatch, "http://localhost/index", nil)
			resp2, _ = http.DefaultClient.Do(req)
			resp, _ = ioutil.ReadAll(resp2.Body)
			fmt.Println(string(resp))
			req, _ = http.NewRequest(http.MethodOptions, "http://localhost/index", nil)
			resp2, _ = http.DefaultClient.Do(req)
			resp, _ = ioutil.ReadAll(resp2.Body)
			fmt.Println(string(resp))
		})
	}()

	app := flygo.GetApp()
	rhandler := func(c *Context) {
		c.Text(fmt.Sprintf("[%s]req", c.Request.Method))
	}
	app.REQUEST("/req", rhandler)

	handler := func(c *Context) {
		c.Text(fmt.Sprintf("[%s]index", c.Request.Method))
	}
	app.GET("/index", handler)
	app.POST("/index", handler)
	app.PUT("/index", handler)
	app.DELETE("/index", handler)
	app.PATCH("/index", handler)
	app.OPTIONS("/index", handler)
	app.Run()
}
```

You can see following outputs:

```
......
[GET]req
[POST]req
[GET]index
[POST]index
[PUT]index
[DELETE]index
[PATCH]index
[OPTIONS]index
```

### Dynamic Route

```go
package main

import (
	"fmt"
	"github.com/billcoding/flygo"
	. "github.com/billcoding/flygo/context"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	go func() {
		time.AfterFunc(time.Second, func() {
			response, _ := http.Get("http://localhost/helloworld")
			resp, _ := ioutil.ReadAll(response.Body)
			fmt.Println(string(resp))
			response, _ = http.Get("http://localhost/user/1000")
			resp, _ = ioutil.ReadAll(response.Body)
			fmt.Println(string(resp))
			response, _ = http.Get("http://localhost/user/afxren")
			resp, _ = ioutil.ReadAll(response.Body)
			fmt.Println(string(resp))
			response, _ = http.Get("http://localhost/user/1100/fteen")
			resp, _ = ioutil.ReadAll(response.Body)
			fmt.Println(string(resp))
			response, _ = http.Get("http://localhost/user/2000/detail")
			resp, _ = ioutil.ReadAll(response.Body)
			fmt.Println(string(resp))
		})
	}()

	flygo.GetApp().GET("/{message}", func(c *Context) {
		c.Text(c.Param("message"))
	}).GET("/user/{id}", func(c *Context) {
		c.Text(c.Param("id"))
	}).GET("/user/{id}/detail", func(c *Context) {
		c.Text(c.Param("id"))
	}).GET("/user/{id}/{name}", func(c *Context) {
		c.Text(c.Param("id") + "-" + c.Param("name"))
	}).Run()
}
```

You can see following outputs:

```
......
helloworld
1000
afxren
1100-fteen
2000
```

### Router Route

```go
package main

import (
	"github.com/billcoding/flygo"
	. "github.com/billcoding/flygo/context"
	. "github.com/billcoding/flygo/router"
)

func main() {
	handler := func(c *Context) {
		c.Text("Group Routed")
	}

	r := NewRouter()
	r.GET("/route", handler)
	r.POST("/route", handler)
	r.PUT("/route", handler)
	r.DELETE("/route", handler)
	r.PATCH("/route", handler)
	r.OPTIONS("/route", handler)

	flygo.GetApp().AddRouter(r).Run()
}
```

### Router Group Route

```go
package main

import (
	"github.com/billcoding/flygo"
	. "github.com/billcoding/flygo/context"
	. "github.com/billcoding/flygo/router"
)

func main() {
	handler := func(c *Context) {
		c.Text("Group Routed")
	}

	rg := NewGroupWithPrefix("/r")

	r := NewRouter()
	r.GET("/route", handler)
	r.POST("/route", handler)
	r.PUT("/route", handler)
	r.DELETE("/route", handler)
	r.PATCH("/route", handler)
	r.OPTIONS("/route", handler)

	rg.Add(r)

	flygo.GetApp().AddRouterGroup(rg).Run()
}
```

### RESTful Controller Route

```go
package main

import (
	"fmt"
	"github.com/billcoding/flygo"
	. "github.com/billcoding/flygo/context"
)

type TestRestController struct {
}

func (t *TestRestController) Prefix() string {
	return "/rest"
}

func (t *TestRestController) GET() func(c *Context) {
	return func(c *Context) {
		c.Text(fmt.Sprintf("GET one : %v", c.RestId()))
	}
}

func (t *TestRestController) GETS() func(c *Context) {
	return func(c *Context) {
		c.Text("GET All")
	}
}

func (t *TestRestController) POST() func(c *Context) {
	return func(c *Context) {
		c.Text("POST")
	}
}

func (t *TestRestController) PUT() func(c *Context) {
	return func(c *Context) {
		c.Text("PUT")
	}
}

func (t *TestRestController) DELETE() func(c *Context) {
	return func(c *Context) {
		c.Text(fmt.Sprintf("DELETE one : %v", c.RestId()))
	}
}

func main() {
	flygo.GetApp().REST(&TestRestController{}).Run()
}
```

## RESTful Controller

### How to define a RESTful Controller

1. Implements interface `rest.Controller`

```
type Controller interface {
	Prefix() string 
	GET() func(c *c.Context) 
	GETS() func(c *c.Context) 
	POST() func(c *c.Context)
	PUT() func(c *c.Context)
	DELETE() func(c *c.Context)
}
```

2. Route your RESTful Controller

```
App.REST(RESTfulControllerPtr)
```

3. Example 

```go
package main

import (
	"fmt"
	"github.com/billcoding/flygo"
	. "github.com/billcoding/flygo/context"
)

type TestRestController struct {
}

func (t *TestRestController) Prefix() string {
	return "/rest"
}

func (t *TestRestController) GET() func(c *Context) {
	return func(c *Context) {
		c.Text(fmt.Sprintf("GET one : %v", c.RestId()))
	}
}

func (t *TestRestController) GETS() func(c *Context) {
	return func(c *Context) {
		c.Text("GET All")
	}
}

func (t *TestRestController) POST() func(c *Context) {
	return func(c *Context) {
		c.Text("POST")
	}
}

func (t *TestRestController) PUT() func(c *Context) {
	return func(c *Context) {
		c.Text("PUT")
	}
}

func (t *TestRestController) DELETE() func(c *Context) {
	return func(c *Context) {
		c.Text(fmt.Sprintf("DELETE one : %v", c.RestId()))
	}
}

func main() {
	flygo.GetApp().REST(&TestRestController{}).Run()
}
```

## Middleware

### How to implement a Middleware

1. Implements interface `middleware.Middleware`

```
type Middleware interface {
	Type() *Type
	Name() string
	Method() Method
	Pattern() Pattern
	Handler() func(c *c.Context)
}
```

2. Use your Middleware

```
App.Use(MiddlewarePtr)
```

3. Example

```go
package main

import (
	"github.com/billcoding/flygo"
	. "github.com/billcoding/flygo/context"
	mw "github.com/billcoding/flygo/middleware"
)

type TestMiddleware struct {
}

func (t *TestMiddleware) Type() *mw.Type {
	return mw.TypeHandler
}

func (t *TestMiddleware) Name() string {
	return "TestMiddleware"
}

func (t *TestMiddleware) Method() mw.Method {
	return "GET"
}

func (t *TestMiddleware) Pattern() mw.Pattern {
	return "/test"
}

func (t *TestMiddleware) Handler() func(c *Context) {
	return func(c *Context) {
		c.Text("Middleware works")
	}
}

func main() {
	flygo.GetApp().Use(&TestMiddleware{}).Run()
}
```

### Embedded implemented Middlewares

* `not_found` Not Found resource handler
* `Logger` Built in Logger implementation
* `recovery` Recover catch handler
* `static` Static resource handler
* `cors` Cors handler
* `uploadFile` Upload files
* `downloadFile` Download files
* `redisauth` Redis simple authentication
* `redistoken` Redis simple authorization
* `session` Session implementation(providers: memory or redis)

### Extra implemented Middlewares

* [Captcha middleware](https://github.com/flygotm/captcha)
* [GZIP compression](https://github.com/flygotm/gzip)
* [Deflate compression](https://github.com/flygotm/deflate)
* [Brotli compression](https://github.com/flygotm/brotli)

## Session Support

### How to use Session Middleware

```
app.UseSession(memory.Provider(), &se.Config{Timeout: 60}, nil)
```

### How to register Session Listener

```
App.UseSession(memory.Provider(), &se.Config{Timeout: 60}, &se.Listener{
    Created: func(s se.Session) {
        log.Println("Created")
    },

    Refreshed: func(s se.Session) {
        log.Println("Refreshed")
    },

    Invalidated: func(s se.Session) {
        log.Println("Invalidated")
    },

    Destroyed: func(s se.Session) {
        log.Println("Destroyed")
    },
})
```

## Binding and Validator

```
type model struct {
	Id1    uint8      `binding:"name(id)" validate:"required(T) min(10) max(1000) message(Id1不合法)"`
	Id2    uint16     `binding:"name(id)" validate:"required(T) min(10) max(1000) message(Id2不合法)"`
	Id3    uint32     `binding:"name(id)" validate:"required(T) min(10) max(1000) message(Id3不合法)"`
	Id4    uint       `binding:"name(id)" validate:"required(T) min(10) max(1000) message(Id4不合法)"`
	Id5    uint64     `binding:"name(id)" validate:"required(T) min(10) max(1000) message(Id5不合法)"`
	Float1 float32    `binding:"name(fl)" validate:"required(T) min(10) max(1000) message(Float1不合法)"`
	Float2 float64    `binding:"name(fl)" validate:"required(T) min(10) max(1000) message(Float2不合法)"`
	Name   string     `binding:"name(name)" validate:"required(T) minlength(2) maxlength(5) message(Name不合法)"`
	In     InnerModel `binding:"name(in)" validate:"required(T)"`
}

type InnerModel struct {
	Id   uint8  `binding:"name(id)" validate:"required(T) min(10) max(1000) message(in.Id1不合法)"`
	Name string `binding:"name(name)" validate:"required(T) minlength(2) maxlength(5) message(in.Name不合法)"`
}

func handler(c *Context) {
	m := model{}
	c.BindWithParamsAndValidate(&m, bind.Param, func() {
		c.JSON(&m)
	})
}
```

## Configuration

### HTTP Server Configuration

#### HTTP Server Max Header Size Configuration

* Code set

```
App.Config.Server.MaxHeaderSize = 1 << 9
```

* YAML config

```yaml
server:
  maxHeaderSize: 10000
```

* Environment variable

```
SERVER_MAX_HEADER_SIZE = 10000
```

#### HTTP Server Read Timeout Configuration

* Code set

```
App.Config.Server.Timeout.Read = time.Minute
```

* YAML config

```yaml
server:
  timeout:
    read: 1m
```

* Environment variable

```
SERVER_READ_TIMEOUT = "1m"
```

#### HTTP Server Read Header Timeout Configuration

* Code set

```
App.Config.Server.Timeout.ReadHeader = time.Minute
```

* YAML config

```yaml
server:
  timeout:
    readHeader: 1m1s
```

* Environment variable

```
SERVER_READ_HEADER_TIMEOUT = "1m1s"
```

#### HTTP Server Write Timeout Configuration

* Code set

```
App.Config.Server.Timeout.Write = time.Minute * 2
```

* YAML config

```yaml
server:
  timeout:
    write: 2m
```

* Environment variable

```
SERVER_WRITE_TIMEOUT = "2m"
```

#### HTTP Server Idle Timeout Configuration

* Code set

```
App.Config.Server.Timeout.Idle = time.Second
```

* YAML config

```yaml
server:
  timeout:
    idle: 1s
```

* Environment variable

```
SERVER_IDLE_TIMEOUT = "1s"
```

### Application Configuration

#### Application Config File Configuration

* Code set

```
App.ConfigFile = "app.yml"
```

* Environment variable

```
FLYGO_CONFIG = "app.yml"
```

#### Application Serve Host Configuration

* Code set

```
App.Config.Flygo.Server.Host = "127.0.0.1"
```

* YAML config

```yaml
flygo:
  server:
    host: 127.0.0.1
```

* Environment variable

```
FLYGO_SERVER_HOST = "127.0.0.1"
```

#### Application Serve Port Configuration

* Code set

```
App.Config.Flygo.Server.Port = 8080
```

* YAML config

```yaml
flygo:
  server:
    port: 8080
```

* Environment variable

```
FLYGO_SERVER_PORT = 8080
```

#### Application Serve TLS Enable Configuration

* Code set

```
App.Config.Flygo.Server.TLS.Enable = true
```

* YAML config

```yaml
flygo:
  server:
    tls:
      enable: T
```

* Environment variable

```
FLYGO_SERVER_TLS_ENABLE = "T"
```

#### Application Serve TLS Cert File Configuration

* Code set

```
App.Config.Flygo.Server.TLS.CertFile = "cert.pem"
```

* YAML config

```yaml
flygo:
  server:
    tls:
      certFile: cert.pem
```

* Environment variable

```
FLYGO_SERVER_TLS_CERT_FILE = "cert.pem"
```

#### Application Serve TLS Key File Configuration

* Code set

```
App.Config.Flygo.Server.TLS.KeyFile = "cert.key"
```

* YAML config

```yaml
flygo:
  server:
    tls:
      keyFile: cert.key
```

* Environment variable

```
FLYGO_SERVER_TLS_KEY_FILE = "cert.key"
```

#### Application Banner Enable Configuration

* Code set

```
App.Config.Flygo.Banner.Enable = true
```

* YAML config

```yaml
flygo:
  banner:
    enable: T
```

* Environment variable

```
FLYGO_BANNER_ENABLE = "T"
```

#### Application Banner Type Configuration

* Code set

```
App.Config.Flygo.Banner.Type = "text"
```

* YAML config

```yaml
flygo:
  banner:
    type: text
```

* Environment variable

```
FLYGO_BANNER_TYPE = "text"
```

#### Application Banner Text Configuration

* Code set

```
App.Config.Flygo.Banner.Text = "This is Banner Title"
```

* YAML config

```yaml
flygo:
  banner:
    text: 'This is Banner Title'
```

* Environment variable

```
FLYGO_BANNER_TEXT = "This is Banner Title"
```

#### Application Banner File Configuration

* Code set

```
App.Config.Flygo.Banner.File = "banner.txt"
```

* YAML config

```yaml
flygo:
  banner:
    file: banner.txt
```

* Environment variable

```
FLYGO_BANNER_FILE = "banner.txt"
```

#### Application Template Enable Configuration

* Code set

```
App.Config.Flygo.Template.Enable = true
```

* YAML config

```yaml
flygo:
  template:
    enable: T
```

* Environment variable

```
FLYGO_TEMPLATE_ENABLE = "T"
```

#### Application Template Cache Enable Configuration

* Code set

```
App.Config.Flygo.Template.Cache = true
```

* YAML config

```yaml
flygo:
  template:
    cache: T
```

* Environment variable

```
FLYGO_TEMPLATE_CACHE = "T"
```

#### Application Template Root Configuration

* Code set

```
App.Config.Flygo.Template.Root = "/to/path/templates"
```

* YAML config

```yaml
flygo:
  template:
    root: /to/path/templates
```

* Environment variable

```
FLYGO_TEMPLATE_ROOT = "/to/path/templates"
```

#### Application Template Suffix Configuration

* Code set

```
App.Config.Flygo.Template.Suffix = ".tpl"
```

* YAML config

```yaml
flygo:
  template:
    suffix: .tpl
```

* Environment variable

```
FLYGO_TEMPLATE_SUFFIX = ".tpl"
```

### Default YAML Configuration File

```yaml
server:
  maxHeaderSize: ''
  timeout:
    read: ''
    readHeader: ''
    write: ''
    idle: ''
flygo:
  server:
    host: localhost
    port: 80
    tls:
      enable: false
      certFile: ''
      keyFile: ''
  banner:
    enable: true
    type: default
    text: ''
    file: ''
  template:
    enable: false
    cache: true
    root: ./templates
    suffix: .html
```

### All Environment Variables Table

|Name|Data type|Optional values|Description|
|---|---|---|---|
|SERVER_MAX_HEADER_SIZE|int|e.g. "102400"|HTTP server max header size in bytes|
|SERVER_READ_TIMEOUT|duration|e.g. "1m"|HTTP server read timeout|
|SERVER_READ_HEADER_TIMEOUT|duration|e.g. "1m10s"|HTTP server read header timeout|
|SERVER_WRITE_TIMEOUT|duration|e.g. "2m"|HTTP server write timeout|
|SERVER_IDLE_TIMEOUT|duration|e.g. "1s"|HTTP server Idle timeout|
|FLYGO_CONFIG|string|e.g. "app.yml"|Yaml config file|
|FLYGO_SERVER_HOST|string|e.g. "127.0.0.1"|Serve host|
|FLYGO_SERVER_PORT|int|e.g. "8080"|Serve port|
|FLYGO_SERVER_TLS_ENABLE|bool|[1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False]|Serve TLS enable|
|FLYGO_SERVER_TLS_CERT_FILE|string|e.g. "cert.pem"|Serve TLS cert file|
|FLYGO_SERVER_TLS_KEY_FILE|string|e.g. "cert.key"|Serve TLS key file|
|FLYGO_BANNER_ENABLE|bool|[1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False]|Banner enable|
|FLYGO_BANNER_TYPE|string|["default", "text", "file"]|Banner type|
|FLYGO_BANNER_TEXT|string|e.g. "Banner banana"|Banner text|
|FLYGO_BANNER_FILE|string|e.g. "/to/path/banner.txt"|Banner file|
|FLYGO_TEMPLATE_ENABLE|bool|[1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False]|Template enable|
|FLYGO_TEMPLATE_CACHE|bool|[1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False]|Template cache enable|
|FLYGO_TEMPLATE_ROOT|string|e.g. "/to/path/templates"|Template root path|
|FLYGO_TEMPLATE_SUFFIX|string|e.g. ".tpl"|Template suffix|