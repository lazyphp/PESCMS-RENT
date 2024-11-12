package DELETE

import (
	"pescms-rent/core/db"
	"pescms-rent/core/route"
	"reflect"

	core "pescms-rent/core/func"

	"github.com/gin-gonic/gin"
)

type Room struct{}

func init() {
	path := Room{}
	route.Register(&path, reflect.TypeOf(path).PkgPath())
}

func (api *Room) Fee(c *gin.Context) {
	roomID := c.Query("roomID")

	var dataJson map[string]interface{}

	if err := c.ShouldBindJSON(&dataJson); err != nil {
		core.Error(c, "获取删除记录失败", 1)
		return
	}

	if dataJson["date"] == nil {
		core.Error(c, "请提交要删除的日期", 1)
		return
	}

	fee_date, err := core.Strtotime(dataJson["date"].(string))
	if err != nil {
		core.Error(c, "请提交正确的日期格式", 1)
		return
	}

	var data map[string]interface{}
	db.DB().Table("pes_room_fee").Where("room_id = ? AND fee_date = ?", roomID, fee_date).Delete(data)

	core.Success(c, "删除成功", nil)
}
