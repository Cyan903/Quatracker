package database

import (
	"context"
	"database/sql"
	"errors"
	"path"

	"github.com/Cyan903/QuaverBuddy/backend/pkg/config"
	"github.com/Cyan903/QuaverBuddy/backend/pkg/log"
)

func GetImage(id int) (string, error) {
	var dir string
	var img string

	c, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	query := Conn.QueryRowContext(c, "SELECT Directory, BackgroundPath FROM Map WHERE id = ? LIMIT 1", id)

	defer cancel()

	if err := query.Scan(&dir, &img); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", ErrBGNotFound
		}

		log.Error.Println("could not scan image GetImage", err)
		return "", err
	}

	if dir == "" || img == "" {
		log.Info.Println("map does not have image", id)
		return "", ErrBGNotFound
	}

	return path.Join(config.Data.GamePath, "Songs", dir, img), nil
}
