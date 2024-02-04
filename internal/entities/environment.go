package entities

type Environment struct {
	Name        string
	Description string

	Characters []Character
	Items      []Item

	SubEnvironments []Environment
}
