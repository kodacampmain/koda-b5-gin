package service

import (
	"context"
	"errors"
	"log"
	"slices"

	"github.com/kodacampmain/koda-b5-gin/internal/dto"
	"github.com/kodacampmain/koda-b5-gin/internal/repository"
	"github.com/kodacampmain/koda-b5-gin/pkg"
)

var ErrInvalidGender = errors.New("invalid gender")

type User struct {
	Id       int    `json:"-"`
	Email    string `json:"email"`
	Password string `json:"pwd"`
	Role     string `json:"role"`
}

var users = make([]User, 0)
var id = 1

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

func (u UserService) Register(newUser User) error {
	hc := pkg.HashConfig{}
	hc.UseRecommended()

	hp, err := hc.GenHash(newUser.Password)
	if err != nil {
		return err
	}

	users = append(users, User{
		Id:       id,
		Email:    newUser.Email,
		Password: hp,
		Role:     newUser.Role,
	})
	id++

	return nil
}

func (u UserService) Login(newUser *User) (bool, error) {
	// log.Println(users)
	var hp string
	for _, v := range users {
		// log.Println(v.Email, newUser.Email)
		if v.Email == newUser.Email {
			hp = v.Password
			newUser.Id = v.Id
			newUser.Role = v.Role
		}
	}
	// log.Println(hp, len(hp))
	if len(hp) == 0 {
		return false, errors.New("email/password is wrong")
	}
	hc := pkg.HashConfig{}
	return hc.ComparePwdAndHash(newUser.Password, hp)
}

func (u UserService) GenJWTToken(user User) (string, error) {
	claims := pkg.NewJWTClaims(user.Id, user.Role)
	return claims.GenToken()

}
