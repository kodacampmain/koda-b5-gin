package repository

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/kodacampmain/koda-b5-gin/internal/dto"
	"github.com/kodacampmain/koda-b5-gin/internal/model"
)

type UserRepo interface {
	GetUsers(ctx context.Context, db DBTX) ([]model.User, error)
	CreateNewUser(ctx context.Context, db DBTX, newUser dto.NewUser, hashedPwd string) (model.User, error)
	GetPwdFromEmail(ctx context.Context, db DBTX, email string) (model.User, error)
	EditProfile(ctx context.Context, db DBTX, profileImg string, id int) (pgconn.CommandTag, error)
	CreateNewProfile(ctx context.Context, db DBTX, userId int) (pgconn.CommandTag, error)
}

type DBTX interface {
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
}

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (u *UserRepository) GetUsers(ctx context.Context, db DBTX) ([]model.User, error) {
	sqlStr := "SELECT id, email, gender, profile_img FROM users"
	rows, err := db.Query(ctx, sqlStr)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer rows.Close()
	var users []model.User
	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.Id, &user.Email, &user.Gender, &user.ProfileImg)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (u *UserRepository) CreateNewUser(ctx context.Context, db DBTX, newUser dto.NewUser, hashedPwd string) (model.User, error) {
	sql := "INSERT INTO users (email, gender, password) VALUES ($1, $2, $3) RETURNING id, email, gender, role"
	values := []any{newUser.Email, newUser.Gender, hashedPwd}
	// return u.db.Exec(ctx, sql, values...)
	row := db.QueryRow(ctx, sql, values...)
	var user model.User
	if err := row.Scan(&user.Id, &user.Email, &user.Gender, &user.Role); err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (u *UserRepository) CreateNewProfile(ctx context.Context, db DBTX, userId int) (pgconn.CommandTag, error) {
	sql := "INSERT INTO profile(user_id) VALUES ($1)"
	return db.Exec(ctx, sql, userId)
}

func (u *UserRepository) GetPwdFromEmail(ctx context.Context, db DBTX, email string) (model.User, error) {
	sql := "SELECT id, password, role FROM users WHERE email=$1"
	values := []any{email}

	var user model.User
	if e := db.QueryRow(ctx, sql, values...).Scan(&user.Id, &user.Password, &user.Role); e != nil {
		return model.User{}, e
	}
	return user, nil
}

func (u *UserRepository) EditProfile(ctx context.Context, db DBTX, profileImg string, id int) (pgconn.CommandTag, error) {
	sql := "UPDATE users SET profile_img=$2 WHERE id=$1"
	values := []any{id, profileImg}

	return db.Exec(ctx, sql, values...)
}
