package graphics

import (
	_ "embed"
	"encoding/json"
	"log"
)

//go:embed resources/config.json
var configData []byte

type Config struct {
	Window struct {
		DefaultWindowWidth  int `json:"defaultWindowWidth"`
		DefaultWindowHeight int `json:"defaultWindowHeight"`
	} `json:"window"`
	Graphics struct {
		TileSize int `json:"tileSize"`
	} `json:"graphics"`
	Input struct {
		ZoomFactor float64 `json:"zoomFactor"`
	} `json:"input"`
}

func LoadConfig() (*Config, error) {

	var cfg Config

	err := json.Unmarshal(configData, &cfg)
	if err != nil {
		log.Fatalf("Error parsing config.json: %v", err)
	}

	return &cfg, nil
}
