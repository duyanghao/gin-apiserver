GinApiServer
===================

GinApiServer is a HTTP apiserver framework based on [Gin](https://github.com/gin-gonic/gin).

## Replacement

You can get another ApiServer instance by executing following command:

```bash
$ grep -rl GinApiServer . | xargs sed -i 's/GinApiServer/youapiserver/g'  
```

## Usage

```bash
$ bash start.sh
$ tailf GinApiServer.log
```

## Structure

```bash
pkg/
├── config
│   ├── config.go
│   ├── key.go
│   ├── model.go
│   └── opt_defs.go
├── controller
│   ├── ping.go
│   ├── todo.go
│   └── version.go
├── log
│   └── log.go
├── middleware
│   ├── basic_auth_middleware.go
├── models
│   └── common.go
├── route
│   └── routes.go
├── service
│   └── todo.go
├── store
└── util
```

* config: configuration
* controller: achieve function by calling relevant service(MVC-Controller)
* service: achieve application-specific logic
* log: log module
* middleware: responsible for auth and common-middleware logic
* models: model(MVC-Model)
* route: gin-route
* store: storage(eg: MySQL, Redis)
* util: basic library

## Feature

* Support configmap reload api
* Support ping-pong health&version check
* Support dump-goroutine-stack-traces

## Image

```bash
$ make dockerfiles.build
```

## 3rd party Implements

* [coredns-dynapi-adapter - coredns dynamic middleware apiserver adapter](https://github.com/duyanghao/coredns-dynapi-adapter)

## Refs

* [GinApiServer Framework](https://duyanghao.github.io/GinApiServer/)
* [how-to-dump-goroutine-stack-traces](https://colobu.com/2016/12/21/how-to-dump-goroutine-stack-traces/)
