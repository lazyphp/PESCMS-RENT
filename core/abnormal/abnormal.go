package abnormal

import (
	"net/http"
	"runtime"

	core "pescms-rent/core/func"

	"github.com/gin-gonic/gin"
)

/**
 * GIN异常处理中间件
 * @return gin.HandlerFunc
 */
func Abnormal() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {

				buf := make([]byte, 4096)
				buf = buf[:runtime.Stack(buf, false)]
				stackTrace := string(buf)

				core.FatalErrorLog(c, "Internal Server Error", err.(error), stackTrace)
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()

		c.Next()
	}
}
