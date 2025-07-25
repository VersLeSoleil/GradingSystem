package main

import (
	"backend/db"
	"backend/function"
	"fmt"
	"net/http"
	"backend/jwt"
	"backend/restful"
)


func RegisterMux(mux *http.ServeMux) {
	// 不需要令牌的接口
	mux.HandleFunc("/login", function.LoginCheck)
	mux.HandleFunc("/register", function.RegisterUser)

	// 需要令牌验证的接口
	mux.Handle("/createPost", jwt.AuthMiddleware(http.HandlerFunc(function.CreatePost)))
	mux.Handle("/getPosts", jwt.AuthMiddleware(http.HandlerFunc(function.GetAllPosts)))
	mux.Handle("/getPostByID", jwt.AuthMiddleware(http.HandlerFunc(function.GetPostByID)))
	mux.Handle("/updatePost", jwt.AuthMiddleware(http.HandlerFunc(function.UpdatePost)))
	mux.Handle("/getPostsByUsername", jwt.AuthMiddleware(http.HandlerFunc(function.GetPostsByUsername)))
	mux.Handle("/updateUserInfo", jwt.AuthMiddleware(http.HandlerFunc(function.ModifyUserInfo)))

	mux.Handle("/getComments", jwt.AuthMiddleware(http.HandlerFunc(function.GetComments)))
	mux.Handle("/addComment", jwt.AuthMiddleware(http.HandlerFunc(function.AddComment)))
	mux.Handle("/deleteComment", jwt.AuthMiddleware(http.HandlerFunc(function.DeleteComment)))
	mux.Handle("/likePost", jwt.AuthMiddleware(http.HandlerFunc(function.AddLikeToPost)))
	mux.Handle("/cancelLikePost", jwt.AuthMiddleware(http.HandlerFunc(function.CancelLikePost)))
	mux.Handle("/checkLikeStatus", jwt.AuthMiddleware(http.HandlerFunc(function.CheckLikeStatus)))
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
	http.ListenAndServe(":8888", restful.CorsMiddleware(mux)) // 启动服务器
}
