package grpc

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/alexdsayandexru/sorm/internal/kafka"
	"github.com/alexdsayandexru/sorm/internal/sorm/models"
	"google.golang.org/grpc/codes"
)

const (
	InvalidArgument = "INVALID_ARGUMENT"
	Ok              = "OK"
	Aborted         = "ABORTED"
)

type EventHandlerResult struct {
	Code    int32
	Message string
	Error   error
}

type EventHandler struct {
	ctx      context.Context
	target   models.IEntity
	producer kafka.IProducer
}

func NewEventHandler(ctx context.Context, target models.IEntity, producer kafka.IProducer) *EventHandler {
	return &EventHandler{
		ctx:      ctx,
		target:   target,
		producer: producer,
	}
}

func (h *EventHandler) Handle() (bool, EventHandlerResult) {
	ok, e := models.Validate(h.target)
	if !ok {
		return false, EventHandlerResult{Code: int32(codes.InvalidArgument), Message: InvalidArgument, Error: errors.New(e.ToJson())}
	}

	ok, err := h.Send()
	if !ok {
		return false, EventHandlerResult{Code: int32(codes.Aborted), Message: Aborted, Error: err}
	}

	return true, EventHandlerResult{Code: int32(codes.OK), Message: Ok}
}

func (h *EventHandler) Send() (bool, error) {
	return true, nil
	bites, _ := json.Marshal(h.target)
	return h.producer.Send(bites)
}
