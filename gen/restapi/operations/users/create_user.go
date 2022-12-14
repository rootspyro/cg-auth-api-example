// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CreateUserHandlerFunc turns a function with the right signature into a create user handler
type CreateUserHandlerFunc func(CreateUserParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateUserHandlerFunc) Handle(params CreateUserParams) middleware.Responder {
	return fn(params)
}

// CreateUserHandler interface for that can handle valid create user params
type CreateUserHandler interface {
	Handle(CreateUserParams) middleware.Responder
}

// NewCreateUser creates a new http.Handler for the create user operation
func NewCreateUser(ctx *middleware.Context, handler CreateUserHandler) *CreateUser {
	return &CreateUser{Context: ctx, Handler: handler}
}

/*
	CreateUser swagger:route POST /users Users createUser

create a new user
*/
type CreateUser struct {
	Context *middleware.Context
	Handler CreateUserHandler
}

func (o *CreateUser) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewCreateUserParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// CreateUserBody create user body
//
// swagger:model CreateUserBody
type CreateUserBody struct {

	// email
	Email string `json:"email,omitempty"`

	// firstname
	Firstname string `json:"firstname,omitempty"`

	// lastname
	Lastname string `json:"lastname,omitempty"`

	// password hash
	PasswordHash string `json:"passwordHash,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}

// Validate validates this create user body
func (o *CreateUserBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create user body based on context it is used
func (o *CreateUserBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateUserBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateUserBody) UnmarshalBinary(b []byte) error {
	var res CreateUserBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
