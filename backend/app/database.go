package app

import "github.com/Cyan903/QuaverBuddy/backend/internal/database"

func (a *App) LoadDB() error {
	return database.LoadDB()
}

func (a *App) GetBestScores(uid, mode, page int) ([]database.ScoreBoard, error) {
	return database.GetBestScores(uid, mode, page, "", "Not Submitted", -1.0)
}

func (a *App) GetRecentScores(uid, mode, page int) ([]database.ScoreBoard, error) {
	return database.GetRecentScores(uid, mode, page, false)
}
