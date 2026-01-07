package service

import (
	"context"
	"errors"
	"log"
	"slices"

	"github.com/kodacampmain/koda-b5-gin/internal/dto"
	"github.com/kodacampmain/koda-b5-gin/internal/repository"
)

var ErrInvalidGender = errors.New("invalid gender")

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

func (u UserService) AddUser(ctx context.Context, newUser dto.NewUser) (dto.User, error) {
	validGender := []string{"L", "P"}
	if !slices.Contains(validGender, newUser.Gender) {
		return dto.User{}, ErrInvalidGender
	}
	data, err := u.userRepository.CreateNewUser(ctx, newUser)
	if err != nil {
		log.Println(err.Error())
		return dto.User{}, err
	}
	response := dto.User{
		Id:     data.Id,
		Gender: data.Gender,
		Name:   data.Name,
	}
	return response, nil

}
