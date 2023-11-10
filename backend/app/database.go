package app

import "github.com/Cyan903/QuaverBuddy/backend/internal/database"

func (a *App) LoadDB() error {
	return database.LoadDB()
}

func (a *App) GetBestScores(uid, mode, page int) ([]database.ScoreBoard, error) {
	return database.GetBestScores(uid, mode, page, "OD 0", "", -1.0)
}

func (a *App) GetRecentScores(uid, mode, page int) ([]database.ScoreBoard, error) {
	return database.GetRecentScores(uid, mode, page, false)
}

func (a *App) GetScoreDetails(score int) (database.ScoreDetails, error) {
	return database.GetScoreDetails(score)
}

func (a *App) GetUsers() (database.Users, error) {
	return database.GetUsers()
}

func (a *App) GetUserJudgements() ([]string, error) {
	return database.GetUserJudgements()
}
