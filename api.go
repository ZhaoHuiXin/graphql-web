package main

import (
	"context"
	"github.com/pkg/errors"
)

func (_ *Resolver) Hello() string {
	return "hello world"
}

func (r *Resolver) GetUser(ctx context.Context, args struct{Input *UserArg}) (*UserResolver, error) {
	user, err := r.wand.getUser(ctx, args.Input.ID)
	if err != nil {
		return nil, errors.Wrap(err, "GetUser")
	}

	res := UserResolver{
		m: *user,
	}
	return &res, nil
}


