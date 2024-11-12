package PUT

import (
	"os"
	"reflect"
	"strings"

	"pescms-rent/core/db"
	core "pescms-rent/core/func"
	"pescms-rent/core/route"

	"github.com/gin-gonic/gin"
)

type Login struct{}

func init() {
	path := Login{}
	route.Register(&path, reflect.TypeOf(path).PkgPath())
}

func (api *Login) FindPw(c *gin.Context) {
	var dataJson map[string]interface{}

	if err := c.ShouldBindJSON(&dataJson); err != nil {
		core.Error(c, "参数错误", 400)
		return
	}

	if dataJson["serverkey"] == nil || dataJson["account"] == nil || dataJson["password"] == nil {
		core.Error(c, "请填写安全密钥、账号或密码", 400)
		return
	}

	submittedKey, _ := dataJson["serverkey"].(string)

	// 读取存储在文件中的密钥
	filePath := "recovery_key.txt"
	storedKey, err := readKeyFromFile(filePath)
	if err != nil {
		core.Error(c, "服务器错误：无法读取密钥文件", 500)
		return
	}

	// 比较前端提交的密钥和存储的密钥是否一致
	if submittedKey != storedKey {
		core.Error(c, "密钥错误", 400)
		return
	}

	passwd, _ := core.GeneratePassword(dataJson["password"].(string))

	db.DB().Table("pes_user").Where("user_account = ?", dataJson["account"]).Updates(map[string]interface{}{"user_password": string(passwd)})

	// 删除文件
	os.Remove(filePath)

	core.Success(c, "密码修改完成", nil)
}

// 读取文件中的密钥
func readKeyFromFile(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(data)), nil
}
