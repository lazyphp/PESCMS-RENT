package main

import (
	"fmt"
	"io"
	"net/http"

	_ "pescms-rent/app"
	"pescms-rent/core/abnormal"
	"pescms-rent/core/route"
	_ "pescms-rent/slice"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	// 初始化 viper 读取配置文件
	viper.SetConfigName("debug") // 配置文件名 (不带扩展名)
	viper.SetConfigType("yaml")  // 配置文件类型
	viper.AddConfigPath(".")     // 配置文件路径

	// 不存在配置文件，就是正式版本运行
	viper.ReadInConfig()

	ginMode := viper.GetString("gin-mod")
	if ginMode == "debug" {
		fmt.Println("当前为调试模式")
		gin.SetMode(gin.DebugMode)
	} else {
		printBanner()
		gin.SetMode(gin.ReleaseMode)
		gin.DefaultWriter = io.Discard // 屏蔽GIN的掉信息输出
	}

	r := gin.Default()

	r.Use(CorsMiddleware())

	r.Use(route.SliceMiddleware("slice/config.yaml"))

	r.Use(AuthorizeAdmin())

	// // 注册异常中间件
	r.Use(abnormal.Abnormal())

	route.Bind(r)

	r.Static("/upload", "./upload")

	r.NoRoute(func(c *gin.Context) {
		path := c.Request.URL.Path

		method := c.Request.Method
		// 返回 404 错误

		var msg string
		if ginMode == "debug" {
			msg = "您" + method + "请求地址：" + path + "不存在！"
		} else {
			msg = "您访问的地址不存在"
		}

		c.JSON(404, gin.H{"code": 404, "msg": msg})
	})

	fmt.Println("服务启动成功")
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 设置允许跨域的源
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

		// 设置允许跨域的方法
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")

		// 设置允许跨域的请求头
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok!")
			c.Abort()
			return
		}

		c.Next()
	}
}

/**
 * 后台鉴权
 */
func AuthorizeAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.Request.URL.Path) >= 5 && c.Request.URL.Path[:5] == "/home" {

			token := c.Request.Header.Get("Authorization")
			if len(token) > 0 {
				_, err := route.ValidateJwt(token)
				if err != nil {
					c.JSON(302, gin.H{"code": 302, "message": "token无效"})
					c.Abort()
				}
			}
			c.Next()

		} else if len(c.Request.URL.Path) == 7 && c.Request.URL.Path[:7] == "/upload" {
			c.JSON(404, gin.H{"code": 404, "msg": "您访问的地址不存在"})
			c.Abort()
		} else {
			c.Next()
		}
	}
}

func printBanner() {
	fmt.Println(`
  _____  ______  _____  _____ __  __  _____ 
 |  __ \|  ____|/ ____|/ ____|  \/  |/ ____|
 | |__) | |__  | (___ | |    | \  / | (___  
 |  ___/|  __|  \___ \| |    | |\/| |\___ \ 
 | |    | |____ ____) | |____| |  | |____) |
 |_|    |______|_____/ \_____|_|  |_|_____/ 
                                            
                                            
   欢迎使用PESCMS RENT房租管理系统
   程序正在启动，请等候加载完成...
  `)
}
