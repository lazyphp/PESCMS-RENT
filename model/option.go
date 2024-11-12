package model

import (
	"pescms-rent/core/db"
)

func GetOption(name string) string {
	var option map[string]interface{}
	db.DB().Table("pes_option").Where("option_name = ?", name).Find(&option)

	if option["option_value"] == nil {
		return ""
	} else {
		return option["option_value"].(string)
	}
}
