package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/Uzazi-Devs/uzazi-monolith/backend/internal/authjwt"
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

	jwksURL := os.Getenv("AUTH_JWKS_URL")
	issuer := strings.TrimSuffix(jwksURL, "/api/auth/jwks")
	verifier, err := authjwt.NewVerifier(ctx, jwksURL, issuer)
	if err != nil {
		log.Fatalf("jwks: %v", err)
	}
	requireAdmin := waitlist.RequireAdmin(verifier)

	waitlistHandler := &waitlist.Handler{Queries: db.New(pool)}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", health)
	mux.HandleFunc("POST /waitlist", waitlist.WithCORS(waitlistHandler.Create))
	mux.HandleFunc("OPTIONS /waitlist", waitlist.WithCORS(waitlistHandler.Create))
	mux.HandleFunc("GET /admin/waitlist", requireAdmin(waitlistHandler.List))
	mux.HandleFunc("POST /admin/waitlist/{id}/accept", requireAdmin(waitlistHandler.Accept))

	log.Fatal(http.ListenAndServe(":"+port, mux))
}
