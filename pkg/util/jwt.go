package util

import (
	"strings"
	"time"

	"gin-gonic/models/sql"
	"gin-gonic/pkg/setting"

	jwt "github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte(setting.JwtSecret)

// Claims is a struct to map a jwt token
type Claims struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	FirsName string `json:"first_name"`
	LastName string `json:"last_name"`
	jwt.StandardClaims
}

// GenerateToken Generate a new token
func GenerateToken(user sql.User) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour * 365)

	claims := Claims{
		user.ID,
		user.Email,
		user.FirstName,
		user.LastName,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// ParseToken return decoded JWT
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

// GetToken return a string with JWT
func GetToken(c []string) (token string) {
	token = c[0]
	splitToken := strings.Split(token, "Bearer")
	token = strings.TrimSpace(splitToken[1])
	return
}
