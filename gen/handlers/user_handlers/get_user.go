package userhandlers

import (
	"auth-api-example/gen/models"
	"auth-api-example/gen/restapi/operations/users"

	"github.com/go-openapi/runtime/middleware"
)

type GetSingleUserImpl struct {}

func NewSingleUserImpl() users.GetUserByUsernameHandler {
	return &GetSingleUserImpl{}
}

func (impl *GetSingleUserImpl) Handle(params users.GetUserByUsernameParams) middleware.Responder {

	if(params.Username == "rootspyro") {

		fakePayload := &models.UserResponse{
			Status: "success",
			Data: &models.User{
				ID: 1,
				Username: "rootspyro",
				Email: "root.spyro@gmail.com",
				Firstname: "Spyridon",
				Lastname: "Mihalopoulos",
			},
		}

		return users.NewGetUserByUsernameOK().WithPayload(fakePayload)

	} else {
		fakePayload := &models.Default{
			Status: "error",
			Data: "Username " + params.Username + " not found!",
		}

		return users.NewGetUserByUsernameDefault(404).WithPayload(fakePayload)
	}

}
