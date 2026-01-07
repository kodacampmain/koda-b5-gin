package service

import (
	"context"

	"github.com/kodacampmain/koda-b5-gin/internal/dto"
	"github.com/kodacampmain/koda-b5-gin/internal/repository"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (u UserService) GetUsers(ctx context.Context) ([]dto.User, error) {
	data, err := u.userRepository.GetUsers(ctx)
	if err != nil {
		return nil, err
	}
	var response []dto.User
	for _, v := range data {
		response = append(response, dto.User{
			Id:     v.Id,
			Name:   v.Name,
			Gender: v.Gender,
		})
	}
	return response, nil
}
