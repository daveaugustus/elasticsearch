package main

import (
	"context"
	"elasticsearch/models"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/olivere/elastic/v7"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

type Post struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func main() {
	users := models.CreateRandomUsers(500)

	for _, v := range users {
		fmt.Printf("%v\n", v)
	}
}

// func randate() time.Time {
// 	min := time.Date(1970, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
// 	max := time.Date(2070, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
// 	delta := max - min

// 	sec := rand.Int63n(delta) + min
// 	return time.Unix(sec, 0)
// }

// func main() {
// 	// Instantiate a client instance of the elastic library
// 	client, err := elastic.NewClient(
// 		elastic.SetSniff(true),
// 		elastic.SetURL("http://localhost:9200"),
// 		elastic.SetHealthcheckInterval(5*time.Second),
// 	)
// 	if err != nil {
// 		fmt.Printf("cannot create client: %s.\n", err.Error())
// 	}

// 	// Declare a Context object for the API calls
// 	ctx := context.Background()
// 	// CreateIndex(ctx, *client)
// 	// posts := GetPosts()
// 	// for _, post := range posts {
// 	// 	if err = SaveDocument(ctx, *client, post); err != nil {
// 	// 		log.Printf("cannot create documents: %v", err)
// 	// 	}
// 	// }
// 	// GetDocuments(ctx, *client)
// 	resp, err := client.GetMapping().Index("post").Pretty(true).Do(ctx)
// 	if err != nil {
// 		log.Println("err: ", err)
// 	}
// 	for i := range resp {
// 		nestedMap, ok := resp[i].(map[string]interface{})
// 		if ok {
// 			log.Printf("Index: %v, Value: %+v\n", i, nestedMap)
// 		}
// 	}
// 	// log.Printf("resp: %+v", resp["post"])
// 	// log.Printf("resp: %T", resp["post"])
// 	// log.Println("resp", resp["map"].(map[string]interface{})["Author"])
// }

// CreateIndex
func CreateIndex(ctx context.Context, client elastic.Client) {
	// // Create a new index and pass the mappings string to the body
	create, err := client.CreateIndex("post").Do(ctx)
	if err != nil {
		log.Fatalf("CreateIndex() ERROR: %v", err)
	} else {
		fmt.Println("CreateIndex():", create)
	}
}

// IndexExists Find out if the Elasticsearch cluster already has the index name there
func IndexExists(ctx context.Context, client elastic.Client) {
	exist, err := client.IndexExists("index").Do(ctx)
	if err != nil {
		log.Fatalf("IndexExists() ERROR: %v", err)
	}
	log.Println("Index exists: ", exist)
}

func SaveDocument(ctx context.Context, client elastic.Client, post Post) error {
	_, err := client.Index().Index("post").Type("posts").
		BodyJson(post).Refresh("true").Do(ctx)

	if err != nil {
		fmt.Printf("Error save document to elastic search: %s. Save log to file\n", err.Error())
	}

	return err
}

func GetPosts() []Post {
	return []Post{
		{
			Title:   "The Black Hat Rust",
			Content: "Well the content is unavailable right now",
			Author:  "Dave A",
		},
		{
			Title: "gRPC: Up and Running",
			Content: `Nowadays software applications are often connected with each other over
			computer networks using inter-process communication technologies.`,
			Author: "Kasun Indrasiri and Danesh Kuruppu",
		},
		{
			Title: "Cloud Native Go",
			Content: `Over the past few years, two infrastructure trends have been happening: Go has been
			increasingly used for infrastructure, in addition to backend; and the infrastructure is
			moving to the cloud.`,
			Author: "Matthew A. Titmus",
		},
		{
			Title: "The TCP/IP Guide",
			Content: `Chances are pretty good that even before you started reading this Guide, you had heard of
			“TCP/IP”—even if you didn't know exactly what it was or even what all those letters stood
			for.`,
			Author: "M. Kozierok",
		},
	}
}

func TotalNoOfDocs(ctx context.Context, client elastic.Client) {
	res, err := client.Search("post").TrackTotalHits(true).Do(ctx)
	if err != nil {
		log.Println(err)
	}
	log.Printf("Response: %+v", res)
}

func GetDocuments(ctx context.Context, client elastic.Client) {
	query := elastic.NewMatchAllQuery()
	searchResult, err := client.Search().
		Index("po*").
		Query(query).
		Pretty(true).
		Do(ctx)
	if err != nil {
		log.Println(err)
	}
	log.Printf("Response: %+v", searchResult.TotalHits())
	log.Printf("Response: %+v", searchResult.Hits.Hits)
}
