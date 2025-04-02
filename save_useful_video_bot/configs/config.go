package configs

import (
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	Messages struct {
		Start            string `yaml:"start"`
		AskName          string `yaml:"ask_name"`
		AskDOB           string `yaml:"ask_dob"`
		InvalidDOBFormat string `yaml:"invalid_dob_format"`
		InvalidDOB       string `yaml:"invalid_dob"`
		WeeksLived       string `yaml:"weeks_lived"`
		DOBAlreadySet    string `yaml:"dob_already_set"`
		ChangeDOB        string `yaml:"change_dob"`
		UnknownCommand   string `yaml:"unknown_command"`
		StartWith        string `yaml:"start_with"`
	} `yaml:"messages"`

	Database struct {
		Path string `yaml:"path"`
	} `yaml:"database"`

	BotToken string
}

func LoadConfig() *Config {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	data, err := os.ReadFile(os.Getenv("CONFIG_PATH"))
	if err != nil {
		log.Fatalf("Error loading main.yml: %v", err)
	}

	var conf Config
	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		log.Fatalf("Error loading main.yml: %v", err)
	}

	conf.BotToken = os.Getenv("BOT_TOKEN")
	if conf.BotToken == "" {
		log.Fatal("Error loading bot token")
	}
	return &conf
}
