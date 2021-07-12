package repository

import "github.com/jmoiron/sqlx"

type LinkRepository struct {
	DB *sqlx.DB
}

func NewLinkRepository(db *sqlx.DB) *LinkRepository {
	return &LinkRepository{DB: db}
}
