package main

import (
	"context"
	"kukuh/go-restful/app"
	"kukuh/go-restful/controller"
	"kukuh/go-restful/exception"
	"kukuh/go-restful/helper"
	"kukuh/go-restful/helper/token"
	"kukuh/go-restful/repository"
	"kukuh/go-restful/service"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validate)
	userController := controller.NewUserController(userService)

	router := httprouter.New()

	router.GET("/api/users", userController.FindAllUser)
	// router.GET("/api/users/:userId", userController.FindUserById)
	router.GET("/api/users/:email", userController.FindUserByEmail)
	router.POST("/api/users", userController.CreateNewUser)
	router.PUT("/api/users/:userId", userController.UpdateUser)
	router.DELETE("/api/users/:userId", userController.DeleteUser)

	router.POST("/api/login", userController.Login)
	router.PUT("/api/auth/users", CheckAuth(userController.UpdateUserOwn))

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}

func CheckAuth(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		header := r.Header.Get("Authorization")

		bearerToken := strings.Split(header, "Bearer ")

		if len(bearerToken) != 2 {
			panic("token lenght error")
		}

		payload, err := token.ValidateJwtToken(bearerToken[1])
		if err != nil {
			helper.PanicIfError(err)
		}

		ctx := context.WithValue(r.Context(), "authId", payload.AuthId)

		next(w, r.WithContext(ctx), ps)
	}
}
