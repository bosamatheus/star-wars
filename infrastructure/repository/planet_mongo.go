package repository

import (
	"github.com/bosamatheus/star-wars/entity"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type PlanetMongoDB struct {
	db         *mgo.Database
	collection string
}

func NewPlanetMongoDB(db *mgo.Database, collection string) *PlanetMongoDB {
	return &PlanetMongoDB{
		db:         db,
		collection: collection,
	}
}

func (r *PlanetMongoDB) Get(id string) (*entity.Planet, error) {
	var planet *entity.Planet
	err := r.db.C(r.collection).FindId(bson.ObjectIdHex(id)).One(&planet)
	return planet, err
}

func (r *PlanetMongoDB) Search(name string) ([]*entity.Planet, error) {
	query := bson.M{"name": &bson.RegEx{Pattern: name, Options: "i"}}
	var planets []*entity.Planet
	err := r.db.C(r.collection).Find(query).All(&planets)
	return planets, err
}

func (r *PlanetMongoDB) List() ([]*entity.Planet, error) {
	var planets []*entity.Planet
	err := r.db.C(r.collection).Find(bson.M{}).All(&planets)
	return planets, err
}

func (r *PlanetMongoDB) Create(e *entity.Planet) (*entity.Planet, error) {
	err := r.db.C(r.collection).Insert(&e)
	return e, err
}

func (r *PlanetMongoDB) Delete(id string) error {
	err := r.db.C(r.collection).RemoveId(bson.ObjectIdHex(id))
	return err
}
