package engine

import (
	"fmt"

	"github.com/reonardoleis/adventurer/internal/dice"
	"github.com/reonardoleis/adventurer/internal/engine/generator"
	"github.com/reonardoleis/adventurer/internal/tables"
	"github.com/reonardoleis/adventurer/internal/utils"
)

var (
	storyFlowStart = &step{
		needsInput: false,
		fn:         generateEnvironment,
	}
)

func generateEnvironment(g *Game, input string) (string, error) {
	generation, err := generator.Environment(g.World)
	if err != nil {
		return "", err
	}

	g.World.SetCurrentEnvironment(generation.ToEntity())

	g.SetNext(&step{
		needsInput: false,
		fn:         generateSituation,
	})
	return g.World.CurrentEnvironment.Present(), nil
}

func generateSituation(g *Game, input string) (string, error) {
	generation, err := generator.Situation(g.World)
	if err != nil {
		return "", err
	}

	g.World.CurrentEnvironment.SendSituationsToHistory()
	situation := generation.ToEntity()
	g.World.CurrentEnvironment.SetNewSituation(situation)

	g.SetNext(&step{
		needsInput: true,
		dynamicInputText: func(g *Game) string {
			return g.MainCharacter.ActionText()
		},
		fn: generateSituationOutcome})
	return g.World.CurrentEnvironment.CurrentSituation.PresentContent(), nil
}

func generateSituationOutcome(g *Game, input string) (string, error) {
	if g.currentSituation >= g.situationsPerEnvironment {
		g.SetNext(&step{
			needsInput: false,
			fn:         generateEnvironment,
		})

		g.currentSituation = 0
	} else {
		g.SetNext(&step{
			needsInput: false,
			fn:         generateSituation,
		})

		g.currentSituation++
	}

	g.World.CurrentEnvironment.SetCurrentSituationDecision(input)

	output := ""

	rollGeneration, err := generator.Decision(g.World)
	if err != nil {
		return "", err
	}

	difficultyLevel, ok := tables.DecisionRolls[rollGeneration.Level]
	if !ok {
		return "", fmt.Errorf("invalid decision level: %s", rollGeneration.Level)
	}

	roll, needed, success := dice.DecisionRoll(difficultyLevel)

	outcomeGeneration, err := generator.DecisionOutcome(
		g.World,
		success,
		rollGeneration.FailStartsBattle,
		rollGeneration.AlwaysStartsBattle)
	if err != nil {
		return "", err
	}

	g.World.CurrentEnvironment.SetCurrentSituationOutcome(outcomeGeneration.Outcome)

	output += utils.SprintfSeparator(
		"[Roll: %d, Needed: %d, Success: %t]\n%s",
		"-", 10,
		roll, needed, success, g.World.CurrentEnvironment.CurrentSituation.Outcome,
	)

	g.World.GenerateStoryLogFromLastN(5)

	return output, nil
}
