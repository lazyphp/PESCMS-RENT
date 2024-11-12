package model

import (
	"encoding/json"
	"html"
)

type Field struct {
	Field_id           int    `gorm:"column:field_id;primaryKey" json:"field_id"`
	Field_model_id     int    `gorm:"column:field_model_id" json:"field_model_id"`
	Field_name         string `gorm:"column:field_name" json:"field_name"`
	Field_display_name string `gorm:"column:field_display_name" json:"field_display_name"`
	Field_type         string `gorm:"column:field_type" json:"field_type"`
	Field_option       string `gorm:"column:field_option" json:"field_option"`
	Field_explain      string `gorm:"column:field_explain" json:"field_explain"`
	Field_default      string `gorm:"column:field_default" json:"field_default"`
	Field_required     int    `gorm:"column:field_required" json:"field_required"`
	Field_listsort     int    `gorm:"column:field_listsort" json:"field_listsort"`
	Field_list         int    `gorm:"column:field_list" json:"field_list"`
	Field_form         int    `gorm:"column:field_form" json:"field_form"`
	Field_status       int    `gorm:"column:field_status" json:"field_status"`
	Field_is_null      int    `gorm:"column:field_is_null" json:"field_is_null"`
	Field_only         int    `gorm:"column:field_only" json:"field_only"`
	Field_action       string `gorm:"column:field_action" json:"field_action"`
	Field_sql_type     string `gorm:"column:field_sql_type" json:"field_sql_type"`
	Field_sql_length   int    `gorm:"column:field_sql_length" json:"field_sql_length"`
}

func (f Field) MarshalJSON() ([]byte, error) {
	type Alias Field
	return json.Marshal(&struct {
		*Alias
		FieldOptionText string `json:"field_option"`
	}{
		Alias:           (*Alias)(&f),
		FieldOptionText: f.GetModelStatusText(),
	})
}

func (f Field) GetModelStatusText() string {

	return html.UnescapeString(f.Field_option)
}
