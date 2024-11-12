package GET

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	core "pescms-rent/core/func"
	"pescms-rent/core/route"
	"pescms-rent/model"

	"github.com/gin-gonic/gin"
)

type Setting struct {
	Content
}

func init() {
	path := Setting{}
	route.Register(&path, reflect.TypeOf(path).PkgPath())
}

/**
 * 软件存活统计
 */
func (api *Setting) Survival(c *gin.Context) {
	// 表单数据
	formData := url.Values{
		"id": {"9"},
	}

	// 发送 POST 请求到 PHP 服务器
	req, err := http.NewRequest(
		"POST",
		"https://www.pescms.com/?g=Api&m=Statistics&a=survival&method=POST",
		strings.NewReader(formData.Encode()),
	)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()
}

func (api *Setting) Install(c *gin.Context) {
	core.Success(c, "ok", gin.H{
		"install": model.GetOption("install"),
	})
}
