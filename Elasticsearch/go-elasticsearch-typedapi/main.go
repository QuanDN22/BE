package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/some"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

func main() {
	// connect to elasticsearch
	// es, err := elasticsearch.NewDefaultClient()
	es, err := elasticsearch.NewTypedClient(elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
		Logger:    &elastictransport.ColorLogger{Output: os.Stdout, EnableRequestBody: true, EnableResponseBody: true},
	})

	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	res, err := es.Info().Do(context.Background())
	if err != nil {
		log.Fatalf("error reading Info request: %s", err)
	}

	if res.Tagline != "You Know, for Search" {
		log.Fatalf("invalid tagline, got: %s", res.Tagline)
	}

	// create an index
	// create an index named test-index
	// and provide a mapping for the field price which will be an integer
	indexName := "test-index"
	// If the index doesn't exist we create it with a mapping.
	if exists, err := es.Indices.Exists(indexName).IsSuccess(context.Background()); !exists && err == nil {
		res, err := es.Indices.Create(indexName).
			Mappings(&types.TypeMapping{
				Properties: map[string]types.Property{
					"price": types.IntegerNumberProperty{},
					"name":  types.KeywordProperty{},
				},
			}).
			Do(context.Background())

		if err != nil {
			log.Fatalf("error creating index test-index: %s", err)
		}

		if !res.Acknowledged && res.Index != indexName {
			log.Fatalf("unexpected error during index creation, got : %#v", res)
		}
	} else if err != nil {
		log.Fatal(err)
	}

	// Indexing a document
	// The standard way of indexing a document is to provide a struct to the Request method,
	// the standard json/encoder will be run on your structure
	// and the result will be sent to Elasticsearch.
	type Document struct {
		Id    int    `json:"id"`
		Name  string `json:"name"`
		Price int    `json:"price"`
		Alt   string `json:"alt"`
	}

	// // Indexing synchronously with refresh.Waitfor, one document at a time
	// for _, document := range []Document{
	// 	{
	// 		Id:    1,
	// 		Name:  "Foo",
	// 		Price: 10,
	// 	},
	// 	{
	// 		Id:    2,
	// 		Name:  "Bar",
	// 		Price: 12,
	// 	},
	// 	{
	// 		Id:    3,
	// 		Name:  "Baz",
	// 		Price: 4,
	// 	},
	// } {
	// 	indexed, err := es.Index(indexName).
	// 		Document(document).
	// 		Id(strconv.Itoa(document.Id)).
	// 		Refresh(refresh.Waitfor).
	// 		Do(context.Background())
	// 	if err != nil {
	// 		log.Fatalf("error indexing document: %s", err)
	// 	}
	// 	if indexed.Result != result.Created {
	// 		log.Fatalf("unexpected result during indexation of document: %v, response: %v", document, indexed)
	// 	}

	// 	document.Alt = fmt.Sprintf("alt_%d", document.Id)
	// 	es.Update(indexName, strconv.Itoa(document.Id)).Doc(document).Do(context.Background())
	// }

	// Retrieving a document
	// Retrieving a document follows the API as part of the argument of the endpoint
	// Check for document existence in index
	if ok, err := es.Get(indexName, "1").IsSuccess(context.Background()); !ok {
		log.Fatalf("could not retrieve document: %s", err)
	} else {
		fmt.Println(ok)
	}

	// Try to retrieve a faulty index name
	if ok, _ := es.Get("non-existent-index", "9999").IsSuccess(context.Background()); ok {
		log.Fatalf("index shouldn't exist")
	}

	// Same faulty index name with error handling
	_, err = es.Get("non-existent-index", "9999").Do(context.Background())
	if !errors.As(err, &types.ElasticsearchError{}) && !errors.Is(err, &types.ElasticsearchError{Status: 404}) {
		log.Fatalf("expected ElasticsearchError, got: %v", err)
	}

	// SEARCH
	// Simple search matching name
	// Building a search query can be done with structs or builder.
	// As an example, let’s search for a document with a field name with a value of Foo in the index named index_name
	res_, err := es.Search().
		Index(indexName). // The targeted index for this search
		Query(&types.Query{
			Match: map[string]types.MatchQuery{
				"name": {Query: "Foo"}, // Match query specifies that name should match Foo
			},
		}).Do(context.Background()) // The query is run with a context.Background and returns the response.

	if err != nil {
		log.Fatalf("error runnning search query: %s", err)
	}

	if res_.Hits.Total.Value == 1 {
		doc := Document{}
		err = json.Unmarshal(res_.Hits.Hits[0].Source_, &doc)
		if err != nil {
			log.Fatalf("cannot unmarshal document: %s", err)
		}
		if doc.Name != "Foo" {
			log.Fatalf("unexpected search result")
		}
	} else {
		log.Fatalf("unexpected search result")
	}

	// Aggregations
	// Aggregate prices with a SumAggregation
	// Given documents with a price field, we run a sum aggregation on index_name
	searchResponse, err := es.Search().
		Index(indexName).
		Size(0). // Sets the size to 0 to retrieve only the result of the aggregation.
		Aggregations(map[string]types.Aggregations{
			"total_prices": { // Specifies the field name on which the sum aggregation runs
				Sum: &types.SumAggregation{
					Field: some.String("price"), // The SumAggregation is part of the Aggregations map.
				},
			},
		}).
		Do(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	for name, agg := range searchResponse.Aggregations {
		if name == "total_prices" {
			switch aggregation := agg.(type) {
			case *types.SumAggregate:
				if aggregation.Value != 26. {
					log.Fatalf("error in aggregation, should be 26, got: %f", aggregation.Value)
				}
			default:
				fmt.Printf("unexpected aggregation: %#v\n", agg)
			}
		}
	}

}
