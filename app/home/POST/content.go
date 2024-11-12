package POST

import (
	"pescms-rent/core/db"
	"pescms-rent/core/route"
	"pescms-rent/model"
	"reflect"

	core "pescms-rent/core/func"

	"github.com/gin-gonic/gin"
)

type Content struct{}

func init() {
	path := Content{}
	route.Register(&path, reflect.TypeOf(path).PkgPath())
}

func (api *Content) Action(c *gin.Context, isJson bool) (int, string, map[string]interface{}, bool) {
	data, modelTable, err := model.HandleData(c)
	if err != nil {
		core.Error(c, err.Error(), err.Status())
		return 0, "", nil, false
	}

	sqlErr := db.DB().Table("pes_" + modelTable).Create(data)

	if sqlErr.Error != nil {
		core.SqlError(c, "添加新内容失败", sqlErr.Error)
		return 0, "", nil, false
	}

	lastInsertID := db.GetLastID()

	if isJson {
		core.Success(c, "新增内容成功", nil)
		return 0, "", nil, true
	} else {
		return lastInsertID, modelTable, nil, true
	}
}
