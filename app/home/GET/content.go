package GET

import (
	"pescms-rent/core/db"
	"pescms-rent/core/route"
	"pescms-rent/model"
	"reflect"
	"strconv"

	core "pescms-rent/core/func"

	"github.com/gin-gonic/gin"
)

type Content struct{}

func init() {
	path := Content{}
	route.Register(&path, reflect.TypeOf(path).PkgPath())
}

/**
 * 获取列表内容
 * @param c GIN结构
 * @param isJson 是否返回内容
 * @return
 */
func (api *Content) Index(c *gin.Context, isJson bool) (int, string, []model.Field, model.Model, []map[string]interface{}, bool) {
	page, _ := strconv.Atoi(c.Query("page"))
	if page >= 1 {
		page -= 1
	}

	modelTable, fieldList, modelInfo, err := model.GetModelBase(c)
	if err != nil {
		core.Error(c, err.Error(), err.Status())
		return 0, "", nil, model.Model{}, nil, false
	}

	// 获取对应模型表数据。
	var pageTotal map[string]interface{}
	var contentList []map[string]interface{}

	db.DB().Table("pes_"+modelTable).Select("COUNT(*) AS total").Take(&pageTotal, nil)

	var listsort string
	listsort = ""
	for _, v := range fieldList {
		if v.Field_name == "listsort" {
			listsort = modelTable + "_listsort ASC,"
			break
		}
	}

	db.DB().Table("pes_"+modelTable).Order(listsort+modelTable+"_id DESC").Limit(modelInfo.Model_page).Offset(page*modelInfo.Model_page).Find(&contentList, nil)

	if isJson {
		core.Success(c, "ok", gin.H{
			"pageTotal":   pageTotal["total"],
			"modelInfo":   modelInfo,
			"contentList": contentList,
			"field":       fieldList,
		})
	}

	return int(pageTotal["total"].(int64)), modelTable, fieldList, modelInfo, contentList, true
}

/**
 * 获取详细的内容
 * @param c GIN结构
 * @param isJson 是否返回内容
 * @return
 */
func (api *Content) Action(c *gin.Context, isJson bool) (string, model.Model, map[string]interface{}, []model.Field, bool) {
	modelTable, fieldList, modelInfo, err := model.GetModelBase(c)
	if err != nil {
		core.Error(c, err.Error(), err.Status())
		return "", model.Model{}, nil, nil, false
	}

	// 获取对应模型表数据。
	var content map[string]interface{}

	id := c.Query("id")
	if len(id) > 0 {
		db.DB().Table("pes_"+modelTable).Take(&content, modelTable+"_id = ?", id)
	}

	if isJson {
		core.Success(c, "ok", gin.H{
			"modelInfo": modelInfo,
			"content":   content,
			"field":     fieldList,
		})
	}

	return modelTable, modelInfo, content, fieldList, true
}
