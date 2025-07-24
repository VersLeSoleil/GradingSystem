package main

import (
	"backend/db"
	"backend/function"
	"fmt"
	"net/http"
)

func RegisterMux(mux *http.ServeMux) {

	mux.HandleFunc("/login", function.LoginCheck)
	mux.HandleFunc("/register", function.RegisterUser)
	mux.HandleFunc("/createPost", function.CreatePost)
	mux.HandleFunc("/getPosts", function.GetAllPosts)
	mux.HandleFunc("/getPostByID", function.GetPostByID)
	mux.HandleFunc("/updatePost", function.UpdatePost)
	mux.HandleFunc("/getPostsByUsername", function.GetPostsByUsername)
	mux.HandleFunc("/UpdateUserInfo", function.ModifyUserInfo)
}

func initDATABase() {

	db.InitDB("config.yaml")

}

func main() {
	initDATABase() // 初始化数据库

	mux := http.NewServeMux() // 创建路由器
	RegisterMux(mux)          // 注册路由
	fmt.Println(db.HashPassword("123"))
	fmt.Println("Server started at :8888")
	http.ListenAndServe(":8888", mux) // 启动服务器
}
