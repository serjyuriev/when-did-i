package controller

import (
	"context"
	"log"
	"net/http"
)

type Controller interface {
	GetHello(w http.ResponseWriter, r *http.Request)
}

func NewController(ctx context.Context) (Controller, error) {
	log.Println("Creating controller...")

	log.Println("Controller was created successfully")
	return &controller{}, nil
}

type controller struct{}

func (c *controller) GetHello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Hello!\n"))
}
