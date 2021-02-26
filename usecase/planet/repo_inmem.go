package planet

import (
	"strings"

	"github.com/bosamatheus/star-wars/entity"
)

type RepoInMem struct {
	m map[entity.ID]*entity.Planet
}

func newRepoInMem() *RepoInMem {
	var m = map[entity.ID]*entity.Planet{}
	return &RepoInMem{
		m: m,
	}
}

func (r *RepoInMem) Get(id entity.ID) (*entity.Planet, error) {
	if r.m[id] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[id], nil
}

func (r *RepoInMem) Search(name string) ([]*entity.Planet, error) {
	var d []*entity.Planet
	for _, j := range r.m {
		if strings.Contains(strings.ToLower(j.Name), name) {
			d = append(d, j)
		}
	}
	if len(d) == 0 {
		return nil, entity.ErrNotFound
	}

	return d, nil
}

func (r *RepoInMem) List() ([]*entity.Planet, error) {
	var d []*entity.Planet
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}

func (r *RepoInMem) Create(e *entity.Planet) (*entity.Planet, error) {
	r.m[e.ID] = e
	return e, nil
}

func (r *RepoInMem) Delete(ID entity.ID) error {
	if r.m[ID] == nil {
		return entity.ErrNotFound
	}
	r.m[ID] = nil
	return nil
}
