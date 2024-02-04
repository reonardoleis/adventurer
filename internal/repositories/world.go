package repositories

import (
	"github.com/reonardoleis/adventurer/internal/database"
	"github.com/reonardoleis/adventurer/internal/entities"
)

func SaveWorld(world *entities.World) error {
	currentSavegame := new(Savegame)
	err := database.GetJsonDatabase().ReadScan(currentSavegame)
	if err != nil {
		return err
	}

	currentSavegame.World = world
	return database.GetJsonDatabase().Write(currentSavegame.JSON())
}

func FindWorld() (*entities.World, bool, error) {
	savegame := new(Savegame)
	err := database.GetJsonDatabase().ReadScan(savegame)
	if err != nil {
		return nil, false, err
	}

	if savegame.World == nil {
		return nil, false, nil
	}

	return savegame.World, true, nil
}
