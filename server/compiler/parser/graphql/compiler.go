package graphqlparser

import (
	"github.com/vektah/gqlparser"
	"github.com/vektah/gqlparser/ast"
)

// CompileGraphQL parses a string and returns an error if the string is not valid.
func CompileGraphQL(files map[string]string) (*ast.Schema, error) {
	sources := constructSources(files)

	schema, err := gqlparser.LoadSchema(sources...)
	if err != nil {
		return nil, err
	}

	return schema, nil
}

func constructSources(files map[string]string) []*ast.Source {
	sources := make([]*ast.Source, len(files))

	index := 0
	for path, content := range files {
		sources[index] = &ast.Source{
			Name:  path,
			Input: content,
		}

		index++
	}

	return sources
}
