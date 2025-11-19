package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestShowAllBanners(t *testing.T) {
	t.Parallel()
	handler := SetupServer()
	AddStartItems(t, handler)
	for _, body := range bannersSlot {
		_ = doRequest(t, handler, http.MethodPost, "/banner/add", body)
	}
	stats := make(map[string]int, len(banners))
	for _, banner := range banners {
		stats[banner["ID"]] = 0
	}

	for idx := 0; idx < 9; idx++ {
		w := doRequest(
			t,
			handler,
			http.MethodGet,
			fmt.Sprintf("/banner?slot_id=%s&group_id=%s", "slot_2", "group_1"),
			struct{}{},
		)
		require.Equal(t, http.StatusOK, w.Code)
		answer := getBannerAnswer{}
		_ = json.NewDecoder(w.Body).Decode(&answer)
		stats[answer.Data["id"]]++
	}
	for k, v := range stats {
		if len(k) > 0 {
			require.Greater(t, v, 0)
		}
	}
}
