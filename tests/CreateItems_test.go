package tests

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateItems(t *testing.T) {
	t.Parallel()
	handler := SetupServer()
	for _, body := range banners {
		w := doRequest(t, handler, http.MethodPost, "/banner", body)
		require.Equal(t, http.StatusCreated, w.Code)
	}

	for _, body := range slots {
		w := doRequest(t, handler, http.MethodPost, "/slot", body)
		require.Equal(t, http.StatusCreated, w.Code)
	}

	for _, body := range groups {
		w := doRequest(t, handler, http.MethodPost, "/group", body)
		require.Equal(t, http.StatusCreated, w.Code)
	}
}
