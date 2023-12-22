package database

import (
	"context"
	"fmt"
	"math"

	"github.com/Cyan903/QuaverBuddy/backend/pkg/api"
	"github.com/Cyan903/QuaverBuddy/backend/pkg/log"
	"github.com/araddon/dateparse"
)

type OverallStats struct {
	Accuracy       float64
	RankedAccuracy float64
	Performance    float64
}

type TotalStats struct {
	Hits  uint64
	Score uint64
}

type PlayStats struct {
	Playcount int
	Passed    int
	Failed    int
}

type LongestMap struct {
	ScoreID int
	MapID   int

	Map struct {
		Title          string
		DifficultyName string
	}

	Total int
}

type LongestStat struct {
	Hits  LongestMap
	Combo LongestMap
}

type MostPlayed struct {
	MapID          int
	Title          string
	DifficultyName string
	PlayCount      int
}

type JudgeCountStat struct {
	CountMarv  int
	CountPerf  int
	CountGreat int
	CountGood  int
	CountOkay  int
	CountMiss  int
}

type GradesCountStat struct {
	AllGrades    map[string]int
	RankedGrades map[string]int
}

func GetOverallStats(uid, mode int) (OverallStats, error) {
	var stats OverallStats
	var acc, racc [100]float64
	var perf [100]float64
	var ai, i int

	c, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	row, err := Conn.QueryContext(c, `
		SELECT
			Score.Accuracy, Score.RankedAccuracy,
			Score.PerformanceRating
		FROM Score join Map WHERE Score.MapMd5 = Map.Md5Checksum
			AND Score.LocalProfileID = ?
			AND Map.Mode = ?
			AND (
				SELECT Id FROM Score WHERE MapMd5 = Map.Md5Checksum
					AND Score.LocalProfileID = ?
				ORDER BY PerformanceRating DESC, RankedAccuracy DESC LIMIT 1
			) == Score.Id
			AND Score.Grade != 5
		ORDER BY Score.PerformanceRating DESC LIMIT 100;
	`, uid, mode, uid)

	defer cancel()

	if err != nil {
		log.Error.Println("could not query overall GetOverallStats", err)
		return stats, err
	}

	for row.Next() {
		if err := row.Scan(
			&acc[i], &racc[i], &perf[i],
		); err != nil {
			log.Error.Println("could not scan GetOverallStats")
			return stats, err
		}

		i++
	}

	// Calculate accuracy
	for _, n := range acc {
		stats.Accuracy += n

		if n != 0.0 {
			ai++
		}
	}

	for _, n := range racc {
		stats.RankedAccuracy += n
	}

	stats.Accuracy /= float64(ai)
	stats.RankedAccuracy /= float64(ai)

	// Calculate performance
	for pi, per := range perf {
		if per == 0 {
			continue
		}

		stats.Performance += per * math.Pow(0.95, float64(pi))
	}

	return stats, nil
}

func GetTotalStats(uid, mode int) (TotalStats, error) {
	var total TotalStats

	c, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	query := Conn.QueryRowContext(c, `
		SELECT
			SUM(CountMarv) + SUM(CountPerf) + SUM(CountGreat) +
			SUM(CountGood) + SUM(CountOkay) + SUM(CountMiss),
			SUM(TotalScore)
		FROM Score join Map WHERE Score.MapMd5 = Map.Md5Checksum
			AND Score.LocalProfileID = ?
			AND Map.Mode = ?
	`, uid, mode)

	defer cancel()

	if err := query.Scan(&total.Hits, &total.Score); err != nil {
		log.Error.Println("could not query total GetTotalStats", err)
		return total, err
	}

	return total, nil
}

func GetPlayStats(uid, mode int) (PlayStats, error) {
	var plays PlayStats

	c, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	pass := Conn.QueryRowContext(
		c, `SELECT COUNT(1) FROM Score join Map WHERE Score.MapMd5 = Map.Md5Checksum AND Score.LocalProfileID = ? AND Map.Mode = ?;`,
		uid, mode,
	)

	fail := Conn.QueryRowContext(
		c, `SELECT COUNT(1) FROM Score join Map WHERE Score.MapMd5 = Map.Md5Checksum AND Score.Grade = 5 AND Score.LocalProfileID = ? AND Map.Mode = ?;`,
		uid, mode,
	)

	defer cancel()

	if err := pass.Scan(&plays.Playcount); err != nil {
		log.Error.Println("could not query passed GetPlayStats", err)
		return plays, err
	}

	if err := fail.Scan(&plays.Failed); err != nil {
		log.Error.Println("could not query failed GetPlayStats", err)
		return plays, err
	}

	plays.Passed = plays.Playcount - plays.Failed
	return plays, nil
}

func GetLongestStats(uid, mode int) (LongestStat, error) {
	var longest LongestStat

	c, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	hits := Conn.QueryRowContext(c, `
		SELECT
			Score.Id as SID, Map.Id as MID,
			Map.Title, Map.DifficultyName, (
				Score.CountMarv + Score.CountPerf + Score.CountGreat +
				Score.CountGood + Score.CountOkay + Score.CountMiss
			) as Hits
		FROM Score join Map WHERE Score.MapMd5 = Map.Md5Checksum
			AND Score.LocalProfileID = ? AND Map.Mode = ?
		ORDER BY Hits DESC LIMIT 1;
	`, uid, mode)

	combo := Conn.QueryRowContext(c, `
		SELECT
			Score.Id as SID, Map.Id as MID,
			Map.Title, Map.DifficultyName, Score.MaxCombo
		FROM Score join Map WHERE Score.MapMd5 = Map.Md5Checksum
			AND Score.LocalProfileID = ? AND Map.Mode = ?
		ORDER BY MaxCombo DESC LIMIT 1;
	`, uid, mode)

	defer cancel()

	if err := hits.Scan(
		&longest.Hits.ScoreID,
		&longest.Hits.MapID,
		&longest.Hits.Map.Title,
		&longest.Hits.Map.DifficultyName,
		&longest.Hits.Total,
	); err != nil {
		log.Error.Println("could not query longest hits GetLongestStats", err)
		return longest, err
	}

	if err := combo.Scan(
		&longest.Combo.ScoreID,
		&longest.Combo.MapID,
		&longest.Combo.Map.Title,
		&longest.Combo.Map.DifficultyName,
		&longest.Combo.Total,
	); err != nil {
		log.Error.Println("could not query longest combo GetLongestStats", err)
		return longest, err
	}

	return longest, nil
}

