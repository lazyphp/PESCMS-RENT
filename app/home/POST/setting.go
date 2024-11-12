package POST

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"reflect"
	"time"

	"pescms-rent/core/db"
	core "pescms-rent/core/func"
	"pescms-rent/core/route"
	"pescms-rent/model"

	"github.com/gin-gonic/gin"
)

type Setting struct{}

func init() {
	path := Setting{}
	route.Register(&path, reflect.TypeOf(path).PkgPath())
}

func (api *Setting) Install(c *gin.Context) {
	install := model.GetOption("install")
	if install == "1" {
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "软件已经安装过，请勿重复安装",
		})
		return
	}

	var dataJson map[string]interface{}

	if err := c.ShouldBindJSON(&dataJson); err != nil {
		core.Error(c, "参数错误", 400)
		return
	}

	if dataJson["account"] == nil || dataJson["password"] == nil {
		core.Error(c, "请填写系统安装的账户和密码", 400)
		return
	}

	account := dataJson["account"].(string)
	pwd, _ := core.GeneratePassword(dataJson["password"].(string))

	data := map[string]interface{}{
		"user_account":    account,
		"user_password":   string(pwd),
		"user_createtime": time.Now().Unix(),
		"user_status":     1,
		"user_listsort":   "0",
	}
	sqlErr := db.DB().Table("pes_user").Create(data)

	if sqlErr.Error != nil {
		core.SqlError(c, "创建管理员账户出错", sqlErr.Error)
		return
	}

	db.DB().Table("pes_option").Where("option_name = ?", "install").Updates(map[string]interface{}{"option_value": "1"})

	// 示例用户的密钥长度设置
	keyLength := 32
	key, err := generateKey(keyLength)
	if err != nil {
		core.Error(c, fmt.Sprintf("生成密钥失败: %v", err), 500)
		return
	}

	// 保存到密钥文件
	filename := "recovery_key.txt"
	err = saveKeyToFile(key, filename)
	if err != nil {
		core.Error(c, fmt.Sprintf("保存密钥到文件失败: %v", err), 500)
		return
	}

	core.Success(c, "安装成功", nil)
}

/**
 * 生成随机密钥
 */
func generateKey(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// 将密钥保存到文件
func saveKeyToFile(key, filename string) error {
	return os.WriteFile(filename, []byte(key), 0o600)
}
