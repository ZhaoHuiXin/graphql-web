package main

import (
	"context"
)

type User struct {
	ID int32	`gorm:"AUTO_INCREMENT;PRIMARY_KEY"`
	Name string `gorm:"type:varchar(16);not null;default ''"`
	Mail string	`gorm:"type:varchar(32);not null;default ''"`
	Password string `gorm:"type:varchar(16);not null;default ''"`
}

type LoginByGraphql struct{
	ID int32
	Password string
}
type UserArg struct{
	ID int32
	Name string
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

func (p *App) GnerateFakeData() (err error){
	p.db.DropTableIfExists(&User{})
	p.db.CreateTable(&User{})
	for _, u := range users{
		if err := p.db.Create(&u).Error; err != nil{
			return err
		}
	}
	return nil
}

func (p *App) getUser(ctx context.Context, id int32) (*User, error){
	var user User
	err := p.db.First(&user, id).Error
	if err != nil{
		return nil, err
	}

	return &user, nil
}

type UserResolver struct{
	app *App
	m User
}

func (u *UserResolver) ID(ctx context.Context) *int32{
	return &u.m.ID
}

func (u *UserResolver) Name(ctx context.Context) *string{
	return &u.m.Name
}

func (u *UserResolver) Mail(ctx context.Context) *string{
	return &u.m.Mail
}

func (u *UserResolver) Password(ctx context.Context) string{
	return *&u.m.Password
}
