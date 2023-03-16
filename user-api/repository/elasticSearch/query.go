package elasticSearch

import (
	"BE-JoanaVidon/user-api/presenter"
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"time"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
)

func QueryUsersByEmail(ctx context.Context, email string) {
	client := ctx.Value(ClientKey).(*elasticsearch.Client)

	var buf bytes.Buffer
	query := map[string]interface{}{
			"query": map[string]interface{}{
				"bool": map[string]interface{}{
					"must": map[string]interface{}{
							"term": map[string]string{
								"email.keyword": email,
						},
					},
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

	var userID string
	if searchRes.Hits.Total.Value > 0 {
		for _, id := range searchRes.Hits.Hits{
			userID = id.Source.ID
		}
		fmt.Printf("✅ User with email %s is the ID: %s\n\n", email, userID)
	}
}