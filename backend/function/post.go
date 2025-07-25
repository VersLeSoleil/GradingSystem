package function

import (
	"backend/db"
	"backend/restful"
	"backend/structTypes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetPostByID(w http.ResponseWriter, r *http.Request) {
	restful.SetCorsHeaders(w)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodGet {
		restful.RespondWithError(w, http.StatusMethodNotAllowed, "只支持 GET 请求")
		return
	}

	postIDStr := r.URL.Query().Get("post_id")
	if postIDStr == "" {
		restful.RespondWithError(w, http.StatusBadRequest, "缺少 post_id 参数")
		return
	}

	var postID int
	if _, err := fmt.Sscanf(postIDStr, "%d", &postID); err != nil {
		restful.RespondWithError(w, http.StatusBadRequest, "post_id 参数必须为整数")
		return
	}

	post, err := db.GetPostByID(postID)
	if err != nil {
		log.Printf("获取帖子失败: %v", err)
		restful.RespondWithError(w, http.StatusInternalServerError, "获取帖子失败")
		return
	}
	if post == nil {
		restful.RespondWithError(w, http.StatusNotFound, "帖子不存在")
		return
	}
	restful.RespondWithSuccess(w, post)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
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

	var post structTypes.PostInfo
	if err := json.Unmarshal(bodyBytes, &post); err != nil {
		log.Printf("JSON 解析失败: %v", err)
		restful.RespondWithError(w, http.StatusBadRequest, "请求体不是合法的 JSON 格式")
		return
	}
	fmt.Println("创建帖子:", post)
	if err := db.CreatedPost(post); err != nil {
		log.Printf("创建帖子失败: %v", err)
		restful.RespondWithError(w, http.StatusInternalServerError, "创建帖子失败")
		return
	}

	restful.RespondWithSuccess(w, "帖子创建成功")
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	restful.SetCorsHeaders(w)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodDelete {
		restful.RespondWithError(w, http.StatusMethodNotAllowed, "只支持 DELETE 请求")
		return
	}

	postIDStr := r.URL.Query().Get("post_id")
	if postIDStr == "" {
		restful.RespondWithError(w, http.StatusBadRequest, "缺少 post_id 参数")
		return
	}

	var postID int
	if _, err := fmt.Sscanf(postIDStr, "%d", &postID); err != nil {
		restful.RespondWithError(w, http.StatusBadRequest, "post_id 参数必须为整数")
		return
	}

	if err := db.DeletePostByID(postID); err != nil {
		log.Printf("删除帖子失败: %v", err)
		restful.RespondWithError(w, http.StatusInternalServerError, "删除帖子失败")
		return
	}

	restful.RespondWithSuccess(w, "帖子删除成功")
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	restful.SetCorsHeaders(w)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPut {
		restful.RespondWithError(w, http.StatusMethodNotAllowed, "只支持 PUT 请求")
		return
	}

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("无法读取请求体: %v", err)
		restful.RespondWithError(w, http.StatusBadRequest, "无法读取请求体")
		return
	}

	var post structTypes.Post
	if err := json.Unmarshal(bodyBytes, &post); err != nil {
		log.Printf("JSON 解析失败: %v", err)
		restful.RespondWithError(w, http.StatusBadRequest, "请求体不是合法的 JSON 格式")
		return
	}

	if err := db.UpdatePostByID(post.PostID, post); err != nil {
		log.Printf("更新帖子失败: %v", err)
		restful.RespondWithError(w, http.StatusInternalServerError, "更新帖子失败")
		return
	}

	restful.RespondWithSuccess(w, "帖子更新成功")
}

func GetPostsByUsername(w http.ResponseWriter, r *http.Request) {
	restful.SetCorsHeaders(w)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodGet {
		restful.RespondWithError(w, http.StatusMethodNotAllowed, "只支持 GET 请求")
		return
	}

	username := r.URL.Query().Get("username")
	if username == "" {
		restful.RespondWithError(w, http.StatusBadRequest, "缺少 username 参数")
		return
	}

	posts, err := db.GetPostByUsername(username)
	if err != nil {
		log.Printf("获取用户帖子失败: %v", err)
		restful.RespondWithError(w, http.StatusInternalServerError, "获取用户帖子失败")
		return
	}

	if len(posts) == 0 {
		restful.RespondWithError(w, http.StatusNotFound, "没有找到该用户的帖子")
		return
	}
	restful.RespondWithSuccess(w, posts)
}

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	restful.SetCorsHeaders(w)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodGet {
		restful.RespondWithError(w, http.StatusMethodNotAllowed, "只支持 GET 请求")
		return
	}
	
	posts, err := db.GetAllPosts()
	if err != nil {
		log.Printf("获取所有帖子失败: %v", err)
		restful.RespondWithError(w, http.StatusInternalServerError, "获取所有帖子失败")
		return
	}

	if len(posts) == 0 {
		restful.RespondWithError(w, http.StatusNotFound, "没有找到任何帖子")
		return
	}
	restful.RespondWithSuccess(w, posts)
}


