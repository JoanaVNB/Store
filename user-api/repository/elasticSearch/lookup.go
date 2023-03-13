package elasticSearch

import (
	//"BE-JoanaVidon/user-api/domain"
	"BE-JoanaVidon/user-api/presenter"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"github.com/elastic/go-elasticsearch/v8"
	"io/ioutil"
	"time"
)

//GET /users/_doc/1/_source?_source_includes=name

func QueryUserByDocumentID(ctx context.Context) {

	client := ctx.Value(ClientKey).(*elasticsearch.Client)

	documentID := "1"
	
	response, err := client.Get("users", documentID)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	defer func() {
		if response.Body != nil {
			_ = response.Body.Close()
		}
	}()

	if response.IsError(){
		log.Fatalf("Document with ID %s not found", documentID)
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
			log.Fatalf("Error reading response: %s", err)
	}
	//fmt.Printf("Response: %s\n", responseBytes)

	type GetUser struct {
		Index string `json:"_index"`
		Type string `json:"_type"`
		ID          string `json:"_id"`
		Version int  `json:"_version"`
		Seq int `json:"_seq_no"`
		Term int `json:"_primary_term"`
		Found bool `json:"found"`
		Source *presenter.GetUser `json:"_source"`
		CreatedAt time.Time `json:"_created_at"`
		UpdatedAt time.Time `json:"_updated_at"`
}

var getUser GetUser
	err = json.Unmarshal(responseBytes, &getUser)
	if err != nil {
    log.Fatalf("Error decoding response: %s", err)
}

/* 	err = json.NewDecoder(response.Body).Decode(&getUser)
	//fmt.Printf("GetResponse: %s", getResponse)
	if err != nil {
		log.Fatalf("Error to decode: %s", err)
	} */
	
	userName := getUser.Source.Name
	fmt.Printf("✅ User with the ID %s: %s \n", documentID, userName)
}


//Response: {"_index":"users","_type":"_doc","_id":"1","_version":17,"_seq_no":146,"_primary_term":3,"found":true,"_source":{"id":"4aaa155a-97e2-4546-95dc-ee8c131226a1","name":"Joana Vidon","cpf":"112..33z7-99","email":"joanavilczlzdon@gmail.com","phone_number":"(21)98108-8057"}}