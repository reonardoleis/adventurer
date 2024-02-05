package engine

import (
	"fmt"
	"strconv"

	"github.com/reonardoleis/adventurer/internal/repositories"
)

var newGameFlow = []*step{
	{needsInput: true, inputText: "What is your name?", fn: setCharacterName},
	{needsInput: true, inputText: "What is your story?", fn: setCharacterStory},
	{repeat: 5, needsInput: true, dynamicInputText: getCurrentAttributeName, fn: setCharacterAttribute},
	{needsInput: true, inputText: "What is the name of your world?", fn: setWorldName},
	{needsInput: true, inputText: "What is the description of your world?", fn: setWorldDescription},
	{fn: saveNewGame},
}

func setCharacterName(g *Game, name string) (string, error) {
	g.characterCreator.SetName(name)
	return "", nil
}

func setCharacterStory(g *Game, story string) (string, error) {
	g.characterCreator.SetStory(story)
	return "", nil
}

func setCharacterAttribute(g *Game, input string) (string, error) {
	val, err := strconv.Atoi(input)
	if err != nil {
		g.Repeat()
		return "Invalid input", err
	}

	err = g.characterCreator.SetAttributeTo(val)
	if err != nil {
		g.Repeat()
		return err.Error(), err
	}

	return fmt.Sprintf("You have %d points left", g.characterCreator.GetPointsLeft()), nil
}

func setWorldName(g *Game, name string) (string, error) {
	g.World.Name = name
	return "", nil
}

func setWorldDescription(g *Game, description string) (string, error) {
	g.World.Description = description
	return "", nil
}

func saveNewGame(g *Game, input string) (string, error) {
	defer g.SetNext(storyFlowStart)
	err := repositories.SaveMainCharacter(g.MainCharacter)
	if err != nil {
		return "", err
	}

	err = repositories.SaveWorld(g.World)
	if err != nil {
		return "", err
	}

	return "Game saved!", nil
}
