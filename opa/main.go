package opa

import (
	"context"
	"github.com/open-policy-agent/opa/rego"
	_ "embed"
)

var(
	//go:embed policies/policies.rego
	_module string
)

type OPA struct {
	Module string
	Query string
}

func NewOPA() OPA {

	return OPA{
		
		Module:	_module,
		Query: "data.policies.allow",
	}
}

func ( opa *OPA ) NewOpaQuery() (rego.PreparedEvalQuery, error){
	ctx := context.Background()
	query, err := rego.New(
		rego.Query(opa.Query),
		rego.Module("policies.rego",opa.Module),
	).PrepareForEval(ctx)
	
	return query, err
} 
