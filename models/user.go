package models

type Address struct {
	Province string `json:"province" bson:"province"`
	City     string `json:"city" bson:"city"`
	Postcode string `json:"postcode" bson:"postcode"`
}

type User struct {
	Name    string  `json:"name" bson:"user_name"`
	Age     int     `json:"age" bson:"user_age"`
	Address Address `json:"address" bson:"user_address"`
}
