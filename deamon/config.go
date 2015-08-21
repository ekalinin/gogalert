package deamon

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Email    EmailConfig `json:"email"`
	Interval int         `json:"interval"`
	Debug    bool        `json:"debug"`
	Pid      string      `json:"pid"`
	LogFile  string      `json:"logfile"`
	Checks   [][]string  `json:"checks"`
}

type EmailConfig struct {
	From  string
	To    string
	Title string
	Body  string
}

func ParseConfig(configPath string) *Config {
	var config Config
	file, err := os.Open(configPath)
	if err != nil {
		panic(fmt.Sprintf("cannot open file %s: %v", configPath, err))
	}
	defer file.Close()
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		panic(fmt.Sprintf("cannot decode file %s: %v", configPath, err))
	}
	return &config
}
