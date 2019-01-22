package main

import (
	"flag"
	log "github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
	"os"
	"time"

	"github.com/gomodule/redigo/redis"
)

type App struct{
	debug bool
	dbDsn string
	redDsn string
	db *gorm.DB
	redis *redis.Pool
}

var DefaultApp *App = NewApp()

func init(){
	dbName := os.Getenv("DBNAME")
	if dbName == ""{
		dbName = "dbname"
	}

	mysqlUrl := os.Getenv("MYSQL_URL")
	if mysqlUrl == ""{
		mysqlUrl = "mysql://user:password@(host:port)/" + dbName
	}
	flag.StringVar(&DefaultApp.dbDsn, "mysql", mysqlUrl, "usage: mysql uri")

	redisUrl := os.Getenv("REDIS_URL")
	if redisUrl == ""{
		redisUrl = "redis://localhost:6379/0"
	}
	flag.StringVar(&DefaultApp.redDsn, "redis", redisUrl, "usage: redis uri")
}

func NewApp() *App{
	return &App{
	}
}

func (p *App) Init(debug bool){
	p.debug = debug
	err := p.OpenDB()
	if err != nil{
		log.WithField("func", "Wand Init").Info("OpenDB: ", err)
	}
	p.OpenRedisPool()
}

func (p *App) OpenRedisPool() {
	p.redis = &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.DialURL(p.redDsn)
		},
	}
}

func (p *App) Redis() redis.Conn {
	return p.redis.Get()
}
