package GET

import (
	"pescms-rent/core/db"
	"pescms-rent/core/route"
	"reflect"

	core "pescms-rent/core/func"

	"github.com/gin-gonic/gin"
)

type Node struct{}

func init() {
	path := Node{}
	route.Register(&path, reflect.TypeOf(path).PkgPath())
}

func (api *Node) Index(c *gin.Context) {
	var node []map[string]interface{}
	db.DB().Table("pes_node").Find(&node)

	// fmt.Println(node)

	core.Success(c, "ok", gin.H{
		"menu": node,
	})
}
