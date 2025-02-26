package controller

import (
	"kukuh/go-restful/helper"
	"kukuh/go-restful/model/domain/web"
	"kukuh/go-restful/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UserController interface {
	CreateNewUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	UpdateUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	DeleteUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindUserById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindUserByEmail(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAllUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params)

	Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	UpdateUserOwn(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (c *UserControllerImpl) CreateNewUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userCreateRequest := web.CreateUserRequest{}
	helper.ReadFromRequestBody(request, &userCreateRequest)

	userResponse := c.UserService.CreateNewUser(request.Context(), userCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (c *UserControllerImpl) UpdateUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userUpdateRequest := web.UpdateUserRequest{}
	helper.ReadFromRequestBody(request, &userUpdateRequest)

	userId := params.ByName("userId")

	userUpdateRequest.Id = userId

	userResponse := c.UserService.UpdateUser(request.Context(), userUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (c *UserControllerImpl) DeleteUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("userId")

	c.UserService.DeleteUser(request.Context(), userId)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (c *UserControllerImpl) FindUserById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("userId")

	userResponse := c.UserService.FindUserById(request.Context(), userId)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (c *UserControllerImpl) FindUserByEmail(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	email := params.ByName("email")

	userResponse := c.UserService.FindUserByEmail(request.Context(), email)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (c *UserControllerImpl) FindAllUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userResponses := c.UserService.FindAllUser(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (c *UserControllerImpl) Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	loginRequest := web.LoginUserRequest{}
	helper.ReadFromRequestBody(request, &loginRequest)

	loginResponse := c.UserService.Login(request.Context(), loginRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   loginResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)

}
func (c *UserControllerImpl) UpdateUserOwn(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := request.Context().Value("authId").(string)

	userUpdateRequest := web.UpdateUserRequest{}
	helper.ReadFromRequestBody(request, &userUpdateRequest)
	userUpdateRequest.Id = userId

	userResponse := c.UserService.UpdateUserOwn(request.Context(), userUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
