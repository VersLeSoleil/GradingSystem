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
	query := `SELECT post_id, type_name, user_name, title, introduction, content, model_types=?, created_date, updated_date, is_public, likes 
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
		if err := rows.Scan(&p.PostID, &p.TypeName, &p.UserName, &p.Title, &p.Introduction, &p.ModelTypes, &p.Content,
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
