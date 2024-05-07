package config

import (
	config "github.com/stvp/go-toml-config"
)

// var (
// 	country            = config.String("country", "Unknown")
// 	atlantaEnabled     = config.Bool("atlanta.enabled", false)
// 	alantaPopulation   = config.Int("atlanta.population", 0)
// 	atlantaTemperature = config.Float("atlanta.temperature", 0)
// )

const CONFIG_PATH = "./config.conf"

func SetVars() {
	config.Parse(CONFIG_PATH)
}
