package whendidi

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/serjyuriev/when-did-i/internal/app/controller"
)

type Server struct {
	c controller.Controller
}

func NewServer(ctx context.Context) (*Server, error) {
	log.Println("Creating WhenDidI server...")

	ctrl, err := controller.NewController(ctx)
	if err != nil {
		log.Println("Unable to initialize controller")
		return nil, err
	}

	log.Println("WhenDidI server was created successfully")
	return &Server{c: ctrl}, nil
}

func (s *Server) Run() error {
	log.Println("Preparing to run WhenDidI server...")

	r := chi.NewRouter()

	r.Get("/hello", s.c.GetHello)

	log.Println("WhenDidI server is listening on \":8080\"")
	return http.ListenAndServe(":8080", r)
}
