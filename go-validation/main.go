package main

import (
	"log"

	"github.com/go-playground/validator/v10"
)

type User struct {
	Name  string `validate:"required,min=1,max=10"`
	Email string `validate:"required,email"`
	Age   int    `validate:"required,gte=18"`
}

func main() {
	u := User{
		// Name:  "TipsGo",
		Email: "tipsgo@.com",
		Age:   17,
	}

	validate := validator.New()
	err := validate.Struct(u)
	if err != nil {
		log.Println("validation failed!")

		for _, err := range err.(validator.ValidationErrors) {
			log.Printf("Field: %s, Error: %s", err.Field(), err.Tag())
		}
	} else {
		log.Println("validation success!")
	}
}
