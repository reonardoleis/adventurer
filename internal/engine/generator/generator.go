package generator

import (
	"encoding/json"
	"fmt"
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

	prompt := fmt.Sprintf(
		PromptGenerateEnvironment,
		[]interface{}{
			world.Name,
			world.Description,
			rand.Intn(6) + 2,
		}...,
	)

	response, err := ai.Generate(
		prompt, nil, 1024, 0.75,
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
	prompt := fmt.Sprintf(
		PromptGenerateSituation,
		[]interface{}{
			world.Name, world.Description,
			mainCharacter.Name + ": " + mainCharacter.Story,
			mainCharacter.Race + " - " + mainCharacter.Class,
			environment.LastSituation(), environment.FailedLastSituation,
			environment.Name + ": " + environment.Description,
			environment.GetCharacterData(),
			startsBattle,
		}...,
	)

	situation, err := ai.Generate(
		prompt, nil, 1024, 0.9,
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
	prompt := fmt.Sprintf(
		PromptGenerateDecision,
		[]interface{}{
			world.Name,
			world.Description,
			environment.Name + ": " + environment.Description,
			situation,
			playerDecision,
		}...,
	)

	response, err := ai.Generate(
		prompt, nil, 1024, 0.5,
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
	prompt := fmt.Sprintf(
		PromptGenerateDecisionOutcome,
		[]interface{}{
			world.Name, world.Description,
			situation, playerDecision,
			playerSucceeded, startedBattle,
		}...,
	)

	response, err := ai.Generate(
		prompt, nil, 1024, 0.5,
	)
	if err != nil {
		return "", err
	}

	return response, nil
}
