package selector

import (
	"click_tune/internal/domain"
	"click_tune/pkg/distribution"
)

func (bg *BannerGroup) NewBeta() {
	bg.Beta = *distribution.NewBeta(bg.Clicks, bg.Shows-bg.Clicks)
}

func (bg *BannerGroup) Sample() float64 {
	bg.UpdateBeta()
	sample := bg.Beta.Sample()
	return sample
}

func (bg *BannerGroup) UpdateBeta() {
	if bg.Shows != bg.Beta.Failed()+bg.Beta.Success() {
		bg.NewBeta()
	}
}

func NewBannerGroup(deps domain.Banner) *BannerGroup {
	if deps.Clicks > deps.Shows {
		deps.Clicks = deps.Shows
	}
	return &BannerGroup{
		IDBanner: deps.ID,
		Shows:    deps.Shows,
		Clicks:   deps.Clicks,
		Beta:     *distribution.NewBeta(deps.Clicks, deps.Shows-deps.Clicks),
	}
}
