package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Configs struct {
	ElasticSearchConfigs ElasticSearchConfig
}

type ElasticSearchConfig struct {
	BaseURL string
	Cert    string
}

func ReadConfigs() *Configs {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Failed to read .env: %v", err)
	}

	// absPath, _ := filepath.Abs("../http_ca.crt")
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(dir)

	elasticSearchConfig := ElasticSearchConfig{
		BaseURL: os.Getenv("ELASTIC_SEARCH.BASE_URL"),
		Cert:    dir + "/http_ca.crt",
	}

	return &Configs{
		ElasticSearchConfigs: elasticSearchConfig,
	}
}
