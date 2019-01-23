package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

func (p *App) OpenDB() (err error) {
	if p.db != nil {
		return
	}
	log.Println(p.dbDsn)
	db, err := gorm.Open("mysql", p.dbDsn + "?charset=utf8&parseTime=True&loc=Local")
	if err != nil{
		return err
	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	p.db = db
	return nil
}


