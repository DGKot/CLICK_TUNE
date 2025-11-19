package storage

import "click_tune/internal/domain"

type Storage interface {
	GetBanners(slotID domain.ID, groupID domain.ID) ([]domain.Banner, error)
	AddBanner(slotID domain.ID, BannerID domain.ID) error
	DeleteBanner(slotID domain.ID, BannerID domain.ID) error
	UpdateShows(slotID domain.ID, BannerID domain.ID, groupID domain.ID) (shows uint, err error)
	UpdateClicks(slotID domain.ID, BannerID domain.ID, groupID domain.ID) (clicks uint, err error)

	CreateSlot(slotID domain.ID, desc string) error
	CreateBanner(BannerID domain.ID, desc string) error
	CreateGroup(groupID domain.ID, desc string) error
}
