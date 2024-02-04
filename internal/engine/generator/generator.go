package generator

import (
	"encoding/json"
	"math/rand"
	"strconv"
	"strings"

	"github.com/reonardoleis/adventurer/internal/ai"
	"github.com/reonardoleis/adventurer/internal/entities"
)

func Environment(
	world *entities.World,
) (*entities.Environment, error) {
	env := new(entities.Environment)

	prompt, err := PromptGenerateEnvironment.
		Chain(Base).Chain(NumberOfCharacters).
		Fill(
			world.Name,
			world.Description,
			rand.Intn(6)+2,
		)
	if err != nil {
		return nil, err
	}

	response, err := ai.Generate(
		prompt.String(), nil, 1024, 1,
	)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(response), env)
	if err != nil {
		return nil, err
	}

	return env, nil
}

func Situation(
	mainCharacter *entities.Character,
	world *entities.World,
	environment *entities.Environment,
	startsBattle bool,
) (string, error) {
	prompt, err := PromptGenerateSituation.
		Chain(Base).
		Chain(MainCharacterInformation).
		Chain(LastSituation).
		Chain(CurrentEnvironment).
		Chain(OtherCharactersInformation).
		Fill(
			startsBattle,
			world.Name, world.Description,
			mainCharacter.Name+mainCharacter.Story,
			mainCharacter.Race+" - "+mainCharacter.Class,
			environment.LastSituation(), environment.FailedLastSituation,
			environment.Name+": "+environment.Description,
			environment.GetCharacterData(),
		)
	if err != nil {
		return "", err
	}

	situation, err := ai.Generate(
		prompt.String(), []string{
			"STORY_LOG: " + world.StoryLog,
		}, 1024, 1,
	)
	if err != nil {
		return "", err
	}

	return situation, nil
}

type DecisionResult struct {
	NeededRoll         int
	FailStartsBattle   bool
	AlwaysStartsBattle bool
}

func Decision(
	world *entities.World,
	environment *entities.Environment,
	situation,
	playerDecision string,
) (d DecisionResult, err error) {
	prompt, err := PromptGenerateDecision.
		Chain(Base).
		Chain(CurrentEnvironment).
		Chain(LastSituation).
		Fill(
			world.Name, world.Description,
			environment.Name+": "+environment.Description,
			situation, environment.FailedLastSituation,
		)
	if err != nil {
		return DecisionResult{}, err
	}

	response, err := ai.Generate(
		prompt.String(), []string{
			"STORY_LOG: " + world.StoryLog,
		}, 1024, 0.5,
	)
	if err != nil {
		return DecisionResult{}, err
	}

	parts := strings.Split(response, ";;")

	neededRoll, err := strconv.Atoi(parts[0])
	if err != nil {
		return DecisionResult{}, err
	}

	alwayStartsBattle, err := strconv.ParseBool(parts[1])
	if err != nil {
		return DecisionResult{}, err
	}

	failStartsBattle, err := strconv.ParseBool(parts[2])
	if err != nil {
		return DecisionResult{}, err
	}

	return DecisionResult{
		NeededRoll:         neededRoll,
		FailStartsBattle:   failStartsBattle,
		AlwaysStartsBattle: alwayStartsBattle,
	}, nil
}

func DecisionOutcome(
	world *entities.World,
	environment *entities.Environment,
	situation,
	playerDecision string,
	playerSucceeded bool,
	startedBattle bool,
) (string, error) {
	prompt, err := PromptGenerateDecisionOutcome.
		Chain(Base).
		Chain(LastSituation).
		Chain(LastPlayerDecision).
		Fill(
			world.Name, world.Description,
			situation, environment.FailedLastSituation,
			playerDecision,
			startedBattle, playerSucceeded,
		)
	if err != nil {
		return "", err
	}

	response, err := ai.Generate(
		prompt.String(), nil, 1024, 0.5,
	)
	if err != nil {
		return "", err
	}

	return response, nil
}
