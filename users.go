package main

import(
	"context"
)

type User struct {
	ID int32	`gorm:"AUTO_INCREMENT;PRIMARY_KEY"`
	Name string `gorm:"type:varchar(16)"`
	Mail string	`gorm:"type:varchar(32)"`
	Password string `gorm:"type:varchar(16)"`
}

type LoginByGraphql struct{
	ID int32
	Password string
}

var users = []User{
	{
		ID:       1,
		Name:     "jack",
		Mail:     "jack@163.com",
		Password: "123456",
	},
	{
		ID:       2,
		Name:     "rose",
		Mail:     "rose@163.com",
		Password: "123456",
	},
}

func (p *Wand) GnerateFakeData() (err error){
	p.db.DropTableIfExists(&User{})
	p.db.CreateTable(&User{})
	for _, u := range users{
		if err := p.db.Create(&u).Error; err != nil{
			return err
		}
	}
	return nil
}

func (p *Wand) getUser(ctx context.Context, id int32) (*User, error){
	var user User
	err := p.db.First(&user, id).Error
	if err != nil{
		return nil, err
	}

	return &user, nil
}