package tests

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddDelBannersToSlot(t *testing.T) {
	handler := SetupServer()
	AddStartItems(t, handler)

	t.Run("Add banner to slot", func(t *testing.T) {
		body := bannersSlot[0]
		w := doRequest(t, handler, http.MethodPost, "/banner/add", body)
		require.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("Delete banner from slot", func(t *testing.T) {
		body := bannersSlot[0]
		w := doRequest(t, handler, http.MethodDelete, "/banner/delete", body)
		require.Equal(t, http.StatusOK, w.Code)
	})
}
