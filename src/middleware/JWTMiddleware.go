package middleware

import (
	"jwtwithgin/src/service"
	"net/http"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the client secret key
		const BEARER_SCHEMA = "Bearer"
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(BEARER_SCHEMA):]
		token, err:= service.NewJWTService().ValidateToken(tokenString)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			fmt.Println( claims)
		} else {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized})
		}
	}