package database

import (
	"context"

	"github.com/Cyan903/QuaverBuddy/backend/pkg/log"
)

type ScoreBoard struct {
	ScoreID string
	MapID   string

	Score struct {
		Name                  string
		DateTime              string
		TotalScore            string
		Grade                 string
		Accuracy              string
		RankedAccuracy        string
		Mods                  string
		PerformanceRating     string
		JudgementWindowPreset string
	} `json:"score"`

	Map struct {
		Directory      string
		Path           string
		BackgroundPath string
		Artist         string
		Title          string
		DifficultyName string
		Creator        string
		RankedStatus   string
	} `json:"map"`
}


// Filter by:
// Mode
// User ID

// Mods
// JudgementWindowPreset
// Before date
// Rank status
func GetScores() ([]ScoreBoard, error) {
	var scores []ScoreBoard

	c, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	row, err := Conn.QueryContext(c, `
		SELECT
			Score.Id as SID,
			Score.Name, Score.DateTime, Score.TotalScore,
			Score.Grade, Score.Accuracy, Score.RankedAccuracy,
			Score.Mods, Score.PerformanceRating, Score.JudgementWindowPreset,

			Map.Id as MID,
			Map.Directory, Map.Path,
			Map.BackgroundPath, Map.Artist, Map.Title,
			Map.DifficultyName, Map.Creator, Map.RankedStatus
		FROM Score join Map WHERE Score.MapMd5 = Map.Md5Checksum AND
			Map.Mode = 2 AND
			Score.LocalProfileID = 1 AND
			(SELECT Id from Score WHERE MapMd5 = Map.Md5Checksum ORDER BY PerformanceRating DESC LIMIT 1) == Score.Id
		ORDER BY Score.PerformanceRating DESC LIMIT 10;
	`)

	defer cancel()

	if err != nil {
		log.Error.Println("could not query scores GetScores", err)
		return scores, err
	}

	for row.Next() {
		var score ScoreBoard

		if err := row.Scan(
			&score.ScoreID,
			&score.Score.Name, &score.Score.DateTime, &score.Score.TotalScore,
			&score.Score.Grade, &score.Score.Accuracy, &score.Score.RankedAccuracy,
			&score.Score.Mods, &score.Score.PerformanceRating, &score.Score.JudgementWindowPreset,

			&score.MapID,
			&score.Map.Directory, &score.Map.Path,
			&score.Map.BackgroundPath, &score.Map.Artist, &score.Map.Title,
			&score.Map.DifficultyName, &score.Map.Creator, &score.Map.RankedStatus,
		); err != nil {
			log.Error.Println("could not scan GetScores", err)

			return scores, err
		}

		scores = append(scores, score)
	}

	return scores, nil
}
