package main

import (
	"fmt"
	"jwtwithgin/src/controller"
	"jwtwithgin/src/middleware"
	"jwtwithgin/src/service"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	var loginsService service.LoginService = service.StaticLoginService()
	var jwtService service.JWTService = service.JWTAuthService()
	var loginController controller.LoginController = controller.LoginHandler(loginsService, jwtService)

	server := gin.New()
	store := cookie.NewStore([]byte("secret-key-897"))
	server.Use(sessions.Sessions("mysession", store))

	server.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{"token": token})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})

	server.Use(middleware.AuthorizeJWT())

	server.GET("/home", func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		fmt.Println(session.Get("name"))
		//user := claims["user"].(string)
		ctx.String(http.StatusOK, "Welcome %s", session.Get("name"))

		//claims := ctx.MustGet("claims").(jwt.MapClaims)
		//user := ctx.MustGet("Authorization").(string)

		//user := claims["user"].(string)
		//fmt.Println(user)
		//ctx.JSON(http.StatusOK, gin.string(user))
	})

	port := 8080
	server.Run(fmt.Sprintf(":%d", port))
}
