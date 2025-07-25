package db

import (
	"backend/structTypes"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func CreatedPost(post structTypes.PostInfo) error {

	query := `INSERT INTO post_table (type_name, user_name, title, introduction, content, model_types, created_date, updated_date, is_public) 
			  VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := DB.Exec(query, post.TypeName, post.UserName, post.Title, post.Introduction, post.Content, post.ModelTypes, time.Now(), time.Now(), post.IsPublic)
	if err != nil {
		return fmt.Errorf("创建帖子失败: %w", err)
	}
	return nil
}

func GetPostByUsername(username string) ([]structTypes.Post, error) {
	query := `SELECT post_id, type_name, user_name, title, introduction, content, model_types, created_date, updated_date, is_public, likes FROM post_table WHERE user_name = ?`
	rows, err := DB.Query(query, username)
	if err != nil {
		return nil, fmt.Errorf("获取用户帖子失败: %w", err)
	}
	defer rows.Close()

	var posts []structTypes.Post
	for rows.Next() {
		var p structTypes.Post
		if err := rows.Scan(&p.PostID, &p.TypeName, &p.UserName, &p.Title, &p.Content, &p.ModelTypes, &p.CreatedDate, &p.UpdatedDate, &p.IsPublic, &p.Likes); err != nil {
			return nil, fmt.Errorf("扫描帖子数据失败: %w", err)
		}
		posts = append(posts, p)
	}
	return posts, nil
}

func UpdatePostByID(postID int, post structTypes.Post) error {
	query := `UPDATE post_table SET type_name = ?, user_name = ?, title = ?, introduction=?, content = ?, model_types=?, updated_date = ?, is_public = ? WHERE post_id = ?`
	_, err := DB.Exec(query, post.TypeName, post.UserName, post.Title, post.Introduction, post.Content, post.ModelTypes, time.Now(), post.IsPublic, postID)
	if err != nil {
		return fmt.Errorf("修改帖子失败: %w", err)
	}
	return nil
}

func DeletePostByID(postID int) error {
	query := `DELETE FROM post_table WHERE post_id = ?`
	_, err := DB.Exec(query, postID)
	if err != nil {
		return fmt.Errorf("删除帖子失败: %w", err)
	}
	return nil
}

func GetPostByID(postID int) (*structTypes.Post, error) {
	query := `SELECT post_id, type_name, user_name, title, introduction, content, model_types, created_date, updated_date, is_public, likes 
			  FROM post_table WHERE post_id = ?`
	row := DB.QueryRow(query, postID)

	var p structTypes.Post
	err := row.Scan(&p.PostID, &p.TypeName, &p.UserName, &p.Title, &p.Introduction, &p.Content, &p.ModelTypes, &p.CreatedDate, &p.UpdatedDate, &p.IsPublic, &p.Likes)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // 帖子不存在
		}
		return nil, fmt.Errorf("获取帖子失败: %w", err)
	}
	return &p, nil
}

func GetAllPosts() ([]structTypes.Post, error) {
	query := `SELECT post_id, type_name, user_name, title, introduction, content, model_types, created_date, updated_date, is_public, likes 
              FROM post_table 
              ORDER BY created_date DESC`

	rows, err := DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("获取所有帖子失败: %w", err)
	}
	defer rows.Close()

	var posts []structTypes.Post
	for rows.Next() {
		var p structTypes.Post
		if err := rows.Scan(&p.PostID, &p.TypeName, &p.UserName, &p.Title, &p.Introduction, &p.Content, &p.ModelTypes,
			&p.CreatedDate, &p.UpdatedDate, &p.IsPublic, &p.Likes); err != nil {
			return nil, fmt.Errorf("扫描帖子数据失败: %w", err)
		}
		posts = append(posts, p)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("处理行数据时出错: %w", err)
	}

	return posts, nil
}

func GetCommentsByPostID(postID int) ([]structTypes.Comment, error) {
	query := `SELECT comment_id, post_id, user_name, content, comment_time FROM post_comment_table WHERE post_id = ?`
	rows, err := DB.Query(query, postID)
	if err != nil {
		return nil, fmt.Errorf("获取帖子评论失败: %w", err)
	}
	defer rows.Close()

	var comments []structTypes.Comment
	for rows.Next() {
		var c structTypes.Comment
		if err := rows.Scan(&c.CommentID, &c.PostID, &c.UserName, &c.Content, &c.CommentTime); err != nil {
			return nil, fmt.Errorf("扫描评论数据失败: %w", err)
		}
		comments = append(comments, c)
	}
	return comments, nil
}

func AddCommentToPost(comment structTypes.Comment) error {
	query := `INSERT INTO post_comment_table (post_id, user_name, content, comment_time) VALUES (?, ?, ?, ?)`
	_, err := DB.Exec(query, comment.PostID, comment.UserName, comment.Content, time.Now())
	if err != nil {
		return fmt.Errorf("添加评论失败: %w", err)
	}
	return nil
}

func DeleteCommentByID(commentID int) error {
	query := `DELETE FROM post_comment_table WHERE comment_id = ?`
	_, err := DB.Exec(query, commentID)
	if err != nil {
		return fmt.Errorf("删除评论失败: %w", err)
	}
	return nil
}

func AddLike(like structTypes.Like) error {
	query := `INSERT INTO post_like_table (post_id, user_name, liked_date) VALUES (?, ?, ?)`
	_, err := DB.Exec(query, like.PostID, like.UserName, time.Now())
	if err != nil {
		return fmt.Errorf("添加点赞失败: %w", err)
	}
	return nil
}

func CancelLike(like structTypes.Like) error {
	query := `DELETE FROM post_like_table WHERE post_id = ? AND user_name = ?`
	_, err := DB.Exec(query, like.PostID, like.UserName)
	if err != nil {
		return fmt.Errorf("取消点赞失败: %w", err)
	}
	return nil
}

func IsLiked(postID int, username string) (structTypes.IsLiked, error) {
	query := `SELECT COUNT(*) FROM post_like_table WHERE post_id = ? AND user_name = ?`
	var likestatus structTypes.IsLiked
	err := DB.QueryRow(query, postID, username).Scan(&likestatus.Liked)
	err = DB.QueryRow(`SELECT COUNT(*) FROM post_like_table WHERE post_id = ?`, postID).Scan(&likestatus.LikeCount)
	if err != nil {
		return structTypes.IsLiked{}, fmt.Errorf("查询点赞状态失败: %w", err)
	}
	return likestatus, nil
}