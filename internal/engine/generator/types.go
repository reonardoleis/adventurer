package generator

import "github.com/reonardoleis/adventurer/internal/entities"

type GeneratedEnvironment struct {
	Name        string `json:"name"`
	Description string `json:"description"`

	Characters []struct {
		Name  string `json:"name"`
		Story string `json:"story"`
		Race  string `json:"race"`
		Class string `json:"class"`
	} `json:"characters"`
}

func (g GeneratedEnvironment) ToEntity() *entities.Environment {
	environment := entities.Environment{
		Name:        g.Name,
		Description: g.Description,
		Situations:  make([]*entities.Situation, 0),
	}

	for _, character := range g.Characters {
		environment.Characters = append(environment.Characters, entities.Character{
			Name:  character.Name,
			Story: character.Story,
		})
	}

	return &environment
}

type GeneratedSituation struct {
	Situation string `json:"situation"`
}

func (g GeneratedSituation) ToEntity() *entities.Situation {
	return &entities.Situation{
		Content: g.Situation,
	}
}

type GeneratedDecisionOutcome struct {
	Outcome       string `json:"outcome"`
	BattleStarted bool   `json:"battle_started"`
}

type GeneratedDecisionRoll struct {
	Level              string `json:"level"`
	FailStartsBattle   bool   `json:"fail_starts_battle"`
	AlwaysStartsBattle bool   `json:"always_starts_battle"`
}
