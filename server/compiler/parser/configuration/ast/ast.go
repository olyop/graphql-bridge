package ast

// Configuration is a struct that contains a configuration language.
type Configuration struct {
	Actions map[int]*Action
}

// Action is a struct that contains a type.
type Action struct {
	Variable string
	Type     string
}
