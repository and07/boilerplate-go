package main

import (
	"context"
	"net/http"

	"github.com/and07/boilerplate-go/internal/pkg/template"
)

func hiHandler(ctx context.Context, tpl *template.Template) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		//span, _ := opentracing.StartSpanFromContext(ctx, "Scratch.hiHandler")
		//defer span.Finish()
		counter.Inc()

		tpl.RenderTemplate(w, "main.html", "Hi")
	}
}

func publicHandle(ctx context.Context, tpl *template.Template) *http.ServeMux {
	rPublic := http.NewServeMux()
	rPublic.HandleFunc("/", hiHandler(ctx, tpl))
	return rPublic
}
