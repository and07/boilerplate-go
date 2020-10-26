package main

import (
	"fmt"

	"github.com/temp-lib/templater"
)

func main() {
	cfg := templater.Config{
		ServiceName:    "TEST",
		User:           "shine",
		IgnorePatterns: []string{".html", ".yml", ".yaml", ".tpl", ".txt", "js.map"},
	}
	templtr := templater.New(cfg)
	err := templtr.BuildService(`./template/{{.ServiceName}}`,
		`./`)
	if err != nil {
		fmt.Println(err)
	}
}
