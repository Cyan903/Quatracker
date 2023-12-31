package database

import (
	"context"

	"github.com/Cyan903/Quatracker/backend/pkg/log"
)

type LocalUsers struct {
	Id       string
	Username string
}

type Users struct {
	Local   []LocalUsers
	Unknown []LocalUsers
}

func userFetch(locals *[]LocalUsers, query string) error {
	c, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	row, err := Conn.QueryContext(c, query)

	defer cancel()

	if err != nil {
		log.Error.Println("could not query users userFetch", err)
		return err
	}

	for row.Next() {
		var local LocalUsers

		if err := row.Scan(
			&local.Id,
			&local.Username,
		); err != nil {
			log.Error.Println("could not scan userFetch", err)
			return err
		}

		*locals = append(*locals, local)
	}

	return nil
}

func GetUsers() (Users, error) {
	var users Users

	// Get local profiles
	if err := userFetch(&users.Local, `SELECT DISTINCT Id, Username FROM UserProfile`); err != nil {
		log.Error.Println("could not get users", err)
		return users, err
	}

	// Get scores with no profiles attached to them
	if err := userFetch(&users.Unknown, `
		SELECT DISTINCT
			LocalProfileId as Id,
			"Unknown" as Username
		FROM Score WHERE
			LocalProfileId NOT IN
			(SELECT Id FROM UserProfile)
	`); err != nil {
		log.Error.Println("could not get unknown users", err)
		return users, err
	}

	return users, nil
}

func GetUserJudgements() ([]string, error) {
	var judgements []string
	c, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	row, err := Conn.QueryContext(c, `SELECT DISTINCT JudgementWindowPreset FROM Score`)

	defer cancel()

	if err != nil {
		log.Error.Println("could not query users GetUserJudgements", err)
		return judgements, err
	}

	for row.Next() {
		var judge string

		if err := row.Scan(&judge); err != nil {
			log.Error.Println("could not scan GetUserJudgements", err)
			return judgements, err
		}

		judgements = append(judgements, judge)
	}

	return judgements, nil
}
