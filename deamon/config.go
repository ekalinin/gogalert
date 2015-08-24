package deamon

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Email     EmailConfig `json:"email"`
	Transport [][]string  `json:"transport"`
	Interval  int         `json:"interval"`
	Debug     bool        `json:"debug"`
	Pid       string      `json:"pid"`
	LogFile   string      `json:"logfile"`
	Checks    [][]string  `json:"checks"`
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

func (this *Config) getTransport(transportType string) string {
	for _, transport := range this.Transport {
		if transport[0] == "smtp" {
			return transport[1] + ":" + transport[2]
		}
	}
	return nil
}
