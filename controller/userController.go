package controller

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"web-demo-interface-test/dao/userDao"
	"web-demo-interface-test/log"
	"web-demo-interface-test/model/user"
)

func RegisterGET(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{"title": "注册页"})
}

func RegisterPOST(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	var userModel user.User
	now := time.Now()
	strTime := now.Format("20060102150405")
	userModel.UserId = strTime + username
	userModel.Username = username
	userModel.Password = password
	if username != "" && password != "" {
		userId := user.SelectUser(username)
		// 先查询username是不是已经被注册
		if userId != "" {
			log.Info("注册失败")
			c.HTML(http.StatusConflict, "register.html", gin.H{"title": "注册页", "msg": "用户名已经被注册"})
		} else {
			if err := user.InsertUser(&userModel); err == nil {
				log.Info("注册成功")
				LoginGET(c)
			} else {
				log.Error("注册失败")
				c.HTML(http.StatusConflict, "register.html", gin.H{"title": "注册页", "msg": "未知错误"})
			}
		}
	} else {
		c.HTML(http.StatusSeeOther, "register.html", gin.H{"title": "注册页", "msg": "请输入内容"})
	}
}

func LoginGET(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{"title": "登录页"})
}

func LoginPOST(c *gin.Context) {
	var nickname string
	username := c.PostForm("username")
	password := c.PostForm("password")
	//验证用户名密码
	loginUser := user.LoginUser(username, password)
	if loginUser.UserId == "" {
		log.Info("登录失败")
		c.HTML(http.StatusBadRequest, "login.html", gin.H{"title": "登录页", "msg": "用户名或密码输入错误"})
	} else {
		if loginUser.Status == 0 {
			err := userDao.QueryRowDB(&nickname, "select nickname from user_info_table where userId=?", loginUser.UserId)
			if err != nil {
				log.Info("登录失败")
				c.HTML(http.StatusBadRequest, "login.html", gin.H{"title": "登录页", "msg": "登录失败，程序内部错误"})
			}
			//设置session
			session := sessions.Default(c)
			session.Set("login_user", loginUser.UserId)
			if err := session.Save(); err != nil {
				c.HTML(http.StatusNotFound, "login.html", gin.H{"title": "登录页", "msg": "session创建失败"})
			}
			log.Info(loginUser.UserId + "登录成功")
			c.HTML(http.StatusOK, "info.html", gin.H{"nickname": nickname})
		} else {
			log.Info("登录失败")
			c.HTML(http.StatusBadRequest, "login.html", gin.H{"title": "登录页", "msg": "用户未注册或无效"})
		}
	}
}

func DeleteGET(c *gin.Context) {
	//获取session，然后删除，使session过期
	session := sessions.Default(c)
	userId := session.Get("login_user")
	userIdStr := fmt.Sprintf("%v", userId)
	_, err := user.DeleteUser(userIdStr)
	session.Clear()
	if err != nil {
		return
	}
	c.HTML(http.StatusOK, "login.html", gin.H{"title": "登录页"})
}

func UpdatePOST(c *gin.Context) {
	var info *user.Info
	nickname := c.PostForm("nickname")
	password := c.PostForm("password")
	newPassword := c.PostForm("newPassword")
	reNewPassword := c.PostForm("reNewPassword")
	//如果nickname不为空，则要修改nickname
	if nickname != "" {
		//从session获取userId，通过userId查询索要修改的相关属性值
		//这里的逻辑有点繁杂，自己理吧
	}
	if password != "" {

	}
	if newPassword != reNewPassword {

	}
	info.NickName = nickname
	info.Password = newPassword
	user.UpdateInfo(info)
}

func LogoutGet(c *gin.Context) {
	//获取session使过期
	session := sessions.Default(c)
	loginUser := session.Get("login_user")
	if loginUser == nil {
		c.HTML(http.StatusNotAcceptable, "login.html", gin.H{"title": "登录页", "msg": "未登录，执行退出操作无效"})
	}
	//session.Delete("login_user")
	session.Clear()
	c.HTML(http.StatusOK, "login.html", gin.H{"title": "登录页"})
}
