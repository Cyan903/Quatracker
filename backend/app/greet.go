package app

import (
	"fmt"

	"github.com/Cyan903/QuaverBuddy/backend/pkg/config"
)

func (a *App) Greet(name string) (string, error) {
	err := config.WriteConfig(name)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Hello %s, It's show time!", config.Data.GamePath), nil
}
