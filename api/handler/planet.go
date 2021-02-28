package handler

import (
	"encoding/json"
	"net/http"

	"github.com/bosamatheus/star-wars/api/presenter"
	"github.com/bosamatheus/star-wars/entity"
	"github.com/bosamatheus/star-wars/usecase/planet"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

func MakePlanetHandlers(r *mux.Router, n negroni.Negroni, service planet.UseCase) {
	r.Handle("/api/v1/planets/{id}", n.With(
		negroni.Wrap(getPlanet(service)),
	)).Methods("GET", "OPTIONS").Name("getPlanet")

	r.Handle("/api/v1/planets/search/{name}", n.With(
		negroni.Wrap(searchPlanets(service)),
	)).Methods("GET", "OPTIONS").Name("searchPlanets")

	r.Handle("/api/v1/planets", n.With(
		negroni.Wrap(listPlanets(service)),
	)).Methods("GET", "OPTIONS").Name("listPlanets")

	r.Handle("/api/v1/planets", n.With(
		negroni.Wrap(createPlanet(service)),
	)).Methods("POST", "OPTIONS").Name("createPlanet")

	r.Handle("/api/v1/planets/{id}", n.With(
		negroni.Wrap(deletePlanet(service)),
	)).Methods("DELETE", "OPTIONS").Name("deletePlanet")
}

func getPlanet(service planet.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		ID := bson.ObjectIdHex(mux.Vars(r)["id"])
		data, err := service.GetPlanet(ID)
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error: " + err.Error()))
			return
		}
		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Not found"))
			return
		}

		toJ := &presenter.Planet{
			ID:      data.ID.Hex(),
			Name:    data.Name,
			Climate: data.Climate,
			Films:   data.Films,
			Terrain: data.Terrain,
		}
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error: " + err.Error()))
		}
	})
}

func searchPlanets(service planet.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var data []*entity.Planet
		var err error
		data, err = service.SearchPlanets(mux.Vars(r)["name"])
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error: " + err.Error()))
			return
		}
		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Not found"))
			return
		}

		var toJ []*presenter.Planet
		for _, d := range data {
			toJ = append(toJ, &presenter.Planet{
				ID:      d.ID.Hex(),
				Name:    d.Name,
				Climate: d.Climate,
				Films:   d.Films,
				Terrain: d.Terrain,
			})
		}
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error: " + err.Error()))
		}
	})
}

func listPlanets(service planet.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var data []*entity.Planet
		var err error
		data, err = service.ListPlanets()
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error: " + err.Error()))
			return
		}
		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Not found"))
			return
		}

		var toJ []*presenter.Planet
		for _, d := range data {
			toJ = append(toJ, &presenter.Planet{
				ID:      d.ID.Hex(),
				Name:    d.Name,
				Climate: d.Climate,
				Films:   d.Films,
				Terrain: d.Terrain,
			})
		}
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error: " + err.Error()))
		}
	})
}

func createPlanet(service planet.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var input struct {
			Name    string `json:"name"`
			Climate string `json:"climate"`
			Terrain string `json:"terrain"`
		}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error: " + err.Error()))
			return
		}

		data, err := service.CreatePlanet(input.Name, input.Climate, input.Terrain)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error: " + err.Error()))
			return
		}
		toJ := &presenter.Planet{
			ID:      data.ID.Hex(),
			Name:    data.Name,
			Climate: data.Climate,
			Films:   data.Films,
			Terrain: data.Terrain,
		}
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error: " + err.Error()))
		}
	})
}

func deletePlanet(service planet.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		ID := bson.ObjectIdHex(mux.Vars(r)["id"])
		err := service.DeletePlanet(ID)
		w.WriteHeader(http.StatusNoContent)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error: " + err.Error()))
		}
	})
}
