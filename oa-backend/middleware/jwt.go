package middleware

import (
	"oa-backend/utils/msg"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var Salt = []byte("PussyJuice")

const TokenExpireDuration = time.Hour * 12

type MyClaims struct {
	EmployeeID int
	jwt.StandardClaims
}

func GenerateToken(employeeID int) (string, int) {
	claims := MyClaims{
		employeeID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "admin",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(Salt))
	if err != nil {
		return "", msg.FAIL
	}
	return tokenStr, msg.SUCCESS
}

func ParseToken(tokenStr string) (*MyClaims, int) {
	token, err := jwt.ParseWithClaims(tokenStr, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return Salt, nil
	})
	if err != nil {
		return nil, msg.ERROR
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid { // 校验token
		return claims, msg.SUCCESS
	}
	return nil, msg.FAIL
}

func CheckToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Token放在Header的token中
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			// 处理 没有token的时候
			msg.Message(c, msg.FAIL, "登录失效")
			c.Abort() // 不会继续停止
			return
		}
		// 解析
		mc, code := ParseToken(token)
		if code != msg.SUCCESS {
			// 处理 解析失败
			msg.Message(c, msg.FAIL, "登录失效")
			c.Abort()
			return
		}
		c.Set("employeeID", mc.EmployeeID)
		c.Next()
	}
}
