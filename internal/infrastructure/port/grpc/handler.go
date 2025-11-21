package grpcport

import (
	"context"

	"github.com/hogartr/go-hexagonal-template/internal/application/usecase"
	"github.com/hogartr/go-hexagonal-template/internal/domain"
	"github.com/hogartr/go-hexagonal-template/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserServer struct {
	proto.UnimplementedUserServiceServer
	createUC *usecase.CreateUserUseCase
	getUc    *usecase.GetUserUseCase
}

func NewUserServer(createUC *usecase.CreateUserUseCase, getUC *usecase.GetUserUseCase) *UserServer {
	return &UserServer{
		createUC: createUC,
		getUc:    getUC,
	}
}

func (s *UserServer) CreateUser(ctx context.Context, req *proto.CreateUserRequest) (*proto.User, error) {
	cmd := usecase.CreateUserCmd{
		Name:  req.Name,
		Email: req.Email,
	}

	user, err := s.createUC.Execute(ctx, cmd)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return domainUserToProto(user), nil
}

func (s *UserServer) GetUser(ctx context.Context, req *proto.GetUserRequest) (*proto.User, error) {
	cmd := usecase.GetUserCmd{
		Id: req.Id,
	}

	user, err := s.getUc.Execute(ctx, cmd)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return domainUserToProto(user), nil
}

func domainUserToProto(u *domain.User) *proto.User {
	return &proto.User{
		Id:        u.GetID().String(),
		Name:      u.GetName(),
		Email:     u.GetEmail(),
		CreatedAt: timestamppb.New(u.GetCreatedAt().ToTime()),
		UpdatedAt: timestamppb.New(u.GetUpdatedAt().ToTime()),
	}
}
