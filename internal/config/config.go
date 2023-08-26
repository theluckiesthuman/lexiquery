package config

import (
	"log"
	"os"
	"regexp"
	"sync"

	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
)

func init() {
	loadConfig()
}

var (
	cfg  config
	once sync.Once
)

const projectDirName = "LexiQuery"

type config struct {
	CollegitaeURL string `env:"COLLEGIATE_URL,required"`
	APIKey        string `env:"API_KEY,required"`
}

func GetCollegiateURL() string {
	return cfg.CollegitaeURL
}

func GetAPIKey() string {
	return cfg.APIKey
}

func loadEnvFile() {
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	_ = godotenv.Load(string(rootPath) + `/.env`)

}

func loadConfig() {
	once.Do(func() {
		//load env file
		loadEnvFile()
		err := env.Parse(&cfg)
		if err != nil {
			log.Fatalf("unable to parse ennvironment variables: %e", err)
		}
	})
}
