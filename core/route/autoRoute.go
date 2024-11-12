package route

import (
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
)

type Route struct {
	path       string         // url路径
	httpMethod string         // http方法 get post
	Method     reflect.Value  // 方法路由
	funcName   string         // 方法名称
	Args       []reflect.Type // 参数类型
}

// 路由集合
var Routes = []Route{}

/**
 * 注册路由
 * @param controller interface{} 控制器
 * @param path string 控制器的地址
 */
func Register(controller interface{}, controllerPath string) bool {
	// 获取路由地址
	rootPath := ""
	if strings.Contains(controllerPath, "/app") {
		// 基于/app目录进行路由地址分割
		splitPath := strings.Split(controllerPath, "/app")
		rootPath = splitPath[len(splitPath)-1]
	}
	// fmt.Println("路由目录地址:", rootPath)

	// 获取控制器的动态类
	controllerType := reflect.TypeOf(controller)
	// 判断是否是指针类型
	if controllerType.Kind() == reflect.Ptr {
		// 解引操作
		controllerType = controllerType.Elem()
	}

	// 获取对应包中[请求类型.控制器名称]
	controllerName := controllerType.String()

	// fmt.Println("获取类型", controllerType)
	// fmt.Println("解引", controllerName)
	// fmt.Println("控制器名称", controllerName)

	httpMethod := "POST" // 路由绑定默认请求方法均为POST
	if strings.Contains(controllerName, ".") {
		splitControllerName := strings.Split(controllerName, ".")

		// fmt.Println("拆分后的控制器请求方法：", strings.ToUpper(splitControllerName[0]))
		// fmt.Println("拆分后的控制器名称：", splitControllerName[1])

		controllerName = "/" + strings.ToLower(splitControllerName[1]) + "/"

		// 覆写正确的请求方法
		httpMethod = strings.ToUpper(splitControllerName[0])

		rootPath = strings.Replace(rootPath, "/"+httpMethod, "", -1)
		// fmt.Println("替换后的基础地址:", rootPath)

	}

	// fmt.Println("控制器方法", controllerName)

	v := reflect.ValueOf(controller)
	for i := 0; i < v.NumMethod(); i++ {
		method := v.Method(i)                              // 方法的指针地址 ?(存疑是内存还是指针地址)
		action := strings.ToLower(v.Type().Method(i).Name) // 当前方法的名称

		// fmt.Println("方法:", method)
		// fmt.Println("方法名:", action)

		url := rootPath + controllerName + action
		// fmt.Println("路由地址", url)

		params := make([]reflect.Type, 0, v.NumMethod())

		// 读取方法中带有的参数类型 (疑似没用的代码)
		for j := 0; j < method.Type().NumIn(); j++ {
			params = append(params, method.Type().In(j))
		}
		// fmt.Println("方法参数类型:", params)

		route := Route{path: url, Method: method, Args: params, httpMethod: httpMethod, funcName: action}
		Routes = append(Routes, route)

		// fmt.Println("总路由:", Routes)

	}
	return true
}

// 绑定路由
func Bind(e *gin.Engine) {
	for _, route := range Routes {

		// 模型管理路由转换
		if strings.HasPrefix(route.path, "/home/content/") {
			route.path = "/home/:model/" + route.funcName
		}

		if route.httpMethod == "GET" {
			e.GET(route.path, match(route.path, route))
		}
		if route.httpMethod == "POST" {
			e.POST(route.path, match(route.path, route))
		}
		if route.httpMethod == "PUT" {
			e.PUT(route.path, match(route.path, route))
		}
		if route.httpMethod == "DELETE" {
			e.DELETE(route.path, match(route.path, route))
		}
	}
}

// 根据path匹配对应的方法
func match(path string, route Route) gin.HandlerFunc {
	return func(c *gin.Context) {
		fields := strings.Split(path, "/")

		// fmt.Println("fields,len(fields)=", fields, len(fields))

		// 此处代码意思是：如果路由地址拆分后长度少于3个，则不属于控制器的路由地址
		if len(fields) < 3 {
			return
		}

		// 判断整套程序记录的路由地址
		if len(Routes) > 0 {

			initArg := len(route.Args)

			arguments := make([]reflect.Value, initArg)
			arguments[0] = reflect.ValueOf(c) // *gin.Context

			if len(route.Args) > 1 {
				arguments[1] = reflect.ValueOf(true) // 兼容Content Index和Action第二个参数
			}

			route.Method.Call(arguments)
		}
	}
}
