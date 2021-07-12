package service

import "github.com/paramonies/ozon-link-shortener/internal/app/repository"

type LinkService struct {
	repo repository.Repository
}

func NewLinkService(repo repository.Repository) *LinkService {
	return &LinkService{repo: repo}
}
