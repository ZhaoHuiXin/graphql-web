package main
import (
	"github.com/pkg/errors"
)

type GraphqlLogin struct {
}

func(_ *GraphqlLogin) Login(args *struct{
	Input *LoginByGraphql
})(string, error){
	for _, user := range users{
		if user.ID == args.Input.ID{
			if user.Password == args.Input.Password{
				token, err := GenerateToken(user)
				if err!= nil{
					return "", errors.Wrap(err, "ERROR WHEN CREATE TOKEN")
				}
				return token, err
			} else{
				return "", errors.New("password is not correct")
			}
		}
	}
	return "", errors.New("User not found.")
}

