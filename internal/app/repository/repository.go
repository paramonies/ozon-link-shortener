package repository

import "github.com/paramonies/ozon-link-shortener/internal/app/model"

type Repository interface {
	GetShortLink(string) string
	CreateLink(string) (model.ClientLink, error)
	GetLongLink(string) (string, error)
}
