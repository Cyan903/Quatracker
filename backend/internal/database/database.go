package database

import (
	"database/sql"
	"errors"
	"path"
	"time"

	"github.com/Cyan903/QuaverBuddy/backend/pkg/config"
	"github.com/Cyan903/QuaverBuddy/backend/pkg/log"
	_ "github.com/mattn/go-sqlite3"
)

var ErrInvalidPath = errors.New("invalid game path")
var DefaultTimeout = 3 * time.Second
var Conn *sql.DB

func LoadDB() error {
	if config.Data.GamePath == "" {
		return ErrInvalidPath
	}

	db, err := sql.Open("sqlite3", path.Join(config.Data.GamePath, "quaver.db"))

	if err != nil {
		log.Error.Println("could not connect to quaver.db", err)
		return err
	}

	Conn = db
	return nil
}
