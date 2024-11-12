package DELETE

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

/**
 * @api {DELETE} /ticket/content 更新内容
 */
func (api *Content) Action(c *gin.Context) {
	modelTable, _, _, _ := model.GetModelBase(c)

	var data map[string]interface{}

	id := c.Query("id")
	if len(id) > 0 {
		db.DB().Table("pes_"+modelTable).Where(modelTable+"_id = ?", id).Delete(data)
	}

	core.Success(c, "删除完毕", nil)
}
