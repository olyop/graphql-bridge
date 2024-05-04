package parser

import (
	"errors"

	"github.com/olyop/graphql-bridge/server/compiler/parser/configuration"
	"github.com/olyop/graphql-bridge/server/compiler/parser/graphql"
)

// ParseSchema parses a string and returns an error if the string is not valid.
func ParseSchema(files map[string]string) (*Schema, error) {
	gqlSchema, err := graphqlparser.CompileGraphQL(files)
	if err != nil {
		return nil, errors.Join(errors.New("failed to parse graphql schema"), err)
	}

	configurationSchema, err := configurationparser.CompileConfiguration(files, gqlSchema)
	if err != nil {
		return nil, errors.Join(errors.New("failed to parse configuration schema"), err)
	}

	schema := &Schema{
		GraphQL:       gqlSchema,
		Configuration: configurationSchema,
	}

	return schema, nil
}
