package userhandlers

import (
	"auth-api-example/gen/models"
	"auth-api-example/gen/restapi/operations/users"
	"database/sql"
	"log"

	"github.com/go-openapi/runtime/middleware"
)

type UsersQueryResponse struct {
	id int64
	username string
	email string
	firstname string
	lastname string
} 

type GetUsersImpl struct {
	dbClient *sql.DB
}

func NewGetUserImpl(db *sql.DB) users.GetUsersHandler{
	return &GetUsersImpl{
		dbClient: db,
	}
}

func ( impl *GetUsersImpl ) Handle(params users.GetUsersParams) middleware.Responder{

	rows, err := impl.dbClient.Query("SELECT * FROM users")

	defer rows.Close()

	if( err != nil ) {

		log.Fatalf("Database Error: %s", err)

		errorPayload := &models.Default{
			Status: "error",
			Data: "Something goes wrong!",
		}

		return users.NewCreateUserDefault(500).WithPayload(errorPayload)

	} else {

		var realPayload []*models.User

		for rows.Next(){

			var user UsersQueryResponse
			
			err := rows.Scan(&user.id, &user.username, &user.email, &user.firstname, &user.lastname);

			if err != nil {
				log.Fatal(err)
			}

			userResponse := models.User{
				ID: user.id,
				Username: user.username,
				Email: user.email,
				Firstname: user.firstname,
				Lastname: user.lastname,
			}
			realPayload = append(realPayload, &userResponse)
		}


		response := &models.UsersResponse{
			Status: "success",
			Data: realPayload,
		}

		return users.NewGetUsersOK().WithPayload(response)
	}
	
}
