package GET

import (
	"pescms-rent/core/db"
	"pescms-rent/core/route"
	"reflect"

	core "pescms-rent/core/func"

	"github.com/gin-gonic/gin"
)

type House struct {
	Content
}

func init() {
	path := House{}
	route.Register(&path, reflect.TypeOf(path).PkgPath())
}

func (api *House) Index(c *gin.Context) {
	total, _, fieldList, modelInfo, contentList, status := api.Content.Index(c, false)
	if !status {
		return
	}

	var room []map[string]interface{}
	db.DB().Table("pes_room").Select("count(room_status) AS total, room_house_id, room_status").Group("room_house_id, room_status").Find(&room)

	roomStatus := make(map[int64]map[int64]int64)
	for _, r := range room {
		houseID := r["room_house_id"].(int64)
		status := r["room_status"].(int64)
		total := r["total"].(int64)

		// 如果该houseID不存在，初始化一个新的map
		if _, exists := roomStatus[houseID]; !exists {
			roomStatus[houseID] = make(map[int64]int64)
		}

		// 设置room_status对应的total
		roomStatus[houseID][status] = total
	}

	core.Success(c, "ok", gin.H{
		"pageTotal":   total,
		"modelInfo":   modelInfo,
		"contentList": contentList,
		"field":       fieldList,
		"roomStatus":  roomStatus,
	})
}
