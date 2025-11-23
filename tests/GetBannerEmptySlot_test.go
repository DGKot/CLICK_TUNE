package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetBannerFromEmptySlot(t *testing.T) {
	t.Parallel()
	handler := SetupServer()
	AddStartItems(t, handler)
	w := doRequest(
		t,
		handler,
		http.MethodGet,
		fmt.Sprintf("/banner?slot_id=%s&group_id=%s", "slot_1", "group_1"),
		struct{}{},
	)
	require.Equal(t, http.StatusNotFound, w.Code)
}
