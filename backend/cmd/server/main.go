package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/Uzazi-Devs/uzazi-monolith/backend/internal/db"
	"github.com/Uzazi-Devs/uzazi-monolith/backend/internal/waitlist"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	ctx := context.Background()
	pool, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("db connect: %v", err)
	}
	defer pool.Close()

	waitlistHandler := &waitlist.Handler{Queries: db.New(pool)}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", health)
	mux.HandleFunc("POST /waitlist", waitlistHandler.Create)
	mux.HandleFunc("GET /admin/waitlist", waitlist.RequireAdmin(waitlistHandler.List))
	mux.HandleFunc("POST /admin/waitlist/{id}/accept", waitlist.RequireAdmin(waitlistHandler.Accept))

	log.Fatal(http.ListenAndServe(":"+port, mux))
}
