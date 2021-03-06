package planet

import (
	"github.com/bosamatheus/star-wars/entity"
)

type Service struct {
	repo   Repository
	client Client
}

func NewService(r Repository, client Client) *Service {
	return &Service{
		repo:   r,
		client: client,
	}
}

func (s *Service) GetPlanet(ID entity.ID) (*entity.Planet, error) {
	return s.repo.Get(ID)
}

func (s *Service) SearchPlanets(name string) ([]*entity.Planet, error) {
	return s.repo.Search(name)
}

func (s *Service) ListPlanets() ([]*entity.Planet, error) {
	return s.repo.List()
}

func (s *Service) CreatePlanet(name, climate, terrain string) (*entity.Planet, error) {
	films, err := s.client.Search(name)
	if err != nil {
		return nil, err
	}
	e, err := entity.NewPlanet(name, climate, terrain, films)
	if err != nil {
		return nil, err
	}
	return s.repo.Create(e)
}

func (s *Service) DeletePlanet(ID entity.ID) error {
	e, err := s.GetPlanet(ID)
	if e == nil {
		return entity.ErrNotFound
	}
	if err != nil {
		return entity.ErrInvalidEntity
	}
	return s.repo.Delete(e.ID)
}
