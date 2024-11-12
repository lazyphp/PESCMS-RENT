package POST

import (
	"pescms-rent/core/db"
	"pescms-rent/core/route"
	"pescms-rent/model"
	"reflect"
	"strings"

	core "pescms-rent/core/func"

	"github.com/gin-gonic/gin"
)

type Field struct{}

func init() {
	path := Field{}
	route.Register(&path, reflect.TypeOf(path).PkgPath())
}

func (api *Field) Action(c *gin.Context) {
	data, _, err := model.HandleData(c)
	if err != nil {
		core.Error(c, err.Error(), err.Status())
		return
	}

	sqlErr := db.DB().Table("pes_field").Create(data)
	if sqlErr.Error != nil {
		core.SqlError(c, "写入字段失败", sqlErr.Error)
		return
	}

	// 生成表字段
	var currentModel model.Model
	db.DB().First(&currentModel, data["field_model_id"])

	tableName := core.SanitizeInput(currentModel.Model_name)
	columnName := tableName + "_" + core.SanitizeInput(data["field_name"].(string))
	fieldType := core.SanitizeInput(data["field_sql_type"].(string))
	fieldLength := core.SanitizeInput(data["field_sql_length"].(string))

	var alterSql string
	if fieldType == "text" {
		alterSql = fieldType
	} else {
		alterSql = fieldType + "(" + fieldLength + ")"
	}

	if data["field_is_null"].(string) == "1" {
		alterSql += " NULL"
	} else {
		alterSql += " NOT NULL DEFAULT ''"
	}

	sqlErr = db.DB().Exec("ALTER TABLE `pes_" + strings.ToLower(tableName) + "` ADD `" + strings.ToLower(columnName) + "` " + alterSql)

	if sqlErr.Error != nil {
		core.SqlError(c, "创建表字段失败", sqlErr.Error)
		return
	}

	core.Success(c, "新增内容成功", nil)
}
