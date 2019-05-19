package main

import (
	"context"
	"log"
	"net/http"
)

func hiHandler(ctx context.Context) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		//span, _ := opentracing.StartSpanFromContext(ctx, "Scratch.hiHandler")
		//defer span.Finish()
		counter.Inc()
		if _, err := w.Write([]byte("hi")); err != nil {
			log.Println(err)
		}
	}
}

func publicHandle(ctx context.Context) *http.ServeMux {
	rPublic := http.NewServeMux()
	rPublic.HandleFunc("/", hiHandler(ctx))
	return rPublic
}
