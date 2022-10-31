package authhandlers

import (
	"auth-api-example/gen/models"
	"auth-api-example/gen/restapi/operations/auth"
	"context"
	"log"

	"github.com/go-openapi/runtime/middleware"
	"github.com/open-policy-agent/opa/rego"
)



// HANDLER
type AuthorizeImpl struct {
	opa rego.PreparedEvalQuery
}

func NewAuthImpl( query rego.PreparedEvalQuery ) auth.AuthorizeRequestHandler {
	return &AuthorizeImpl{
		opa: query,
	}
}

func ( impl *AuthorizeImpl ) Handle( parameters auth.AuthorizeRequestParams ) middleware.Responder{

	body := parameters.Body

	ctx := context.TODO()

	results, err := impl.opa.Eval(ctx, rego.EvalInput(body));


	if ( err != nil ) {
		log.Fatalln(err)
		errPayload := &models.Default{
			Status: "error",
			Data: "Something goes wrong",
		}
		return auth.NewAuthorizeRequestDefault(500).WithPayload(errPayload)
	} 


	response := &models.AuthResponse{
		Status: "sucess",
		Data: &models.AuthResponseData{
			Result: "not allowed",
		},
	}

	if(results.Allowed()){
		response.Data.Result = "allowed"
	}

	return auth.NewAuthorizeRequestOK().WithPayload(response)

}
