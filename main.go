package main

import (
	"fmt"
	"os"

	"github.com/temp-lib/templater"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("usage : app serviceName user")
		return
	}
	cfg := templater.Config{
		ServiceName:    os.Args[1],
		User:           os.Args[2],
		IgnorePatterns: []string{".html", ".yml", ".yaml", ".tpl", ".txt", "js.map"},
	}
	templtr := templater.New(cfg)
	err := templtr.BuildService(`./template/{{.ServiceName}}`,
		`./`)
	if err != nil {
		fmt.Println(err)
	}
}
