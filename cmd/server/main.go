package server

import (
	"net/http"
	"podcaster/internal/handlers"
	"podcaster/internal/repository"
	"podcaster/pkg/logger"
	postgres "podcaster/pkg/storage/postges"

	"github.com/gorilla/mux"
)

func Start() {

	log := logger.New()

	db, err := postgres.New(*log)

	if err != nil {
		log.Fatal().Err(err).Msg("Error while connecting to database")
	}

	episodeRepo := repository.NewGormEpisodeRepository(db)

	router := mux.NewRouter()
	handler := handlers.New(&episodeRepo)

	router.HandleFunc("/", handler.HandleHome).Methods("GET")

	// serve files in public
	router.PathPrefix("/public/").Handler(
		http.StripPrefix("/public/", http.FileServer(http.Dir("public"))),
	)

	log.Info().Msgf("Listening on %v\n", "localhost:8080")
	http.ListenAndServe(":8080", router)
}
