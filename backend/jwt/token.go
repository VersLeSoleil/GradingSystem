package jwt

import (
	"backend/restful"
	"encoding/json"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKeyAccess = []byte(os.Getenv("JWT_SECRET_ACCESS"))
var jwtKeyRefresh = []byte(os.Getenv("JWT_SECRET_REFRESH"))

const (
	AccessTokenExpireDuration  = 15 * time.Minute
	RefreshTokenExpireDuration = 7 * 24 * time.Hour
)

func GenerateTokens(username string, userrole string) (accessToken string, refreshToken string, err error) {
	accessClaims := jwt.MapClaims{
		"username": username,
		"role":     userrole,
		"exp":      time.Now().Add(AccessTokenExpireDuration).Unix(),
	}
	accessTokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessToken, err = accessTokenObj.SignedString(jwtKeyAccess)
	if err != nil {
		return
	}

	refreshClaims := jwt.MapClaims{
		"username": username,
		"role":     userrole,
		"exp":      time.Now().Add(RefreshTokenExpireDuration).Unix(),
	}
	refreshTokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshToken, err = refreshTokenObj.SignedString(jwtKeyRefresh)
	return
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			restful.RespondWithError(w, http.StatusUnauthorized, "未提供令牌")
			return
		}
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKeyAccess, nil
		})
		if err != nil || !token.Valid {
			restful.RespondWithError(w, http.StatusUnauthorized, "令牌无效")
			return
		}
		// 可以把 claims 存到 context 里，供后续 handler 使用
		next.ServeHTTP(w, r)
	})
}

func RefreshHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("refresh_token")
	if err != nil {
		http.Error(w, "No refresh token", http.StatusUnauthorized)
		return
	}

	refreshToken := cookie.Value
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKeyRefresh, nil
	})

	if err != nil || !token.Valid {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	username, _ := claims["username"].(string)
	role, _ := claims["role"].(string)

	accessToken, _, err := GenerateTokens(username, role)
	if err != nil {
		http.Error(w, "Token generation failed", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"access_token": accessToken,
	})
}
