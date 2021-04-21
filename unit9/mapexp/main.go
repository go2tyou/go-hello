package main

import (
	"fmt"
)

func modifyUser(users map[string]map[string]string, name string)  {
	// v, ok := users[name]
	if users[name] != nil {
		users[name]["pwd"] = "888888"
	} else {
		users[name] = make(map[string]string, 2)
		users[name]["pwd"] = "888888"
		users[name]["nickname"] = "nick~" + name
	}
}

func main()  {
	users := make(map[string]map[string]string, 10)
	users["smith"] = make(map[string]string)
	users["smith"]["pwd"] = "999999"
	users["smith"]["nickname"] = "sb"
	modifyUser(users, "tom")
	modifyUser(users, "marry")
	modifyUser(users, "smith")
	fmt.Println("users=", users)
}