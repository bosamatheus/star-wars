package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPlanet(t *testing.T) {
	p, err := NewPlanet("Tatooine", "arid", "desert", 5)
	assert.Nil(t, err)
	assert.Equal(t, p.Name, "Tatooine")
	assert.NotNil(t, p.ID)
	assert.Equal(t, p.Climate, "arid")
	assert.Equal(t, p.Terrain, "desert")
	assert.Equal(t, p.Films, 5)
}

func TestPlanetValidate(t *testing.T) {
	type test struct {
		name    string
		climate string
		terrain string
		films   int
		want    error
	}
	tests := []test{
		{
			name:    "Tatooine",
			climate: "arid",
			terrain: "desert",
			films:   5,
			want:    nil,
		},
		{
			name:    "",
			climate: "",
			terrain: "",
			films:   0,
			want:    ErrInvalidEntity,
		},
	}

	for _, tc := range tests {
		_, err := NewPlanet(tc.name, tc.climate, tc.terrain, tc.films)
		assert.Equal(t, err, tc.want)
	}
}
