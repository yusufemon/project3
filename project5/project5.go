package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/yusufemon/project3/users"
)

func main() {
	method := os.Args[1]
	user, err := usersd.New()
	checkErr(err)
	switch method {
	case "get":
		fmt.Println(user.Get())
	case "insert":
		id := getId()
		name, balance := getNameAndBalance()
		fmt.Println(user.Insert(id, name, balance))
	case "update":
		id := getId()
		name, balance := getNameAndBalance()
		fmt.Println(user.Update(id, name, balance))
	case "delete":
		id := getId()
		fmt.Println(user.Delete(id))
	}
}

func getId() int {
	id, err := strconv.Atoi(os.Args[2])
	checkErr(err)

	return id
}

func getNameAndBalance() (string, int) {
	name := os.Args[3]
	balance, err := strconv.Atoi(os.Args[4])
	checkErr(err)

	return name, balance
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
