package main

import (
	"context"

	"github.com/pkg/errors"
)

func (p *App) GnerateFakeData() (err error){
	p.db.DropTableIfExists(&User{}, &Book{})
	p.db.CreateTable(&User{}, &Book{})
	for _, u := range users{
		if err := p.db.Create(&u).Error; err != nil{
			return err
		}
	}
	for _,b := range books{
		if err := p.db.Create(&b).Error; err != nil{
			return err
		}
	}
	return nil
}

func (_ *Resolver) Hello() string {
	return "hello world"
}

func (r *Resolver) GetUser(ctx context.Context, args struct{Input *UserArg}) (*UserResolver, error) {
	user, err := r.app.getUser(ctx, args.Input.ID)
	if err != nil {
		return nil, errors.Wrap(err, "GetUser")
	}

	res := UserResolver{
		app: r.app,
		m: *user,
	}
	return &res, nil
}



