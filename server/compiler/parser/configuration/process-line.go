package configurationparser

import "strings"

func processLine(sourceLine string, state *parseState) {
	line := sourceLine[strings.Index(sourceLine, "### ")+4:]

	if !state.isAction {
		state.isAction = true

		// define action
		if line[0] == '@' {
			for i := 0; i < len(line); i++ {
				if line[i] == ' ' {
					break
				}

				state.actions[state.actionIndex].Type += string(line[i])
			}
		} else {
			for i := 0; i < len(line); i++ {
				if line[i] == ' ' {
					break
				}

				state.actions[state.actionIndex].Variable += string(line[i])
			}
		}

		return
	}

	if strings.TrimRight(line, " ") == "}" {
		state.isAction = false

		// end action

		return
	}

	// process JSON
}