func GetComments(w http.ResponseWriter, r *http.Request) {
	restful.SetCorsHeaders(w)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodGet {
		restful.RespondWithError(w, http.StatusMethodNotAllowed, "只支持 GET 请求")
		return
	}

	postIDStr := r.URL.Query().Get("post_id")
	if postIDStr == "" {
		restful.RespondWithError(w, http.StatusBadRequest, "缺少 post_id 参数")
		return
	}

	var postID int
	if _, err := fmt.Sscanf(postIDStr, "%d", &postID); err != nil {
		restful.RespondWithError(w, http.StatusBadRequest, "post_id 参数必须为整数")
		return
	}

	comments, err := db.GetCommentsByPostID(postID)
	if err != nil {
		log.Printf("获取帖子评论失败: %v", err)
		restful.RespondWithError(w, http.StatusInternalServerError, "获取帖子评论失败")
		return
	}

	if len(comments) == 0 {
		restful.RespondWithError(w, http.StatusNotFound, "没有找到该帖子的评论")
		return
	}
	
	restful.RespondWithSuccess(w, comments)
}

func AddComment(w http.ResponseWriter, r *http.Request) {
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

	var comment structTypes.Comment
	if err := json.Unmarshal(bodyBytes, &comment); err != nil {
		log.Printf("JSON 解析失败: %v", err)
		restful.RespondWithError(w, http.StatusBadRequest, "请求体不是合法的 JSON 格式")
		return
	}

	if err := db.AddCommentToPost(comment); err != nil {
		log.Printf("添加评论失败: %v", err)
		restful.RespondWithError(w, http.StatusInternalServerError, "添加评论失败")
		return
	}

	restful.RespondWithSuccess(w, "评论添加成功")
}

func DeleteComment(w http.ResponseWriter, r *http.Request) {
	restful.SetCorsHeaders(w)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodDelete {
		restful.RespondWithError(w, http.StatusMethodNotAllowed, "只支持 DELETE 请求")
		return
	}

	commentIDStr := r.URL.Query().Get("comment_id")
	if commentIDStr == "" {
		restful.RespondWithError(w, http.StatusBadRequest, "缺少 comment_id 参数")
		return
	}

	var commentID int
	if _, err := fmt.Sscanf(commentIDStr, "%d", &commentID); err != nil {
		restful.RespondWithError(w, http.StatusBadRequest, "comment_id 参数必须为整数")
		return
	}

	if err := db.DeleteCommentByID(commentID); err != nil {
		log.Printf("删除评论失败: %v", err)
		restful.RespondWithError(w, http.StatusInternalServerError, "删除评论失败")
		return
	}

	restful.RespondWithSuccess(w, "评论删除成功")
}

func AddLikeToPost(w http.ResponseWriter, r *http.Request) {
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

	var like structTypes.Like
	if err := json.Unmarshal(bodyBytes, &like); err != nil {
		log.Printf("JSON 解析失败: %v", err)
		restful.RespondWithError(w, http.StatusBadRequest, "请求体不是合法的 JSON 格式")
		return
	}

	if err := db.AddLike(like); err != nil {
		log.Printf("添加点赞失败: %v", err)
		restful.RespondWithError(w, http.StatusInternalServerError, "添加点赞失败")
		return
	}

	restful.RespondWithSuccess(w, "点赞成功")
}

func CancelLikePost(w http.ResponseWriter, r *http.Request) {
	restful.SetCorsHeaders(w)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodDelete {
		restful.RespondWithError(w, http.StatusMethodNotAllowed, "只支持 DELETE 请求")
		return
	}

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("无法读取请求体: %v", err)
		restful.RespondWithError(w, http.StatusBadRequest, "无法读取请求体")
		return
	}

	var like structTypes.Like
	if err := json.Unmarshal(bodyBytes, &like); err != nil {
		log.Printf("JSON 解析失败: %v", err)
		restful.RespondWithError(w, http.StatusBadRequest, "请求体不是合法的 JSON 格式")
		return
	}

	if err := db.CancelLike(like); err != nil {
		log.Printf("取消点赞失败: %v", err)
		restful.RespondWithError(w, http.StatusInternalServerError, "取消点赞失败")
		return
	}

	restful.RespondWithSuccess(w, "取消点赞成功")
}

func CheckLikeStatus(w http.ResponseWriter, r *http.Request){
    restful.SetCorsHeaders(w)
    if r.Method == http.MethodOptions {
        w.WriteHeader(http.StatusOK)
        return
    }
    if r.Method != http.MethodGet {
        restful.RespondWithError(w, http.StatusMethodNotAllowed, "只支持检查点赞请求")
        return
    }

    postIDStr := r.URL.Query().Get("post_id")
    userName := r.URL.Query().Get("user_name")
    if postIDStr == "" || userName == "" {
        restful.RespondWithError(w, http.StatusBadRequest, "缺少参数")
        return
    }
    var postID int
    if _, err := fmt.Sscanf(postIDStr, "%d", &postID); err != nil {
        restful.RespondWithError(w, http.StatusBadRequest, "post_id 参数必须为整数")
        return
    }
    isLiked, err := db.IsLiked(postID, userName)
    if err != nil {
        log.Printf("查询点赞状态失败: %v", err)
        restful.RespondWithError(w, http.StatusInternalServerError, "查询点赞状态失败")
        return
    }
    restful.RespondWithSuccess(w, isLiked)
}
