package GET

import (
	"pescms-rent/core/route"
	"reflect"
)

type Model struct{}

func init() {
	path := Model{}
	route.Register(&path, reflect.TypeOf(path).PkgPath())
}
