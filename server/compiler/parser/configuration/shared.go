package configurationparser

import (
	"github.com/olyop/graphql-bridge/server/compiler/parser/configuration/ast"
)

type parseState struct {
	line            string
	actionIndex     int
	isAction        bool
	graphqlLocation string // GLOBAL | TYPE | ARGUMENT
	actions         map[int]*ast.Action
}
