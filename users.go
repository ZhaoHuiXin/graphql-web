package main

import (
	"context"
	"github.com/pkg/errors"
)

type User struct {
	ID int32	`gorm:"AUTO_INCREMENT;PRIMARY_KEY"`
	Name string `gorm:"type:varchar(16);not null;default ''"`
	Mail string	`gorm:"type:varchar(32);not null;default ''"`
	Password string `gorm:"type:varchar(16);not null;default ''"`
	Books []Book
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



func (p *App) getUser(ctx context.Context, id int32) (*User, error){
	var user User
	err := p.db.First(&user, id).Error
	if err != nil{
		return nil, err
	}

	return &user, nil
}

func (p *App) GetUserBooks(ctx context.Context, id int32)([]Book, error){
	var b []Book
	err := p.db.Where("user_id = ?", id).Find(&b).Error
	if err != nil{
		return nil, err
	}
	return b, nil
}

type UserResolver struct{
	app *App
	m User
}

func (u *UserResolver) Id(ctx context.Context) *int32{
	return &u.m.ID
}

func (u *UserResolver) Name(ctx context.Context) *string{
	return &u.m.Name
}

func (u *UserResolver) Mail(ctx context.Context) *string{
	return &u.m.Mail
}

func (u *UserResolver) Password(ctx context.Context) string{
	return u.m.Password
}

func (u *UserResolver) Books(ctx context.Context) (*[]*BookResolver, error){
	books, err := u.app.GetUserBooks(ctx, u.m.ID)
	if err != nil{
		return nil, errors.Wrap(err, "Books")
	}

	r := make([]*BookResolver, len(books))
	for i := range books{
		r[i] = &BookResolver{
			app: u.app,
			m: books[i],
		}
	}
	return &r, nil
}
