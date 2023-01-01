package repository

import (
	"context"
	"database/sql"
	"errors"
	"sclearn/learn-golang/helper"
	"sclearn/learn-golang/model/domain"
)

type UserRepositoryImpl struct {

}

func NewUserRepository() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, user domain.User) domain.User  {
	SQL := "INSERT INTO users(email, password) VALUES(?, ?)"
	result, err := tx.ExecContext(ctx, SQL, user.Email, user.Password)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)


	user.Id = int(id)
	return user
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User  {
	SQL := "UPDATE users SET email = ?, password = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, user.Email, user.Password, user.Id)
	helper.PanicIfError(err)

	return user
}

func (repository *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, user domain.User) {
	SQL := "DELETE FROM users WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, user.Id)
	helper.PanicIfError(err)
}

func (repository *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.User, error)  {
	SQL := "SELECT id, email, password FROM users WHERE id = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, SQL, id)
	helper.PanicIfError(err)
	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Email, &user.Password)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("User Not Found")
	}
}

func (repository *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.User  {
	SQL := "SELECT id, email, password FROM users"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		user := domain.User{}
		err := rows.Scan(&user.Id, &user.Email, &user.Password)
		helper.PanicIfError(err)

		users = append(users, user)
	}

	return users
}