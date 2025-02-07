package service

import (
	"context"
	"database/sql"
	"kukuh/go-restful/exception"
	"kukuh/go-restful/helper"
	"kukuh/go-restful/helper/token"
	"kukuh/go-restful/model/domain"
	"kukuh/go-restful/model/domain/web"
	"kukuh/go-restful/repository"

	"github.com/go-playground/validator/v10"
)

type UserService interface {
	CreateNewUser(ctx context.Context, request web.CreateUserRequest) web.UserResponse
	UpdateUser(ctx context.Context, request web.UpdateUserRequest) web.UserResponse
	DeleteUser(ctx context.Context, userId int)
	FindUserById(ctx context.Context, userId int) web.UserResponse
	FindUserByEmail(ctx context.Context, email string) web.UserResponse
	FindAllUser(ctx context.Context) []web.UserResponse

	Login(ctx context.Context, request web.LoginUserRequest) web.LoginUserResponse
	UpdateUserOwn(ctx context.Context, request web.UpdateUserRequest) web.UserResponse
}

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, DB *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (s *UserServiceImpl) CreateNewUser(ctx context.Context, request web.CreateUserRequest) web.UserResponse {
	err := s.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	password, err := helper.HashPassword(request.Password)
	helper.PanicIfError(err)

	user := domain.User{
		Email:    request.Email,
		Password: password,
		Name:     request.Name,
	}

	user = s.UserRepository.Save(ctx, tx, user)

	userResponse := web.UserResponse{
		Email: user.Email,
		Name:  user.Name,
	}

	return userResponse
}

func (s *UserServiceImpl) UpdateUser(ctx context.Context, request web.UpdateUserRequest) web.UserResponse {
	err := s.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	password, err := helper.HashPassword(request.Password)

	user, err := s.UserRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	user.Email = request.Email
	user.Password = password
	user.Name = request.Name

	user = s.UserRepository.Update(ctx, tx, user)

	userResponse := web.UserResponse{
		Email: user.Email,
		Name:  user.Name,
	}

	return userResponse
}

func (s *UserServiceImpl) DeleteUser(ctx context.Context, userId int) {
	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	_, err = s.UserRepository.FindById(ctx, tx, userId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	s.UserRepository.Delete(ctx, tx, userId)
}

func (s *UserServiceImpl) FindUserById(ctx context.Context, userId int) web.UserResponse {
	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := s.UserRepository.FindById(ctx, tx, userId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	userResponse := web.UserResponse{
		Email: user.Email,
		Name:  user.Name,
	}

	return userResponse
}

func (s *UserServiceImpl) FindUserByEmail(ctx context.Context, email string) web.UserResponse {
	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := s.UserRepository.FindByEmail(ctx, tx, email)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	userResponse := web.UserResponse{
		Email: user.Email,
		Name:  user.Name,
	}

	return userResponse
}

func (s *UserServiceImpl) FindAllUser(ctx context.Context) []web.UserResponse {
	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	users := s.UserRepository.FindAll(ctx, tx)

	var userResponses []web.UserResponse
	for _, user := range users {
		userResponse := web.UserResponse{
			Email: user.Email,
			Name:  user.Name,
		}
		userResponses = append(userResponses, userResponse)
	}

	return userResponses
}

func (s *UserServiceImpl) Login(ctx context.Context, request web.LoginUserRequest) web.LoginUserResponse {
	err := s.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := s.UserRepository.FindByEmail(ctx, tx, request.Email)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	err = helper.CheckPasswordHash(request.Password, user.Password)
	helper.PanicIfError(err)

	token, err := token.GenerateJwtToken(user.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	loginResponse := web.LoginUserResponse{
		Name:  user.Name,
		Email: user.Email,
		Token: token,
	}

	return loginResponse
}

func (s *UserServiceImpl) UpdateUserOwn(ctx context.Context, request web.UpdateUserRequest) web.UserResponse {
	err := s.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := s.UserRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	password, err := helper.HashPassword(request.Password)
	helper.PanicIfError(err)

	user.Email = request.Email
	user.Password = password
	user.Name = request.Name

	user = s.UserRepository.Update(ctx, tx, user)

	userResponse := web.UserResponse{
		Email: user.Email,
		Name:  user.Name,
	}

	return userResponse
}
