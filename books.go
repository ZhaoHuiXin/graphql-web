package main

import "context"

type Book struct{
	ID int32 `gorm:"AUTO_INCREMENT;PRIMARY_KEY"`
	Title string `gorm:"type:varchar(32);not null;default ''"`
	Author string `gorm:"type:varchar(32);not null;default ''"`
	PulishAt string `gorm:"type:varchar(10);not null;default ''"`
	BorrowAt string `gorm:"type:varchar(10);not null;default ''"`
	BackAt string `gorm:"type:varchar(10);not null;default ''"`
	UserId int32 `gorm:"type:int(11);not null;default 0"`
}

var books = []Book{
	{
		Title:"The Old Man and the Sea",
		Author:"Ernest Miller Hemingway",
		PulishAt:"1952-09-01",
		UserId:1,
	},
	{
		Title:"Le Rouge et le Noir",
		Author:"Stendhal",
		PulishAt:"1830-11",
		UserId:1,
	},
}

type BookResolver struct {
	app *App
	m Book
}

func (b *BookResolver)Title(ctx context.Context) *string{
	return &b.m.Title
}

func (b *BookResolver)ID(ctx context.Context) *int32{
	return &b.m.ID
}

func (b *BookResolver)Author(ctx context.Context) *string{
	return &b.m.Author
}