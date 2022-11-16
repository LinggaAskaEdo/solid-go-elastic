package repository

import "github.com/linggaaskaedo/solid-go-elastic/internal/search/model"

type SearchHits struct {
	Hits Hits `json:"hits"`
}

type Hits struct {
	Hits []*Hit `json:"hits"`
}

type Hit struct {
	Source *model.Employee `json:"_source"`
}
