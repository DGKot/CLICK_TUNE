package selector

import (
	"click_tune/internal/domain"
	"testing"

	"github.com/stretchr/testify/require"
)

type storage struct{}

func (s *storage) GetBannersForGroup(groupID domain.ID) []domain.Banner {
	switch groupID {
	case "1":
		return []domain.Banner{}
	default:
		return []domain.Banner{
			{
				ID:     "0",
				Shows:  7,
				Clicks: 3,
			},
			{
				ID:     "1",
				Shows:  2,
				Clicks: 0,
			},
			{
				ID:     "2",
				Shows:  15,
				Clicks: 1,
			},
			{
				ID:     "3",
				Shows:  10,
				Clicks: 8,
			},
			{
				ID:     "4",
				Shows:  22,
				Clicks: 12,
			},
			{
				ID:     "5",
				Shows:  17,
				Clicks: 14,
			},
			{
				ID:     "6",
				Shows:  6,
				Clicks: 2,
			},
		}
	}
}

func getPopularBanner(g []domain.Banner) domain.ID {
	maxVal := 0.00
	var res domain.ID
	for _, v := range g {
		av := float64(v.Clicks) * 100 / float64(v.Shows)
		if av > maxVal {
			maxVal = av
			res = v.ID
		}
	}
	return res
}

func TestBanner(t *testing.T) {
	storage := &storage{}
	selector := NewSelector()
	t.Run("Base test", func(t *testing.T) {
		bannerID, err := selector.Banner(storage.GetBannersForGroup("2"))
		require.NoError(t, err)
		require.NotEmpty(t, bannerID)
	})
	t.Run("Most popular", func(t *testing.T) {
		mp := make(map[string]int)
		for idx := 0; idx < 150; idx++ {
			bannerID, err := selector.Banner(storage.GetBannersForGroup("2"))
			require.NoError(t, err)
			mp[string(bannerID)]++
		}
		require.NotEmpty(t, mp)
		popularBanner := getPopularBanner(storage.GetBannersForGroup("2"))
		m := 0
		var popularResult string
		for k, v := range mp {
			if v > m {
				m = v
				popularResult = k
			}
		}
		require.Equal(t, string(popularBanner), popularResult)
	})
	t.Run("Empty list", func(t *testing.T) {
		bannerID, err := selector.Banner(storage.GetBannersForGroup("1"))
		require.EqualError(t, err, ErrEmptyBannersForGroup.Error())
		require.True(t, bannerID.IsZero())
	})
	// TODO Тесты на вывод при изменямых данных
}
