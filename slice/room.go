package slice

import (
	"encoding/json"
	"fmt"
	"net/http"

	"pescms-rent/core/route"
	"pescms-rent/model"

	core "pescms-rent/core/func"

	"github.com/gin-gonic/gin"
	"github.com/iancoleman/orderedmap"
)

type Room struct{}

func init() {
	route.RegSlice(&Room{})
}

/**
 * 获取房间列表前置判断
 */
func (api *Room) Base(c *gin.Context) {
	houseID := c.Query("houseID")

	if houseID == "" {
		core.Error(c, "请提交房子的ID", 1)
		return
	}

	allowedMethods := map[string]bool{
		http.MethodPost: true,
		http.MethodPut:  true,
	}

	if allowedMethods[c.Request.Method] {
		model.PrePayload(c, func(data map[string]interface{}) error {
			feeTypeJson := orderedmap.New()

			for _, v := range data["fee_rules"].([]interface{}) {

				v := v.(string)

				feeType := data[v].(string)
				if len(feeType) == 0 || feeType == "" {
					core.Error(c, fmt.Sprintf("请填写%s的费用", v), 1)
				}

				feeTypeJson.Set(v, feeType)

			}

			jsonData, err := json.Marshal(feeTypeJson)
			if err != nil {
				core.Error(c, "转换计费类型时出错", 1)
			}

			data["cost"] = string(jsonData)
			return nil
		})
	}
}

/**
 * 房间收租请求前置判断
 */
func (api *Room) Fee(c *gin.Context) {
	id := c.Query("roomID")
	if id == "" {
		core.Error(c, "请提交房间的ID", 1)
		return
	}
}
