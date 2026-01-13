package service

import (
	"context"
	"errors"

	"github.com/kodacampmain/koda-b5-gin/internal/dto"
	"github.com/kodacampmain/koda-b5-gin/internal/err"
	"github.com/kodacampmain/koda-b5-gin/internal/repository"
	"github.com/kodacampmain/koda-b5-gin/pkg"
)

var ErrInvalidGender = errors.New("invalid gender")

// type User struct {
// 	Id       int    `json:"-"`
// 	Email    string `json:"email"`
// 	Password string `json:"pwd"`
// 	Role     string `json:"role"`
// }

// var users = make([]User, 0)
// var id = 1

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
			Email:  v.Email,
			Gender: v.Gender,
		})
	}
	return response, nil
}

// func (u UserService) AddUser(ctx context.Context, newUser dto.NewUser) (dto.User, error) {
// 	validGender := []string{"L", "P"}
// 	if !slices.Contains(validGender, newUser.Gender) {
// 		return dto.User{}, ErrInvalidGender
// 	}
// 	data, err := u.userRepository.CreateNewUser(ctx, newUser)
// 	if err != nil {
// 		log.Println(err.Error())
// 		return dto.User{}, err
// 	}
// 	response := dto.User{
// 		Id:     data.Id,
// 		Gender: data.Gender,
// 		Name:   data.Name,
// 	}
// 	return response, nil
// }

func (u UserService) UpdateImage(ctx context.Context, profileImg string, id int) error {
	cmd, e := u.userRepository.EditProfile(ctx, profileImg, id)
	if e != nil {
		return e
	}
	if cmd.RowsAffected() == 0 {
		return err.ErrNoRowsUpdated
	}
	return nil
}

func (u UserService) Register(ctx context.Context, newUser dto.NewUser) (dto.User, error) {
	hc := pkg.HashConfig{}
	hc.UseRecommended()

	hp, err := hc.GenHash(newUser.Password)
	if err != nil {
		return dto.User{}, err
	}

	// users = append(users, User{
	// 	Id:       id,
	// 	Email:    newUser.Email,
	// 	Password: hp,
	// 	Role:     newUser.Role,
	// })
	// id++
	data, e := u.userRepository.CreateNewUser(ctx, newUser, hp)
	if e != nil {
		return dto.User{}, e
	}
	user := dto.User{
		Email:  data.Email,
		Gender: data.Gender,
		Role:   data.Role,
		Id:     data.Id,
	}
	return user, nil
}

func (u UserService) Login(ctx context.Context, email string, password string) (dto.User, error) {
	// log.Println(users)
	// var hp string
	// for _, v := range users {
	// 	// log.Println(v.Email, newUser.Email)
	// 	if v.Email == newUser.Email {
	// 		hp = v.Password
	// 		newUser.Id = v.Id
	// 		newUser.Role = v.Role
	// 	}
	// }
	// log.Println(hp, len(hp))
	user, e := u.userRepository.GetPwdFromEmail(ctx, email)
	if e != nil {
		return dto.User{}, e
	}
	// if len(hp) == 0 {
	// 	return false, errors.New("email/password is wrong")
	// }
	hc := pkg.HashConfig{}
	_, e = hc.ComparePwdAndHash(password, user.Password)
	if e != nil {
		return dto.User{}, e
	}
	return dto.User{
		Id:   user.Id,
		Role: user.Role,
	}, nil
}

func (u UserService) GenJWTToken(user dto.User) (string, error) {
	claims := pkg.NewJWTClaims(user.Id, user.Role)
	return claims.GenToken()

}
