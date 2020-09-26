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
  - [cors](https://gitee.com/flygotm/cors) : A Cors middleware for flygo
  - [redisauth](https://gitee.com/flygotm/redisauth) : A simple authentication with redis middleware for flygo
  - [redistoken](https://gitee.com/flygotm/redistoken) : A simple authorization with redis middleware for flygo
  - [uploadfile](https://gitee.com/flygotm/uploadfile) : A upload file middleware for flygo
  - [downfile](https://gitee.com/flygotm/downfile) : A download file middleware for flygo

- Session Provider
  - [cookiprovider](https://gitee.com/flygotm/cookieprovider) : A default provider for flygo
  - [redisprovider](https://gitee.com/flygotm/redisprovider) : A session provider using redis for flygo

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

## Versions

https://github.com/billcoding/flygo/releases

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

## Configuration

| Name            | Key               | Config                                      | Default     | Description                       |
| --------------- | ----------------- | ------------------------------------------- | ----------- | --------------------------------- |
| WebRoot         | global.webRoot    | `app.SetWebRoot(WEB_ROOT)`                  | `./`        | The app webroot                   |
| StaticEnable    | static.enable     | `app.SetStaticEnable(STATIC_ENABLE)`        | `false`     | Static resource enable            |
| StaticPattern   | static.pattern    | `app.SetStaticPattern(STATIC_PATTERN)`      | `/static`   | Static resource request URL       |
| StaticPrefix    | static.prefix     | `app.SetStaticPrefix(STATIC_PREFIX)`        | `static`    | Static resource local file prefix |
| StaticCache     | static.cache      | `app.SetStaticCache(STATIC_CACHE)`          | `true`      | Static resource enable cache      |
| ViewPrefix      | view.prefix       | `app.SetViewPrefix(VIEW_PREFIX)`            | `templates` | View local file prefix            |
| ViewSuffix      | view.suffix       | `app.SetViewSuffix(VIEW_SUFFIX)`            | `html`      | View local file suffix            |
| ViewCache       | view.cache        | `app.SetViewCache(VIEW_CACHE)`              | `true`      | View enable cache                 |
| ValidateErrCode | validate.err.code | `app.SetValidateErrCode(Validate_ERR_CODE)` | `1`         | Route field validate err code     |

## Environment Variables

| Name               | Description                      |
| ------------------ | -------------------------------- |
| FLYGO_HOST         | bind address                     |
| FLYGO_PORT         | bind port                        |
| FLYGO_WEB_ROOT     | app webroot                      |
| FLYGO_STATIC_CACHE | app enable static resource cache |
| FLYGO_VIEW_CACHE   | app view cache                   |

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

## Default Handler Config

- Not found resource

```
 app.DefaultHandler(HANDLER)
```

- Not supported request

```
app.RequestNotSupportedHandler(HANDLER)
```

- Favicon ico handler

```
app.FaviconIco()
```

_The favicon ico file should be stored at `$WEB_ROOT/$STATIC_PREFIX/favicon.ico`_

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
