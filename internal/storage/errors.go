package storage

import "errors"

var (
	ErrEmptySlot           = errors.New("slot is empty")
	ErrWrongSlot           = errors.New("wrong slot id")
	ErrWrongGroup          = errors.New("wrong group id")
	ErrWrongBanner         = errors.New("wrong banner id")
	ErrSlotCreate          = errors.New("slot is already exist")
	ErrBannerCreate        = errors.New("banner is already exist")
	ErrGroupCreate         = errors.New("group is already exist")
	ErrClicksMoreThanShows = errors.New("clicks update, clicks more than shows")
)
