package model

import "pescms-rent/core/db"

/**
 * 基于房间ID获取对应房间信息
 */
func GetRoomWithID(id string) map[string]interface{} {
	var room map[string]interface{}
	db.DB().Table("pes_room").Where("room_id = ?", id).Find(&room)

	return room
}
