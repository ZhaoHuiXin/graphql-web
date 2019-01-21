package main

import (
	"flag"
	log "github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
	"net/http"
	"os"
	"time"

	"github.com/gomodule/redigo/redis"
)

type Wand struct{
	debug bool
	dbDsn string
	redDsn string
	db *gorm.DB
	mux *http.ServeMux
	redis *redis.Pool
}

var DefaultWand *Wand = NewWand()

func init(){
	dbName := os.Getenv("DBNAME")
	if dbName == ""{
		dbName = "dbname"
	}

	mysqlUrl := os.Getenv("MYSQL_URL")
	if mysqlUrl == ""{
		mysqlUrl = "mysql://user:password@(host:port)/" + dbName
	}
	flag.StringVar(&DefaultWand.dbDsn, "mysql", mysqlUrl, "usage: mysql uri")

	redisUrl := os.Getenv("REDIS_URL")
	if redisUrl == ""{
		redisUrl = "redis://localhost:6379/0"
	}
	flag.StringVar(&DefaultWand.redDsn, "redis", redisUrl, "usage: redis uri")
}

func NewWand() *Wand{
	return &Wand{
		mux: http.NewServeMux(),
	}
}

func (p *Wand) Init(debug bool){
	p.debug = debug
	err := p.OpenDB()
	if err != nil{
		log.WithField("func", "Wand Init").Info("OpenDB: ", err)
	}
	p.OpenRedisPool()
}

func (p *Wand) OpenRedisPool() {
	p.redis = &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.DialURL(p.redDsn)
		},
	}
}

func (p *Wand) Redis() redis.Conn {
	return p.redis.Get()
}
