package app

import (
	"encoding/json"

	"github.com/Cyan903/Quatracker/backend/pkg/config"
	"github.com/Cyan903/Quatracker/backend/pkg/log"
)

func (a *App) GetConfig() config.Config {
	return config.Data
}

func (a *App) SetGamePath(path string) error {
	cfg, err := config.ConfigCopy(&config.Data)

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

func (a *App) SetMainMode(mode bool) error {
	cfg, err := config.ConfigCopy(&config.Data)

	if err != nil {
		log.Error.Println("could not clone config", err)
		return err
	}

	cfg.MainMode = mode
	js, err := json.Marshal(cfg)

	if err != nil {
		log.Error.Println("could not marshal mode", err)
		return err
	}

	return config.WriteConfig(string(js))
}
