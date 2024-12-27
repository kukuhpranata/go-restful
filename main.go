package main

import (
	"kukuh/go-restful/app"
	"kukuh/go-restful/controller"
	"kukuh/go-restful/helper"
	"kukuh/go-restful/repository"
	"kukuh/go-restful/service"
	"net/http"

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

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
