package core

// HTTPError 是一个包含错误消息和HTTP状态码的自定义类型
type HTTPError struct {
	Code int    // HTTP状态码
	Msg  string // 错误消息
}

//error 接口
func (e *HTTPError) Error() string {
	return e.Msg
}

// Status 返回状态码
func (e *HTTPError) Status() int {
	return e.Code
}

// NewHTTPError 创建一个新的HTTPError实例
func SetError(msg string, code int) *HTTPError {
	return &HTTPError{
		Msg:  msg,
		Code: code,
	}
}
