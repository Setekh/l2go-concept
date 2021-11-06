package game

import "l2go-concept/internal/game/storage"

type dependencyDef struct {
	GameTime TimeController
	Storage  storage.GameStorage
}

type DependencyManager interface {
	GetTimeController() TimeController
	GetStorage() storage.GameStorage
}

func (d *dependencyDef) GetTimeController() TimeController {
	return d.GameTime
}

func (d *dependencyDef) GetStorage() storage.GameStorage {
	return d.Storage
}

func CreateDependencyManager(controller TimeController, storage storage.GameStorage) DependencyManager {
	return &dependencyDef{
		GameTime: controller,
		Storage:  storage,
	}
}
