package main

import "math/rand"

type Account struct {
	Id          int    `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	PhoneNumber int64  `json:"phoneNumber"`
	Balance     int64  `json:"balance"`
}

func NewAccount(firstNameParam, LastNameParam string) *Account {
	return &Account{
		Id:          rand.Intn(10000),
		FirstName:   firstNameParam,
		LastName:    LastNameParam,
		PhoneNumber: int64(rand.Intn(1000000)),
	}
}
