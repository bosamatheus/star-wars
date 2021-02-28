package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/bosamatheus/star-wars/api/handler"
	"github.com/bosamatheus/star-wars/api/middleware"
	"github.com/bosamatheus/star-wars/config"
	"github.com/bosamatheus/star-wars/infrastructure/repository"
	"github.com/bosamatheus/star-wars/pkg/auth"
	"github.com/bosamatheus/star-wars/usecase/planet"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

func main() {
	var err error
	// DB
	session, err := mgo.Dial(config.MONGODB_HOST)
	if err != nil {
		log.Fatal(err)
	}
	dbMongo := session.DB(config.MONGODB_STAR_WARS)

	// Repositories
	planetRepo := repository.NewPlanetMongoDB(dbMongo, config.MONGODB_COLLECTION)
	StarWarsClient := repository.NewStarWarsClient(config.SWAPI_BASE_URL)

	// Services
	authService, err := auth.NewAuthService(config.API_SECRET)
	planetService := planet.NewService(planetRepo, StarWarsClient)

	// Handlers
	middlewareAuth := negroni.New(
		negroni.HandlerFunc(middleware.Cors),
		negroni.HandlerFunc(middleware.Auth(authService)),
		negroni.NewLogger(),
	)
	middleware := negroni.New(
		negroni.HandlerFunc(middleware.Cors),
		negroni.NewLogger(),
	)

	// Routers
	r := mux.NewRouter()
	handler.MakeAuthHandlers(r, *middleware, authService)
	handler.MakePlanetHandlers(r, *middlewareAuth, planetService)
	http.Handle("/", r)
	r.HandleFunc("/api/v1/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)
	srv := &http.Server{
		Addr:           ":" + strconv.Itoa(config.API_PORT),
		ErrorLog:       logger,
		Handler:        context.ClearHandler(http.DefaultServeMux),
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
