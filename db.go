package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//type DB struct{
//	DB *gorm.DB
//}

func (p *App) OpenDB() (err error) {
	if p.db != nil {
		return
	}
	db, err := gorm.Open("mysql", "root:zhx123456@/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil{
		return err
	}
	p.db = db
	return nil
}


