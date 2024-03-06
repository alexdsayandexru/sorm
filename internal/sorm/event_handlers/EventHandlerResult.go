package event_handlers

const (
	InvalidArgument = "INVALID_ARGUMENT"
	Ok              = "OK"
)

type EventHandlerResult struct {
	Code    int32
	Message string
	Error   error
}
