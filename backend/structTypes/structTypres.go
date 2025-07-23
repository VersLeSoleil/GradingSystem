package structTypes

import (
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
	UserID       int
	UserName     string
	PasswordHash string
	Sex          string
	Birthday     time.Time
	Avatar       string
	Role         string
	Email        string
	Phone        string
	Resume       string
	CreatedDate  time.Time
}
