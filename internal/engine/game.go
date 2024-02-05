package engine

import (
	"github.com/reonardoleis/adventurer/internal/engine/character_creation"
	"github.com/reonardoleis/adventurer/internal/entities"
	"github.com/reonardoleis/adventurer/internal/repositories"
)

type step struct {
	repeat           int
	needsInput       bool
	inputText        string
	dynamicInputText func(g *Game) string
	fn               func(g *Game, input string) (string, error)
}

type Game struct {
	MainCharacter            *entities.Character
	World                    *entities.World
	step                     *step
	situationsPerEnvironment int
	currentSituation         int
	next                     []*step
	characterCreator         *character_creation.CharacterCreator
}

func NewGame() *Game {
	character := new(entities.Character)
	character.Attributes = entities.NewAttributes()
	world := new(entities.World)
	world.MainCharacter = character
	characterCreator := character_creation.NewCharacterCreator(character)
	return &Game{
		MainCharacter:            character,
		World:                    world,
		characterCreator:         characterCreator,
		step:                     newGameFlow[0],
		next:                     newGameFlow[1:],
		situationsPerEnvironment: 3,
	}
}

func LoadGame() (*Game, bool, error) {
	character, found, err := repositories.FindCharacter()
	if err != nil {
		return nil, false, err
	}

	if !found {
		return nil, false, nil
	}

	world, found, err := repositories.FindWorld()
	if err != nil {
		return nil, false, err
	}

	if !found {
		return nil, false, nil
	}

	world.MainCharacter = character

	return &Game{
		MainCharacter:            character,
		World:                    world,
		step:                     storyFlowStart,
		situationsPerEnvironment: 15,
	}, true, nil
}

func (g *Game) NeedsInput() bool {
	return g.step.needsInput
}

func (g *Game) GetInputText() string {
	if g.step.dynamicInputText != nil {
		return g.step.dynamicInputText(g)
	}
	return g.step.inputText
}

func (g *Game) GetStepFunction() func(g *Game, input string) (string, error) {
	return g.step.fn
}

func (g *Game) NextStep() {
	if g.step.repeat > 0 {
		g.step.repeat--
		return
	}

	if len(g.next) > 0 {
		g.step = g.next[0]
		if len(g.next) > 1 {
			g.next = g.next[1:]
		} else {
			g.next = nil
		}
	}
}

func (g *Game) HasNextStep() bool {
	return g.next != nil || g.step.repeat > 0
}

func (g *Game) Repeat() {
	g.step.repeat++
}

func (g *Game) SetNext(next *step) {
	g.next = append([]*step{next}, g.next...)
}
