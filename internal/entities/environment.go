package entities

import "strings"

type Environment struct {
	Name        string
	Description string

	Characters []Character
	Items      []Item

	SubEnvironments []Environment

	FailedLastSituation bool

	Situations []string
}

func (e Environment) LastSituation() string {
	if len(e.Situations) == 0 {
		return ""
	}

	return e.Situations[len(e.Situations)-1]
}

func (e *Environment) AddSituation(situation string) {
	e.Situations = append(e.Situations, situation)
}

func (e Environment) Present() string {
	return "You are in " + e.Name + ". " + e.Description
}

func (e Environment) GetCharacterData() string {
	var data []string
	for _, c := range e.Characters {
		t := c.Name + ": " + c.Story + ";"
		if c.IsHostile {
			t += "Hostile"
		} else {
			t += "Not_hostile"
		}
		data = append(data, t)
	}

	return strings.Join(data, ";;;")
}
