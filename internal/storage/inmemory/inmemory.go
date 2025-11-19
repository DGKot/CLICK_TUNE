package inmemory

import (
	"click_tune/internal/domain"
	"click_tune/internal/storage"
)

func NewStorage() *Inmemory {
	return &Inmemory{
		Slots:        make(Item),
		Groups:       make(Item),
		Banners:      make(Item),
		SlotsBanners: make(map[domain.ID]SlotsBannersItem),
	}
}

func (inm *Inmemory) GetBanners(slotID domain.ID, groupID domain.ID) ([]domain.Banner, error) {
	if !inm.CheckSlotID(slotID) {
		return nil, storage.ErrWrongSlot
	}
	if !inm.CheckGroupID(groupID) {
		return nil, storage.ErrWrongGroup
	}
	bannerGroup, ok := inm.SlotsBanners[slotID]
	if !ok {
		return nil, storage.ErrEmptySlot
	}
	banners := make([]domain.Banner, 0, len(inm.SlotsBanners[slotID]))
	for bannerID, bannerStats := range bannerGroup {
		banners = append(banners, domain.Banner{
			ID:     bannerID,
			Shows:  bannerStats[groupID].Shows,
			Clicks: bannerStats[groupID].Click,
		})
	}
	return banners, nil
}

func (inm *Inmemory) CheckSlotID(id domain.ID) bool {
	_, ok := inm.Slots[id]
	return ok
}

func (inm *Inmemory) CheckGroupID(id domain.ID) bool {
	_, ok := inm.Groups[id]
	return ok
}

func (inm *Inmemory) CheckBannerID(id domain.ID) bool {
	_, ok := inm.Banners[id]
	return ok
}

func (inm *Inmemory) AddBanner(slotID domain.ID, bannerID domain.ID) error {
	if !inm.CheckSlotID(slotID) {
		return storage.ErrWrongSlot
	}
	groups := inm.GetGroups()
	inm.SlotsBanners[slotID][bannerID] = make(Stats)
	for _, group := range groups {
		inm.SlotsBanners[slotID][bannerID][group] = StatsBannerGroup{
			Shows: 0,
			Click: 0,
		}
	}
	return nil
}

func (inm *Inmemory) GetGroups() []domain.ID {
	groups := make([]domain.ID, 0, len(inm.Groups))
	for group := range inm.Groups {
		groups = append(groups, group)
	}
	return groups
}

func (inm *Inmemory) DeleteBanner(slotID domain.ID, bannerID domain.ID) error {
	if !inm.CheckSlotID(slotID) {
		return storage.ErrWrongSlot
	}
	if !inm.CheckBannerID(bannerID) {
		return storage.ErrWrongBanner
	}
	delete(inm.SlotsBanners[slotID], bannerID)
	return nil
}

func (inm *Inmemory) UpdateShows(slotID domain.ID, bannerID domain.ID, groupID domain.ID) (shows uint, err error) {
	if !inm.CheckSlotID(slotID) {
		return 0, storage.ErrWrongSlot
	}
	if !inm.CheckBannerID(bannerID) {
		return 0, storage.ErrWrongBanner
	}
	if !inm.CheckGroupID(groupID) {
		return 0, storage.ErrWrongGroup
	}
	shows = inm.SlotsBanners[slotID][bannerID][groupID].Shows + 1
	inm.SlotsBanners[slotID][bannerID][groupID] = StatsBannerGroup{
		Shows: shows,
		Click: inm.SlotsBanners[slotID][bannerID][groupID].Click,
	}
	return shows, nil
}

func (inm *Inmemory) UpdateClicks(slotID domain.ID, bannerID domain.ID, groupID domain.ID) (clicks uint, err error) {
	if !inm.CheckSlotID(slotID) {
		return 0, storage.ErrWrongSlot
	}
	if !inm.CheckBannerID(bannerID) {
		return 0, storage.ErrWrongBanner
	}
	if !inm.CheckGroupID(groupID) {
		return 0, storage.ErrWrongGroup
	}
	clicks = inm.SlotsBanners[slotID][bannerID][groupID].Click + 1
	if clicks > inm.SlotsBanners[slotID][bannerID][groupID].Shows {
		return 0, storage.ErrClicksMoreThanShows
	}
	inm.SlotsBanners[slotID][bannerID][groupID] = StatsBannerGroup{
		Shows: inm.SlotsBanners[slotID][bannerID][groupID].Shows,
		Click: clicks,
	}
	return clicks, nil
}

func (inm *Inmemory) CreateSlot(slotID domain.ID, desc string) error {
	if inm.CheckSlotID(slotID) {
		return storage.ErrSlotCreate
	}

	inm.Slots[slotID] = struct{ Description string }{
		Description: desc,
	}
	inm.SlotsBanners[slotID] = make(SlotsBannersItem)
	return nil
}

func (inm *Inmemory) CreateBanner(bannerID domain.ID, desc string) error {
	if inm.CheckBannerID(bannerID) {
		return storage.ErrBannerCreate
	}
	inm.Banners[bannerID] = struct{ Description string }{
		Description: desc,
	}
	return nil
}

func (inm *Inmemory) CreateGroup(groupID domain.ID, desc string) error {
	if inm.CheckGroupID(groupID) {
		return storage.ErrGroupCreate
	}
	inm.Groups[groupID] = struct{ Description string }{
		Description: desc,
	}
	return nil
}
