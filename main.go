package main

import (
	"fmt"
	"jwtwithgin/src/controller"
	"jwtwithgin/src/middleware"
	"jwtwithgin/src/service"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	var loginsService service.LoginService = service.StaticLoginService()
	var jwtService service.JWTService = service.JWTAuthService()
	var loginController controller.LoginController = controller.LoginHandler(loginsService, jwtService)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	//Getenv CORS_SITE, and split it with comma, and assign it to allowOrigins
	corsSite := os.Getenv("CORS_SITE")
	//filter the corsSite, and remove the space
	corsSite = strings.ReplaceAll(corsSite, " ", "")
	key := os.Getenv("SESSION_KEY")
	server := gin.New()
	config := cors.DefaultConfig()
	config.AllowOrigins = strings.Split(corsSite, ",")
	//config.AllowOrigins = []string{"http://192.168.2.55", "http://localhost:3000"}
	fmt.Println(config.AllowOrigins)
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	server.Use(cors.New(config))

	store := cookie.NewStore([]byte(key))
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

	server.POST("/home", func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		fmt.Println(session.Get("name"))
		//user := claims["user"].(string)
		//ctx.String(http.StatusOK, "Welcome %s", session.Get("name"))
		ctx.JSON(http.StatusOK, gin.H{"Gouser": session.Get("name")})
		//ctx.JSON(http.StatusOK, gin.H{"Gouser": "name"})

		//claims := ctx.MustGet("claims").(jwt.MapClaims)
		//user := ctx.MustGet("Authorization").(string)

		//user := claims["user"].(string)
		//fmt.Println(user)
		//ctx.JSON(http.StatusOK, gin.string(user))
	})

	port := 8080
	server.Run(fmt.Sprintf(":%d", port))
}
