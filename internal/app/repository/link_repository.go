package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/paramonies/ozon-link-shortener/internal/app/model"
	"github.com/paramonies/ozon-link-shortener/internal/app/utils"
)

const (
	LINKTABLE = "links"
)

type LinkRepository struct {
	DB *sqlx.DB
}

func NewLinkRepository(db *sqlx.DB) *LinkRepository {
	return &LinkRepository{DB: db}
}

func (r *LinkRepository) GetShortLink(url string) string {
	selectQuery := fmt.Sprintf("SELECT short_id FROM %s WHERE long_url = $1", LINKTABLE)
	row := r.DB.QueryRow(selectQuery, url)
	var result string
	if err := row.Scan(&result); err != nil {
		return ""
	}
	return result
}

func (r *LinkRepository) CreateLink(url string) (model.ClientLink, error) {
	tx, err := r.DB.Begin()
	if err != nil {
		return model.ClientLink{}, err
	}

	var id int
	createQuery := fmt.Sprintf("INSERT INTO %s (long_url) VALUES ($1) RETURNING id", LINKTABLE)
	row := tx.QueryRow(createQuery, url)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return model.ClientLink{}, err
	}

	shortUrl := utils.Convert(id, url)
	updateQuery := fmt.Sprintf("UPDATE %s SET short_id = $1 WHERE id = $2", LINKTABLE)
	_, err = tx.Exec(updateQuery, shortUrl, id)
	if err != nil {
		tx.Rollback()
		return model.ClientLink{}, err
	}

	return model.ClientLink{Url: shortUrl}, tx.Commit()
}

func (r *LinkRepository) GetLongLink(id string) (string, error) {
	selectQuery := fmt.Sprintf("SELECT long_url FROM %s WHERE short_id = $1", LINKTABLE)
	row := r.DB.QueryRow(selectQuery, id)
	var result string
	if err := row.Scan(&result); err != nil {
		switch {
		case err == sql.ErrNoRows:
			return result, errors.New("long link not found")
		default:
			return result, err
		}
	}
	return result, nil
}
