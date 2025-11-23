package selector

import "click_tune/internal/domain"

func NewSelector() *Selector {
	return &Selector{}
}

func (s *Selector) Banner(banners []domain.Banner) (domain.ID, error) {
	if len(banners) == 0 {
		return "", ErrEmptyBannersForGroup
	}
	var sample float64
	var bannerID domain.ID
	for _, banner := range banners {
		bGroup := NewBannerGroup(banner)
		bSample := bGroup.Sample()
		if bSample > sample {
			sample = bSample
			bannerID = banner.ID
		}
	}
	return bannerID, nil
}
