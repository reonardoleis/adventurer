package terminal

import (
	"fmt"
	"strconv"

	"github.com/reonardoleis/adventurer/internal/engine/character_creation"
	"github.com/reonardoleis/adventurer/internal/engine/generator"
	"github.com/reonardoleis/adventurer/internal/entities"
	"github.com/reonardoleis/adventurer/internal/repositories"
)

type TerminalFrontend struct {
	Character        *entities.Character
	Environment      *entities.Environment
	Generator        *generator.Generator
	characterCreator *character_creation.CharacterCreator
}

func NewTerminal() *TerminalFrontend {
	character := new(entities.Character)
	environment := new(entities.Environment)
	generator := new(generator.Generator)
	characterCreator := character_creation.NewCharacterCreator(character)
	return &TerminalFrontend{
		Character:        character,
		Environment:      environment,
		Generator:        generator,
		characterCreator: characterCreator,
	}
}

func (t *TerminalFrontend) Run() error {
	var ok bool
	var err error

	characterCreation := []func() (bool, error){
		t.createCharacter,
		t.createAttributes,
		t.buyAttributes,
	}

	character, found, err := repositories.FindCharacter()
	if err != nil {
		return err
	}

	if found && character.CreationFinished {
		t.Character = character
		t.characterCreator.Character = t.Character
	} else {
		for i := 0; i < len(characterCreation); i++ {
			ok, err = characterCreation[i]()
			if err != nil {
				return err
			}

			if !ok {
				i--
			}
		}
	}

	return nil
}

func (t *TerminalFrontend) createCharacter() (bool, error) {
	show("Lets start creating your character!", Green, true)
	name := askFor("Name", Yellow)
	story := askFor("Story", Yellow)
	race := askFor("Race", Yellow)
	class := askFor("Class", Yellow)

	t.characterCreator.CreateCharacter(name, story, race, class)

	err := repositories.SaveMainCharacter(t.Character)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (t *TerminalFrontend) createAttributes() (bool, error) {
	show("Rolling attributes...", Green, true)
	t.characterCreator.RollAttributes()

	show("Your attribute roll is:", Green, false)
	show(t.characterCreator.AttributeRoll, Yellow, true)

	rollCounter := t.characterCreator.AttributeRoll

	attributeNames := t.Character.Attributes.GetNames()
	attributeMap := make(map[string]int)
	for i := 0; i < len(attributeNames); i++ {
		attributeName := attributeNames[i]
		show(fmt.Sprintf("You have %d points left to distribute", rollCounter), Green, true)
		res := askFor(attributeName, Yellow)

		attributeValue, err := strconv.Atoi(res)
		if err != nil {
			show("Invalid input", Red, true)
			i--
			continue
		}

		if attributeValue > rollCounter {
			show("You don't have enough points", Red, true)
			i--
			continue
		}

		rollCounter -= attributeValue
		attributeMap[attributeName] = attributeValue
	}

	ok := t.characterCreator.SetAttributes(
		attributeMap["Strength"],
		attributeMap["Dexterity"],
		attributeMap["Constitution"],
		attributeMap["Intelligence"],
		attributeMap["Wisdom"],
		attributeMap["Charisma"],
	)

	if !ok {
		show("Invalid attribute distribution", Red, true)
		return false, nil
	}

	err := repositories.SaveMainCharacter(t.Character)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (t *TerminalFrontend) buyAttributes() (bool, error) {
	show("Buying attributes...", Green, true)
	show("You have 27 points to distribute", Green, true)

	buyPoints := 27
	attributeNames := t.Character.Attributes.GetNames()

	attributeMap := make(map[string]int)
	for i := 0; i < len(attributeNames); i++ {
		show(fmt.Sprintf("You have %d buy points left to distribute", buyPoints), Green, true)
		res := askFor("Set "+attributeNames[i]+" to", Yellow)

		attributeValue, err := strconv.Atoi(res)
		if err != nil {
			show("Invalid input", Red, true)
			i--
			continue
		}

		cost := attributeValue - 8

		if cost > buyPoints {
			show("You don't have enough points", Red, true)
			i--
			continue
		}

		buyPoints -= cost
		attributeMap[attributeNames[i]] = attributeValue
	}

	ok := t.characterCreator.BuyAttributes(
		attributeMap["Strength"],
		attributeMap["Dexterity"],
		attributeMap["Constitution"],
		attributeMap["Intelligence"],
		attributeMap["Wisdom"],
		attributeMap["Charisma"],
	)

	if !ok {
		show("Invalid attribute distribution", Red, true)
		return false, nil
	}

	err := repositories.SaveMainCharacter(t.Character)
	if err != nil {
		return false, err
	}

	return true, nil
}
