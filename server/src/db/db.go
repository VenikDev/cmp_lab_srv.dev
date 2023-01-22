package db

import (
	"errors"
	"github.com/go-pg/pg/v10"
)

var (
	Database *pg.DB
)

func ConnectToDB(opt *pg.Options) error {
	Database = pg.Connect(opt)
	if Database == nil {
		return errors.New("no connection to DB")
	}

	return nil
}

// GetCites Получить все города
func GetCites(db *pg.DB) ([]string, error) {
	if db == nil {
		return nil, errors.New("DB is nil")
	}
	var cites []string
	_, err := db.Query(&cites, `SELECT name FROM city`)
	if err != nil {
		return nil, err
	}
	return cites, nil
}
