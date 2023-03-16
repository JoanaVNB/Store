package elasticSearch

import (
	"BE-JoanaVidon/user-api/domain"
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esutil"
)

func IndexUsersAsDocuments(ctx context.Context) {

	users := ctx.Value(UserKey).([]domain.User)
	client := ctx.Value(ClientKey).(*elasticsearch.Client)

	
/* 	for documentID, document := range users {
		res, err := client.Index("users", esutil.NewJSONReader(document),
			client.Index.WithDocumentID(strconv.Itoa(documentID)))
		if err != nil {
			log.Fatalf("Error getting response: %s", err)
		}
		if res != nil{
			fmt.Printf("✅ Users indexed on Elasticsearch")
		}
	} */
	

	bulkIndexer, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Index:      "users",
		Client:     client,
		NumWorkers: 5,
	})
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	for documentID, document := range users {
		err = bulkIndexer.Add(
			ctx,
			esutil.BulkIndexerItem{
				Action:     "index",
				DocumentID: strconv.Itoa(documentID),
				Body:       esutil.NewJSONReader(document),
			},
		)
		if err != nil {
			log.Fatalf("Error to index: %s", err)
		}
	}

	bulkIndexer.Close(ctx)
	biStats := bulkIndexer.Stats()
	fmt.Printf("✅ Users indexed on ElasticSearch: %d \n", biStats.NumIndexed)

}
