package app

import (
	"github.com/Cyan903/QuaverBuddy/backend/internal/database"
	"github.com/Cyan903/QuaverBuddy/backend/pkg/api"
	"github.com/Cyan903/QuaverBuddy/backend/pkg/log"
)

func (a *App) LoadDB() error {
	return database.LoadDB()
}

// Scores
func (a *App) GetBestScores(uid, mode, page int, judgement, mstatus string, ln float64) ([]database.ScoreBoard, error) {
	if mstatus != "" && api.RevertRankedStatus(mstatus) == -1 {
		log.Warning.Println("invalid rank status passed", mstatus)
	}

	return database.GetBestScores(uid, mode, page, judgement, mstatus, ln)
}

func (a *App) GetRecentScores(uid, mode, page int, failed bool) (database.CountedScoreBoard, error) {
	var counted database.CountedScoreBoard
	total, err := database.GetTotalRecent(uid, mode, page, failed)

	if err != nil {
		log.Warning.Println("GetRecentScores failed to count")
		return counted, err
	}

	scores, err := database.GetRecentScores(uid, mode, page, failed)

	if err != nil {
		log.Warning.Println("GetRecentScores failed to get scores")
		return counted, err
	}

	counted.Total = total
	counted.Scores = scores
	return counted, nil
}

func (a *App) GetScoreDetails(id int) (database.ScoreDetails, error) {
	return database.GetScoreDetails(id)
}

// Users
func (a *App) GetUsers() (database.Users, error) {
	return database.GetUsers()
}

func (a *App) GetUserJudgements() ([]string, error) {
	return database.GetUserJudgements()
}
