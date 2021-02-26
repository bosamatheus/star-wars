package planet

import (
	"testing"
	"time"

	"github.com/bosamatheus/star-wars/entity"
	"github.com/stretchr/testify/assert"
)

func newFixturePlanet() *entity.Planet {
	return &entity.Planet{
		ID:        entity.NewID(),
		Name:      "Tatooine",
		Climate:   "arid",
		Terrain:   "desert",
		Films:     1,
		CreatedAt: time.Now(),
	}
}

func Test_Create(t *testing.T) {
	repo := newRepoInMem()
	client := newClientInMem()
	m := NewService(repo, client)
	p := newFixturePlanet()
	_, err := m.CreatePlanet(p.Name, p.Climate, p.Terrain)
	assert.Nil(t, err)
	assert.False(t, p.CreatedAt.IsZero())
}

func Test_SearchAndList(t *testing.T) {
	repo := newRepoInMem()
	client := newClientInMem()
	m := NewService(repo, client)
	p1 := newFixturePlanet()
	p2 := newFixturePlanet()
	p2.Name = "Alderaan"

	p, _ := m.CreatePlanet(p1.Name, p1.Climate, p1.Terrain)
	_, _ = m.CreatePlanet(p2.Name, p2.Climate, p2.Terrain)

	t.Run("search", func(t *testing.T) {
		c, err := m.SearchPlanets("tatooine")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(c))
		assert.Equal(t, "Tatooine", c[0].Name)

		c, err = m.SearchPlanets("foo")
		assert.Equal(t, entity.ErrNotFound, err)
		assert.Nil(t, c)
	})
	t.Run("list all", func(t *testing.T) {
		all, err := m.ListPlanets()
		assert.Nil(t, err)
		assert.Equal(t, 2, len(all))
	})
	t.Run("get", func(t *testing.T) {
		saved, err := m.GetPlanet(p.ID)
		assert.Nil(t, err)
		assert.Equal(t, p1.Name, saved.Name)
	})
}

func TestDelete(t *testing.T) {
	repo := newRepoInMem()
	client := newClientInMem()
	m := NewService(repo, client)
	p1 := newFixturePlanet()
	p2 := newFixturePlanet()
	p2, _ = m.CreatePlanet(p2.Name, p2.Climate, p2.Terrain)

	err := m.DeletePlanet(p1.ID)
	assert.Equal(t, entity.ErrNotFound, err)

	err = m.DeletePlanet(p2.ID)
	assert.Nil(t, err)
	_, err = m.GetPlanet(p2.ID)
	assert.Equal(t, entity.ErrNotFound, err)
}
