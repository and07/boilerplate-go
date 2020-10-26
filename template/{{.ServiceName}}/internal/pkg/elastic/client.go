package elastic

import (
	"context"
	"fmt"
	"log"

	"github.com/olivere/elastic"
	"github.com/opentracing/opentracing-go"
)

// NewElasticClient ...
func NewElasticClient(ctx context.Context, url string) {
	span, _ := opentracing.StartSpanFromContext(ctx, "NewElasticClient")
	defer span.Finish()
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		log.Fatal(err)
	}

	// Ping the Elasticsearch server to get e.g. the version number
	info, code, err := client.Ping(url).Do(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)
}
