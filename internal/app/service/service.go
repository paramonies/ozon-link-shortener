package service

import "github.com/paramonies/ozon-link-shortener/internal/app/model"

type Service interface {
	GetShortLink(string) string
	CreateLink(string) (model.ClientLink, error)
}
