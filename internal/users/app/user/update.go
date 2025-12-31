package user

import (
	"github.com/xiaohangshuhub/xiaohangshu/internal/users/domain/account"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// UpdateCmd 更新用户命令,包含更新用户所需的信息
type UpdateCmd struct {
	ID       string  `json:"id"`
	Nikename *string `json:"nikename"`
}

type UpdateCmdHandler struct {
	*account.Manager
	*gorm.DB
	*zap.Logger
}

func NewUpdateCmdHandler(m *account.Manager, repo *gorm.DB, zap *zap.Logger) *UpdateCmdHandler {
	return &UpdateCmdHandler{
		Manager: m,
		DB:      repo,
		Logger:  zap,
	}
}

func (h *UpdateCmdHandler) Handle(cmd UpdateCmd) (bool, error) {

	return true, nil
}
