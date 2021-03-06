package planet

import "github.com/bosamatheus/star-wars/entity"

type Reader interface {
	Get(id entity.ID) (*entity.Planet, error)
	Search(name string) ([]*entity.Planet, error)
	List() ([]*entity.Planet, error)
}

type Writer interface {
	Create(e *entity.Planet) (*entity.Planet, error)
	Delete(id entity.ID) error
}

type Repository interface {
	Reader
	Writer
}

type Client interface {
	Search(name string) (int, error)
}

type UseCase interface {
	GetPlanet(id entity.ID) (*entity.Planet, error)
	SearchPlanets(name string) ([]*entity.Planet, error)
	ListPlanets() ([]*entity.Planet, error)
	CreatePlanet(name, climate, terrain string) (*entity.Planet, error)
	DeletePlanet(id entity.ID) error
}
