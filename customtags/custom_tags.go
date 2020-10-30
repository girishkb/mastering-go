package main

import (
	"fmt"
)

// Name of the struct tag used in examples

type User struct {
	Id    int    `validate:"number,min=1,max=1000"`
	Name  string `validate:"string,min=2,max=10"`
	Bio   string `validate:"string"`
	Email string `validate:"email"`
}

func main() {
	user := User{
		Id:    0,
		Name:  "superlongstring",
		Bio:   "",
		Email: "foobar",
	}
	fmt.Println("Errors:")
	for i, err := range validateStruct(user) {
		fmt.Printf("\t%d. %s\n", i+1, err.Error())
	}

	/*usertest := UserTest{
		Id:    1,
		Name:  "John Doe",
		Email: "john@example",
	}
	t := reflect.TypeOf(usertest)
	fmt.Println("Type:", t.Name())
	fmt.Println("Kind:", t.Kind())
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get(tagName)
		fmt.Printf("%d. %v (%v), tag: '%v'\n", i+1, field.Name, field.Type.Name(), tag)
	}*/
}
