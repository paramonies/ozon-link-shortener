package service

import (
	"github.com/paramonies/ozon-link-shortener/internal/app/model"
	"github.com/paramonies/ozon-link-shortener/internal/app/repository"
)

type LinkService struct {
	repo repository.Repository
}

func NewLinkService(repo repository.Repository) *LinkService {
	return &LinkService{repo: repo}
}

func (s *LinkService) CreateLink(input string) (model.ClientLink, error) {
	return s.repo.CreateLink(input)
}

func (s *LinkService) GetShortLink(url string) string {
	return s.repo.GetShortLink(url)
}
