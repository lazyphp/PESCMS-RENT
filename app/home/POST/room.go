package POST

import (
	"encoding/json"
	"net/http"
	"reflect"
	"strconv"
	"time"

	"pescms-rent/core/db"
	"pescms-rent/core/route"
	"pescms-rent/model"

	core "pescms-rent/core/func"

	"github.com/gin-gonic/gin"
)

type Room struct{}

func init() {
	path := Room{}
	route.Register(&path, reflect.TypeOf(path).PkgPath())
}

func (api *Room) Fee(c *gin.Context) {
	var data map[string]interface{}
	if err := c.ShouldBindJSON(&data); err != nil {
		core.Error(c, err.Error(), http.StatusBadRequest)
		return
	}

	fee_date, _ := core.Strtotime(data["date"].(string))

	room := model.GetRoomWithID(data["roomID"].(string))
	if room == nil {
		core.Error(c, "房间不存在", http.StatusBadRequest)
		return
	}

	var roomCost map[string]interface{}
	err := json.Unmarshal([]byte(room["room_cost"].(string)), &roomCost)
	if err != nil {
		core.Error(c, "读取房间费用失败", http.StatusBadRequest)
		return
	}

	feeRules := make(map[string]interface{})

	feeReules, _ := model.GetFeeRules()

	for _, item := range feeReules {
		feeRules[item["fee_rules_name"].(string)] = item
	}

	tx := db.DB().Begin()

	for name, unitPrice := range roomCost {

		unitPrice, _ = strconv.ParseFloat(unitPrice.(string), 64)

		value, _ := strconv.ParseFloat(data[name].(string), 64)

		var getLastFee map[string]interface{}
		db.DB().Table("pes_room_fee").Where("room_id = ? AND fee_name = ? ", room["room_id"], name).Order("fee_date DESC").Find(&getLastFee)
		// 如果没有最近一次信息，就是0开始
		var fee_value string
		if len(getLastFee) == 0 {
			fee_value = "0"
		} else {
			fee_value = getLastFee["fee_value"].(string)
		}

		lastValue, _ := strconv.ParseFloat(fee_value, 64)

		used := value - lastValue
		rule, _ := feeRules[name].(map[string]interface{})

		var price float64
		if rule["fee_rules_price_type"] == int64(1) {
			price = used * unitPrice.(float64)
		} else {
			price, _ = strconv.ParseFloat(data[name].(string), 64)
		}

		// 插入房间使用费用
		addFee := map[string]interface{}{
			"room_id":         room["room_id"],
			"fee_name":        name,
			"fee_value":       data[name],
			"fee_price":       price,
			"fee_used":        used,
			"fee_date":        fee_date,
			"fee_create_time": int32(time.Now().Unix()),
		}

		tx.Table("pes_room_fee").Create(&addFee)

	}

	// 插入租金记录
	tx.Table("pes_room_fee").Create(&map[string]interface{}{
		"room_id":         room["room_id"],
		"fee_name":        "租金",
		"fee_value":       room["room_rent"],
		"fee_price":       room["room_rent"],
		"fee_used":        room["room_rent"],
		"fee_date":        fee_date,
		"fee_create_time": int32(time.Now().Unix()),
	})

	tx.Commit()

	core.Success(c, "添加成功", nil)
}
