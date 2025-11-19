package service

import (
	"click_tune/internal/domain"
	"errors"
)

type Deps struct {
	Storage  Storage
	Selector Selector
}

func NewService(deps Deps) *Service {
	return &Service{
		storage:  deps.Storage,
		selector: deps.Selector,
	}
}

func checkIDs(err error, ids ...domain.ID) error {
	for _, id := range ids {
		if id.IsZero() {
			return errors.Join(err, ErrEmptyID)
		}
	}
	return nil
}

func (s *Service) GetBannersForGroup(slotID domain.ID, groupID domain.ID) (domain.ID, error) {
	err := checkIDs(ErrGetBanner, slotID, groupID)
	if err != nil {
		return "", err
	}
	banners, err := s.storage.GetBanners(slotID, groupID)
	if err != nil {
		return "", errors.Join(ErrGetBanner, err)
	}
	banner, err := s.selector.Banner(banners)
	if err != nil {
		return "", errors.Join(ErrGetBanner, err)
	}
	err = s.UpdateShows(slotID, groupID, banner)
	if err != nil {
		return "", errors.Join(ErrGetBanner, err)
	}
	return banner, nil
}

func (s *Service) UpdateShows(slotID domain.ID, groupID domain.ID, bannerID domain.ID) error {
	// TODO Отправка события в kafka
	_, err := s.storage.UpdateShows(slotID, bannerID, groupID)
	return err
}

func (s *Service) UpdateClicks(slotID domain.ID, groupID domain.ID, bannerID domain.ID) error {
	// TODO Отправка события в kafka
	_, err := s.storage.UpdateClicks(slotID, bannerID, groupID)
	return err
}

func (s *Service) AddBanner(slotID domain.ID, bannerID domain.ID) error {
	err := checkIDs(ErrAddBannerSlot, slotID, bannerID)
	if err != nil {
		return err
	}
	return s.storage.AddBanner(slotID, bannerID)
}

func (s *Service) DeleteBanner(slotID domain.ID, bannerID domain.ID) error {
	err := checkIDs(ErrDeleteBannerSlot, slotID, bannerID)
	if err != nil {
		return err
	}
	return s.storage.DeleteBanner(slotID, bannerID)
}

func (s *Service) CreateSlot(slotID domain.ID, desc string) error {
	err := checkIDs(ErrCreateSlot, slotID)
	if err != nil {
		return err
	}
	return s.storage.CreateSlot(slotID, desc)
}

func (s *Service) CreateBanner(bannerID domain.ID, desc string) error {
	err := checkIDs(ErrCreateBanner, bannerID)
	if err != nil {
		return err
	}
	return s.storage.CreateBanner(bannerID, desc)
}

func (s *Service) CreateGroup(groupID domain.ID, desc string) error {
	err := checkIDs(ErrCreateGroup, groupID)
	if err != nil {
		return err
	}
	return s.storage.CreateGroup(groupID, desc)
}
