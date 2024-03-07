package grpc

import (
	"context"
	"fmt"
	"github.com/alexdsayandexru/sorm/gen"
	"github.com/alexdsayandexru/sorm/internal/kafka"
	"github.com/alexdsayandexru/sorm/internal/sorm/factory"
	"github.com/alexdsayandexru/sorm/internal/sorm/models"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/anypb"
	"net"
)

type UserDataManagementServerImpl struct {
	port int
	sorm.UnimplementedUserDataManagementServer
}

func NewServer(port int) *UserDataManagementServerImpl {
	return &UserDataManagementServerImpl{port: port}
}

func (s *UserDataManagementServerImpl) Run() error {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return err
	}
	server := grpc.NewServer()
	sorm.RegisterUserDataManagementServer(server, s)
	err = server.Serve(listen)
	return err
}

func Handle(ctx context.Context, target models.IEntity) (bool, EventHandlerResult) {
	ok, result := NewEventHandler(
		ctx,
		target,
		kafka.GetProducer("localhost", 9092, "idp.sorm.user-registration.0"),
	).Handle()
	return ok, result
}

func (UserDataManagementServerImpl) RegisterUser(ctx context.Context, request *sorm.RegisterUserRequest) (*sorm.RegisterUserResponse, error) {
	target := factory.NewRegisterUser(request)
	ok, result := Handle(ctx, target)

	if !ok {
		return &sorm.RegisterUserResponse{Code: result.Code, Message: result.Message, Details: []*anypb.Any{{Value: []byte(result.Error.Error())}}},
			status.Errorf(codes.Code(result.Code), result.Message, result.Error.Error())
	}

	return &sorm.RegisterUserResponse{Code: result.Code, Message: result.Message}, nil
}

func (UserDataManagementServerImpl) LoginUser(ctx context.Context, request *sorm.LoginUserRequest) (*sorm.LoginUserResponse, error) {
	target := factory.NewLoginUser(request)
	ok, result := Handle(ctx, target)

	if !ok {
		return &sorm.LoginUserResponse{Code: result.Code, Message: result.Message, Details: []*anypb.Any{{Value: []byte(result.Error.Error())}}},
			status.Errorf(codes.Code(result.Code), result.Message, result.Error.Error())
	}

	return &sorm.LoginUserResponse{Code: result.Code, Message: result.Message}, nil
}

func (UserDataManagementServerImpl) LogoutUser(ctx context.Context, request *sorm.LogoutUserRequest) (*sorm.LogoutUserResponse, error) {
	target := factory.NewLogoutUser(request)
	ok, result := Handle(ctx, target)

	if !ok {
		return &sorm.LogoutUserResponse{Code: result.Code, Message: result.Message, Details: []*anypb.Any{{Value: []byte(result.Error.Error())}}},
			status.Errorf(codes.Code(result.Code), result.Message, result.Error.Error())
	}

	return &sorm.LogoutUserResponse{Code: result.Code, Message: result.Message}, nil
}

func (UserDataManagementServerImpl) DeleteUserAccount(ctx context.Context, request *sorm.DeleteUserAccountRequest) (*sorm.DeleteUserAccountResponse, error) {
	target := factory.NewDeleteUserAccount(request)
	ok, result := Handle(ctx, target)

	if !ok {
		return &sorm.DeleteUserAccountResponse{Code: result.Code, Message: result.Message, Details: []*anypb.Any{{Value: []byte(result.Error.Error())}}},
			status.Errorf(codes.Code(result.Code), result.Message, result.Error.Error())
	}

	return &sorm.DeleteUserAccountResponse{Code: result.Code, Message: result.Message}, nil
}

func (UserDataManagementServerImpl) UpdateUserData(ctx context.Context, request *sorm.UpdateUserDataRequest) (*sorm.UpdateUserDataResponse, error) {
	target := factory.NewUpdateUserData(request)
	ok, result := Handle(ctx, target)

	if !ok {
		return &sorm.UpdateUserDataResponse{Code: result.Code, Message: result.Message, Details: []*anypb.Any{{Value: []byte(result.Error.Error())}}},
			status.Errorf(codes.Code(result.Code), result.Message, result.Error.Error())
	}

	return &sorm.UpdateUserDataResponse{Code: result.Code, Message: result.Message}, nil
}

func (UserDataManagementServerImpl) DeleteUserAccountAdmin(context.Context, *sorm.DeleteUserAccountAdminRequest) (*sorm.DeleteUserAccountAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserAccountAdmin not implemented")
}

func (UserDataManagementServerImpl) UpdateUserDataAdmin(context.Context, *sorm.UpdateUserDataAdminRequest) (*sorm.UpdateUserDataAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserDataAdmin not implemented")
}

func (UserDataManagementServerImpl) DeleteAccount(context.Context, *sorm.DeleteAccountRequest) (*sorm.DeleteAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAccount not implemented")
}

func (UserDataManagementServerImpl) UserAccountRecovery(context.Context, *sorm.UserAccountRecoveryRequest) (*sorm.UserAccountRecoveryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserAccountRecovery not implemented")
}

func (UserDataManagementServerImpl) DirectoryData(context.Context, *sorm.DirectoryDataRequest) (*sorm.DirectoryDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DirectoryData not implemented")
}
