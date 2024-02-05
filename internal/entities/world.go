package entities

import (
	"fmt"

	"github.com/reonardoleis/adventurer/internal/ai"
)

type World struct {
	Name               string
	MainCharacter      *Character
	Description        string
	CurrentEnvironment *Environment
	EnvironmentHistory []*Environment
	StoryLog           string
}

func (w *World) JSON() []byte {
	return []byte(`{"name":"` + w.Name + `","description":"` + w.Description + `"}`)
}

func (w World) IsValid() bool {
	return w.Name != "" && w.Description != ""
}

func (w *World) GenerateStoryLogFromLastN(n int) error {
	if n > len(w.EnvironmentHistory) {
		n = len(w.EnvironmentHistory)
	}

	changed := false

	for i := len(w.EnvironmentHistory) - n; i < len(w.EnvironmentHistory); i++ {
		if w.EnvironmentHistory[i] == nil {
			continue // check why this is happening
		}

		changed = true
		w.StoryLog += fmt.Sprintf(
			"On %s these things happened:",
			w.EnvironmentHistory[i].Name,
		)
		for _, s := range w.EnvironmentHistory[i].Situations {
			w.StoryLog += fmt.Sprintf(
				"\n- Situation(%s),Decision(%s),Outcome(%s)",
				s.Content,
				s.Decision,
				s.Outcome,
			)
		}
	}

	if changed {
		summarized, err := ai.Generate(
			"Summarize this into a small story that will be used to maintain context:\n\n"+w.StoryLog,
			nil,
			1024,
			0.75,
		)
		if err != nil {
			return err
		}

		w.StoryLog = summarized
	}

	return nil
}

func (w World) EnvironmentHistorySummary() string {
	var count int
	var summary string
	for _, e := range w.EnvironmentHistory {
		if count >= 3 {
			break
		}

		summary += e.Name + ": " + e.Description + "\n"
		count++
	}
	return summary
}

func (w *World) SetCurrentEnvironment(environment *Environment) {
	w.CurrentEnvironment = environment
	w.EnvironmentHistory = append(w.EnvironmentHistory, environment)
}
