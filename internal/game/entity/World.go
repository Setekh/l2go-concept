package entity

type Entity uint32

const (
	TypePlayer = iota
	TypePet
	TypeNpc
	TypeItem
)

type dummyComponent struct {
	aaa uint32
}

type world struct {
	components []interface{}
}

type World interface {
	CreateEntity(components ...interface{}) Entity
}

func createWorld() World {
	world := &world{
		components: make([]interface{}, 200),
	}

	var gg = make([]dummyComponent, 0)
	world.components = append(world.components, gg)
	world.components[0] = nil

	return world
}

func (w *world) CreateEntity(components ...interface{}) Entity {
	var choseIndex = 0

	//for i, components := range w.components {
	//
	//}

	return Entity(choseIndex)
}
