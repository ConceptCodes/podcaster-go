package handlers

import (
	"net/http"
	"podcaster/internal/model"
	"podcaster/internal/repository"
	"podcaster/views"
)

type Handler struct {
	episodeRepository repository.EpisodeRepository
}

func New(episodeRepository repository.EpisodeRepository) *Handler {
	return &Handler{
		episodeRepository: episodeRepository,
	}
}

func (h *Handler) HandleHome(w http.ResponseWriter, r *http.Request) {
	// TODO: add episodes and config data

	cfg := model.Config{
		Title:           "Sample Podcast",
		Host:            "Johhy Appleseed",
		ImgUrl:          "",
		Tagline:         "",
		Desc:            "",
		SpotifyUrl:      "",
		ApplePodcastUrl: "",
	}

	views.Index(cfg).Render(r.Context(), w)
}
