package factory

import (
	sorm "github.com/alexdsayandexru/sorm/gen"
	"github.com/alexdsayandexru/sorm/internal/sorm/models"
)

func NewLoginUser(s *sorm.LoginUserRequest) *models.LoginUser {
	return &models.LoginUser{}
}
