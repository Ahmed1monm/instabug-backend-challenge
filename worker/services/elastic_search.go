package services

import (
	"context"

	es "github.com/elastic/go-elasticsearch/v8"
)

var typedClient, _ = es.NewTypedClient(es.Config{
	Addresses: []string{"http://elasticsearch:9200"},
})

func SetupElasticSearch() error {
	_, err := typedClient.Indices.Create("messages").Do(context.Background())

	if err != nil {
		return err
	}

	return nil
}

func Index(id string, body []byte) error {
	_, err := typedClient.Search().Index("messages").AllowPartialSearchResults(true).Do(context.Background())

	if err != nil {
		return err
	}

	return nil
}
