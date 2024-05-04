package configurationparser

func combineSourceFiles(sourceFiles map[string]string) string {
	contents := ""

	for _, fileContent := range sourceFiles {
		contents += fileContent
	}

	return contents
}
