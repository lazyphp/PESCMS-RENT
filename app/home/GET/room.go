package GET

import (
	"reflect"

	"pescms-rent/core/db"
	"pescms-rent/core/route"
	"pescms-rent/model"

	core "pescms-rent/core/func"

	"github.com/gin-gonic/gin"
)

type Room struct {
	Content
}

func init() {
	path := Room{}
	route.Register(&path, reflect.TypeOf(path).PkgPath())
}

func (api *Room) Index(c *gin.Context) {
	total, _, fieldList, modelInfo, _, status := api.Content.Index(c, false)

	if !status {
		return
	}

	var contentList []map[string]interface{}
	db.DB().Table("pes_room").Where("room_house_id = ?", c.Query("houseID")).Order("room_listsort ASC, room_id DESC").Find(&contentList)

	feeReules, _ := model.GetFeeRules()

	core.Success(c, "ok", gin.H{
		"pageTotal":   total,
		"modelInfo":   modelInfo,
		"contentList": contentList,
		"field":       fieldList,
		"feeReules":   feeReules,
	})
}

func (api *Room) Action(c *gin.Context) {
	_, modelInfo, content, fieldList, status := api.Content.Action(c, false)

	if !status {
		return
	}

	feeReules, _ := model.GetFeeRules()

	core.Success(c, "ok", gin.H{
		"modelInfo": modelInfo,
		"content":   content,
		"field":     fieldList,
		"feeReules": feeReules,
	})
}

func (api *Room) Fee(c *gin.Context) {
	id := c.Query("roomID")

	var feeList []model.Room_fee
	db.DB().Table("pes_room_fee").Select("*, strftime('%Y-%m', fee_date, 'unixepoch') AS fee_date").Where("room_id = ?", id).Order("fee_date DESC").Find(&feeList)

	feeReules, _ := model.GetFeeRules()

	core.Success(c, "ok", gin.H{
		"room":      model.GetRoomWithID(id),
		"feeList":   feeList,
		"feeReules": feeReules,
	})
}
