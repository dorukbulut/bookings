package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/tsawler/bookings-app/internal/config"
	"testing"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
		//do nothing ; test passed
	default:
		t.Error(fmt.Sprintf("Expected type *chi.Mux, but got %T", v))

	}
}
