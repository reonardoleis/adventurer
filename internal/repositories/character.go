package repositories

import (
	"github.com/reonardoleis/adventurer/internal/database"
	"github.com/reonardoleis/adventurer/internal/entities"
)

func SaveMainCharacter(character *entities.Character) error {
	currentSavegame := new(Savegame)
	err := database.GetJsonDatabase().ReadScan(currentSavegame)
	if err != nil {
		return err
	}

	currentSavegame.MainCharacter = character
	return database.GetJsonDatabase().Write(currentSavegame.JSON())
}

func FindCharacter() (*entities.Character, bool, error) {
	savegame := new(Savegame)
	err := database.GetJsonDatabase().ReadScan(savegame)
	if err != nil {
		return nil, false, err
	}

	if savegame.MainCharacter == nil {
		return nil, false, nil
	}

	return savegame.MainCharacter, true, nil
}
