# flygo

## Introduction

A simple and lightweight web framework, pure native and no third dependencies.

## Features

- Pure native
- No third dependencies
- Support variables route
- Support env variables
- Support filter & interceptor
- Field preset & validation
- Session supports
- Session listener supports
- Go/Template supports
- Middleware supports

## Middlewares

  - [cors](https://github.com/flygotm/cors)
  
  - [redisauth](https://github.com/flygotm/redisauth)
  
  - [redistoken](https://github.com/flygotm/redistoken)
  
  - [uploadfile](https://github.com/flygotm/uploadfile)
  
  - [downfile](https://github.com/flygotm/downfile)
  
  - [captcha](https://github.com/flygotm/captcha)
  
  - [gzip](https://github.com/flygotm/gzip)
  
  - ......

## Extensions
      
  - [cookiprovider](https://github.com/flygotm/cookieprovider)
  
  - [redisprovider](https://github.com/flygotm/redisprovider)
  
  - ......

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

## Usage

```
package main

import (
	"github.com/billcoding/flygo"
)

func main() {
	app := flygo.NewApp()

	app.Get("/", func(context Context) {
		context.Text("index")
	})

	app.Run()
}
```

## Yml Configuration
```yaml
flygo:
  server:
    host: localhost
    port: 80
    webRoot: ''
    contextPath: ''
  banner:
    enable: true
    type: default
    text: ''
    file: ''
  static:
    enable: false
    pattern: static
    prefix: static
    cache: false
    favicon:
      enable: false
    mimes:
      css: text/css;charset=utf-8
      json: application/json;charset=utf-8
      jpg: image/jpg
      png: image/png
      gif: image/gif
      ico: image/x-icon
  view:
    enable: false
    prefix: templates
    suffix: html
    cache: false
  template:
    enable: false
    delims:
      left: '{{'
      right: '}}'
  validate:
    code: 1
  log:
    type: stdout
    file:
      out: out.log
      err: err.log
```

## Environment Variables

```
flggoConfig = "FLYGO_CONFIG" //env config file

flygoDevDebug = "FLYGO_DEV_DEBUG" //env dev debug

flygoServerHost        = "FLYGO_SERVER_HOST"         //env server host
flygoServerPort        = "FLYGO_SERVER_PORT"         //env server port
flygoServerContextPath = "FLYGO_SERVER_CONTEXT_PATH" //env server contextPath
flygoServerWebRoot     = "FLYGO_SERVER_WEB_ROOT"     //env server webRoot

flygoBannerEnable = "FLYGO_BANNER_ENABLE" //env banner enable
flygoBannerType   = "FLYGO_BANNER_TYPE"   //env banner type
flygoBannerText   = "FLYGO_BANNER_TEXT"   //env banner text
flygoBannerFile   = "FLYGO_BANNER_ENABLE" //env banner file

flygoServerTlsEnable   = "FLYGO_SERVER_TLS_ENABLE"    //env server tls enable
flygoServerTlsCertFile = "FLYGO_SERVER_TLS_CERT_FILE" //env server tls cert file
flygoServerTlsKeyFile  = "FLYGO_SERVER_TLS_KEY_FILE"  //env server tls key file

flygoStaticEnable  = "FLYGO_STATIC_ENABLE"  //env static enable
flygoStaticPattern = "FLYGO_STATIC_PATTERN" //env static pattern
flygoStaticPrefix  = "FLYGO_STATIC_PREFIX"  //env static prefix
flygoStaticCache   = "FLYGO_STATIC_CACHE"   //env static cache

flygoStaticFaviconEnable = "FLYGO_STATIC_FAVICON_ENABLE" //env static favicon enable

flygoViewEnable = "FLYGO_VIEW_ENABLE" //env view enable
flygoViewPrefix = "FLYGO_VIEW_PREFIX" //env view prefix
flygoViewSuffix = "FLYGO_VIEW_SUFFIX" //env view suffix
flygoViewCache  = "FLYGO_VIEW_CACHE"  //env view cache

flygoTemplateEnable      = "FLYGO_TEMPLATE_ENABLE"       //env template
flygoTemplateDelimsLeft  = "FLYGO_TEMPLATE_DELIMS_LEFT"  //env template delims left
flygoTemplateDelimsRight = "FLYGO_TEMPLATE_DELIMS_RIGHT" //env template delims right

flygoSessionEnable  = "FLYGO_SESSION_ENABLE"  //env session enable
flygoSessionTimeout = "FLYGO_SESSION_TIMEOUT" //env session timeout
```

## Middleware

- Middleware (based on route)

  _Must implements flygo.Middleware interface_

```
app.Use(middlewares ...Middleware)
```

- FilterMiddleware (based on filter)

  _Must implements flygo.FilterMiddleware interface_

```
app.UseFilter(filtermiddlewares ...FilterMiddleware)
```

- InterceptorMiddleware (based on interceptor)

  _Must implements flygo.InterceptorMiddleware interface_

```
app.UseInterceptor(interceptorMiddlewares ...InterceptorMiddleware)
```

## Route Types

- Pattern route

```
/index/id
```

- Variables route

```
/index/{id}/{name}
```

## Handler Methods

- All

```
app.Req(PATTERN, HANDLER, FIELD...)
```

- GET

```
app.Get(PATTERN, HANDLER, FIELD...)
```

- POST

```
app.Post(PATTERN, HANDLER, FIELD...)
```

- DELETE

```
app.Delete(PATTERN, HANDLER, FIELD...)
```

- PUT

```
app.Put(PATTERN, HANDLER, FIELD...)
```

- PATCH

```
app.Patch(PATTERN, HANDLER, FIELD...)
```

## Filter Config

- Before filter

```
app.BeforeFilter(PATTERN, HANDLER)
```

- After filter

```
app.AfterFilter(PATTERN, HANDLER)
```

## Interceptor Config

- Get prev results

```
context.Response.GetData()
```

- Before interceptor

```
app.BeforeInterceptor(PATTERN, HANDLER)
```

- After interceptor

```
app.AfterInterceptor(PATTERN, HANDLER)
```

## Field Preset

- First set part enable

```
field.Preset()
```

- Default value

```
field.DefaultVal(DEFAULT_VAL)
```

- Concat multiple values

```
field.Concat(true)
```

- Concat rune

```
field.ConcatRune(CONCAT_RUNE)
```

- Split single value

```
field.Split(true)
```

- Split rune

```
field.SplitRune(SPLIT_RUNE)
```

## Field Validation

- First set part enable

```
field.Validate()
```

- Min number value

```
field.Min(MIN_VALUE)
```

- Max number value

```
field.Max(MAX_VALUE)
```

- Min length string

```
field.MinLength(MIN_LENGTH)
```

- Max length string

```
field.MaxLength(MAX_LENGTH)
```

- Optional values

```
field.Enums(VALUE...)
```

- Regex validation

```
field.Regex(PATTERN)
```

## Multipart Support

- First parse the request

```
context.ParseMultipart(MEMORY_SIZE)
```

- Get a multipart file

```
context.GetMultipartFile(FILE_NAME)
```

- Get some multipart files

```
context.GetMultipartFiles(FILE_NAME)
```

- Get multipart file original filename

```
multipartFile.Filename()
```

- Get multipart file size in byte

```
multipartFile.Size()
```

- Get multipart file MIME type

```
multipartFile.ContentType()
```

- Get formdata request header

```
multipartFile.Header(NAME)
```

- Get formdata request all headers

```
multipartFile.Headers()
```

- Copy to local disk

```
multipartFile.Copy(DISK_PATH)
```