package app

import (
	"encoding/json"

	"github.com/Cyan903/QuaverBuddy/backend/pkg/config"
	"github.com/Cyan903/QuaverBuddy/backend/pkg/log"
)

func ccopy(orig *config.Config) (*config.Config, error) {
	og, err := json.Marshal(orig)

	if err != nil {
		return nil, err
	}

	clone := config.Config{}

	if err = json.Unmarshal(og, &clone); err != nil {
		return nil, err
	}

	return &clone, nil
}

func (a *App) GetConfig() config.Config {
	return config.Data
}

func (a *App) SetGamePath(path string) error {
	cfg, err := ccopy(&config.Data)

	if err != nil {
		log.Error.Println("could not clone config", err)
		return err
	}

	cfg.GamePath = path
	js, err := json.Marshal(cfg)

	if err != nil {
		log.Error.Println("could not marshal game path", err)
		return err
	}

	return config.WriteConfig(string(js))
}

// TODO: Add main gamemode to config.json
