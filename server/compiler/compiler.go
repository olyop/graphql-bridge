package compiler

import (
	"fmt"

	"github.com/olyop/graphql-bridge/server/compiler/parser"
)

// CompileSchema parses a string and returns an error if the string is not valid.
func CompileSchema(files map[string]string) error {
	parsedSchema, err := parser.ParseSchema(files)
	if err != nil {
		return err
	}

	fmt.Println(parsedSchema)

	return nil
}
