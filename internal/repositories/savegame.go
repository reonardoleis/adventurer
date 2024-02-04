package repositories

import (
	"encoding/json"

	"github.com/reonardoleis/adventurer/internal/entities"
)

type Savegame struct {
	MainCharacter      *entities.Character   `json:"main_character"`
	CurrentEnvironment *entities.Environment `json:"current_environment"`
}

func (s Savegame) JSON() []byte {
	j, _ := json.Marshal(s)
	return j
}
