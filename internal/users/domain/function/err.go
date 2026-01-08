package function

import "errors"

// Error 定义了功能领域(30000-39999)的错误类型
type Error struct {
	Err  error
	Code int
}

func (e Error) Error() string {
	return e.Err.Error()
}

var (
	ErrUnknown               = &Error{Code: 30000, Err: errors.New("unknown error")}            // 未知错误
	ErrNameEmpty             = &Error{Code: 30001, Err: errors.New("function name is empty")}   // 功能名称为空
	ErrRouteEmpty            = &Error{Code: 30002, Err: errors.New("route is empty")}           // 路由为空
	ErrIconEmpty             = &Error{Code: 30003, Err: errors.New("icon is empty")}            // 图标为空
	ErrDescEmpty             = &Error{Code: 30004, Err: errors.New("description is empty")}     // 描述为空
	ErrFunctionAlreadyExists = &Error{Code: 30005, Err: errors.New("funcation already exists")} // 功能已存在
	ErrFunctionNotFound      = &Error{Code: 30006, Err: errors.New("funcation not found")}      // 功能不存在
)
