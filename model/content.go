package model

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"pescms-rent/core/db"
	"strings"
	"time"

	core "pescms-rent/core/func"

	"github.com/gin-gonic/gin"
	"github.com/iancoleman/orderedmap"
)

type Content struct{}

/**
 * 获取模型基础信息
 */
func GetModelBase(c *gin.Context) (string, []Field, Model, *core.HTTPError) {
	modelTable := c.Param("model")
	if len(modelTable) == 0 {
		modelTable = c.Query("model")
	}

	if len(modelTable) == 0 {
		return "", nil, Model{}, core.SetError("模型名称不能为空", 1)
	}

	checkModel, err := CheckModel(modelTable)

	if err != nil {
		return "", nil, Model{}, err
	} else {
		// 获取对应模型的字段
		fieldList := GetModelField(checkModel.Model_id)

		return modelTable, fieldList, checkModel, nil
	}
}

/**
 * 检查模型是否存在
 */
func CheckModel(modelTable string) (Model, *core.HTTPError) {
	var checkModel Model

	db.DB().Where("model_name = ?", modelTable).First(&checkModel)

	if len(checkModel.Model_name) == 0 {
		return Model{}, core.SetError(fmt.Sprintf("%s 模型不存在", modelTable), 404)
	} else {
		return checkModel, nil
	}
}

/**
 * 获取模型字段
 */
func GetModelField(modelId int) []Field {
	var fieldList []Field

	db.DB().Where("field_model_id = ?", modelId).Order("field_listsort ASC, field_id DESC").Find(&fieldList)

	return fieldList
}

/**
 * 处理数据
 */
func HandleData(c *gin.Context) (map[string]interface{}, string, *core.HTTPError) {
	var dataJson map[string]interface{}

	if err := c.ShouldBindJSON(&dataJson); err != nil {
		return nil, "", core.SetError(err.Error(), 1)
	}

	modelTable, fieldList, modelInfo, err := GetModelBase(c)
	if err != nil {
		return nil, "", err
	}

	data := make(map[string]interface{})

	// 组装更新的SQL数据
	for _, item := range fieldList {

		if item.Field_status == 0 {
			continue
		}

		if item.Field_required == 1 && dataJson[item.Field_name] == "" {
			return nil, "", core.SetError(fmt.Sprintf("'%s'为必填项", item.Field_display_name), 1)
		}

		if !strings.Contains(item.Field_action, c.Request.Method) {
			continue
		}

		switch item.Field_type {
		case "multiple", "checkbox", "imgs", "files", "videos": // 字符串需要进行切割处理

			data[modelInfo.Model_name+"_"+item.Field_name] = core.SliceToString(dataJson[item.Field_name].([]interface{}))

		case "option":

			data[modelInfo.Model_name+"_"+item.Field_name] = ConvertOptionToJson(dataJson[item.Field_name].([]interface{}))

		case "date":
			// 定义一个时间字符串
			timeString := dataJson[item.Field_name].(string) + " 00:00:00"

			// 定义时间格式
			layout := "2006-01-02 15:04:05"
			loc, _ := time.LoadLocation("Local")
			// 将字符串转换为时间
			t, err := time.ParseInLocation(layout, timeString, loc)
			if err != nil {
				return nil, "", core.SetError(fmt.Sprintf("error parsing time: %s", err), 1)
			}

			// 将时间转换为时间戳
			timestamp := t.Unix()

			data[modelInfo.Model_name+"_"+item.Field_name] = timestamp

		default:
			data[modelInfo.Model_name+"_"+item.Field_name] = dataJson[item.Field_name]

		}

	}

	return data, modelTable, nil
}

/**
 * 转换选项为 JSON 字符串
 */
func ConvertOptionToJson(option []interface{}) string {
	orderedMap := orderedmap.New()

	for _, item := range option {
		itemMap := item.(map[string]interface{})
		key := itemMap["key"].(string)
		value := itemMap["value"].(string)

		if key == "" || value == "" {
			continue
		}

		orderedMap.Set(key, value)
	}

	// 转换为 JSON 字符串
	jsonData, err := json.Marshal(orderedMap)
	if err != nil {
		// fmt.Println("Error:", err)
		return ""
	}

	return string(jsonData)
}

/**
 * 预处理和覆盖对应的请求体
 * @param c *gin.Context
 * @param modifyFunc func(map[string]interface{}) error 处理数据的业务逻辑
 */
func PrePayload(c *gin.Context, modifyFunc func(map[string]interface{}) error) {
	// 首先读取原始请求体
	var data map[string]interface{}
	if err := c.ShouldBindJSON(&data); err != nil {
		core.Error(c, err.Error(), http.StatusBadRequest)
		return
	}

	// 执行回调函数，修改数据
	if err := modifyFunc(data); err != nil {
		core.Error(c, err.Error(), http.StatusInternalServerError)
		return
	}

	// 将修改后的数据重新编码为 JSON 字符串
	modifiedData, err := json.Marshal(data)
	if err != nil {
		core.Error(c, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}

	// 将修改后的 JSON 数据写回到请求体中
	c.Request.Body = io.NopCloser(bytes.NewBuffer(modifiedData))
}
