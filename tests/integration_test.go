package tests

import (
	"bytes"
	"click_tune/internal/http/rest"
	"click_tune/internal/selector"
	"click_tune/internal/service"
	"click_tune/internal/storage/inmemory"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var (
	banners = []map[string]string{
		{
			"ID":          "banner_1",
			"description": "desc_1",
		},
		{
			"ID":          "banner_2",
			"description": "desc_2",
		},
		{
			"ID":          "banner_3",
			"description": "desc_3",
		},
		{
			"ID":          "banner_4",
			"description": "desc_4",
		},
		{
			"ID":          "banner_5",
			"description": "desc_5",
		},
	}
	slots = []map[string]string{
		{
			"ID":          "slot_1",
			"description": "desc_1",
		},
		{
			"ID":          "slot_2",
			"description": "desc_2",
		},
	}
	groups = []map[string]string{
		{
			"ID":          "group_1",
			"description": "desc_1",
		},
		{
			"ID":          "group_2",
			"description": "desc_2",
		},
	}
	bannersSlot = []map[string]string{
		{
			"slotId":   "slot_1",
			"bannerId": "banner_1",
		},
		{
			"slotId":   "slot_2",
			"bannerId": "banner_1",
		},
		{
			"slotId":   "slot_2",
			"bannerId": "banner_2",
		},
		{
			"slotId":   "slot_2",
			"bannerId": "banner_3",
		},
		{
			"slotId":   "slot_2",
			"bannerId": "banner_4",
		},
		{
			"slotId":   "slot_2",
			"bannerId": "banner_5",
		},
	}
)

type getBannerAnswer struct {
	Success bool              `json:"success"`
	Data    map[string]string `json:"data"`
	Error   string            `json:"error"`
}

type bannerForGroup struct {
	SlotID   string `json:"slotId"`
	GroupID  string `json:"groupId"`
	BannerID string `json:"bannerId"`
}

func AddStartItems(t *testing.T, handler http.Handler) {
	t.Helper()
	for _, body := range banners {
		_ = doRequest(t, handler, http.MethodPost, "/banner", body)
	}

	for _, body := range slots {
		_ = doRequest(t, handler, http.MethodPost, "/slot", body)
	}

	for _, body := range groups {
		_ = doRequest(t, handler, http.MethodPost, "/group", body)
	}
}

func SetupServer() http.Handler {
	selector := selector.NewSelector()
	storage := inmemory.NewStorage()

	serviceDeps := service.Deps{
		Storage:  storage,
		Selector: selector,
	}
	service := service.NewService(serviceDeps)

	server := rest.NewServer(rest.ServerDeps{
		Service:      service,
		Host:         "localhost",
		Port:         "8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  30 * time.Second,
	})

	handler := server.Handler()
	return handler
}

func doRequest(
	t *testing.T,
	handler http.Handler,
	method string,
	url string,
	body any,
) *httptest.ResponseRecorder {
	t.Helper()
	data, _ := json.Marshal(body)
	req := httptest.NewRequest(method, url, bytes.NewReader(data))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	return w
}
