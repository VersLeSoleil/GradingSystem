package structTypes

import (
	"database/sql"
	"time"
)

type Config struct {
	DatabaseConnection struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"database_connection"`
	DatabaseNames struct {
		LocalSql string `yaml:"LocalSql"`
	} `yaml:"database_names"`
}

// 用户登录信息结构体
type UserLogin struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type UserRegister struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type UserInfo struct {
	UserName    string    `json:"user_name"`
	Password    string    `json:"password"`
	Sex         string    `json:"sex"`
	Birthday    string    `json:"birthday"`
	Avatar      string    `json:"avatar"`
	Role        string    `json:"role"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	Resume      string    `json:"resume"`
	CreatedDate time.Time `json:"created_date"`
}

// 用户信息结构体
type User struct {
	UserName     string
	PasswordHash string
	Sex          string
	Birthday     time.Time
	Avatar       sql.NullString
	Role         string
	Email        sql.NullString
	Phone        sql.NullString
	Resume       sql.NullString
	CreatedDate  time.Time
}

type PostInfo struct {
	TypeName     string    `json:"type_name"`
	UserName     string    `json:"user_name"`
	Title        string    `json:"title"`
	Introduction string    `json:"introduction"`
	Content      string    `json:"content"`
	ModelTypes   string    `json:"model_types"`
	CreatedDate  time.Time `json:"created_date"`
	UpdatedDate  time.Time `json:"updated_date"`
	IsPublic     bool      `json:"is_public"`
}

type Post struct {
	PostID       int       `json:"post_id"`
	TypeName     string    `json:"type_name"`
	UserName     string    `json:"user_name"`
	Title        string    `json:"title"`
	Introduction string    `json:"introduction"`
	Content      string    `json:"content"`
	ModelTypes   string    `json:"model_types"`
	CreatedDate  time.Time `json:"created_date"`
	UpdatedDate  time.Time `json:"updated_date"`
	IsPublic     bool      `json:"is_public"`
	Likes        int       `json:"likes"`
}

type Comment struct {
	CommentID   int       `json:"comment_id"`
	PostID      int       `json:"post_id"`
	UserName    string    `json:"user_name"`
	Content     string    `json:"content"`
	CommentTime time.Time `json:"comment_time"`
}

type Like struct {
	PostID    int       `json:"post_id"`
	UserName  string    `json:"user_name"`
	LikedDate time.Time `json:"liked_date"`
}

type IsLiked struct {
	Liked    bool `json:"liked"`
	LikeCount int  `json:"like_count"`
}
