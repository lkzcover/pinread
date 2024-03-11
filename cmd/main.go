package main

import (
	"context"
	"github.com/lkzcover/pinread/internal/pkg/db"
	"github.com/lkzcover/pinread/internal/pkg/env"
	"github.com/lkzcover/pinread/internal/pkg/tg"
	"log"
	"sync"
)

func main() {
	err := env.Load()
	if err != nil {
		log.Fatalf("load env error: %s", err)
	}

	ctx := context.Background()

	dbConn, err := db.NewDBConn(ctx)
	if err != nil {
		log.Fatalf("startup db connection error: %s", err)
	}
	defer dbConn.Close()

	wg := sync.WaitGroup{}

	wg.Add(1)

	tg.StartListener(dbConn)

	wg.Wait()
}
