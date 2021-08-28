package main

import (
	"log"

	"github.com/maiaaraujo5/consul-golang/app/fx"
	"github.com/maiaaraujo5/gostart/application"
)

func main() {
	providers := application.Start.
		WithCustomProvider(fx.Module()).
		Build()

	err := application.Run(providers)
	if err != nil {
		log.Fatal(err)
	}
}
