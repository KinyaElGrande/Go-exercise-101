package includes

import (
	"encoding/json"
	"os"
)

//Config load emojis
type Config struct {
	Player   string `json:"player"`
	Ghost    string `json:"ghost"`
	Wall     string `json:"wall"`
	Dot      string `json:"dot"`
	Pill     string `json:"pill"`
	Death    string `json:"death"`
	Space    string `json:"space"`
	UseEmoji bool   `json:"use_emoji"`
}

var cfg Config

func loadConfig(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}

	defer f.Close()

	decoder := json.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		return err
	}

	return nil
}
