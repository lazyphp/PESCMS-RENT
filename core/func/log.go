package core

import (
	"os"
	"path/filepath"
	"time"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func init() {
	// 初始化日志
	log = logrus.New()

	// 设置日志级别
	log.SetLevel(logrus.InfoLevel)

	// 创建日志目录
	createLogDir()
}

func createLogDir() {
	logDir := "logs"
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		os.Mkdir(logDir, os.ModePerm)
	}

	// 检查日志目录中的子目录
	fileInfos, err := os.ReadDir(logDir)
	if err != nil {
		logrus.Fatal("Failed to read log directory:", err)
	}

	// 获取当前日期
	currentDate := time.Now()

	// 遍历子目录
	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			// 解析子目录名称为日期
			dirDate, err := time.Parse("2006-01-02", fileInfo.Name())
			if err != nil {
				logrus.Warn("Failed to parse directory name:", err)
				continue
			}

			// 计算子目录日期与当前日期的差距
			daysDiff := currentDate.Sub(dirDate).Hours() / 24

			// 如果超过7天，则删除该目录
			if daysDiff > 7 {
				err := os.RemoveAll(filepath.Join(logDir, fileInfo.Name()))
				if err != nil {
					logrus.Warn("Failed to remove directory:", err)
				} else {
					logrus.Info("Removed directory:", fileInfo.Name())
				}
			}
		}
	}

	// 根据年月日创建子目录
	logSubDir := time.Now().Format("2006-01-02")
	logPath := filepath.Join(logDir, logSubDir)
	if _, err := os.Stat(logPath); os.IsNotExist(err) {
		os.MkdirAll(logPath, os.ModePerm)
	}

	// 创建日志文件
	logFileName := filepath.Join(logPath, "logfile.log")
	file, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o666)
	if err != nil {
		logrus.Fatal("Failed to open log file:", err)
	}

	// 设置日志输出到文件
	log.Out = file
}
