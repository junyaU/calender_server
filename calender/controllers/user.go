package controllers

import (
	"calender/models"
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) CreateUser() {
	name := this.GetString("name")
	email := this.GetString("email")
	password := this.GetString("password")
	o := orm.NewOrm()

	hash, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	hashPassword := string(hash)

	user := models.User{}
	user.Name = name
	user.Email = email
	user.Password = hashPassword
	if _, err := o.Insert(&user); err != nil {
		return
	}

	token := CreateJwt(user.Id, user.Name)
	this.Data["json"] = token
	this.ServeJSON()
}

func (this *UserController) Login() {
	email := this.GetString("email")
	password := this.GetString("password")

	user := models.User{Email: email}
	o := orm.NewOrm()
	if err := o.Read(&user, "Email"); err != nil {
		return
	}
	passwordError := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if passwordError != nil {
		return
	}

	token := CreateJwt(user.Id, user.Name)
	this.Data["json"] = token
	this.ServeJSON()
}

func (this *UserController) EmailCheck() {
	email := this.Ctx.Input.Param(":email")
	user := models.User{Email: email}
	o := orm.NewOrm()

	if err := o.Read(&user, "Email"); err != nil {
		this.Data["json"] = "ok"
		this.ServeJSON()
		return
	}
}

func CreateJwt(id int64, name string) string {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = id
	claims["iss"] = beego.AppConfig.String("appname")
	claims["name"] = name
	claims["exp"] = time.Now().Add(time.Hour * 730).Unix()
	claims["iat"] = time.Now()

	tokenString, _ := token.SignedString([]byte(beego.AppConfig.String("JwtSecretKey")))
	return tokenString
}

func VerificationToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("error!")
		}
		return []byte(beego.AppConfig.String("JwtSecretKey")), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}

func TokenValidation(tokenString string) error {
	token, err := VerificationToken(tokenString)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}

	return nil
}
