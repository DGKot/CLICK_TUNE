package inmemory

import (
	"click_tune/internal/domain"
	"click_tune/internal/storage"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func initTestInmemory() *Inmemory {
	inm := NewStorage()
	inm.CreateSlot("1-slot", "Desc")
	inm.CreateSlot("2-slot", "Desc")

	inm.CreateGroup("child", "child group")
	inm.CreateGroup("man", "man's group")
	inm.CreateGroup("woman", "woman's group")
	inm.CreateBanner("1-banner", "Best banner")
	inm.CreateBanner("2-banner", "Best banner")
	inm.CreateBanner("3-banner", "Best banner")

	inm.AddBanner("1-slot", "2-banner")
	inm.AddBanner("2-slot", "1-banner")
	inm.AddBanner("2-slot", "3-banner")
	return inm
}

func sortBanner(banners []domain.Banner) []domain.Banner {
	sort.Slice(banners, func(i, j int) bool {
		return banners[i].ID < banners[j].ID
	})
	return banners
}

func TestInmemory(t *testing.T) {
	t.Run("Base", func(t *testing.T) {
		inm := initTestInmemory()
		excepted := []domain.Banner{
			{
				ID:     "1-banner",
				Shows:  0,
				Clicks: 0,
			},
			{
				ID:     "3-banner",
				Shows:  0,
				Clicks: 0,
			},
		}
		banners, err := inm.GetBanners("2-slot", "man")
		require.NoError(t, err)
		banners = sortBanner(banners)
		require.Equal(t, excepted, banners)
	})

	t.Run("Delete banner", func(t *testing.T) {
		inm := initTestInmemory()
		err := inm.DeleteBanner("2-slot", "1-banner")
		require.NoError(t, err)

		excepted := []domain.Banner{
			{
				ID:     "3-banner",
				Shows:  0,
				Clicks: 0,
			},
		}
		banners, err := inm.GetBanners("2-slot", "man")
		require.NoError(t, err)
		require.Equal(t, excepted, banners)
	})

	t.Run("Update shows", func(t *testing.T) {
		inm := initTestInmemory()
		shows, err := inm.UpdateShows("2-slot", "1-banner", "man")
		require.NoError(t, err)
		require.Equal(t, uint(1), shows)
		inm.UpdateShows("2-slot", "1-banner", "man")

		excepted := []domain.Banner{
			{
				ID:     "1-banner",
				Shows:  2,
				Clicks: 0,
			},
			{
				ID:     "3-banner",
				Shows:  0,
				Clicks: 0,
			},
		}

		banners, err := inm.GetBanners("2-slot", "man")
		require.NoError(t, err)
		banners = sortBanner(banners)
		require.Equal(t, excepted, banners)
	})

	t.Run("Update clicks", func(t *testing.T) {
		inm := initTestInmemory()
		inm.UpdateShows("2-slot", "1-banner", "man")
		inm.UpdateShows("2-slot", "1-banner", "man")
		clicks, err := inm.UpdateClicks("2-slot", "1-banner", "man")
		require.NoError(t, err)
		require.Equal(t, uint(1), clicks)

		expected := []domain.Banner{
			{
				ID:     "1-banner",
				Shows:  2,
				Clicks: 1,
			},
			{
				ID:     "3-banner",
				Shows:  0,
				Clicks: 0,
			},
		}

		banners, err := inm.GetBanners("2-slot", "man")
		require.NoError(t, err)
		banners = sortBanner(banners)
		require.Equal(t, expected, banners)
	})

	t.Run("Update clicks (more than shows)", func(t *testing.T) {
		inm := initTestInmemory()
		inm.UpdateShows("2-slot", "1-banner", "man")
		inm.UpdateClicks("2-slot", "1-banner", "man")
		_, err := inm.UpdateClicks("2-slot", "1-banner", "man")
		require.EqualError(t, err, storage.ErrClicksMoreThanShows.Error())
	})

	t.Run("Get groups", func(t *testing.T) {
		inm := initTestInmemory()
		groups := inm.GetGroups()
		ids := make([]string, len(groups))
		for i, id := range groups {
			ids[i] = string(id)
		}
		sort.Strings(ids)

		expected := []string{
			"child", "man", "woman",
		}
		require.Equal(t, expected, ids)
	})
}
