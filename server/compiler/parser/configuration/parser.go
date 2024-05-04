package configurationparser

import (
	"bufio"
	"strings"

	configurationAST "github.com/olyop/graphql-bridge/server/compiler/parser/configuration/ast"
	"github.com/olyop/graphql-bridge/server/utils"
	graphqlAST "github.com/vektah/gqlparser/ast"
)

// CompileConfiguration parses a GraphQL schema
func CompileConfiguration(files map[string]string, graphQL *graphqlAST.Schema) (*configurationAST.Configuration, error) {
	source := combineSourceFiles(files)
	lineScanner := bufio.NewScanner(strings.NewReader(source))

	state := &parseState{
		line:            "",
		graphqlLocation: "",
		actions:         make(map[int]*configurationAST.Action),
		isAction:        false,
	}

	for lineScanner.Scan() {
		state.line = lineScanner.Text()

		if !isLanguageLine(state.line) {
			continue
		}

		processLine(state.line, state)
	}

	utils.DebugStruct(state.actions)

	configuration := &configurationAST.Configuration{
		Actions: state.actions,
	}

	return configuration, nil
}

func isLanguageLine(line string) bool {
	lineTrimmed := strings.TrimSpace(line)

	poundCount := 0

	for _, char := range lineTrimmed {
		if poundCount == 1 && char == ' ' {
			return false
		}

		if poundCount == 3 {
			if char == ' ' {
				return true
			}

			return false
		}

		if char == '#' {
			poundCount++
		}
	}

	return false
}
