package repository

import (
	"context"
	"database/sql"
	"sclearn/learn-golang/model/domain"
)

type UserRepository interface{
	Create(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Delete(ctx context.Context, tx *sql.Tx, user domain.User)
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.User, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.User
}