package POST

import (
	"pescms-rent/core/db"
	"pescms-rent/core/route"
	"pescms-rent/model"
	"reflect"

	core "pescms-rent/core/func"

	"github.com/gin-gonic/gin"
)

type Model struct{}

func init() {
	path := Model{}
	route.Register(&path, reflect.TypeOf(path).PkgPath())
}

func (api *Model) Action(c *gin.Context) {
	data, _, err := model.HandleData(c)
	if err != nil {
		core.Error(c, err.Error(), err.Status())
		return
	}

	sqlErr := db.DB().Table("pes_model").Create(data)
	lastInsertID := db.GetLastID()
	if sqlErr.Error != nil {
		core.SqlError(c, "写入模型表失败", sqlErr.Error)
		return
	}

	// 写入模型字段
	field := []model.Field{
		{Field_model_id: lastInsertID, Field_name: "status", Field_display_name: "状态", Field_type: "radio", Field_option: "{\"禁用\":\"0\",\"启用\":\"1\"}", Field_default: "1", Field_required: 1, Field_listsort: 100, Field_list: 1, Field_form: 1, Field_status: 1, Field_is_null: 0, Field_only: 0, Field_action: "POST,PUT", Field_sql_type: "int", Field_sql_length: 11},
		{Field_model_id: lastInsertID, Field_name: "listsort", Field_display_name: "排序", Field_type: "text", Field_default: "0", Field_required: 1, Field_listsort: 98, Field_list: 1, Field_form: 1, Field_status: 1, Field_is_null: 0, Field_only: 0, Field_action: "POST,PUT", Field_sql_type: "int", Field_sql_length: 11},
		{Field_model_id: lastInsertID, Field_name: "createtime", Field_display_name: "创建时间", Field_type: "date", Field_required: 1, Field_listsort: 99, Field_list: 1, Field_form: 1, Field_status: 1, Field_is_null: 0, Field_only: 0, Field_action: "POST,PUT", Field_sql_type: "varchar", Field_sql_length: 255},
	}
	sqlErr = db.DB().Create(&field)
	if sqlErr.Error != nil {
		core.SqlError(c, "创建模型基础字段失败", sqlErr.Error)
		return

	}

	// 创建表
	sqlStatement := `
	CREATE TABLE IF NOT EXISTS pes_` + data["model_name"].(string) + ` (
    ` + data["model_name"].(string) + `_id INTEGER PRIMARY KEY AUTOINCREMENT,
    ` + data["model_name"].(string) + `_status INTEGER NOT NULL DEFAULT 0,
    ` + data["model_name"].(string) + `_listsort INTEGER NOT NULL DEFAULT 0,
    ` + data["model_name"].(string) + `_createtime INTEGER NOT NULL DEFAULT 0
	);
`
	sqlErr = db.DB().Exec(sqlStatement)
	if sqlErr.Error != nil {
		core.SqlError(c, "创建模型表失败", sqlErr.Error)
		return

	}

	//@todo 缺少了基础权限添加

	core.Success(c, "新增模型完成", nil)
}
