package app

import "github.com/Cyan903/QuaverBuddy/backend/internal/database"

func (a *App) LoadDB() error {
	return database.LoadDB()
}

func (a *App) GetScores() ([]string, error) {
	return database.GetScores()
}
