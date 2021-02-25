package entity

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Planet struct {
	ID        bson.ObjectId `bson:"_id"`
	Name      string        `bson:"name"`
	Climate   string        `bson:"climate"`
	Terrain   string        `bson:"terrain"`
	Films     int           `bson:"films"`
	CreatedAt time.Time     `bson:"created_at"`
}

func NewPlanet(name, climate, terrain string, films int) (*Planet, error) {
	e := &Planet{
		ID:        bson.NewObjectId(),
		Name:      name,
		Climate:   climate,
		Terrain:   terrain,
		Films:     films,
		CreatedAt: time.Now(),
	}
	err := e.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}
	return e, nil
}

func (e *Planet) Validate() error {
	if e.Name == "" || e.Climate == "" || e.Terrain == "" {
		return ErrInvalidEntity
	}
	return nil
}
