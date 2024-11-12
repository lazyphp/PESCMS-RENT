package route

import (
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
)

// 定义切片配置结构体
type SliceStruct struct {
	Method interface{} `yaml:"method"` // 允许为字符串或数组
	Routes []string    `yaml:"routes"`
	Files  []string    `yaml:"files"`
	Ignore []string    `yaml:"ignore"`
}

type SliceConfig struct {
	Slices map[string]SliceStruct `yaml:"slices"`
}

type SliceFunc struct {
	Path     string         // 切片文件路径
	Method   reflect.Value  // 切片方法
	funcName string         // 方法名称
	Args     []reflect.Type // 切片参数
}

var SliceFuncList = []SliceFunc{}

// 定义中间件函数
func SliceMiddleware(configPath string) gin.HandlerFunc {
	// 加载 YAML 配置文件
	data, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Printf("无法读取切片YAML文件: %v\n", err)
		return nil
	}

	// 解析 YAML 数据
	var config SliceConfig
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		fmt.Printf("解包切片YAML文件出错: %v\n", err)
		return nil
	}

	return func(c *gin.Context) {
		path := c.Request.URL.Path
		method := c.Request.Method

		for _, slice := range config.Slices {
			// 方法匹配逻辑
			if isMatchMethod(slice.Method, method) {

				// 检查是否应该忽略此路由
				if isIgnored(slice.Ignore, path) {
					continue
				}

				// 遍历定义的路由进行匹配
				for _, route := range slice.Routes {
					if matchRoute(route, path) {
						// 匹配切片路由成功，执行对应的文件方法
						for _, file := range slice.Files {
							executeFile(file, c)
						}
					}
				}
			}
		}
		// 继续执行下一个中间件或请求处理器
		c.Next()
	}
}

// 检查请求方法是否匹配
func isMatchMethod(methods interface{}, currentMethod string) bool {
	switch v := methods.(type) {
	case string:
		return strings.ToLower(v) == "any" || strings.ToLower(v) == strings.ToLower(currentMethod)
	case []interface{}:
		for _, method := range v {
			if strings.ToUpper(method.(string)) == strings.ToUpper(currentMethod) {
				return true
			}
		}
	}
	return false
}

// 路由匹配校验
func matchRoute(route, path string) bool {
	// 替换路径模式中的 :param 为正则表达式的捕获组 ([^/]+)
	routePattern := route
	routePattern = strings.ReplaceAll(routePattern, `:g`, `([^/]+)`)
	routePattern = strings.ReplaceAll(routePattern, `:m`, `([^/]+)`)
	routePattern = strings.ReplaceAll(routePattern, `:a`, `([^/]+)`)

	// 构造完整的正则表达式，匹配整个路径
	routePattern = "^" + routePattern + "$"

	// 打印调试信息
	// fmt.Printf("routePattern: %s\n", routePattern)
	// fmt.Printf("path: %s\n", path)

	// 编译正则表达式
	re, err := regexp.Compile(routePattern)
	if err != nil {
		fmt.Printf("路由正则转换失败 %v\n", err)
		return false
	}

	// 使用正则表达式匹配路径
	match := re.MatchString(path)

	// 打印匹配结果
	// fmt.Printf("result: %v\n", match)

	return match
}

// 检查请求路径是否应该被忽略
func isIgnored(ignoreRoutes []string, path string) bool {
	for _, route := range ignoreRoutes {
		if matchRoute(route, path) {
			return true
		}
	}
	return false
}

/**
 * 执行切片文件方法
 */
func executeFile(file string, c *gin.Context) {
	splitFile := strings.Split(file, ".")
	for _, obj := range SliceFuncList {

		if obj.Path != splitFile[0] && obj.funcName != splitFile[1] {
			continue
		}

		initArg := len(obj.Args)
		arguments := make([]reflect.Value, initArg)
		arguments[0] = reflect.ValueOf(c)

		obj.Method.Call(arguments)
	}
}

/**
 * 注册切片方法
 */
func RegSlice(slice interface{}) {
	sliceType := reflect.TypeOf(slice)

	if sliceType.Kind() == reflect.Ptr {
		// 解引操作
		sliceType = sliceType.Elem()
	}

	slicePath := sliceType.String()
	if strings.Contains(slicePath, ".") {
		splitSliceName := strings.Split(slicePath, ".")

		slicePath = splitSliceName[0] + "/" + strings.ToLower(splitSliceName[1])
	}

	v := reflect.ValueOf(slice)
	for i := 0; i < v.NumMethod(); i++ {
		method := v.Method(i)
		action := v.Type().Method(i).Name

		params := make([]reflect.Type, 0, v.NumMethod())
		for j := 0; j < method.Type().NumIn(); j++ {
			params = append(params, method.Type().In(j))
		}

		sliceList := SliceFunc{Path: slicePath, Method: method, Args: params, funcName: action}
		SliceFuncList = append(SliceFuncList, sliceList)
	}
}
