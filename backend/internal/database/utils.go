package database

import (
	"context"
	"time"

	"github.com/Cyan903/QuaverBuddy/backend/pkg/log"
	"github.com/araddon/dateparse"
)

// Scores
func GetTotalRecent(uid, mode int, hideFailed bool) (int, error) {
	var total int
	var failed = -1

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

// Stats
func GraphDateEarliest(uid, mode int) (time.Time, error) {
	var date time.Time
	var undate string

	// Fetch dates
	c, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	earliest := Conn.QueryRowContext(c, `
		SELECT Score.DateTime FROM Score join Map WHERE Score.MapMd5 = Map.Md5Checksum
			AND Score.LocalProfileID = ?
			AND Map.Mode = ?
		ORDER BY Score.Id ASC LIMIT 1;
	`, uid, mode)

	defer cancel()

	if err := earliest.Scan(&undate); err != nil {
		log.Error.Println("could not query earliest GraphDateRange", err)
		return date, err
	}

	// Parse dates
	date, err := dateparse.ParseAny(undate)

	if err != nil {
		log.Error.Println("could not parse date GraphDateRange", err)
		return date, err
	}

	return date, nil
}

func GraphCreateRange(earliest time.Time, dates map[string]int) {
	end := time.Now().AddDate(0, 1, 0)

	for d := earliest; !d.After(end); d = d.AddDate(0, 1, 0) {
		dates[d.Format("2006/01")] = 0
	}
}
