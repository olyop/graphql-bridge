package main

import (
	"log"

	"github.com/olyop/graphql-bridge/server/compiler"

	"github.com/olyop/graphql-bridge/server/sourcereader"
)

func main() {
	contents := sourcereader.ReadFolder("./my-schema")

	err := compiler.CompileSchema(contents)
	if err != nil {
		log.Fatal(err)
	}
}
