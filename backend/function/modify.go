package function

import (
	"backend/db"
	"backend/restful"
	"backend/structTypes"
	"encoding/json"
	"net/http"
	"log"
)

func ModifyUserInfo(w http.ResponseWriter, r *http.Request) {
	restful.SetCorsHeaders(w)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		restful.RespondWithError(w, http.StatusMethodNotAllowed, "只支持 POST 请求")
		return
	}

	var user structTypes.UserInfo

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Println("JSON解析失败:", err)
		restful.RespondWithError(w, http.StatusBadRequest, "请求体不是合法的 JSON 格式")
		return
	}

	if err := db.UpdateUserInfo(user); err != nil {
		restful.RespondWithError(w, http.StatusInternalServerError, "注册失败: "+err.Error())
		return
	}

	restful.RespondWithSuccess(w, "修改成功")
}

