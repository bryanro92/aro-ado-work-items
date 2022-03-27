package main

import (
	"context"
	"log"

	"github.com/bryanro92/aro-ado-work-items/pkg/work"
)

func main() {
	ctx := context.Background()
	if err := work.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
