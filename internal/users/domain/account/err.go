package account

import "errors"

type Error struct {
	Err  error
	Code int
}

func (e Error) Error() string {
	return e.Err.Error()
}

var (
	ErrUnknown           = Error{Code: 10000, Err: errors.New("unknown error")}       // 未知错误
	ErrPhoneEmpty        = Error{Code: 10001, Err: errors.New("phone is empty")}      // 登录名为空
	ErrPasswordEmpty     = Error{Code: 10002, Err: errors.New("password is empty")}   // 密码为空
	ErrCreatebyEmpty     = Error{Code: 10003, Err: errors.New("create_by is	 empty")} // 创建人为空
	ErrUserNotFound      = Error{Code: 10004, Err: errors.New("user not found")}      // 用户不存在
	ErrUserAlreadyExists = Error{Code: 10005, Err: errors.New("user already exists")} // 用户已存在
	ErrPhoneAlreadyExist = Error{Code: 10006, Err: errors.New("phone already exist")} // 手机号已存在
	ErrEmailAlreadyExist = Error{Code: 10007, Err: errors.New("email already exist")} // 邮箱已存在
	ErrPasswordInvalid   = Error{Code: 10008, Err: errors.New("password invalid")}    // 密码不正确
	ErrUserCreateFailed  = Error{Code: 10009, Err: errors.New("user create failed")}  // 用户创建失败
	ErrUserUpdateFailed  = Error{Code: 10010, Err: errors.New("user update failed")}  // 用户更新失败
	ErrUserDeleteFailed  = Error{Code: 10011, Err: errors.New("user delete failed")}  // 用户删除失败
)
