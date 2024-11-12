package GET

import (
	"pescms-rent/core/db"
	"pescms-rent/core/route"
	"reflect"

	core "pescms-rent/core/func"

	"github.com/gin-gonic/gin"
)

type Field struct {
	Content
}

func init() {
	path := Field{}
	route.Register(&path, reflect.TypeOf(path).PkgPath())
}

func (api *Field) Index(c *gin.Context) {
	modelID := c.Query("modelID")

	if modelID == "" {
		core.Error(c, "请提交模型ID", 1)
		return
	}

	_, _, fieldList, modelInfo, _, status := api.Content.Index(c, false)

	if !status {
		return
	}

	// 获取对应模型表数据。
	var contentList []map[string]interface{}

	db.DB().Table("pes_field").Where("field_model_id = ?", modelID).Order("field_listsort ASC, field_id DESC").Find(&contentList, nil)

	core.Success(c, "ok", gin.H{
		"pageTotal":   1,
		"modelInfo":   modelInfo,
		"contentList": contentList,
		"field":       fieldList,
	})
}

func (api *Field) Action(c *gin.Context) {
	modelTable, modelInfo, _, fieldList, status := api.Content.Action(c, false)

	if !status {
		return
	}

	// 获取对应模型表数据。
	var content map[string]interface{}

	id := c.Query("id")
	if len(id) > 0 {
		db.DB().Table("pes_"+modelTable).Take(&content, modelTable+"_id = ?", id)
	} else {
		content = make(map[string]interface{})
		content["field_model_id"] = c.Query("modelID")
	}

	core.Success(c, "ok", gin.H{
		"modelInfo": modelInfo,
		"content":   content,
		"field":     fieldList,
	})
}
