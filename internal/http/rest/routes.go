package rest

import (
	"net/http"
)

func (h *Handler) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /banner", h.GetBannerForGroup)
	mux.HandleFunc("POST /update/shows", h.UpdateShows)
	mux.HandleFunc("POST /update/clicks", h.UpdateClicks)
	mux.HandleFunc("POST /banner/add", h.AddBannerToSlot)
	mux.HandleFunc("DELETE /banner/delete", h.DeleteBannerFromSlot)
	mux.HandleFunc("POST /banner", h.CreateBanner)
	mux.HandleFunc("POST /slot", h.CreateSlot)
	mux.HandleFunc("POST /group", h.CreateGroup)

	return mux
}
