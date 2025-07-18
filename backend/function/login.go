package function

import (
	"backend/db"
	"backend/jwt"
	"backend/restful"
	"backend/structTypes"
	"encoding/json"
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
	log.Printf("接收到的原始数据: %s", string(bodyBytes))

	var userL structTypes.UserLogin

	if err := json.Unmarshal(bodyBytes, &userL); err != nil {
		log.Printf("JSON 解析失败: %v", err)
		restful.RespondWithError(w, http.StatusBadRequest, "请求体不是合法的 JSON 格式")
		return
	}

	var user *structTypes.User

	user, err = db.GetUserByUsername(userL.UserName)

	if !db.CheckPasswordHash(userL.Password, user.PasswordHash) {
		log.Println("密码验证失败")
		restful.RespondWithError(w, http.StatusUnauthorized, "用户名或密码错误")
		return
	}

	accessToken, refreshToken, err := jwt.GenerateTokens(user.UserName, user.Role)

	resp := map[string]interface{}{
		"user":          user,
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}
	// var userInfoBack structTypes.UserInfo
	// userInfoBack = structTypes.CopyUserToUserInfo(user)
	log.Printf("用户 %s 登录成功", userL.UserName)
	restful.RespondWithSuccess(w, resp)
}
