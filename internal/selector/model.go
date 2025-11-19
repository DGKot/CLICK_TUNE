package selector

import (
	"click_tune/internal/domain"
	"click_tune/pkg/distribution"
)

type Selector struct{}

type BannerGroup struct {
	IDBanner domain.ID
	IDGroup  domain.ID
	Shows    uint
	Clicks   uint
	Beta     distribution.Beta
}
