package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

func main() {
	// connect to elasticsearch server
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	// The Elasticsearch server information
	res_, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res_.Body.Close()

	fmt.Println(res_)

	// Index a document
	// doc := `{"title": "Go and Elasticsearch", "content": "A tutorial on how to use Go and Elasticsearch together"}`
	// req := esapi.IndexRequest{
	// 	Index:      "test-index",
	// 	DocumentID: "1",
	// 	Body:       strings.NewReader(doc),
	// 	Refresh:    "true",
	// }

	// res, err := req.Do(context.Background(), es.Transport)
	// if err != nil {
	// 	log.Fatalf("Error indexing document: %s", err)
	// }
	// defer res.Body.Close()

	// fmt.Println(res)

	// Searching data in Elasticsearch
	query := `{"query": {"match": {"title": "Go"}}}`
	req := esapi.SearchRequest{
		Index: []string{"articles"},
		Body:  strings.NewReader(query),
	}

	res, err := req.Do(context.Background(), es.Transport)
	if err != nil {
		log.Fatalf("Error searching documents: %s", err)
	}
	defer res.Body.Close()

	fmt.Println(res)
}
