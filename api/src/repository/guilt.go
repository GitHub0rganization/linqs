package repository

import (
	"github.com/GitHub0rganization/WebApp/model"
	"github.com/jmoiron/sqlx"
)

func AllGuilt(db *sqlx.DB) ([]model.Guilt, error) {
	guilts := make([]model.Guilt, 0)
	if err := db.Select(&guilts, `SELECT id, name, code, description, example FROM guilt`); err != nil {
		return nil, err
	}
	return guilts, nil
}

func GetGuilt(db *sqlx.DB, id int64) (*model.Guilt, error) {
	guilt := model.Guilt{}
	if err := db.Get(&guilt, `
SELECT id, name, code, description, example FROM guilt WHERE id = ?
`, id); err != nil {
		return nil, err
	}
	return &guilt, nil
}
