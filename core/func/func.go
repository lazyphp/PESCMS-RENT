package core

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/araddon/dateparse"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

/**
 * GeneratePassword 创建哈希值密码
 * @param  {string} password 密码
 * @return {[]byte}
 * @return {error}
 */
func GeneratePassword(password string) ([]byte, error) {
	// bcrypt.GenerateFromPassword 使用默认的cost（10）生成哈希值
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return hash, nil
}

/**
 * CompareHashAndPassword 比较存储的哈希值和用户输入的密码是否匹配
 * @param  {[]byte} hash     哈希值
 * @param  {string} password 密码
 * @return {error}
 */
func CompareHashAndPassword(hash []byte, password string) error {
	err := bcrypt.CompareHashAndPassword(hash, []byte(password))
	if err != nil {
		return err
	}
	return nil
}

/**
 * SqlError 数据库错误
 * @param  {gin.Context} c   gin
 * @param  {string}      msg 错误消息
 * @param  {error}       err sql错误
 */
func SqlError(c *gin.Context, msg string, sqlError error) {
	log.Error(sqlError)
	Error(c, msg, 500)
}

/**
 * FatalErrorLog 致命错误日志
 * @param  {gin.Context} c         gin
 * @param  {string}      msg       错误消息
 * @param  {error}       err       错误
 * @param  {string}      stackTrace 堆栈跟踪
 */
func FatalErrorLog(c *gin.Context, msg string, err error, stackTrace string) {
	log.Error(err)
	log.Error(stackTrace)
	Error(c, msg, 500)
}

/**
 * Error gin的JSON错误处理
 * @param  {gin.Context} c    gin
 * @param  {string}      msg  错误消息
 * @param  {int}         code 错误码
 */
func Error(c *gin.Context, msg string, code int) {
	var status int
	switch code {
	case 301, 302:
		status = http.StatusMovedPermanently
	case 404:
		status = http.StatusNotFound
	case 500, 501, 502:
		status = http.StatusInternalServerError
	default:
		status = http.StatusOK
	}

	c.JSON(status, gin.H{
		"code": code,
		"msg":  msg,
	})
	c.Abort()
}

/**
 * Success gin的JSON成功处理 （状态码默认为0）
 * @param  {gin.Context} c    gin
 * @param  {string}      msg  消息
 * @param  {interface{}} data 数据
 */
func Success(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  msg,
		"data": data,
	})
	c.Abort()
}

/**
 * SanitizeInput 移除非法字符
 * @param  {string} input 输入字符串
 * @return {string}
 */
func SanitizeInput(input string) string {
	// 使用正则表达式移除非法字符
	reg := regexp.MustCompile("[^a-zA-Z0-9_]")
	sanitized := reg.ReplaceAllString(input, "")
	return sanitized
}

/**
 * SliceToString 将切片转换为字符串 示例：1,2,3,4
 * @param  {[]interface{}} slice 切片
 * @return {string}
 */
func SliceToString(slice []interface{}) string {
	var sliceStr string
	for i, v := range slice {
		sliceStr += v.(string)
		if i < len(slice)-1 {
			sliceStr += ","
		}
	}

	return sliceStr
}

/**
 * MergeMaps 合并多个map
 * @param  {map[string]interface{}}
 * @return {map[string]interface{}}
 */
func MergeMaps(maps ...map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}

/**
 * GenerateRandomFilename 生成随机文件名
 * @param  {int} length 文件名长度
 * @return {string}
 */
func GenerateRandomFilename(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

/**
 * ReplaceAll 替换字符串
 * @param  {string}                 s            字符串
 * @param  {map[string]string}      replacements 替换规则
 * @return {string}
 */
func StrReplace(s string, replacements map[string]string) string {
	for old, new := range replacements {
		s = strings.ReplaceAll(s, old, new)
	}
	return s
}

func SearchValueInJsonAndReturnKey(jsonStr string, value string) string {
	if len(jsonStr) == 0 {
		return ""
	} else {
		var data map[string]interface{}

		err := json.Unmarshal([]byte(jsonStr), &data)
		if err != nil {
			log.Error("json string is invalid, jsonStr: ", jsonStr)
			return ""
		}

		for k, v := range data {
			if v == value {
				return k
			}
		}
		return ""
	}
}

/**
 * Strtotime 将日期字符串转换为 UNIX 时间戳
 * @param  {string} date 日期字符串
 * @return {int64}
 * @return {error}
 */
func Strtotime(date string) (int64, error) {
	t, err := dateparse.ParseAny(date)
	if err != nil {
		return 0, err
	}

	// 将 time.Time 对象转换为 UNIX 时间戳
	timestamp := t.Unix()

	return timestamp, nil
}