func GetMostPlayed(uid, mode, page int) ([]MostPlayed, error) {
	var mpmaps []MostPlayed

	c, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	row, err := Conn.QueryContext(c, `
		SELECT DISTINCT
			Map.Id, Map.Title, Map.DifficultyName, Map.TimesPlayed
		FROM Score join Map WHERE Score.MapMd5 = Map.Md5Checksum
			AND Score.LocalProfileID = ?
			AND Map.Mode = ?
		ORDER BY Map.TimesPlayed DESC LIMIT ?, 10;
	`, uid, mode, page)

	defer cancel()

	if err != nil {
		log.Error.Println("could not query maps GetMostPlayed", err)
		return mpmaps, err
	}

	for row.Next() {
		var mpmap MostPlayed

		if err := row.Scan(
			&mpmap.MapID,
			&mpmap.Title,
			&mpmap.DifficultyName,
			&mpmap.PlayCount,
		); err != nil {
			log.Error.Println("could not scan GetMostPlayed", err)
			return mpmaps, err
		}

		mpmaps = append(mpmaps, mpmap)
	}

	return mpmaps, nil
}

func GetJudgesCountStats(uid, mode int) (JudgeCountStat, error) {
	var judges JudgeCountStat

	c, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	query := Conn.QueryRowContext(c, `
		SELECT
			SUM(CountMarv), SUM(CountPerf), SUM(CountGreat),
			SUM(CountGood), SUM(CountOkay), SUM(CountMiss)
		FROM Score join Map WHERE Score.MapMd5 = Map.Md5Checksum
			AND Score.LocalProfileID = ?
			AND Map.Mode = ?
	`, uid, mode)

	defer cancel()

	if err := query.Scan(
		&judges.CountMarv,
		&judges.CountPerf,
		&judges.CountGreat,
		&judges.CountGood,
		&judges.CountOkay,
		&judges.CountMiss,
	); err != nil {
		log.Error.Println("could not scan GetJudgesCountStats", err)
		return judges, err
	}

	return judges, nil
}

func GetGradeCount(uid, mode int, pbOnly bool) (GradesCountStat, error) {
	var grades GradesCountStat
	var pb = "AND Map.Md5Checksum != ?"

	grades.AllGrades = map[string]int{
		"X":  0,
		"SS": 0,
		"S":  0,
		"A":  0,
		"B":  0,
		"C":  0,
		"D":  0,
		"F":  0,
	}

	grades.RankedGrades = map[string]int{
		"X":  0,
		"SS": 0,
		"S":  0,
		"A":  0,
		"B":  0,
		"C":  0,
		"D":  0,
		"F":  0,
	}

	if pbOnly {
		pb = `
			AND (
				SELECT Id FROM Score WHERE MapMd5 = Map.Md5Checksum
					AND Score.LocalProfileID = ?
				ORDER BY PerformanceRating DESC, RankedAccuracy DESC LIMIT 1
			) == Score.Id
			AND Score.Grade != 5
		`
	}

	query := fmt.Sprintf(`
		SELECT
			Score.Accuracy, Score.RankedAccuracy, Score.Grade
		FROM Score join Map WHERE Score.MapMd5 = Map.Md5Checksum
			AND Score.LocalProfileID = ?
			AND Map.Mode = ?
			%s
	`, pb)

	c, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	row, err := Conn.QueryContext(c, query, uid, mode, uid)

	defer cancel()

	if err != nil {
		log.Error.Println("could not query accuracy GetGradeCount", err)
		return grades, err
	}

	for row.Next() {
		var acc float64
		var racc float64
		var grade int

		if err := row.Scan(&acc, &racc, &grade); err != nil {
			log.Error.Println("could not scan GetGradeCount", err)
			return grades, err
		}

		grades.AllGrades[api.ConvertGrade(acc, grade)]++
		grades.RankedGrades[api.ConvertGrade(racc, grade)]++

	}

	return grades, nil
}

func CountPlaycountDates(uid, mode int, dates map[string]int) error {
	c, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	row, err := Conn.QueryContext(c, `
		SELECT Score.DateTime FROM Score join Map WHERE Score.MapMd5 = Map.Md5Checksum
			AND Score.LocalProfileID = ?
			AND Map.Mode = ?
	`, uid, mode)

	defer cancel()

	if err != nil {
		log.Error.Println("could not query playcount CountPlaycountDates", err)
		return err
	}

	for row.Next() {
		var datetime string

		if err := row.Scan(&datetime); err != nil {
			log.Error.Println("could not scan CountPlaycountDates", err)
			return err
		}

		t, err := dateparse.ParseAny(datetime)

		if err != nil {
			log.Warning.Println("invalid date CountPlaycountDates", err)
			log.Warning.Println("date:", datetime)

			continue
		}

		short := fmt.Sprintf("%d/%d", t.Year(), t.Month())
		dates[short]++
	}

	return nil
}
