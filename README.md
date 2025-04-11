# Ceres


The Comunion business backend service.

## Projects

We have chosen the [ego](https://github.com/gotomicro/ego) to support the config and other common behaviors of web project.
Fistly, you have to read the documents of Ego seriously.

The user interface API we will support as HTTP APIs, and the eth event will be trigger from other service by gRpc interfaces.

Pleas see more informations at [Implementation Notes](implementaton.md)

## Description
Comunion Backend Service
### References
#### EGO
[HomePage](https://github.com/gotomicro/ego)

#### GORM
[HomePage](https://gorm.io/)

## Quick Start

- run docker-compose

```shell
 docker-compose -f docker-compose.yml up -d
```

- init database

```shell
./hack/run database init
```

- copy the config file

```shell
cp ./hack/config/config.toml ./hack/config/config.dev.toml
```

- running the app

```shell
./hack/run start
```

## Project structure
```
├── config (Config)
│   ├── local.toml
├── logs (Log)
├── pkg (Source codes)
│   ├── invoker
│   ├── model (DB models)
│   │   ├── dto
│   │   ├── mysql
│   │   └── transport
│   └── router (API routers)
├── .gitignore
├── go.mod (modules)
├── go.sum
├── main.go (main)
├── README.md
```

## Manual to devlop
### router
```go
// pkg/router/router.go
r.GET("/api/enums/:id", core.Handle(api.EnumInfo))
```
### handler
```go
// pkg/router/api/enum.go
func EnumInfo(c *core.Context) {
	id := cast.ToInt(c.Param("id"))
	if id == 0 {
		c.JSONE(1, "bad request", nil)
		return
	}
	info, _ := mysql.EnumInfo(invoker.Db, id)
	c.JSONOK(info)
}
```
### model
```go
// pkg/model/mysql/enum.go
type Enum struct {
	Id int `gorm:"AUTO_INCREMENT;comment:'id'"`
	GroupKey string `gorm:"not null;comment:'unique key'"`
	GroupTitle string `gorm:"not null;comment:'group title'"`
	Key int `gorm:"not null;comment:'key'"`
	Title string `gorm:"not null;comment:'title'"`
	Ctime int64 `gorm:"not null;comment:'created at'"`
	Utime int64 `gorm:"not null;comment:'updated at'"`
	Dtime int64 `gorm:"not null;comment:'deleted at'"`
	
}
```
### query
```go
// pkg/model/mysql/enum.go
func EnumInfo(db *gorm.DB, paramId int) (resp Enum, err error) {
	var sql = "`id`= ?"
	var binds = []interface{}{paramId}

	if err = db.Model(Enum{}).Where(sql, binds...).First(&resp).Error; err != nil {
		invoker.Logger.Error("enum info error", zap.Error(err))
		return
	}
	return
}
```
### dto (optional)
```go
// pkg/model/dto/enum.go
type EnumCreate struct {
	Id int `json:"id" binding:""` // id
	GroupKey string `json:"groupKey" binding:""` // group key
	GroupTitle string `json:"groupTitle" binding:""` // group title
	Key int `json:"key" binding:""` // key
	Title string `json:"title" binding:""` // title
}
```

