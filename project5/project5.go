package main

import (
	"fmt"
	"os"
	"strconv"
	"log"
	
	"github.com/yusufemon/project3/users"
)

func main() {
	method := os.Args[1]
	switch method {
		case "get":
			fmt.Println(users.Get())
		case "insert":
			id := getId()
			name, balance := getNameAndBalance()
			fmt.Println(users.Insert(id, name, balance))
		case "update":
			id := getId()
			name, balance := getNameAndBalance()
			fmt.Println(users.Update(id, name, balance))
		case "delete":
			id := getId()
			fmt.Println(users.Delete(id))
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