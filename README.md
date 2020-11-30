# flygo

## Introduction

A simple and lightweight web framework, pure native and no third dependencies.

## Features

- Pure native
- No third dependencies
- Support variables route
- Support env variables for deploymenting
- Support filter & interceptor
- Field preset & validation
- Session supports
- Session listener supports
- Go/Template supports

## New

- Support [Middleware](#middleware)

  - logger : Built in Logger
  - [cors](https://github.com/flygotm/cors) : A Cors middleware for flygo
  - [redisauth](https://github.com/flygotm/redisauth) : A simple authentication with redis middleware for flygo
  - [redistoken](https://github.com/flygotm/redistoken) : A simple authorization with redis middleware for flygo
  - [uploadfile](https://github.com/flygotm/uploadfile) : A upload file middleware for flygo
  - [downfile](https://github.com/flygotm/downfile) : A download file middleware for flygo

- Session Provider
  - [cookiprovider](https://github.com/flygotm/cookieprovider) : A default provider for flygo
  - [redisprovider](https://github.com/flygotm/redisprovider) : A session provider using redis for flygo

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
      - css: text/css;charset=utf-8
      - json: application/json;charset=utf-8
      - jpg: image/jpg
      - png: image/png
      - gif: image/gif
      - ico: image/x-icon
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
  session:
    enable: false
  validate:
    err:
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
flygoHost   = "FLYGO_HOST"   //env host
flygoPort   = "FLYGO_PORT"   //env port

flygoContextPath = "FLYGO_CONTEXT_PATH" //env contextPath
flygoWebRoot     = "FLYGO_WEB_ROOT"     //env webRoot

flygoBannerEnable = "FLYGO_BANNER_ENABLE" //env banner enable
flygoBannerType   = "FLYGO_BANNER_TYPE"   //env banner type
flygoBannerText   = "FLYGO_BANNER_TEXT"   //env banner text
flygoBannerFile   = "FLYGO_BANNER_ENABLE" //env banner file

flygoTlsEnable   = "FLYGO_TLS_ENABLE"    //env tls enable
flygoTlsCertFile = "FLYGO_TLS_CERT_FILE" //env tls cert file
flygoTlsKeyFile  = "FLYGO_TLS_KEY_FILE"  //env tls key file

flygoStaticEnable  = "FLYGO_STATIC_ENABLE"  //env static enable
flygoStaticPattern = "FLYGO_STATIC_PATTERN" //env static pattern
flygoStaticPrefix  = "FLYGO_STATIC_PREFIX"  //env static prefix
flygoStaticCache   = "FLYGO_STATIC_CACHE"   //env static cache

flygoStaticFaviconEnable = "FLYGO_STATIC_FAVICON_ENABLE" //env static favicon enable

flygoViewEnable = "FLYGO_VIEW_ENABLE" //env view enable
flygoViewPrefix = "FLYGO_VIEW_PREFIX" //env view prefix
flygoViewSuffix = "FLYGO_VIEW_SUFFIX" //env view suffix
flygoViewCache  = "FLYGO_VIEW_CACHE"  //env view cache

flygoTemplateEnable     = "FLYGO_TEMPLATE_ENABLE"      //env template
flygoTemplateDelimLeft  = "FLYGO_TEMPLATE_DELIM_LEFT"  //env template delim left
flygoTemplateDelimRight = "FLYGO_TEMPLATE_DELIM_RIGHT" //env template delim right

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

## Global Cache

- Set data

```
app.SetCache(KEY,DATA)
```

- Get data

```
app.GetCache(KEY)
```

- Remove data

```
app.RemoveCache(KEY)
```

- Clear caches

```
app.ClearCaches()
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
