package userhandlers

import (
	"auth-api-example/gen/models"
	"auth-api-example/gen/restapi/operations/users"

	"github.com/go-openapi/runtime/middleware"
)

type GetUsersImpl struct {}

func NewGetUserImpl() users.GetUsersHandler{
	return &GetUsersImpl{}
}

func ( impl *GetUsersImpl ) Handle(params users.GetUsersParams) middleware.Responder{

	fakePayload := &models.UsersResponse{
		Status: "success",
		Data: []*models.User{
			{
				ID: 1,
				Username: "rootspyro",
				Email: "root.spyro@gmail.com",
				Firstname: "Spyridon",
				Lastname: "Mihalopoulos",
			},
		},
	}

	return users.NewGetUsersOK().WithPayload(fakePayload)
}
