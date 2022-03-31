package clickhouse

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/ClickHouse/clickhouse-go"
	"github.com/opentracing/opentracing-go"
)

// NewClickhouseClient ...
func NewClickhouseClient(ctx context.Context, dataSourceName string) {
	span, _ := opentracing.StartSpanFromContext(ctx, "NewClickhouseClient")
	defer span.Finish()

	connect, err := sql.Open("clickhouse", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	if err := connect.Ping(); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			fmt.Printf("[%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		} else {
			fmt.Println(err)
		}
		return
	}
}
