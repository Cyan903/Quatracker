package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Cyan903/QuaverBuddy/backend/pkg/api"
	"github.com/Cyan903/QuaverBuddy/backend/pkg/log"
)

type ScoreBoard struct {
	ScoreID int
	MapID   int

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
	}

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
	}
}

type ScoreDetails struct {
	ScoreID int
	MapID   int

	Score struct {
		LocalProfileId        int
		Name                  string
		DateTime              string
		TotalScore            int
		Grade                 string
		MaxCombo              int
		Mods                  string
		Mode                  int
		ScrollSpeed           int
		PauseCount            int
		PerformanceRating     float64
		JudgementWindowPreset string

		JudgedHits struct {
			Accuracy       float64
			RankedAccuracy float64
			CountMarv      int
			CountPerf      int
			CountGreat     int
			CountGood      int
			CountOkay      int
			CountMiss      int
		}

		JudgementConfig struct {
			CountMarv  int
			CountPerf  int
			CountGreat int
			CountGood  int
			CountOkay  int
			CountMiss  int
		}

		Versions struct {
			RatingVersion     string
			DifficultyVersion string
		}
	}

	Map struct {
		Directory      string
		Path           string
		BackgroundPath string
		Artist         string
		Title          string
		DifficultyName string
		Creator        string
		RankedStatus   string

		DifficultyInfo struct {
			SongLength int
			BPM        float64
			LNPercent  float64
			Rating     float64
		}

		ModeInfo struct {
			DifficultyProcessorVersion string
			HasScratchKey              bool
		}
	}
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
					AND Score.LocalProfileID = ? %[1]s %[2]s
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
		mode, uid, uid,

		// Optional
		judgementWindowPreset, api.RevertRankedStatus(RankedStatus),
		judgementWindowPreset, api.RevertRankedStatus(RankedStatus), LNPercent,

		page*DefaultPaginate,
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
	`, mode, uid, failed, page*DefaultPaginate)

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

func GetScoreDetails(id int) (ScoreDetails, error) {
	var details ScoreDetails
	var grade int
	var status int

	c, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	query := Conn.QueryRowContext(c, `
		SELECT
			Score.Id as SID,
			Score.LocalProfileId, Score.Name, Score.DateTime, Score.TotalScore,
			Score.Grade, Score.Accuracy, Score.MaxCombo, Score.CountMarv,
			Score.CountPerf, Score.CountGreat, Score.CountGood, Score.CountOkay, Score.CountMiss,
			Score.Mods, Score.Mode, Score.ScrollSpeed, Score.PauseCount, Score.PerformanceRating,
			Score.RatingProcessorVersion, Score.DifficultyProcessorVersion, Score.JudgementWindowPreset,
			Score.JudgementWindowMarv, Score.JudgementWindowPerf, Score.JudgementWindowGreat, Score.JudgementWindowGood,
			Score.JudgementWindowOkay, Score.JudgementWindowMiss, Score.RankedAccuracy,

			Map.Id as MID,
			Map.Directory, Map.Path, Map.BackgroundPath, Map.Artist, Map.Title, Map.DifficultyName, Map.Creator, Map.RankedStatus,
			Map.SongLength, Map.Bpm, Map.DifficultyProcessorVersion,
			((LongNoteCount*1.0) / (RegularNoteCount+LongNoteCount)*1.0) as LN,
			Map.HasScratchKey, Map.Difficulty10x
		FROM Score join Map WHERE 
			Score.MapMd5 = Map.Md5Checksum AND
			Score.Id = ?
	`, id)

	defer cancel()

	if err := query.Scan(
		&details.ScoreID,
		&details.Score.LocalProfileId, &details.Score.Name, &details.Score.DateTime, &details.Score.TotalScore,
		&grade, &details.Score.JudgedHits.Accuracy, &details.Score.MaxCombo, &details.Score.JudgedHits.CountMarv,
		&details.Score.JudgedHits.CountPerf, &details.Score.JudgedHits.CountGreat, &details.Score.JudgedHits.CountGood, &details.Score.JudgedHits.CountOkay, &details.Score.JudgedHits.CountMiss,
		&details.Score.Mods, &details.Score.Mode, &details.Score.ScrollSpeed, &details.Score.PauseCount, &details.Score.PerformanceRating,
		&details.Score.Versions.RatingVersion, &details.Score.Versions.DifficultyVersion, &details.Score.JudgementWindowPreset,
		&details.Score.JudgementConfig.CountMarv, &details.Score.JudgementConfig.CountPerf, &details.Score.JudgementConfig.CountGreat, &details.Score.JudgementConfig.CountGood,
		&details.Score.JudgementConfig.CountOkay, &details.Score.JudgementConfig.CountMiss, &details.Score.JudgedHits.RankedAccuracy,

		&details.MapID,
		&details.Map.Directory, &details.Map.Path, &details.Map.BackgroundPath, &details.Map.Artist, &details.Map.Title, &details.Map.DifficultyName, &details.Map.Creator, &status,
		&details.Map.DifficultyInfo.SongLength, &details.Map.DifficultyInfo.BPM, &details.Map.ModeInfo.DifficultyProcessorVersion,
		&details.Map.DifficultyInfo.LNPercent,
		&details.Map.ModeInfo.HasScratchKey, &details.Map.DifficultyInfo.Rating,
	); err != nil && err != sql.ErrNoRows {
		log.Error.Println("could not scan GetScoreDetails", err)
		return details, err
	}

	details.Score.Grade = api.ConvertGrade(details.Score.JudgedHits.Accuracy, grade)
	details.Map.RankedStatus = api.ConvertRankedStatus(status)
	return details, nil
}
