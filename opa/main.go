package opa

import (
	"context"

	"github.com/open-policy-agent/opa/rego"
)

type OPA struct {
	Module string
	Query string
}

func NewOPA() OPA {

	_module := `
	package cgapi

	default allow := false 

	allow = true {
		input.request.method == "GET"
	}
	` 

	return OPA{
		Module:	_module,
		Query: "data.cgapi.allow",
	}
}

func ( opa *OPA ) NewOpaQuery() (rego.PreparedEvalQuery, error){
	ctx := context.Background()
	query, err := rego.New(
		rego.Query(opa.Query),
		rego.Module("cgapi.rego",opa.Module),
	).PrepareForEval(ctx)
	
	return query, err
} 
