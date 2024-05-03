package repository

import (
	"cats-social/model/database"
	"context"
)

type CatRepositoryInterface interface {
	GetCatById(id int) (response database.Cat, err error)
	CreateCat(ctx context.Context, data database.Cat) (err error)
}

type UserRepositoryInterface interface {
	GetUserByEmail(ctx context.Context, email string) (response database.User, err error)
	CreateUser(ctx context.Context, data database.User) (err error)
}
