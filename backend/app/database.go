package app

import (
	"github.com/Cyan903/Quatracker/backend/internal/database"
	"github.com/Cyan903/Quatracker/backend/pkg/api"
	"github.com/Cyan903/Quatracker/backend/pkg/log"
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
	total, err := database.GetTotalRecent(uid, mode, failed)

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

// Stats
func (a *App) GetOverallStats(uid, mode int) (database.OverallStats, error) {
	total, err := database.GetTotalRecent(uid, mode, true)

	if err != nil {
		log.Warning.Println("GetOverallStats - GetRecentScores failed to count")
		return database.OverallStats{}, err
	}

	if total <= 0 {
		log.Warning.Println("GetOverallStats - no scores found")
		return database.OverallStats{}, database.ErrNoScores
	}

	return database.GetOverallStats(uid, mode)
}

func (a *App) GetTotalStats(uid, mode int) (database.TotalStats, error) {
	total, err := database.GetTotalRecent(uid, mode, false)

	if err != nil {
		log.Warning.Println("GetTotalStats - GetRecentScores failed to count")
		return database.TotalStats{}, err
	}

	if total <= 0 {
		log.Warning.Println("GetTotalStats - no scores found")
		return database.TotalStats{}, database.ErrNoScores
	}

	return database.GetTotalStats(uid, mode)
}

func (a *App) GetPlayStats(uid, mode int) (database.PlayStats, error) {
	return database.GetPlayStats(uid, mode)
}

func (a *App) GetLongestStats(uid, mode int) (database.LongestStat, error) {
	total, err := database.GetTotalRecent(uid, mode, false)

	if err != nil {
		log.Warning.Println("GetLongestStats - GetRecentScores failed to count")
		return database.LongestStat{}, err
	}

	if total <= 0 {
		log.Warning.Println("GetLongestStats - no scores found")
		return database.LongestStat{}, database.ErrNoScores
	}

	return database.GetLongestStats(uid, mode)
}

func (a *App) GetMostPlayed(uid, mode, page int) ([]database.MostPlayed, error) {
	return database.GetMostPlayed(uid, mode, page)
}

func (a *App) GetJudgesCountStats(uid, mode int) (database.JudgeCountStat, error) {
	total, err := database.GetTotalRecent(uid, mode, false)

	if err != nil {
		log.Warning.Println("GetJudgesCountStats - GetRecentScores failed to count")
		return database.JudgeCountStat{}, err
	}

	if total <= 0 {
		log.Warning.Println("GetJudgesCountStats - no scores found")
		return database.JudgeCountStat{}, database.ErrNoScores
	}

	return database.GetJudgesCountStats(uid, mode)
}

func (a *App) GetGradeCount(uid, mode int, pb bool) (database.GradesCountStat, error) {
	return database.GetGradeCount(uid, mode, pb)
}

func (a *App) GetUserPlayGraph(uid, mode int) (map[string]int, error) {
	dates := make(map[string]int)
	total, err := database.GetTotalRecent(uid, mode, false)

	// Has scores?
	if err != nil {
		log.Warning.Println("GetUserPlayGraph - GetRecentScores failed to count")
		return dates, err
	}

	if total <= 0 {
		log.Warning.Println("GetUserPlayGraph - no scores found")
		return dates, database.ErrNoScores
	}

	// Fetch/calculate range
	earliest, err := database.GraphDateEarliest(uid, mode)

	if err != nil {
		log.Error.Println("could not get earliest date GetUserPlayGraph")
		return dates, err
	}

	database.GraphCreateRange(earliest, dates)

	if err := database.CountPlaycountDates(uid, mode, dates); err != nil {
		log.Error.Println("could not calculate playcount GetUserPlayGraph")
		return dates, err
	}

	return dates, nil
}
