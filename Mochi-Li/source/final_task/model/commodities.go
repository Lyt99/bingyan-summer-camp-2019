package model

import (
	"net/url"
)

type commodity struct {
	ID            int32
	Title         string
	Desc          string
	Category      string
	Price         float64
	Picture       url.URL
	PromulgatorID string
}
