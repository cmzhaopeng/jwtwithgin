package main

import (
	"fmt"
	"jwtwithgin/src/controller"
	"jwtwithgin/src/service"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func main() {
	var loginsService service.LoginService = service.StaticLoginService()
	var jwtService service.JWTService = service.JWTAuthService()
	var loginController controller.LoginController = controller.LoginHandler(loginsService, jwtService)

	server := gin.New()
	server.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{"token": token})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})

	//server.Use(middleware.AuthorizeJWT())

	server.GET("/home", func(ctx *gin.Context) {
		fmt.Printf("home called")
		const BEARER_SCHEMA = "Bearer"
		authHeader := ctx.GetHeader("Authorization")
		tokenString := authHeader[len(BEARER_SCHEMA):]
		token, err := service.JWTAuthService().ValidateToken(tokenString)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			fmt.Println(claims)
			//user := claims["user"].(string)
			ctx.String(http.StatusOK, "Welcome %s", claims["name"])
		} else {
			fmt.Println(err)
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}

		//claims := ctx.MustGet("claims").(jwt.MapClaims)
		//user := ctx.MustGet("Authorization").(string)

		//user := claims["user"].(string)
		//fmt.Println(user)
		//ctx.JSON(http.StatusOK, gin.string(user))
	})

	port := 8080
	server.Run(fmt.Sprintf(":%d", port))
}
