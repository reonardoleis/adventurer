package terminal

import (
	"fmt"
	"strconv"

	"github.com/reonardoleis/adventurer/internal/dice"
	"github.com/reonardoleis/adventurer/internal/engine/character_creation"
	"github.com/reonardoleis/adventurer/internal/engine/generator"
	"github.com/reonardoleis/adventurer/internal/entities"
	"github.com/reonardoleis/adventurer/internal/repositories"
	"github.com/reonardoleis/adventurer/internal/tables"
)

type TerminalFrontend struct {
	Character        *entities.Character
	Environment      *entities.Environment
	World            *entities.World
	characterCreator *character_creation.CharacterCreator
}

func NewTerminal() *TerminalFrontend {
	character := new(entities.Character)
	environment := new(entities.Environment)
	world := new(entities.World)
	characterCreator := character_creation.NewCharacterCreator(character)
	return &TerminalFrontend{
		Character:        character,
		Environment:      environment,
		World:            world,
		characterCreator: characterCreator,
	}
}

func (t *TerminalFrontend) handlerRunner(handlers []func() (bool, error)) error {
	var ok bool
	var err error

	for i := 0; i < len(handlers); i++ {
		ok, err = handlers[i]()
		if err != nil {
			return err
		}

		if !ok {
			i--
		}
	}

	return nil
}

func (t *TerminalFrontend) Run() error {
	var err error

	characterCreation := []func() (bool, error){
		t.createCharacter,
		// t.createAttributes,
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
		err = t.handlerRunner(characterCreation)
		if err != nil {
			return err
		}
	}

	worldCreation := []func() (bool, error){
		t.createWorld,
	}

	world, found, err := repositories.FindWorld()
	if err != nil {
		return err
	}

	if found && world.IsValid() {
		t.World = world
	} else {
		err = t.handlerRunner(worldCreation)
		if err != nil {
			return err
		}
	}

	for {
		show("Generating environment...", Green, true)
		t.Environment, err = generator.Environment(
			t.World,
		)
		if err != nil {
			return err
		}

		show(t.Environment.Present(), Blue, true)

		for i := 0; i < 5; i++ {
			show("Generating situation...", Green, true)
			situation, err := generator.Situation(
				t.Character,
				t.World,
				t.Environment,
				true,
			)
			if err != nil {
				return err
			}

			show(situation, Yellow, true)

			t.Environment.AddSituation("Situation: " + situation + "\n")

			t.World.EnvironmentHistory = append(t.World.EnvironmentHistory, t.Environment)
			t.World.GenerateStoryLogFromLastN(3)

			playerDecision := askFor("What do you do?", Yellow)
			for len(playerDecision) == 0 {
				show("Invalid input", Red, true)
				playerDecision = askFor("What do you do?", Yellow)
			}

			decisionResult, err := generator.Decision(
				t.World,
				t.Environment,
				t.Environment.LastSituation(),
				playerDecision,
			)
			if err != nil {
				return err
			}

			t.Environment.Situations[len(t.Environment.Situations)-1] = "Player decision (highly important): " + playerDecision + "\n"
			roll := dice.Roll(1, 20).Value()
			success := roll.Get() >= decisionResult.NeededRoll
			startsBattle := decisionResult.AlwaysStartsBattle || (!success && decisionResult.FailStartsBattle)

			outcome, err := generator.DecisionOutcome(
				t.World,
				t.Environment,
				t.Environment.LastSituation(),
				playerDecision,
				success,
				startsBattle,
			)
			if err != nil {
				return err
			}

			t.Environment.Situations[len(t.Environment.Situations)-1] = "Player decision outcome: " + outcome + "\n"

			if startsBattle {
				if decisionResult.FailStartsBattle {
					show("Needs to roll at least: "+strconv.Itoa(decisionResult.NeededRoll), Green, true)
					roll.Present()
					t.Environment.FailedLastSituation = true
				}

				show(outcome, Blue, true)

				show("You are in a battle!", Red, true)

				t.Environment.FailedLastSituation = false
			} else {
				show("Needs to roll at least: "+strconv.Itoa(decisionResult.NeededRoll), Green, true)
				show(roll.Present(), Yellow, true)
				if success {
					show("You succeeded!", Green, true)
				} else {
					show("You failed!", Red, true)
				}

				show(outcome, Blue, true)

				t.Environment.FailedLastSituation = !success
			}
		}

		show("Moving to a new environment...", Green, true)
	}
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

		cost := tables.AttributeCost(attributeValue)

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

func (t *TerminalFrontend) createWorld() (bool, error) {
	show("Lets start creating your world!", Green, true)
	if t.World.Name == "" {
		t.World.Name = askFor("World name", Yellow)
		if t.World.Name == "" {
			show("World name is required", Red, true)
			return false, nil
		}
	}

	if t.World.Description == "" {
		t.World.Description = askFor("World description", Yellow)
		if t.World.Description == "" {
			show("World description is required", Red, true)
			return false, nil
		}
	}

	err := repositories.SaveWorld(t.World)
	if err != nil {
		return false, err
	}

	return true, nil
}
