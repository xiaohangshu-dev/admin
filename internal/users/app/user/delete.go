package user

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// DeleteCmd 删除用户命令,包含删除用户所需的信息
type DeleteCmd struct {
	ID string `json:"id"`
}

type DeleteCmdHandler struct {
	*gorm.DB
	*zap.Logger
}

func NewDeleteCmdHandler(repo *gorm.DB, zap *zap.Logger) *DeleteCmdHandler {
	return &DeleteCmdHandler{
		DB:     repo,
		Logger: zap,
	}
}

func (h *DeleteCmdHandler) Handle(cmd DeleteCmd) (bool, error) {

	return true, nil
}
