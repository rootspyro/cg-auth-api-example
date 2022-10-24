package main

import (
	"auth-api-example/db"
	"auth-api-example/gen/restapi"
	"auth-api-example/gen/restapi/operations"
	userhandlers "auth-api-example/handlers/user_handlers"
	"auth-api-example/middlewares"
	"flag"
	"log"
	"os"
	"strconv"

	"github.com/go-openapi/loads"
	"github.com/joho/godotenv"
)



func main() {

	//load env variables
	godotenv.Load()

	// create portFlag 
	serverPort, err := strconv.Atoi(os.Getenv("PORT"))
	
	var portFlag = flag.Int("port", serverPort, "Port to run this service on")

	// load embedded swagger file
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
			log.Fatalln(err)
	}


	// create new service API
	api := operations.NewCgAuthAPIAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	//DB client

	clientBuilder := db.NewClientBuilder(
		os.Getenv("PG_HOST"),
		os.Getenv("PG_PORT"),
		os.Getenv("PG_DB"),
		os.Getenv("PG_USER"), 
		os.Getenv("PG_PASSW"),
	)

	dbClient := clientBuilder.BuildSqlClient()
	defer dbClient.Close()

	//handlers
	api.UsersGetUsersHandler = userhandlers.NewGetUserImpl(dbClient)
	api.UsersGetUserByUsernameHandler = userhandlers.NewSingleUserImpl()

	//middlerwares
	api.AddMiddlewareFor("GET", "/users", middlewares.EnsureValidToken())
	api.AddMiddlewareFor("GET", "/users/{username}", middlewares.EnsureValidToken())

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
