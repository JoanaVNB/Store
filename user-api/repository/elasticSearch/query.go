package elasticSearch

import (
	"BE-JoanaVidon/user-api/presenter"
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"time"
	"github.com/elastic/go-elasticsearch/v8"
)

func QueryUsersByEmail(ctx context.Context, email string) {
	client := ctx.Value(ClientKey).(*elasticsearch.Client)

	var buf bytes.Buffer
	query := map[string]interface{}{
			"query": map[string]interface{}{
					"match": map[string]interface{}{
							"email": email,
					},
			},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
			log.Fatalf("Error encoding query: %s", err)
	}

	response, err := client.Search(
			client.Search.WithContext(ctx),
			client.Search.WithIndex("users"),
			client.Search.WithBody(&buf),
			client.Search.WithTrackTotalHits(true),
			client.Search.WithPretty(),
	)
	if err != nil {
			log.Fatalf("Error searching for users: %s", err)
	}

	defer func() {
			if response.Body != nil {
					_ = response.Body.Close()
			}
	}()

	if response.IsError() {
			log.Fatalf("Error searching for users: %s", response.Status())
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
			log.Fatalf("Error reading response: %s", err)
	}
	
	type hits struct{
		Index string `json:"_index"`
		Type string `json:"_type"`
		ID   string `json:"_id"`
		Score float32  `json:"_score"`
		Source *presenter.GetUser `json:"_source"`
		CreatedAt time.Time `json:"_created_at"`
		UpdatedAt time.Time `json:"_updated_at"`
	}

	type searchResult struct {
    Hits struct {
        Total struct {
            Value    int    `json:"value"`
            Relation string `json:"relation"`
        } `json:"total"`
        MaxScore float64 `json:"max_score"`
        Hits     []hits  `json:"hits"`
    } `json:"hits"`
}


	var searchRes searchResult

		err = json.Unmarshal(responseBytes, &searchRes)
	if err != nil {
    log.Fatalf("Error decoding response: %s", err)
}

	for _, hit := range searchRes.Hits.Hits {
			log.Printf("User with email %s found: ID %s", email, hit.ID)
	}
}
