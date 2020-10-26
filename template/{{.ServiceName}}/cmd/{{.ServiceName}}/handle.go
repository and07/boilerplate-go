package main

import (
	"context"
	"net/http"

	"github.com/{{.User}}/{{.ServiceName}}/internal/pkg/template"
)

func hiHandler(ctx context.Context, tpl *template.Template, tmplData interface{}) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		//span, _ := opentracing.StartSpanFromContext(ctx, "Scratch.hiHandler")
		//defer span.Finish()
		counter.Inc()

		tpl.RenderTemplate(w, "main.html", tmplData)
	}
}

func publicHandle(ctx context.Context, tpl *template.Template, tmplData interface{}) *http.ServeMux {
	rPublic := http.NewServeMux()
	rPublic.HandleFunc("/", hiHandler(ctx, tpl, tmplData))
	return rPublic
}
