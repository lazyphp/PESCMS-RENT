package slice

import (
	"pescms-rent/core/db"
	core "pescms-rent/core/func"
	"pescms-rent/core/route"
	"pescms-rent/model"

	"github.com/gin-gonic/gin"
)

type User struct{}

func init() {
	route.RegSlice(&User{})
}

func (api *User) Action(c *gin.Context) {
	model.PrePayload(c, func(data map[string]interface{}) error {
		if len(data["password"].(string)) > 0 {
			passwd, _ := core.GeneratePassword(data["password"].(string))
			data["password"] = string(passwd)
		} else {
			if c.Request.Method == "PUT" {
				var user model.User
				userID := c.Query("id")
				db.DB().Where("user_id = ?", userID).First(&user)
				data["password"] = user.User_password
			}
		}
		return nil
	})
}
