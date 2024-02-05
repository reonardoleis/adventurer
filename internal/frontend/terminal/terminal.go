package terminal

import (
	"time"

	"github.com/reonardoleis/adventurer/internal/engine"
)

type TerminalFrontend struct {
	Game *engine.Game
}

func NewTerminal() *TerminalFrontend {
	game, saved, err := engine.LoadGame()
	if err != nil {
		panic(err)
	}

	if !saved {
		show("Welcome to Adventurer!", Green, true)
		game = engine.NewGame()
	} else {
		show("Welcome back to Adventurer!", Green, true)
	}

	return &TerminalFrontend{Game: game}
}

func (t *TerminalFrontend) Run() error {
	for {
		var input string
		if t.Game.NeedsInput() {
			input = askFor(t.Game.GetInputText(), Green)
		}

		stepFunction := t.Game.GetStepFunction()
		output, err := stepFunction(t.Game, input)
		if err != nil {
			if output != "" {
				show(output, Red, true)
			} else {
				return err
			}
		}

		if output != "" && err == nil {
			show(output, Green, true)
		}

		if !t.Game.HasNextStep() {
			time.Sleep(2 * time.Second)
			break
		}

		t.Game.NextStep()
	}

	return nil
}
