package db

import (
	"github.com/glebarez/sqlite"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

/**
 * 初始化数据库
 * @return *gorm.DB
 */
func InitDB() *gorm.DB {
	var logMode logger.LogLevel
	switch viper.GetString("sqlist-mod") {
	case "Error":
		logMode = logger.Error
	case "Warn":
		logMode = logger.Warn
	case "Info":
		logMode = logger.Info
	default:
		logMode = logger.Silent // 默认使用 Silent 等级
	}

	// 使用 SQLite3 数据库，数据库文件为 "data.db"
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "pes_", // 表前缀
			SingularTable: true,   // 关闭复数声明
		},
		Logger: logger.Default.LogMode(logMode), // 日志等级
	})
	if err != nil {
		panic("数据库连接失败")
	}

	return db
}

/**
*
* 执行GORM数据库操作
* @return *gorm.DB
 */
func DB() *gorm.DB {
	if db == nil {
		db = InitDB()
	}
	return db
}

/**
 * 获取最后插入ID
 * @return int64
 */
func GetLastID() int {
	var lastInsertID int64
	DB().Raw("SELECT last_insert_rowid()").Row().Scan(&lastInsertID)
	return int(lastInsertID)
}
