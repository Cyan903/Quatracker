package config

import (
	"encoding/json"
	"errors"

	"github.com/Cyan903/QuaverBuddy/backend/pkg/log"
)

var ErrMissingItem = errors.New("missing field")

func validate(cjson Config) error {
	if cjson.GamePath == "" {
		log.Error.Println("GamePath missing field")
		return ErrMissingItem
	}

	return nil
}

func ConfigCopy(orig *Config) (*Config, error) {
	og, err := json.Marshal(orig)

	if err != nil {
		return nil, err
	}

	clone := Config{}

	if err = json.Unmarshal(og, &clone); err != nil {
		return nil, err
	}

	return &clone, nil
}
