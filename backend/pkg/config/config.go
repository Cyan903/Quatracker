package config

import (
	"encoding/json"
	"os"
	"path"

	"github.com/Cyan903/Quatracker/backend/pkg/log"
	"github.com/adrg/xdg"
)

type Config struct {
	GamePath string `json:"GamePath"`
	MainMode bool   `json:"MainMode"`
}

var cfgpath = path.Join("Quatracker", "config.json")
var Data Config

func CreateConfig() string {
	cfg, err := xdg.ConfigFile(cfgpath)

	if err != nil {
		log.Error.Println("Fatal - could not obtain config file", err)
		os.Exit(1)
	}

	fresh, err := os.Create(cfg)
	jsoncfg, jerr := json.MarshalIndent(&Config{}, "", "    ")

	if err != nil {
		log.Error.Println("Fatal - could not create config file", err)
		os.Exit(1)
	}

	if jerr != nil {
		log.Error.Println("Fatal - could not marshal json", jerr)
		os.Exit(1)
	}

	if _, err := fresh.Write(jsoncfg); err != nil {
		log.Error.Println("Fatal - could not write to fresh config", err)
		os.Exit(1)
	}

	fresh.Close()
	return cfg
}

func LoadConfig() Config {
	cpath, err := xdg.SearchConfigFile(cfgpath)

	if err != nil {
		log.Info.Printf("%s was not found. Creating a new one...\n", cfgpath)
		cpath = CreateConfig()
	}

	cdata, err := os.ReadFile(cpath)
	data := Config{}

	if err != nil {
		log.Error.Println("Fatal - could not read config file", err)
		os.Exit(1)
	}

	if err := json.Unmarshal(cdata, &data); err != nil {
		log.Error.Println("Fatal - invalid config file", err)

		if err := os.Remove(cpath); err != nil {
			log.Error.Println("Fatal - could not remove broken json config", err)
		}

		os.Exit(1)
	}

	return data
}

func WriteConfig(cfg string) error {
	cjson := Config{}
	cpath, cerr := xdg.SearchConfigFile(cfgpath)

	// Find file
	if err := json.Unmarshal([]byte(cfg), &cjson); err != nil {
		log.Error.Println("invalid json", err)
		return err
	}

	if cerr != nil {
		log.Error.Println("Fatal - could not find config file", cerr)
		os.Exit(1)
	}

	// Validate
	if err := validate(cjson); err != nil {
		log.Error.Println("validate error", err)
		return err
	}

	// Write to config, update data
	saved, err := json.MarshalIndent(&cjson, "", "    ")

	if err != nil {
		log.Error.Println("invalid json", err)
		return err
	}

	if err := os.WriteFile(cpath, []byte(saved), 0666); err != nil {
		log.Error.Println("could not save config", err)
		return err
	}

	Data = cjson
	return nil
}
