package database

import (
	"context"

	"github.com/Cyan903/QuaverBuddy/backend/pkg/log"
)

func GetTotalRecent(
	uid, mode, page int,

	// Optional
	hideFailed bool,
) (int, error) {
	var total int

	failed := -1

	if hideFailed {
		failed = 5
	}

	// Query
	c, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	query := Conn.QueryRowContext(c, `
		SELECT
			COUNT(Score.Id) as Total
		FROM Score join Map WHERE Score.MapMd5 = Map.Md5Checksum
			AND Score.LocalProfileID = ?
			AND Map.Mode = ?
			AND Score.Grade != ?
	`, uid, mode, failed)

	defer cancel()

	if err := query.Scan(&total); err != nil {
		log.Error.Println("could not query scores GetTotalRecent", err)
		return total, err
	}

	return total, nil
}
