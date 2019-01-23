package main
import (
	"github.com/pkg/errors"
)

type GraphqlLogin struct {
	app *App
}

func(p *GraphqlLogin) Login(args *struct{
	Input *LoginByGraphql
})(string, error){
	var u User
	err := p.app.db.Where("id = ? AND password = ?", args.Input.ID, args.Input.Password).Find(&u).Error
	if err != nil{
		return "", errors.Wrap(err, "no such user or password is incorrect.")
	}
	token, err := GenerateToken(u)
	if err!= nil{
		return "", errors.Wrap(err, "ERROR WHEN CREATE TOKEN")
	}
	return token, err
}

