package repository

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/paramonies/ozon-link-shortener/internal/app/model"
	"github.com/paramonies/ozon-link-shortener/internal/app/utils"
	"github.com/stretchr/testify/assert"
)

func TestLinkRepository_GetShortLink(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' while opening a stub database connection", err)
	}

	defer mockDB.Close()
	db := sqlx.NewDb(mockDB, "sqlmock")
	r := NewLinkRepository(db)

	tests := []struct {
		name  string
		mock  func()
		input string
		want  string
	}{
		{
			name: "OK",
			mock: func() {
				rows := sqlmock.NewRows([]string{"short_id"}).AddRow("beeLDlFcPz")
				mock.ExpectQuery("SELECT short_id FROM links").WithArgs("http://test.ru").WillReturnRows(rows)
			},
			input: "http://test.ru",
			want:  "beeLDlFcPz",
		},
		{
			name: "Link not found",
			mock: func() {
				rows := sqlmock.NewRows([]string{"short_id"})
				mock.ExpectQuery("SELECT short_id FROM links").WithArgs("http://test.ru/not/found").WillReturnRows(rows)
			},
			input: "http://test.ru/not/found",
			want:  "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mock()
			got := r.GetShortLink(test.input)
			assert.Equal(t, test.want, got)
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}

}

func TestLinkRepository_CreateLink(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' while opening a stub database connection", err)
	}

	defer mockDB.Close()
	db := sqlx.NewDb(mockDB, "sqlmock")
	r := NewLinkRepository(db)

	tests := []struct {
		name    string
		mock    func()
		input   string
		want    model.ClientLink
		wantErr bool
	}{
		{
			name: "OK",
			mock: func() {
				mock.ExpectBegin()
				rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
				mock.ExpectQuery("INSERT INTO links").
					WithArgs("http://test.ru").WillReturnRows(rows)

				shortUrl := utils.Convert(1, "http://test.ru")

				mock.ExpectExec("UPDATE links SET").WithArgs(shortUrl, 1).WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			input:   "http://test.ru",
			want:    model.ClientLink{Url: "mFU1iiuKLp"},
			wantErr: false,
		},
		{
			name: "Start transaction fail",
			mock: func() {
				mock.ExpectBegin().WillReturnError(errors.New("start transaction fail"))
			},
			wantErr: true,
		},
		{
			name: "Insert roll back",
			mock: func() {
				mock.ExpectBegin()
				mock.ExpectQuery("INSERT INTO links").
					WithArgs("http://test.ru").WillReturnError(errors.New("insert roll back"))
			},
			input:   "http://test.ru",
			wantErr: true,
		},
		{
			name: "Update roll back",
			mock: func() {
				mock.ExpectBegin()
				rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
				mock.ExpectQuery("INSERT INTO links").
					WithArgs("http://test.ru").WillReturnRows(rows)

				shortUrl := utils.Convert(1, "http://test.ru")

				mock.ExpectExec("UPDATE links SET").WithArgs(shortUrl, 1).WillReturnError(errors.New("Update roll back"))
			},
			input:   "http://test.ru",
			wantErr: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mock()

			got, err := r.CreateLink(test.input)
			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.want, got)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestLinkRepository_GetLongLink(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' while opening a stub database connection", err)
	}

	defer mockDB.Close()
	db := sqlx.NewDb(mockDB, "sqlmock")
	r := NewLinkRepository(db)

	tests := []struct {
		name    string
		mock    func()
		input   string
		want    string
		wantErr bool
	}{
		{
			name: "OK",
			mock: func() {
				rows := sqlmock.NewRows([]string{"long_url"}).AddRow("http://test.ru")
				mock.ExpectQuery("SELECT long_url FROM links").WithArgs("beeLDlFcPz").WillReturnRows(rows)
			},
			input:   "beeLDlFcPz",
			want:    "http://test.ru",
			wantErr: false,
		},
		{
			name: "Long link not found",
			mock: func() {
				rows := sqlmock.NewRows([]string{"long_url"})
				mock.ExpectQuery("SELECT long_url FROM links").WithArgs("beeLDlFcPZ").WillReturnRows(rows)
			},
			input:   "beeLDlFcPZ",
			want:    "",
			wantErr: true,
		},
		{
			name: "Internal server error",
			mock: func() {
				mock.ExpectQuery("SELECT long_url FROM links").WithArgs("beeLDlFcPZ").WillReturnError(errors.New("Internal server error"))
			},
			input:   "beeLDlFcPZ",
			want:    "",
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mock()

			got, err := r.GetLongLink(test.input)
			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.want, got)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}

}
