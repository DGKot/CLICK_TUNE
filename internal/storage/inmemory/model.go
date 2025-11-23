package inmemory

import "click_tune/internal/domain"

type Item map[domain.ID]struct {
	Description string
}

type SlotsBannersItem map[domain.ID]Stats

type Stats map[domain.ID]StatsBannerGroup

type StatsBannerGroup struct {
	Shows uint
	Click uint
}

type Inmemory struct {
	Groups       Item
	Banners      Item
	Slots        Item
	SlotsBanners map[domain.ID]SlotsBannersItem
}
