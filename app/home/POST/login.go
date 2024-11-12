package POST

import (
	"encoding/json"
	"errors"
	"net/http"
	"reflect"

	"pescms-rent/core/db"
	"pescms-rent/core/route"

	core "pescms-rent/core/func"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Login struct{}

func init() {
	path := Login{}
	route.Register(&path, reflect.TypeOf(path).PkgPath())
}

type LoginRequest struct {
	Username string `json:"Username" binding:"required"`
	Password string `json:"Password" binding:"required"`
}

type ApiError struct {
	Field string
	Msg   string
}

func (api *Login) Login(c *gin.Context) {
	var request LoginRequest
	// 表单基础校验
	if err := c.ShouldBindJSON(&request); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]ApiError, len(ve))
			for i, fe := range ve {
				out[i] = ApiError{fe.Field(), msgForTag(fe.Tag())}
			}
			c.JSON(http.StatusBadRequest, gin.H{"msg": out})
			return
		}
	}

	var user map[string]interface{}
	// 查询账号信息
	db.DB().Table("pes_user").Where("user_account = ?", request.Username).Find(&user)

	if len(user) <= 0 {
		core.Error(c, "账号不存在", 1)
		return
	}

	// 开始校验密码
	password := []byte(user["user_password"].(string))
	err := core.CompareHashAndPassword(password, request.Password)
	if err != nil {
		core.Error(c, "密码错误", 1)
		return
	}

	// // 生成JWT token
	user["user_password"] = ""
	jsonData, _ := json.Marshal(user)
	token, _ := route.CreateJwt(user["user_account"].(string), string(jsonData))

	core.Success(c, "登录成功", gin.H{
		"token": token,
	})
}

func (api *Login) ValidateToken(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	_, err := route.ValidateJwt(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Token无效"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "token持续有效"})
}

func msgForTag(tag string) string {
	switch tag {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email"
	}
	return ""
}
