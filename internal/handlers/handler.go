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
		Host:            "Johnny Appleseed",
		ImgUrl:          "via.placeholder.com/400",
		Tagline:         "A really witty tagline to test the UI",
		Desc:            "In this show, Eric and Wes dig deep to get to the facts with guests who have been labeled villains by a society quick to judge, without actually getting the full story. Tune in every Thursday to get to the truth with another misunderstood outcast as they share the missing context in their tragic tale.",
		SpotifyUrl:      "https://google.com",
		ApplePodcastUrl: "https://google.com",
	}

	views.Index(cfg).Render(r.Context(), w)
}
