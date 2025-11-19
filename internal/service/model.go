package service

import (
	"click_tune/internal/domain"
)

type Storage interface {
	GetBanners(slotID domain.ID, groupID domain.ID) ([]domain.Banner, error)
	AddBanner(slotID domain.ID, bannerID domain.ID) error
	DeleteBanner(slotID domain.ID, bannerID domain.ID) error
	UpdateShows(slotID domain.ID, bannerID domain.ID, groupID domain.ID) (shows uint, err error)
	UpdateClicks(slotID domain.ID, bannerID domain.ID, groupID domain.ID) (clicks uint, err error)
	CreateSlot(slotID domain.ID, desc string) error
	CreateBanner(bannerID domain.ID, desc string) error
	CreateGroup(groupID domain.ID, desc string) error
}

type Selector interface {
	Banner(banners []domain.Banner) (domain.ID, error)
}

type Service struct {
	storage  Storage
	selector Selector
}
