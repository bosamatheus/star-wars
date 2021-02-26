package planet

import (
	"strings"
)

type ClientInMem struct {
	m map[string]int
}

func newClientInMem() *ClientInMem {
	var m = map[string]int{
		"Tatooine": 5,
		"Alderaan": 2,
		"Yavin IV": 1,
		"Hoth":     1,
		"Dagobah":  3,
	}
	return &ClientInMem{
		m: m,
	}
}

func (r *ClientInMem) Search(name string) (int, error) {
	var films int
	for k, v := range r.m {
		if strings.Contains(strings.ToLower(k), strings.ToLower(name)) {
			films = v
		}
	}
	return films, nil
}
