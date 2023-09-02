package router

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"html/template"
	"time"
	"web-demo-interface-test/controller"
	"web-demo-interface-test/log"
	"web-demo-interface-test/middleware"
)

// SetupRouter 设置路由
func SetupRouter() *gin.Engine {
	r := gin.New()

	// 接管Gin的日志
	r.Use(log.GinLogger(log.Logger))

	//这个要写在路由定义之前，初始化session存储引擎，cookie方式
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mySession", store))

	// 设置时间格式
	r.SetFuncMap(template.FuncMap{
		"timeStr": func(timeStamp int64) string {
			return time.Unix(timeStamp, 0).Format("2023-04-24 12:00:00")
		},
	})

	// 配置模板
	r.LoadHTMLGlob("./templates/*")

	// 登录注册，无需认证
	{
		r.GET("/register", controller.RegisterGET)
		r.POST("/register", controller.RegisterPOST)

		r.GET("/login", controller.LoginGET)
		r.GET("/", controller.LoginGET)
		r.POST("/login", controller.LoginPOST)
	}

	// 需要认证的路由
	{
		routerGroup := r.Group("/api", middleware.BasicAuth())
		//退出登录
		routerGroup.GET("/logout", controller.LogoutGet)
		//修改信息
		routerGroup.POST("/change", controller.UpdatePOST)
		//注销
		routerGroup.GET("/delete", controller.DeleteGET)
	}
	return r
}
