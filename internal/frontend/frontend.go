package frontend

import "github.com/reonardoleis/adventurer/internal/frontend/terminal"

type FrontendType int

const (
	Terminal FrontendType = iota
)

type Frontend interface {
	Run() error
}

func Get(frontendType FrontendType) Frontend {
	switch frontendType {
	case Terminal:
		return terminal.NewTerminal()
	default:
		return terminal.NewTerminal()
	}
}
