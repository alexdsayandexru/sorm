package event_handlers

import (
	"context"
	"errors"
	"github.com/alexdsayandexru/sorm/internal/sorm/models"
	"github.com/alexdsayandexru/sorm/internal/sorm/validation"
	"google.golang.org/grpc/codes"
)

type RegisterUserEventHandler struct {
	ctx    context.Context
	target *models.RegisterUser
}

func NewRegisterUserEventHandler(ctx context.Context, target *models.RegisterUser) *RegisterUserEventHandler {
	return &RegisterUserEventHandler{
		ctx:    ctx,
		target: target,
	}
}

func (h *RegisterUserEventHandler) Handle() (bool, EventHandlerResult) {
	rules := validation.GetRegisterUserValidationRules(h.target)
	ok, e := validation.ValidateRules(rules)
	if !ok {
		return false, EventHandlerResult{
			Code:    int32(codes.InvalidArgument),
			Message: InvalidArgument,
			Error:   errors.New(e.ToJson()),
		}
	}
	return true, EventHandlerResult{
		Code:    int32(codes.OK),
		Message: Ok,
	}
}
