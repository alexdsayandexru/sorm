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

func Handle(ctx context.Context, target models.IEntity) (*sorm.DataManagementResponse, error) {
	ok, result := NewEventHandler(
		ctx,
		target,
		models.GetValidator(),
		kafka.GetProducer("localhost", 9092, "idp.sorm.user-registration.0"),
	).Handle()

	if !ok {
		return &sorm.DataManagementResponse{Code: result.Code, Message: result.Message, Details: []*anypb.Any{{Value: []byte(result.Error.Error())}}},
			status.Errorf(codes.Code(result.Code), result.Message, result.Error.Error())
	}

	return &sorm.DataManagementResponse{Code: result.Code, Message: result.Message}, nil
}

func (UserDataManagementServerImpl) RegisterUser(ctx context.Context, request *sorm.RegisterUserRequest) (*sorm.DataManagementResponse, error) {
	target := factory.NewRegisterUser(request)
	return Handle(ctx, target)
}

func (UserDataManagementServerImpl) LoginUser(ctx context.Context, request *sorm.LoginUserRequest) (*sorm.DataManagementResponse, error) {
	target := factory.NewLoginUser(request)
	return Handle(ctx, target)
}

func (UserDataManagementServerImpl) LogoutUser(ctx context.Context, request *sorm.LogoutUserRequest) (*sorm.DataManagementResponse, error) {
	target := factory.NewLogoutUser(request)
	return Handle(ctx, target)
}

func (UserDataManagementServerImpl) DeleteUserAccount(ctx context.Context, request *sorm.DeleteUserAccountRequest) (*sorm.DataManagementResponse, error) {
	target := factory.NewDeleteUserAccount(request)
	return Handle(ctx, target)
}

func (UserDataManagementServerImpl) UpdateUserData(ctx context.Context, request *sorm.UpdateUserDataRequest) (*sorm.DataManagementResponse, error) {
	target := factory.NewUpdateUserData(request)
	return Handle(ctx, target)
}

func (UserDataManagementServerImpl) DeleteUserAccountAdmin(ctx context.Context, request *sorm.DeleteUserAccountAdminRequest) (*sorm.DataManagementResponse, error) {
	target := factory.NewDeleteUserAccountAdmin(request)
	return Handle(ctx, target)
}

func (UserDataManagementServerImpl) UpdateUserDataAdmin(ctx context.Context, request *sorm.UpdateUserDataAdminRequest) (*sorm.DataManagementResponse, error) {
	target := factory.NewUpdateUserDataAdmin(request)
	return Handle(ctx, target)
}

func (UserDataManagementServerImpl) DeleteAccount(ctx context.Context, request *sorm.DeleteAccountRequest) (*sorm.DataManagementResponse, error) {
	target := factory.NewDeleteAccount(request)
	return Handle(ctx, target)
}

func (UserDataManagementServerImpl) UserAccountRecovery(ctx context.Context, request *sorm.UserAccountRecoveryRequest) (*sorm.DataManagementResponse, error) {
	target := factory.NewUserAccountRecovery(request)
	return Handle(ctx, target)
}

func (UserDataManagementServerImpl) DirectoryData(ctx context.Context, request *sorm.DirectoryDataRequest) (*sorm.DataManagementResponse, error) {
	target := factory.NewDirectoryData(request)
	return Handle(ctx, target)
}
