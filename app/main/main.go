package main

import (
	"fmt"
	"github.com/yusufemon/project3/app/users"
	"log"
)

func main() {
	user, err := users.New()

	data, err := user.Get()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)

	err = user.Update(2, "budi", 2000)
	if err != nil {
		panic(err)
	}
	fmt.Println("updated")

	data, err = user.Get()
	if err != nil {
		panic(err)
	}
	fmt.Println(data)

	err = user.Delete(2)
	if err != nil {
		panic(err)
	}
	fmt.Println("OK")

	data, err = user.Get()
	if err != nil {
		panic(err)
	}
	fmt.Println(data)

	err = user.Insert(2, "budi", 20000)
	if err != nil {
		panic(err)
	}
	fmt.Println("OK")

	data, err = user.Get()
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
}
