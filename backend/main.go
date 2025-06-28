package main

import (
	"backend/db"
	"fmt"
)

func main() {
	db.InitDB("config.yaml")
	user, err := db.GetUserByUsername("bob22")
	if err != nil {
		panic(err)
	}
	if user == nil {
		fmt.Println("未找到 admin 用户")
	} else {
		fmt.Printf("User: %s, Role: %s, Email: %s\n", user.UserName, user.Role, user.Email)
	}

}
