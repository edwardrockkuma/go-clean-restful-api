package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthorizationMiddleware(ctx *gin.Context) {
	authHeader := ctx.Request.Header.Get("Authorization")
	jToken := strings.TrimPrefix(authHeader, "Bearer ")
	err := validateJwt(jToken)

	if err != nil {
		ctx.Copy().AbortWithStatus(http.StatusUnauthorized)
		return
	}
}

func validateJwt(jToken string) error {
	_, err := jwt.Parse(jToken, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if ok == false {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte("secret"), nil
	})

	return err
}

func LoginHandle(ctx *gin.Context) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
	})

	ss, err := token.SignedString([]byte("secret"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": ss,
	})
}
