GinApiServer
===================

GinApiServer is ApiServer Framework based on [gin](https://github.com/gin-gonic/gin).

## Replacement

You can get another ApiServer instance by executing following command:

```bash
grep -rl GinApiServer . | xargs sed -i 's/GinApiServer/youapiserver/g'  
```

## Usage

```bash
bash start.sh
tailf GinApiServer.log
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
* Support ping-pong health check
* Support version get

## Image

```bash
make dockerfiles.build
```
