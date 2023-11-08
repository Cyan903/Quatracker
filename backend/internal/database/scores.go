package database

import (
	"context"
	"fmt"

	"github.com/Cyan903/QuaverBuddy/backend/pkg/api"
	"github.com/Cyan903/QuaverBuddy/backend/pkg/log"
)

type ScoreBoard struct {
	ScoreID string
	MapID   string

	Score struct {
		Name                  string
		DateTime              string
		TotalScore            int
		Grade                 string
		Accuracy              float64
		RankedAccuracy        float64
		Mods                  string
		PerformanceRating     float64
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
		LNPercent      float64
	} `json:"map"`
}

// TODO: Optional filter by:
//   - Mods
//   - Before date
func GetBestScores(
	uid, mode, page int,

	// Optional
	judgementWindowPreset string,
	RankedStatus string,
	LNPercent float64,

) ([]ScoreBoard, error) {
	var scores []ScoreBoard

	// Configure optional
	optionalJudge := "AND Score.Id != ?"
	optionalRanked := "AND Score.Id != ?"
	optionalLNPercent := "AND Score.Id != ?"

	if judgementWindowPreset != "" {
		optionalJudge = "AND Score.JudgementWindowPreset = ?"
	}

	if RankedStatus != "" {
		optionalRanked = "AND Map.RankedStatus = ?"
	}

	if LNPercent > -1.0 {
		optionalLNPercent = "AND LN >= ?"
	}

	query := fmt.Sprintf(`
		SELECT
			Score.Id as SID,
			Score.Name, Score.DateTime, Score.TotalScore,
			Score.Grade, Score.Accuracy, Score.RankedAccuracy,
			Score.Mods, Score.PerformanceRating, Score.JudgementWindowPreset,

			Map.Id as MID,
			Map.Directory, Map.Path,
			Map.BackgroundPath, Map.Artist, Map.Title,
			Map.DifficultyName, Map.Creator, Map.RankedStatus,
			((LongNoteCount*1.0) / (RegularNoteCount+LongNoteCount)*1.0) as LN
		FROM Score join Map WHERE Score.MapMd5 = Map.Md5Checksum AND
			Map.Mode = ? AND
			Score.LocalProfileID = ? AND
			(
				SELECT Id from Score WHERE MapMd5 = Map.Md5Checksum
					%[1]s %[2]s
				ORDER BY PerformanceRating DESC LIMIT 1
			) == Score.Id

			AND Score.Grade != 5
			%[1]s %[2]s %[3]s
		ORDER BY Score.PerformanceRating DESC LIMIT ?, 10;
	`, optionalJudge, optionalRanked, optionalLNPercent)

	// Query
	c, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	row, err := Conn.QueryContext(
		c, query,
		mode, uid,

		// Optional
		judgementWindowPreset, api.RevertRankedStatus(RankedStatus),
		judgementWindowPreset, api.RevertRankedStatus(RankedStatus), LNPercent,

		page,
	)

	defer cancel()

	if err != nil {
		log.Error.Println("could not query scores GetBestScores", err)
		return scores, err
	}

	for row.Next() {
		var score ScoreBoard
		var grade int
		var status int

		if err := row.Scan(
			&score.ScoreID,
			&score.Score.Name, &score.Score.DateTime, &score.Score.TotalScore,
			&grade, &score.Score.Accuracy, &score.Score.RankedAccuracy,
			&score.Score.Mods, &score.Score.PerformanceRating, &score.Score.JudgementWindowPreset,

			&score.MapID,
			&score.Map.Directory, &score.Map.Path,
			&score.Map.BackgroundPath, &score.Map.Artist, &score.Map.Title, &score.Map.DifficultyName,
			&score.Map.Creator, &status, &score.Map.LNPercent,
		); err != nil {
			log.Error.Println("could not scan GetBestScores", err)

			return scores, err
		}

		score.Score.Grade = api.ConvertGrade(score.Score.Accuracy, grade)
		score.Map.RankedStatus = api.ConvertRankedStatus(status)
		scores = append(scores, score)
	}

	return scores, nil
}

func GetRecentScores(
	uid, mode, page int,

	// Optional
	hideFailed bool,
) ([]ScoreBoard, error) {
	var scores []ScoreBoard

	// Configure optional
	failed := -1

	if hideFailed {
		failed = 5
	}

	// Query
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
			Map.DifficultyName, Map.Creator, Map.RankedStatus,
			((LongNoteCount*1.0) / (RegularNoteCount+LongNoteCount)*1.0) as LN
		FROM Score join Map WHERE Score.MapMd5 = Map.Md5Checksum AND
			Map.Mode = ? AND
			Score.LocalProfileID = ? AND
			Score.Grade != ?
		ORDER BY Score.Id DESC LIMIT ?, 10;
	`, mode, uid, failed, page)

	defer cancel()

	if err != nil {
		log.Error.Println("could not query scores GetRecentScores", err)
		return scores, err
	}

	for row.Next() {
		var score ScoreBoard
		var grade int
		var status int

		if err := row.Scan(
			&score.ScoreID,
			&score.Score.Name, &score.Score.DateTime, &score.Score.TotalScore,
			&grade, &score.Score.Accuracy, &score.Score.RankedAccuracy,
			&score.Score.Mods, &score.Score.PerformanceRating, &score.Score.JudgementWindowPreset,

			&score.MapID,
			&score.Map.Directory, &score.Map.Path,
			&score.Map.BackgroundPath, &score.Map.Artist, &score.Map.Title, &score.Map.DifficultyName,
			&score.Map.Creator, &status, &score.Map.LNPercent,
		); err != nil {
			log.Error.Println("could not scan GetRecentScores", err)

			return scores, err
		}

		score.Score.Grade = api.ConvertGrade(score.Score.Accuracy, grade)
		score.Map.RankedStatus = api.ConvertRankedStatus(status)
		scores = append(scores, score)
	}

	return scores, nil
}
