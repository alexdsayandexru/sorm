package grpc

import (
	"context"
	"fmt"
	"github.com/alexdsayandexru/sorm/gen"
	"github.com/alexdsayandexru/sorm/internal/sorm/event_handlers"
	"github.com/alexdsayandexru/sorm/internal/sorm/factory"
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

func (UserDataManagementServerImpl) RegisterUser(ctx context.Context, request *sorm.RegisterUserRequest) (*sorm.RegisterUserResponse, error) {
	target := factory.CreateRegisterUser(request)

	ok, result := event_handlers.NewRegisterUserEventHandler(ctx, target).Handle()
	if !ok {
		return &sorm.RegisterUserResponse{Code: result.Code, Message: result.Message, Details: []*anypb.Any{{Value: []byte(result.Error.Error())}}},
			status.Errorf(codes.Code(result.Code), result.Message, result.Error.Error())
	}

	return &sorm.RegisterUserResponse{Code: result.Code, Message: result.Message}, nil
}

func (UserDataManagementServerImpl) LoginUser(context.Context, *sorm.LoginUserRequest) (*sorm.LoginUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
}

func (UserDataManagementServerImpl) LogoutUser(context.Context, *sorm.LogoutUserRequest) (*sorm.LogoutUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogoutUser not implemented")
}

func (UserDataManagementServerImpl) DeleteUserAccount(context.Context, *sorm.DeleteUserAccountRequest) (*sorm.DeleteUserAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserAccount not implemented")
}

func (UserDataManagementServerImpl) UpdateUserData(context.Context, *sorm.UpdateUserDataRequest) (*sorm.UpdateUserDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserData not implemented")
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
