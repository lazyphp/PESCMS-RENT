package model

type User struct {
	User_id         int    `gorm:"primaryKey;column:user_id"`
	User_account    string `gorm:"column:user_account"`
	User_password   string `gorm:"column:user_password"`
	User_name       string `gorm:"column:user_name"`
	User_status     string `gorm:"column:user_status"`
	User_createtime string `gorm:"column:user_createtime"`
	User_listsort   int    `gorm:"column:user_listsort"`
}
