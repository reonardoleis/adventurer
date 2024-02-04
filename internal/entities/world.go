package entities

import (
	"fmt"

	"github.com/reonardoleis/adventurer/internal/ai"
)

type World struct {
	Name        string
	Description string

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

	for i := len(w.EnvironmentHistory) - n; i < len(w.EnvironmentHistory); i++ {
		for _, s := range w.EnvironmentHistory[i].Situations {
			w.StoryLog += fmt.Sprintf(
				"On %s (%s), a thing that happened was: %s\n",
				w.EnvironmentHistory[i].Name,
				w.EnvironmentHistory[i].Description,
				s,
			)
		}
	}

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
	return nil
}
