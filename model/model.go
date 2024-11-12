package model

type Model struct {
	Model_id     int    `gorm:"column:model_id;primaryKey" json:"model_id"`
	Model_name   string `gorm:"column:model_name" json:"model_name"`
	Model_title  string `gorm:"column:model_title" json:"model_title"`
	Model_status int    `gorm:"column:model_status" json:"model_status"`
	Model_search int    `gorm:"column:model_search" json:"model_search"`
	Model_attr   int    `gorm:"column:model_attr" json:"model_attr"`
	Model_page   int    `gorm:"column:model_page" json:"model_page"`
}

// func GetModel(id int) *Model {
// 	return &Model{}
// }
