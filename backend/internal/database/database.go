package database

import (
	"context"
	"database/sql"
	"errors"
	"path"
	"time"

	"github.com/Cyan903/Quatracker/backend/pkg/config"
	"github.com/Cyan903/Quatracker/backend/pkg/log"
	_ "github.com/mattn/go-sqlite3"
)

const DefaultTimeout = 3 * time.Second
const DefaultPaginate = 10
const DefaultMiniPaginate = 5

var ErrInvalidPath = errors.New("invalid game path")
var ErrBGNotFound = errors.New("background file not found")
var ErrNoScores = errors.New("no scores")
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
	return TestConnection()
}

func TestConnection() error {
	var testItem int

	c, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	query := Conn.QueryRowContext(c, "SELECT 1 FROM QuaverSettings LIMIT 1")

	defer cancel()

	if err := query.Scan(&testItem); err != nil && err != sql.ErrNoRows {
		log.Error.Println("error testing connection", err)
		return err
	}

	return nil
}
