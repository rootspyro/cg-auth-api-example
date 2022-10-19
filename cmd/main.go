package main

import (
	"auth-api-example/gen/restapi"
	"auth-api-example/gen/restapi/operations"
	"flag"
	"log"

	userhandlers "auth-api-example/gen/handlers/user_handlers"

	"github.com/go-openapi/loads"
)


var portFlag = flag.Int("port", 3000, "Port to run this service on")

func main() {

	// load embedded swagger file
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
			log.Fatalln(err)
	}

	// create new service API
	api := operations.NewCgAuthAPIAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	//handlers
	api.UsersGetUsersHandler = userhandlers.NewGetUserImpl()
	api.UsersGetUserByUsernameHandler = userhandlers.NewSingleUserImpl()

	// parse flags
	flag.Parse()
	// set the port this service will be run on
	server.Port = *portFlag

	// TODO: Set Handle

	// serve API
	if err := server.Serve(); err != nil {
			log.Fatalln(err)
	}
}
