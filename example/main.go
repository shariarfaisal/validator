package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/shariarfaisal/validator"
)

type Address struct {
	Label string `json:"label" v:"enum:Home,Office"`
	Street string `json:"street" v:"required"`
	City string `json:"city" v:"required"`
}

type User struct {
	Name  string `json:"name" v:"required;min:3;max:20"`
	Age   int    `json:"age" v:"required;min:18;max:60"`
	Email string `json:"email" v:"required;email"`
	Addresses []Address `json:"addresses" v:"required;min:1;max:3"`
	DateOfBirth string `json:"dateOfBirth" v:"required;date"`
	Gender string `json:"gender" v:"include:male,female"`
}

func main() {
	user := User{
		Name:  "test",
		Age:   17,
		Email: "example@gmail.com",
		Addresses: []Address{
			{
				Label: "Home",
			},
		},
		Gender: "mole",
	}

	// Validate the user
	ok, errors := validator.Validate(user)
	if !ok {
		fmt.Println(json.NewEncoder(os.Stdout).Encode(errors))
	}

	// Email validation using function
	if isValid := validator.IsValidEmail("test.email"); !isValid {
		fmt.Println("Email is not valid")
	}

}
