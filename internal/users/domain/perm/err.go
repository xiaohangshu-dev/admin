package perm

import "errors"

// Error 定义了功能领域(20000-29999)的错误类型
type Error struct {
	Err  error
	Code int
}

func (e Error) Error() string {
	return e.Err.Error()
}

var (
	ErrUnknown               = &Error{Code: 20000, Err: errors.New("unknown error")}             // 未知错误
	ErrNameEmpty             = &Error{Code: 20001, Err: errors.New("permission name is empty")}  // 权限名称为空
	ErrRouteEmpty            = &Error{Code: 20002, Err: errors.New("route is empty")}            // 路由为空
	ErrIconEmpty             = &Error{Code: 20003, Err: errors.New("icon is empty")}             // 图标为空
	ErrDescEmpty             = &Error{Code: 20004, Err: errors.New("description is empty")}      // 描述为空
	ErrFunctionAlreadyExists = &Error{Code: 20005, Err: errors.New("permission already exists")} // 权限已存在
	ErrFunctionNotFound      = &Error{Code: 20006, Err: errors.New("permission not found")}      // 权限不存在
	ErrRoleAlreadyExists     = &Error{Code: 20007, Err: errors.New("role already exists")}       // 角色已存在
	ErrRoleNotFound          = &Error{Code: 20008, Err: errors.New("role not found")}            // 角色不存在
)
