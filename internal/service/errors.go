package service

import "errors"

var (
	ErrEmptyID          = errors.New("id is empty")
	ErrGetBanner        = errors.New("get banner")
	ErrAddBannerSlot    = errors.New("add banner to slot")
	ErrDeleteBannerSlot = errors.New("delete banner to slot")
	ErrCreateSlot       = errors.New("create slot")
	ErrCreateBanner     = errors.New("create banner")
	ErrCreateGroup      = errors.New("create group")
)
