package rest

import (
	"click_tune/internal/domain"
	"click_tune/internal/service"
	"encoding/json"
	"net/http"
)

type Handler struct {
	service *service.Service
}

type answerID struct {
	ID domain.ID `json:"id"`
}

type bannerForGroup struct {
	SlotID   domain.ID `json:"slotId"`
	GroupID  domain.ID `json:"groupId"`
	BannerID domain.ID `json:"bannerId"`
}

type bannerForSlot struct {
	SlotID   domain.ID `json:"slotId"`
	BannerID domain.ID `json:"bannerId"`
}

type item struct {
	ID          domain.ID `json:"id"`
	Description string    `json:"description"`
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) GetBannerForGroup(w http.ResponseWriter, r *http.Request) {
	slotID := domain.ID(r.URL.Query().Get("slot_id"))
	groupID := domain.ID(r.URL.Query().Get("group_id"))

	bannerID, err := h.service.GetBannersForGroup(slotID, groupID)
	if err != nil {
		writeError(w, http.StatusNotFound, err.Error())
		return
	}
	writeSuccess(w, http.StatusOK, answerID{ID: bannerID})
}

func (h *Handler) UpdateShows(w http.ResponseWriter, r *http.Request) {
	banner := &bannerForGroup{}
	err := json.NewDecoder(r.Body).Decode(&banner)
	defer r.Body.Close()
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	err = h.service.UpdateShows(banner.SlotID, banner.GroupID, banner.BannerID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeSuccess(w, http.StatusOK, nil)
}

func (h *Handler) UpdateClicks(w http.ResponseWriter, r *http.Request) {
	banner := &bannerForGroup{}
	err := json.NewDecoder(r.Body).Decode(&banner)
	defer r.Body.Close()
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	err = h.service.UpdateClicks(banner.SlotID, banner.GroupID, banner.BannerID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeSuccess(w, http.StatusOK, nil)
}

func (h *Handler) AddBannerToSlot(w http.ResponseWriter, r *http.Request) {
	banner := &bannerForSlot{}
	err := json.NewDecoder(r.Body).Decode(&banner)
	defer r.Body.Close()
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	err = h.service.AddBanner(banner.SlotID, banner.BannerID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeSuccess(w, http.StatusOK, nil)
}

func (h *Handler) DeleteBannerFromSlot(w http.ResponseWriter, r *http.Request) {
	banner := &bannerForSlot{}
	err := json.NewDecoder(r.Body).Decode(&banner)
	defer r.Body.Close()
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	err = h.service.DeleteBanner(banner.SlotID, banner.BannerID)
	if err != nil {
		writeError(w, http.StatusNotFound, err.Error())
		return
	}
	writeSuccess(w, http.StatusOK, nil)
}

func (h *Handler) CreateSlot(w http.ResponseWriter, r *http.Request) {
	item := &item{}
	err := json.NewDecoder(r.Body).Decode(&item)
	defer r.Body.Close()
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	err = h.service.CreateSlot(item.ID, item.Description)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeSuccess(w, http.StatusCreated, nil)
}

func (h *Handler) CreateBanner(w http.ResponseWriter, r *http.Request) {
	item := &item{}
	err := json.NewDecoder(r.Body).Decode(&item)
	defer r.Body.Close()
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	err = h.service.CreateBanner(item.ID, item.Description)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeSuccess(w, http.StatusCreated, nil)
}

func (h *Handler) CreateGroup(w http.ResponseWriter, r *http.Request) {
	item := &item{}
	err := json.NewDecoder(r.Body).Decode(&item)
	defer r.Body.Close()
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	err = h.service.CreateGroup(item.ID, item.Description)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeSuccess(w, http.StatusCreated, nil)
}
