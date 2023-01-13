package models

type Address struct {
	Province string
	City     string
	Postcode string
}

type User struct {
	Name    string
	Age     int
	Address Address
}
