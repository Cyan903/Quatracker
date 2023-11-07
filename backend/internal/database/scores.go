package database

import (
	"context"

	"github.com/Cyan903/QuaverBuddy/backend/pkg/log"
)

func GetScores() ([]string, error) {
	var scores []string

	c, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	row, err := Conn.QueryContext(c, "SELECT id FROM Score LIMIT 10")

	defer cancel()

	if err != nil {
		log.Error.Println("could not query scores GetScores", err)
		return scores, err
	}

	for row.Next() {
		var score string

		if err := row.Scan(&score); err != nil {
			log.Error.Println("could not scan GetScores", err)

			return scores, err
		}

		scores = append(scores, score)
	}

	return scores, nil
}
