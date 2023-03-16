package elasticSearch

import (
	"github.com/elastic/go-elasticsearch/v8"
)


type ElasticSeach struct{
	Client *elasticsearch.Client
}

type contextKey struct {
	Key int
}

var UserKey contextKey = contextKey{Key: 1}
var ClientKey contextKey = contextKey{Key: 2}

func NewElasticSearch (Client *elasticsearch.Client,) *ElasticSeach{
	return &ElasticSeach{
		Client:   Client,
	}
}

type MovieRaw struct {
	Title string   `json:"title"`
	Year  int      `json:"year"`
}



/* func Elastic() (context.Context){
	cert, _ := ioutil.ReadFile(*cacert)
	newClient, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{
			"https://localhost:9200",
		},
			Username: "adm",
  		Password: "Pass123!",
			CACert: cert, //custom certificate authority
	})
	if err != nil{
		log.Fatalf("Error creating the client: %s", err)
	}

	es.New
	res, err := es.Info()
	if err != nil{
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {// Check response status
    log.Fatalf("Error: %s", res.String())
  }

	return context.WithValue(ctx, domain.ClientKey, newClient)
 */
/* 
	var (
    r  map[string]interface{}
    wg sync.WaitGroup
  )
	// Deserialize the response into a map.
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
	log.Fatalf("Error parsing the response body: %s", err)
  }
  // Print client and server version numbers.
  log.Printf("Client: %s", elasticsearch.Version)
  log.Printf("Server: %s", r["version"].(map[string]interface{})["number"])
  log.Println(strings.Repeat("~", 37)) */

/* }

	res, err = es.Index(
		"Store",
		strings.NewReader(`{"title" : "User"}`),
		es.Index.WithDocumentID(idUser),
		es.Index.WithRefresh("true"),
	)
		if err != nil{
		log.Fatalf("ERROR: %s", err)
	}
} */
