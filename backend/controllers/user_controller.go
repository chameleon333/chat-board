package controllers

import (
	"log"
	"net/http"
	"work/models"
	"work/session"

	"github.com/gin-gonic/gin"
)

var user models.User

func LoginFormRoute(c *gin.Context) {
	c.HTML(200, "login.html", gin.H{})
}

func Login(c *gin.Context) {

	// バリデーション処理
	if err := c.Bind(&user); err != nil {
		log.Println("ログインできませんでした")
		c.HTML(http.StatusBadRequest, "signup.html", gin.H{"err": err})
		c.Abort()
	}

	log.Println("ログインできました")
	isExist, _ := models.IsUserExist(user.Email)

	if isExist {
		user, _ := models.GetUser(user.Email, user.Password)
		session.Login(c, user.Email)
	}
	c.Redirect(302, "/list")
}

func Logout(c *gin.Context) {
	session.Logout(c)
	c.Redirect(http.StatusMovedPermanently, "/login")
}

func SignUpFormRoute(c *gin.Context) {
	c.HTML(200, "signup.html", gin.H{})
}

func SignUp(c *gin.Context) {
	var user models.User
	// バリデーション処理
	if err := c.Bind(&user); err != nil {
		c.HTML(http.StatusBadRequest, "signup.html", gin.H{"err": err})
		c.Abort()
	} else {
		email := c.PostForm("email")
		password := c.PostForm("password")
		// 登録ユーザーが重複していた場合にはじく処理
		if err := user.CreateUser(email, password); err != nil {
			c.HTML(http.StatusBadRequest, "signup.html", gin.H{"Err": err})
		}
		c.Redirect(302, "/list")
	}
}
