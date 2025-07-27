package function

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"backend/restful"
)

func PredictHandler(w http.ResponseWriter, r *http.Request) {

	restful.SetCorsHeaders(w)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	if r.Method != http.MethodPost {
		restful.RespondWithError(w, http.StatusMethodNotAllowed, "只支持 POST 请求")
		return
	}

	// 限制上传大小
	err := r.ParseMultipartForm(100 << 20)
	if err != nil {
		http.Error(w, "解析表单失败: "+err.Error(), http.StatusBadRequest)
		return
	}

	images := r.MultipartForm.File["images"]
	masks := r.MultipartForm.File["masks"]

	if len(images) != len(masks) {
		http.Error(w, "上传的图像和掩膜数量不一致", http.StatusBadRequest)
		return
	}

	// 创建转发用的 multipart 表单
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	for i := 0; i < len(images); i++ {
		imageFile, err := images[i].Open()
		if err != nil {
			http.Error(w, "读取图像失败: "+err.Error(), http.StatusInternalServerError)
			return
		}
		defer imageFile.Close()

		maskFile, err := masks[i].Open()
		if err != nil {
			http.Error(w, "读取掩膜失败: "+err.Error(), http.StatusInternalServerError)
			return
		}
		defer maskFile.Close()

		// 每对图像和掩膜以序号命名发送给 Python 端
		imageField := fmt.Sprintf("images")
		maskField := fmt.Sprintf("masks")

		imagePart, err := writer.CreateFormFile(imageField, images[i].Filename)
		if err != nil {
			http.Error(w, "创建图像字段失败", http.StatusInternalServerError)
			return
		}
		io.Copy(imagePart, imageFile)

		maskPart, err := writer.CreateFormFile(maskField, masks[i].Filename)
		if err != nil {
			http.Error(w, "创建掩膜字段失败", http.StatusInternalServerError)
			return
		}
		io.Copy(maskPart, maskFile)

		fmt.Println("添加到转发请求中：", images[i].Filename, masks[i].Filename)
	}

	writer.Close()

	// 发给 Python 模型服务
	resp, err := http.Post("http://localhost:5000/predict", writer.FormDataContentType(), &buf)
	if err != nil {
		http.Error(w, "调用模型失败: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()


	// 设置响应头并返回 Python 响应
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}
