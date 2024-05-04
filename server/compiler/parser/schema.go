package parser

import (
	configurationAST "github.com/olyop/graphql-bridge/server/compiler/parser/configuration/ast"
	graphqlAST "github.com/vektah/gqlparser/ast"
)

// Schema is a struct that contains a graphql schema and a configuration language.
type Schema struct {
	GraphQL       *graphqlAST.Schema
	Configuration *configurationAST.Configuration
}
