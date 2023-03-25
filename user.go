package main

import "math/rand"

type User struct {
	userID    int
	userName  string
	legalName string
}

func (u *User) Create() *User {
	return &User{userID: rand.Int(), userName: "jinzhu", legalName: "Jinzhu"}
}
