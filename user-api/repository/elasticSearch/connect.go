package elasticSearch

import (
	"context"
	"log"
	"github.com/elastic/go-elasticsearch/v8"
	"fmt"
)

func ConnectionWithElasticSearch(ctx context.Context) context.Context {
	newClient, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
/* 			Username: "adm",
			Password: "Pass123!", */
		})
	if err != nil{
		log.Fatalf("Error creating the client: %s", err)
	}
	fmt.Printf("✅ Client connected\n")
	return context.WithValue(ctx, ClientKey, newClient)
}
