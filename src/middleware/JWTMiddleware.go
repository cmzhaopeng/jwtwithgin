package middleware

import (
	"fmt"
	"jwtwithgin/src/service"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the client secret key
		fmt.Println("AuthorizeJWT")
		fmt.Println("--------------------------------------")

		session := sessions.Default(c)
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")
		fmt.Println(authHeader)
		tokenString := authHeader[len(BEARER_SCHEMA):]
		fmt.Println(tokenString)
		token, err := service.JWTAuthService().ValidateToken(tokenString)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			username := claims["name"].(string)
			session.Set("name", username)
			session.Save()
			fmt.Println(claims)
		} else {
			fmt.Println(err)
			fmt.Print("Invalid token")
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
