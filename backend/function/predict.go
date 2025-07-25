package function

import (
	"backend/restful"
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)


func HandelPredict(w http.ResponseWriter, r *http.Request) {
	restful.SetCorsHeaders(w)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		restful.RespondWithError(w, http.StatusMethodNotAllowed, "只支持 POST 请求")
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		restful.RespondWithError(w, http.StatusBadRequest, "文件上传失败: "+err.Error())
		return
	}
	defer file.Close()

	// 2. 保存临时图像文件
    tempFile, err := os.CreateTemp("", "upload-*.jpg")
    io.Copy(tempFile, file)
    tempFile.Close()

    // 3. 将图像发给 Python 推理服务
    body := &bytes.Buffer{}
    writer := multipart.NewWriter(body)
    imageFile, _ := os.Open(tempFile.Name())
    part, _ := writer.CreateFormFile("image", filepath.Base(tempFile.Name()))
    io.Copy(part, imageFile)
    writer.Close()

    resp, err := http.Post("http://localhost:5000/predict", writer.FormDataContentType(), body)
    if err != nil {
        http.Error(w, "调用模型失败", 500)
        return
    }
    defer resp.Body.Close()

    // 4. 返回 Python 服务的预测结果给前端
    io.Copy(w, resp.Body)
}