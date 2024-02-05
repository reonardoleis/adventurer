package generator

import (
	"encoding/json"

	"github.com/reonardoleis/adventurer/internal/ai"
	"github.com/reonardoleis/adventurer/internal/entities"
)

func Environment(
	world *entities.World,
) (GeneratedEnvironment, error) {
	prompt, err :=
		Base.
			Fill(
				world.Name,
				world.Description,
				world.StoryLog,
				world.MainCharacter.Name,
				world.MainCharacter.Story,
			).
			Chain(CreateEnvironment).
			Fill(
				world.EnvironmentHistorySummary(),
			).
			Chain(Language("Portuguese")).
			Finish()
	if err != nil {
		return GeneratedEnvironment{}, err
	}

	generatedEnvironment := new(GeneratedEnvironment)
	resp, err := ai.Generate(
		prompt,
		nil,
		1024,
		0.75,
	)
	if err != nil {
		return GeneratedEnvironment{}, err
	}

	err = json.Unmarshal([]byte(resp), generatedEnvironment)
	if err != nil {
		return GeneratedEnvironment{}, err
	}

	return *generatedEnvironment, nil
}

func Situation(
	world *entities.World,
) (GeneratedSituation, error) {
	prompt, err :=
		Base.
			Fill(
				world.Name,
				world.Description,
				world.StoryLog,
				world.MainCharacter.Name,
				world.MainCharacter.Story,
			).
			Chain(CreateSituation).
			Fill(
				world.CurrentEnvironment.GetCurrentSituationData(),
				world.CurrentEnvironment.Description,
			).
			Chain(Language("Portuguese")).
			Finish()
	if err != nil {
		return GeneratedSituation{}, err
	}

	generatedSituation := new(GeneratedSituation)
	resp, err := ai.Generate(
		prompt,
		nil,
		1024,
		0.75,
	)
	if err != nil {
		return GeneratedSituation{}, err
	}

	err = json.Unmarshal([]byte(resp), generatedSituation)
	if err != nil {
		return GeneratedSituation{}, err
	}

	return *generatedSituation, nil
}

func Decision(
	world *entities.World,
) (d GeneratedDecisionRoll, err error) {
	prompt, err :=
		Base.
			Fill(
				world.CurrentEnvironment.Name,
				world.CurrentEnvironment.Description,
				world.StoryLog,
				world.MainCharacter.Name,
				world.MainCharacter.Story,
			).
			Chain(CreateRollForDecision).
			Fill(
				world.CurrentEnvironment.CurrentSituation.Content,
				world.CurrentEnvironment.CurrentSituation.Decision,
			).
			Finish()
	if err != nil {
		return GeneratedDecisionRoll{}, err
	}

	resp, err := ai.Generate(
		prompt,
		nil,
		1024,
		0.75,
	)

	if err != nil {
		return GeneratedDecisionRoll{}, err
	}

	err = json.Unmarshal([]byte(resp), &d)
	if err != nil {
		return GeneratedDecisionRoll{}, err
	}

	return d, nil
}

func DecisionOutcome(
	world *entities.World,
	succeeded bool,
	failStartsBattle bool,
	alwaysStartsBattle bool,
) (GeneratedDecisionOutcome, error) {
	prompt, err :=
		Base.
			Fill(
				world.CurrentEnvironment.Name,
				world.CurrentEnvironment.Description,
				world.StoryLog,
				world.MainCharacter.Name,
				world.MainCharacter.Story,
			).
			Chain(CreateDecisionOutcome).
			Fill(
				world.CurrentEnvironment.CurrentSituation.Content,
				world.CurrentEnvironment.CurrentSituation.Decision,
				succeeded,
				failStartsBattle,
				alwaysStartsBattle,
			).
			Chain(Language("Portuguese")).
			Finish()
	if err != nil {
		return GeneratedDecisionOutcome{}, err
	}

	generatedDecisionOutcome := new(GeneratedDecisionOutcome)
	resp, err := ai.Generate(
		prompt,
		nil,
		1024,
		0.75,
	)
	if err != nil {
		return GeneratedDecisionOutcome{}, err
	}

	err = json.Unmarshal([]byte(resp), generatedDecisionOutcome)
	if err != nil {
		return GeneratedDecisionOutcome{}, err
	}

	return *generatedDecisionOutcome, nil
}
