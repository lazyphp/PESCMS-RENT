package PUT

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
 * @api {put} /ticket/content 更新内容
 */
func (api *Content) Action(c *gin.Context) {
	data, modelTable, err := model.HandleData(c)
	if err != nil {
		core.Error(c, err.Error(), err.Status())
		return
	}

	id := c.Query("id")
	if len(id) > 0 {
		res := db.DB().Table("pes_"+modelTable).Where(modelTable+"_id = ?", id).Updates(data)
		if res.Error != nil {
			core.SqlError(c, "内容更新失败", res.Error)
			return
		} else if res.RowsAffected <= 0 {
			core.Error(c, "更新内容没有变化", 1)
			return
		}
	}

	core.Success(c, "内容更新完毕", nil)
}

func (api *Content) Sort(c *gin.Context) {
	var dataJson map[string]interface{}

	if err := c.ShouldBindJSON(&dataJson); err != nil {
		core.Error(c, "请提交排序内容", 1)
		return
	}

	sortSlice, ok := dataJson["sort"].([]interface{})
	if !ok || len(sortSlice) <= 0 {
		core.Error(c, "请提交排序值", 1)
		return
	}

	modelTable, _, _, err := model.GetModelBase(c)
	if err != nil {
		core.Error(c, err.Error(), err.Status())
		return
	}

	for _, v := range sortSlice {
		data, _ := v.(map[string]interface{})

		sqlData := map[string]interface{}{
			modelTable + "_listsort": data["sort"],
		}
		db.DB().Table("pes_"+modelTable).Where(modelTable+"_id = ?", data["id"]).Updates(sqlData)
	}

	core.Success(c, "内容更新排序值成功", nil)
}
