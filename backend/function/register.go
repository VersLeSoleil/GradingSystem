package function

import (
	"backend/db"
	"backend/restful"
	"backend/structTypes"
	"encoding/json"
	"fmt"
	"net/http"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	restful.SetCorsHeaders(w)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		restful.RespondWithError(w, http.StatusMethodNotAllowed, "只支持 POST 请求")
		return
	}

	var user structTypes.UserRegister

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		restful.RespondWithError(w, http.StatusBadRequest, "请求体不是合法的 JSON 格式")
		return
	}
	//哈希加密注册密码
	fmt.Println("注册用户:",user.Password)
	hashedPwd, err := db.HashPassword(user.Password)
	if err != nil {
		restful.RespondWithError(w, http.StatusInternalServerError, "注册时密码加密失败: "+err.Error())
		return
	}
	user.Password = hashedPwd

	if err := db.CreateUser(user); err != nil {
		restful.RespondWithError(w, http.StatusInternalServerError, "注册失败: "+err.Error())
		return
	}

	restful.RespondWithSuccess(w, "注册成功")
}
