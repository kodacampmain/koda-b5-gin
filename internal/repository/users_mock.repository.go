package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/kodacampmain/koda-b5-gin/internal/dto"
	"github.com/kodacampmain/koda-b5-gin/internal/model"
)

type userRepositoryMock struct{}

func NewUserRepositoryMock() *userRepositoryMock {
	return &userRepositoryMock{}
}

func (u userRepositoryMock) GetUsers(ctx context.Context) ([]model.User, error) {
	image := "a.jpg"
	return []model.User{
		{
			Id:         1,
			Gender:     "P",
			Email:      "mail@mail.com",
			ProfileImg: &image,
			Password:   "blabla",
			Role:       "admin",
		},
	}, nil
}
func (u userRepositoryMock) CreateNewUser(ctx context.Context, newUser dto.NewUser, hashedPwd string) (model.User, error) {
	return model.User{}, errors.New("failed to create user")
}
func (u userRepositoryMock) GetPwdFromEmail(ctx context.Context, email string) (model.User, error) {
	return model.User{}, errors.New("failed to get user")
}
func (u userRepositoryMock) EditProfile(ctx context.Context, profileImg string, id int) (pgconn.CommandTag, error) {
	return pgconn.NewCommandTag("UPDATE 0"), nil
}
