package services

import (
	"context"

	"gorm.io/gorm"
	pb "order-service/pb"
	"order-service/src/models"
	"order-service/src/repositories"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
	repo *repositories.UserRepository
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		repo: repositories.NewUserRepository(db),
	}
}

func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {
	user, err := s.repo.GetByID(req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.User{
		Id:    uint32(user.ID),
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.User, error) {
	user := &models.User{
		Name:  req.Name,
		Email: req.Email,
	}
	if err := s.repo.Create(user); err != nil {
		return nil, err
	}
	return &pb.User{
		Id:    uint32(user.ID),
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (s *UserService) GetAllUsers(ctx context.Context, req *pb.Empty) (*pb.Users, error) {
	users, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	pbUsers := &pb.Users{}
	for _, user := range users {
		pbUser := &pb.User{
			Id:    uint32(user.ID),
			Name:  user.Name,
			Email: user.Email,
		}
		pbUsers.Users = append(pbUsers.Users, pbUser)
	}
	return pbUsers, nil
}
