package db

import (
	"backend/structTypes"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// 注册新用户
func CreateUser(tempUser structTypes.UserRegister) error {
	hash, err := HashPassword(tempUser.Password)
	if err != nil {
		return err
	}
	query := `INSERT INTO user_table (user_name, user_password_hash) 
			  VALUES (?, ?)`

	_, err = DB.Exec(query, tempUser.UserName, hash)
	return err
}

// 根据用户名获取用户信息
func GetUserByUsername(username string) (*structTypes.User, error) {
	query := `SELECT user_name, user_password_hash, sex, birthday, avatar, role, email, phone, resume, created_date FROM user_table WHERE user_name = ?`
	row := DB.QueryRow(query, username)
	var u structTypes.User
	err := row.Scan(&u.UserName, &u.PasswordHash, &u.Sex, &u.Birthday, &u.Avatar, &u.Role, &u.Email, &u.Phone, &u.Resume, &u.CreatedDate)
	if err != nil {
		fmt.Println("GetUserByUsername error:", err)
		return nil, err
	}
	return &u, nil
}

// 获取用户登录信息用于验证
func GetUserLoginInfo(username string) (*structTypes.UserLogin, error) {
	var user structTypes.UserLogin
	err := DB.QueryRow("SELECT user_name, user_password_hash FROM user_table WHERE user_name = ?", username).
		Scan(&user.UserName, &user.Password)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// 设置密码重置 token
func SetResetToken(username, token string) error {
	query := `UPDATE user_table SET reset_token = ?, reset_date = NOW() WHERE user_name = ?`
	_, err := DB.Exec(query, token, username)
	return err
}

// 验证重置 token
func ValidateResetToken(username, token string) (bool, error) {
	query := `SELECT reset_token, reset_date FROM user_table WHERE user_name = ?`
	var dbToken sql.NullString
	var dbDate sql.NullTime

	err := DB.QueryRow(query, username).Scan(&dbToken, &dbDate)
	if err != nil {
		return false, err
	}

	// 例如验证是否过期（30分钟）
	if dbToken.Valid && dbDate.Valid && dbToken.String == token {
		if time.Since(dbDate.Time) <= 30*time.Minute {
			return true, nil
		}
	}
	return false, nil
}

// 修改密码（在用户提供 token 验证通过之后）
func UpdatePassword(username, newPasswordHash string) error {
	query := `UPDATE user_table SET user_password_hash = ?, reset_token = NULL, reset_date = NULL WHERE user_name = ?`
	_, err := DB.Exec(query, newPasswordHash, username)
	return err
}

// 修改用户信息
func UpdateUserInfo(tempUser structTypes.UserInfo) error {
	query := `UPDATE user_table SET sex = ?, birthday = ?, avatar = ?, email = ?, phone = ?, resume = ? WHERE user_name = ?`
	_, err := DB.Exec(query, tempUser.Sex, tempUser.Birthday, tempUser.Avatar, tempUser.Email, tempUser.Phone, tempUser.Resume, tempUser.UserName)
	return err
}
