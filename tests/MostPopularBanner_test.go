package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMostPopularBanner(t *testing.T) {
	// t.Parallel()
	handler := SetupServer()
	AddStartItems(t, handler)
	for _, body := range bannersSlot {
		_ = doRequest(t, handler, http.MethodPost, "/banner/add", body)
	}
	t.Run("Update shows and clicks", func(t *testing.T) {
		b1Shows := 78
		b1Clicks := 73
		b2Shows := 83
		b2Clicks := 65

		body := bannerForGroup{
			BannerID: "banner_1",
			GroupID:  "group_2",
			SlotID:   "slot_2",
		}

		for idx := 0; idx < b1Shows; idx++ {
			w := doRequest(t, handler, http.MethodPost, "/update/shows", body)
			require.Equal(t, http.StatusOK, w.Code)
		}

		for idx := 0; idx < b1Clicks; idx++ {
			w := doRequest(t, handler, http.MethodPost, "/update/clicks", body)
			require.Equal(t, http.StatusOK, w.Code)
		}

		body = bannerForGroup{
			BannerID: "banner_2",
			GroupID:  "group_2",
			SlotID:   "slot_2",
		}

		for idx := 0; idx < b2Shows; idx++ {
			w := doRequest(t, handler, http.MethodPost, "/update/shows", body)
			require.Equal(t, http.StatusOK, w.Code)
		}

		for idx := 0; idx < b2Clicks; idx++ {
			w := doRequest(t, handler, http.MethodPost, "/update/clicks", body)
			require.Equal(t, http.StatusOK, w.Code)
		}
	})

	t.Run("Most popular banners", func(t *testing.T) {
		stats := make(map[string]int)
		for i := 0; i < 200; i++ {
			w := doRequest(
				t,
				handler,
				http.MethodGet,
				fmt.Sprintf("/banner?slot_id=%s&group_id=%s", "slot_2", "group_2"),
				struct{}{},
			)
			require.Equal(t, http.StatusOK, w.Code)
			answer := getBannerAnswer{}
			_ = json.NewDecoder(w.Body).Decode(&answer)
			stats[answer.Data["id"]]++
		}
		var maxCount int
		var maxID string
		for k, v := range stats {
			if v > maxCount {
				maxID = k
				maxCount = v
			}
		}
		require.Equal(t, "banner_1", maxID)
	})
}
