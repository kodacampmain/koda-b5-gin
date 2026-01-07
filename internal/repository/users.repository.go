package repository

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kodacampmain/koda-b5-gin/internal/model"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u UserRepository) GetUsers(ctx context.Context) ([]model.User, error) {
	sqlStr := "SELECT name, id, gender FROM users"
	rows, err := u.db.Query(ctx, sqlStr)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	var users []model.User
	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.Name, &user.Id, &user.Gender)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
