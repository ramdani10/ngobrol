package commons

import (
	"github.com/gomodule/redigo/redis"
	"github.com/imansuparman/ngobrol/config"
	"github.com/kitabisa/perkakas/v2/log"
	"github.com/kitabisa/perkakas/v2/metrics/influx"
	"github.com/pkgz/websocket"
	"gopkg.in/gorp.v2"
)

// Options common option for all object that needed
type Options struct {
	Config    config.Provider
	DbMysql   *gorp.DbMap
	DbPostgre *gorp.DbMap
	CachePool *redis.Pool
	Influx    *influx.Client
	Logger    *log.Logger
	Websocket *websocket.Server
}
