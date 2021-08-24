package services

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"invest/app/config"
	"invest/app/models"
	"invest/app/utils"
	"strings"
	"time"
)

type AuthService struct {}

func (l *AuthService) AuthenticateUser (input utils.LoginInput) (models.User, error){
	email := input.Email
	password := input.Password

	var userService = UserService{}

	user, err := userService.FindOneByEmail(email)
	if err != nil {
		return user, errors.New("Email not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err!= nil{
		return user, err
	}
	return user,nil
}

func (l *AuthService) AuthenticateAdmin (input utils.AdminLoginInput) (models.Admin, error){
	username := input.Username
	password := input.Password

	var adminService = AdminService{}

	admin, err := adminService.FindOneByUsername(username)
	if err != nil {
		return admin, errors.New("Username not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password))
	if err!= nil{
		return admin, err
	}
	return admin,nil
}

func (l *AuthService) CreateToken(user models.User) (string, error) {
	payload := jwt.MapClaims{}

	payload["name"] = user.Name
	payload["id"] = user.ID
	payload["email"] = user.Email
	payload["phonenumber"] = user.Phonenumber
	payload["display_pict"] = user.DisplayPict
	payload["is_verified"] = user.IsVerified
	payload["role"] = "user"
	payload["exp"] = time.Now().Add(time.Hour * 36).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := at.SignedString([]byte(config.AppConfig.JwtKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (l *AuthService) CreateAdminToken(admin models.Admin) (string, error) {
	payload := jwt.MapClaims{}

	payload["name"] = admin.Name
	payload["id"] = admin.ID
	payload["email"] = admin.Email
	payload["username"] = admin.Username
	payload["role"] = "admin"
	payload["hak_akses"] = admin.HakAkses
	payload["exp"] = time.Now().Add(time.Hour * 36).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := at.SignedString([]byte(config.AppConfig.JwtKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (l *AuthService) ExtractToken(bearToken string) string {
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func (l *AuthService) VerifyToken (tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.AppConfig.JwtKey), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}