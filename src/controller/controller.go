package controller

import (
	"jwtwithgin/src/dto"
	"jwtwithgin/src/service"

	"github.com/gin-gonic/gin"
)

//login controller interface
type LoginController interface {
	Login(ctx *gin.Context) string
}

type loginController struct {
	loginService service.LoginService
	jwtService   service.JWTService
}

func LoginHandler(loginService service.LoginService, jwtService service.JWTService) LoginController {
	return &loginController{
		loginService: loginService,
		jwtService:   jwtService,
	}
}

func (c *loginController) Login(ctx *gin.Context) string {
	var credential dto.LoginCredentials
	err := ctx.ShouldBind(&credential)
	if err != nil {
		return "no data found"
	}
	isAuthenticated := c.loginService.LoginUser(credential.Email, credential.Password)
	if isAuthenticated {
		return c.jwtService.GenerateToken(credential.Email, true)
	}
	return ""
}
