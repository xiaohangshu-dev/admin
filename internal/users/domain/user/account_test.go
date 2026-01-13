package user

import (
	"testing"

	"github.com/google/uuid"
)

// Test_newAccount_should_return_account_and_nil_error 测试 newAccount() 方法是否能正确返回 Account 实例和 nil 错误
func Test_newAccount_should_return_account_and_nil_error(t *testing.T) {

	// 准备测试数据
	username := "admin"
	nickname := "管理员"
	password := "123456"
	avatar := "http://www.example.com/avatar.jpg"

	// 调用 newAccount() 方法
	account, error := newAccount(username, nickname, avatar, password, uuid.New(), Male)

	// 断言
	if error != nil || account == nil {
		t.Error("newAccount() error should be nil")
	}
	if account.Username != username {
		t.Error("account.Username should be ", username)
	}
	if account.Nickname != nickname {
		t.Error("account.Nickname should be ", nickname)
	}
	if account.Avatar != avatar {
		t.Error("account.Avatar should be ", avatar)
	}
	if account.Gender != Male {
		t.Error("account.Gender should be ", Male)
	}
	if account.Salt != "" {
		t.Error("account.Salt should be empty")
	}
	if account.Pwd != "" {
		t.Error("account.Pwd should be empty")
	}
	if account.CreateBy == uuid.Nil {
		t.Error("account.CreateBy should not be nil")
	}
	if account.CreatedAt.IsZero() {
		t.Error("account.CreatedAt should not be zero")
	}
}
