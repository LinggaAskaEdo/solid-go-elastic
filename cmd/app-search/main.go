package main

import (
	"crypto/x509"
	"io/ioutil"
	"log"

	"github.com/linggaaskaedo/solid-go-elastic/configs"
	"github.com/linggaaskaedo/solid-go-elastic/internal/search/handler"
	"github.com/linggaaskaedo/solid-go-elastic/internal/search/repository"
	"github.com/linggaaskaedo/solid-go-elastic/internal/search/service"
)

var (
	clusterURLs = []string{"https://localhost:9200"}
	username    = "elastic"
	password    = "elactic"
)

func main() {
	// read a configs instance
	configs := configs.ReadConfigs()

	caCert, err := ioutil.ReadFile(configs.ElasticSearchConfigs.Cert)
	if err != nil {
		log.Fatal(err)
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// construct a repository
	repository := repository.NewElasticSearch(configs.ElasticSearchConfigs.BaseURL, caCertPool)

	// do a health check
	// if error abort operation immediately
	checkHealth(repository)

	// construct services
	searchService := service.NewSearchService(repository)
	syncService := service.NewSyncService(repository)
	constructService := service.NewConstructService(repository)

	// initiate an index
	constructService.CreateIndex()

	// construct a rest handler
	rest := handler.NewRest(searchService, syncService)

	// listen and serve server
	rest.ListenAndServe()
}

func checkHealth(repository repository.Repository) {
	if err := repository.CheckHealth(); err != nil {
		log.Fatalf("missing elasticsearch connection: %v", err)
	}
}
