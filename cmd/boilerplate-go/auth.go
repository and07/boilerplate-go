package main

import (
	"context"
	"net/http"
	"sort"

	"github.com/and07/boilerplate-go/internal/pkg/template"
)

// ProviderIndex ...
type ProviderIndex struct {
	Providers    []string
	ProvidersMap map[string]string
}

func authHandler(ctx context.Context, tpl *template.Template) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		//span, _ := opentracing.StartSpanFromContext(ctx, "Scratch.hiHandler")
		//defer span.Finish()
		m := make(map[string]string)
		//m["facebook"] = "Facebook"
		m["google"] = "Google"
		//m["apple"] = "Apple"

		var keys []string
		for k := range m {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		providerIndex := &ProviderIndex{Providers: keys, ProvidersMap: m}

		tpl.RenderTemplate(w, "auth.html", providerIndex)

	}
}
