package main

import (
	"fmt"
	"github.com/yusufemon/project3/users"
)

func main() {
	// fmt.Printf(stringutil.Reverse("Hello, world.\n"))
	fmt.Println(users.Get())
	fmt.Println(users.Update(2,"budi",2000))
	fmt.Println(users.Get())
	fmt.Println(users.Delete(2))
	fmt.Println(users.Get())
	fmt.Println(users.Insert(2,"budi",20000))
	fmt.Println(users.Get())
}
