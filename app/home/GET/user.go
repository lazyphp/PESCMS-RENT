package GET

import (
	"pescms-rent/core/route"
	"reflect"

	core "pescms-rent/core/func"

	"github.com/gin-gonic/gin"
)

type User struct {
	Content
}

func init() {
	path := User{}
	route.Register(&path, reflect.TypeOf(path).PkgPath())
}

func (api *User) Action(c *gin.Context) {
	modelTable, modelInfo, content, fieldList, status := api.Content.Action(c, false)

	if !status {
		return
	}

	delete(content, "user_password")

	core.Success(c, "ok", gin.H{
		"modelTable": modelTable,
		"modelInfo":  modelInfo,
		"content":    content,
		"field":      fieldList,
	})
}
