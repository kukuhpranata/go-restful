package repository

import (
	"context"
	"database/sql"
	"errors"
	"kukuh/go-restful/helper"
	"kukuh/go-restful/model/domain"
)

type UserRepository interface {
	Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Delete(ctx context.Context, tx *sql.Tx, userId int)
	FindById(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error)
	FindByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.User, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.User
}

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (r UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	query := "INSERT INTO users(email, password, name) values (?, ?, ?)"
	result, err := tx.ExecContext(ctx, query, user.Email, user.Password, user.Name)
	helper.PanicIfError(err)

	userId, err := result.LastInsertId()
	helper.PanicIfError(err)

	user.Id = int(userId)
	return user
}

func (r UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	query := "UPDATE users set email = ?, password = ?, name = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, user.Email, user.Password, user.Name, user.Id)
	helper.PanicIfError(err)

	return user
}

func (r UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, userId int) {
	query := "DELETE from users WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, userId)
	helper.PanicIfError(err)

	helper.PanicIfError(err)
}

func (r UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error) {
	query := "SELECT * FROM users WHERE id = ?"
	rows, err := tx.QueryContext(ctx, query, userId)
	helper.PanicIfError(err)

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Email, &user.Password, &user.Name)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("user not found")
	}
}

func (r UserRepositoryImpl) FindByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.User, error) {
	query := "SELECT * FROM users WHERE email = ?"
	rows, err := tx.QueryContext(ctx, query, email)
	helper.PanicIfError(err)

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Email, &user.Password, &user.Name)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("user not found")
	}
}

func (r UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.User {
	query := "SELECT * FROM users"
	rows, err := tx.QueryContext(ctx, query)
	helper.PanicIfError(err)

	var users []domain.User
	if rows.Next() {
		user := domain.User{}
		err := rows.Scan(&user.Id, &user.Email, &user.Password, &user.Name)
		helper.PanicIfError(err)
		users = append(users, user)
	}
	return users
}
