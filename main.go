package main

import (
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

type Person struct {
	Name string `validate:"required"`
	Age  int    `validate:"gte=0,lte=150"`
}

func main() {
	user := User{
		ID:    1,
		Name:  "John",
		Age:   25,
		Email: "john@example.com",
	}
	fmt.Println("user:", user)

	jsonData, err := json.Marshal(user)
	if err != nil {
		fmt.Println("JSON marshal error:", err)
		return
	}
	fmt.Println("jsonData:", string(jsonData))

	validate := validator.New()

	person := Person{
		Name: "",
		Age:  25,
	}

	err = validate.Struct(person)
	if err != nil {
		fmt.Println("Validation error:", err)
	} else {
		fmt.Println("Person is valid:", person)
	}
}
