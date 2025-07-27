package function

import (
	"backend/db"
	"backend/jwt"
	"backend/restful"
	"backend/structTypes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// 登录处理函数
func LoginCheck(w http.ResponseWriter, r *http.Request) {
	log.Println("接收到登录请求")
	restful.SetCorsHeaders(w)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		restful.RespondWithError(w, http.StatusMethodNotAllowed, "只支持 POST 请求")
		return
	}

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("无法读取请求体: %v", err)
		restful.RespondWithError(w, http.StatusBadRequest, "无法读取请求体")
		return
	}

	var userL structTypes.UserLogin

	if err := json.Unmarshal(bodyBytes, &userL); err != nil {
		log.Printf("JSON 解析失败: %v", err)
		restful.RespondWithError(w, http.StatusBadRequest, "请求体不是合法的 JSON 格式")
		return
	}
	var user *structTypes.User
	user, err = db.GetUserByUsername(userL.UserName)
	fmt.Println("获取用户信息:", user)
	if err != nil || user == nil {
		log.Println("用户名不存在")
		restful.RespondWithError(w, http.StatusUnauthorized, "用户名或密码错误")
		return
	}

	if !db.CheckPasswordHash(userL.Password, user.PasswordHash) {
		log.Println("密码验证失败")
		restful.RespondWithError(w, http.StatusUnauthorized, "用户名或密码错误")
		return
	}

	accessToken, refreshToken, err := jwt.GenerateTokens(user.UserName, user.Role)
	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		HttpOnly: true,
		Secure:   false, //!!
		Path:     "/",
		SameSite: http.SameSiteStrictMode,
		MaxAge:   7 * 24 * 60 * 60,
	})
	user.PasswordHash = "" // 清除密码哈希，避免泄露
	resp := map[string]interface{}{
		"user":         user,
		"access_token": accessToken,
	}

	log.Printf("用户 %s 登录成功", userL.UserName)
	restful.RespondWithSuccess(w, resp)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
    // 设置一个过期的 Cookie 删除 refresh_token
    http.SetCookie(w, &http.Cookie{
        Name:     "refresh_token",
        Value:    "",
        Path:     "/",
        Domain:   "localhost",
        MaxAge:   -1,
        HttpOnly: true,
        Secure:   false, // 生产环境建议设为 true
    })

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"message": "logout success"}`))
}