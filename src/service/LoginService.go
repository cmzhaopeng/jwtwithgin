package service

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type LoginService interface {
	LoginUser(email string, password string) bool
}

type loginInformation struct {
	email    string
	password string
}

func StaticLoginService() LoginService {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return &loginInformation{
			email:    "testing@139.com",
			password: "LiriTesJ",
		}
	}
	return &loginInformation{
		email:    "",
		password: os.Getenv("APP_KEY"),
	}
}

func (info *loginInformation) LoginUser(email string, password string) bool {
	//return info.email == email && info.password == password
	//verify the app key. This is a static key. We should verify the user stored in the database on the future.
	return info.password == password
}
