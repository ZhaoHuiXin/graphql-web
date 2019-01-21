package main

import (
	"context"
)

func (_ *Resolver) Hello() string {
	return "hello world"
}

func (r *Resolver) GetUser(ctx context.Context, args struct{}) string {
	return "hello world"
}


