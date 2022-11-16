package configs

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Configs struct {
	ElasticSearchConfigs ElasticSearchConfig
}

type ElasticSearchConfig struct {
	BaseURL string
	CACert  []byte
}

func ReadConfigs() *Configs {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Failed to read .env: %v", err)
	}

	absPath, _ := filepath.Abs("../http_ca.crt")
	cert, _ := ioutil.ReadFile(absPath)

	elasticSearchConfig := ElasticSearchConfig{
		BaseURL: os.Getenv("ELASTIC_SEARCH.BASE_URL"),
		CACert:  cert,
	}

	return &Configs{
		ElasticSearchConfigs: elasticSearchConfig,
	}
}
