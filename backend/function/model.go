package function
import (
	"bytes"
	"io"
	"net/http"
	"mime/multipart"
)
func PredictHandler(w http.ResponseWriter, r *http.Request) {
	// 限制最大上传大小（比如 10MB）
	r.ParseMultipartForm(10 << 20)

	// 获取前端上传的文件
	imageFile, imageHeader, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "图像上传失败: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer imageFile.Close()

	maskFile, maskHeader, err := r.FormFile("mask")
	if err != nil {
		http.Error(w, "掩膜上传失败: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer maskFile.Close()

	// 创建 multipart/form-data 请求，转发给 Python Flask 模型服务
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	imagePart, err := writer.CreateFormFile("image", imageHeader.Filename)
	if err != nil {
		http.Error(w, "创建图像字段失败", http.StatusInternalServerError)
		return
	}
	io.Copy(imagePart, imageFile)

	maskPart, err := writer.CreateFormFile("mask", maskHeader.Filename)
	if err != nil {
		http.Error(w, "创建掩膜字段失败", http.StatusInternalServerError)
		return
	}
	io.Copy(maskPart, maskFile)

	writer.Close()

	// 调用 Python 模型服务
	resp, err := http.Post("http://localhost:5000/predict", writer.FormDataContentType(), &buf)
	if err != nil {
		http.Error(w, "调用模型失败: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// 设置响应类型
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)

	// 把 Python 的响应转发给前端
	io.Copy(w, resp.Body)
}
