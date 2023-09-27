package main

import (
	"context"
	"log"
	"time"

	whendidi "github.com/serjyuriev/when-did-i/internal/app/server"
)

func main() {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()

	srv, err := whendidi.NewServer(ctx)
	if err != nil {
		log.Println("Unable to initialize server, aborting")
		return
	}

	log.Fatal(srv.Run())
}
