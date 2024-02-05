package entities

import (
	"fmt"
	"strings"

	"github.com/reonardoleis/adventurer/internal/utils"
)

type Situation struct {
	Content  string
	Decision string
	Outcome  string
}

type Environment struct {
	Name        string `json:"name"`
	Description string `json:"description"`

	Characters []Character `json:"characters"`
	Items      []Item

	CurrentSituation    *Situation
	FailedLastSituation bool

	Situations []*Situation
}

func (e Environment) LastSituation() *Situation {
	if len(e.Situations) == 0 {
		return &Situation{}
	}

	return e.Situations[len(e.Situations)-1]
}

func (e *Environment) SetNewSituation(situation *Situation) {
	if e.CurrentSituation != nil {
		e.Situations = append(e.Situations, e.CurrentSituation)
	}
	e.CurrentSituation = situation
}

func (e *Environment) SendSituationsToHistory() {
	if e.CurrentSituation != nil {
		e.Situations = append(e.Situations, e.CurrentSituation)
	}
}

func (e *Environment) SetCurrentSituationDecision(decision string) {
	e.CurrentSituation.Decision = decision
}

func (e *Environment) SetCurrentSituationOutcome(outcome string) {
	e.CurrentSituation.Outcome = outcome
}

func (e Environment) Present() string {
	return utils.SprintfSeparator(
		"[Environment: %s]\n%s",
		"-", 10,
		e.Name, e.Description,
	)
}

func (s Situation) PresentContent() string {
	return utils.SprintfSeparator(
		"%s",
		"-", 10,
		s.Content,
	)
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

func (e Environment) GetSituationsData() string {
	data := make([]string, 0, len(e.Situations))
	for _, s := range e.Situations {
		data = append(data, fmt.Sprintf(
			"Situation:%s|Decision:%s|Outcome:%s",
			s.Content, s.Decision, s.Outcome,
		))
	}

	return strings.Join(data, ";;;")
}

func (e Environment) GetCurrentSituationData() string {
	if e.CurrentSituation == nil {
		return ""
	}

	return fmt.Sprintf(
		"Situation:%s|Decision:%s|Outcome:%s",
		e.CurrentSituation.Content, e.CurrentSituation.Decision, e.CurrentSituation.Outcome,
	)
}
