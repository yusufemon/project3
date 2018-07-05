package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/yusufemon/project3/app/users"
)

func main() {
	method := os.Args[1]
	user, err := users.New()
	checkErr(err)
	switch method {
	case "get":
		data, err := user.Get()
		checkErr(err)
		fmt.Println(data)
	case "insert":
		id := getId()
		name, balance := getNameAndBalance()
		err := user.Insert(id, name, balance)
		checkErr(err)
		fmt.Println("OK")
	case "update":
		id := getId()
		name, balance := getNameAndBalance()
		err := user.Update(id, name, balance)
		checkErr(err)
		fmt.Println("OK")
	case "delete":
		id := getId()
		err := user.Delete(id)
		checkErr(err)
		fmt.Println("OK")
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
