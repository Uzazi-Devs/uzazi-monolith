package main

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/Uzazi-Devs/uzazi-monolith/backend/internal/ai"
	"github.com/Uzazi-Devs/uzazi-monolith/backend/internal/auth"
	"github.com/Uzazi-Devs/uzazi-monolith/backend/internal/community"
	"github.com/Uzazi-Devs/uzazi-monolith/backend/internal/db"
	"github.com/Uzazi-Devs/uzazi-monolith/backend/internal/health"
	"github.com/Uzazi-Devs/uzazi-monolith/backend/internal/media"
	"github.com/Uzazi-Devs/uzazi-monolith/backend/internal/shared"
)

func main() {
	log := shared.NewLogger()
	cfg := shared.Load()
	ctx := context.Background()

	pool, err := shared.NewPool(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Error("db connect", "err", err)
		os.Exit(1)
	}
	defer pool.Close()

	queries := db.New(pool)

	// Modules are wired together as Go interfaces — no inter-module HTTP.
	var (
		healthSvc    health.Service       = health.NewService(queries)
		communitySvc community.Service    = community.NewService(queries)
		mediaSvc     media.Service        = media.NewService(queries)
		aiProvider   ai.InferenceProvider = ai.NewProvider(cfg.AIProvider)
	)
	_, _, _ = communitySvc, mediaSvc, aiProvider // wired here, exposed as needed

	// Auth verifier needs the auth-service JWKS. If it isn't up yet, keep
	// running so /healthz works — protected routes just won't mount.
	verifier, err := auth.NewVerifier(ctx, cfg.AuthJWKSURL, queries)
	if err != nil {
		log.Warn("auth verifier not ready (auth-service down?)", "err", err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})

	// Example protected route: verifies a BetterAuth JWT, then uses the
	// health module through its interface.
	if verifier != nil {
		mux.Handle("GET /api/health-records", verifier.Middleware(http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				user, _ := auth.UserFromContext(r.Context())
				recs, err := healthSvc.ListForUser(r.Context(), user.ID)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				w.Header().Set("Content-Type", "application/json")
				_ = json.NewEncoder(w).Encode(recs)
			})))
	}

	srv := &http.Server{
		Addr:              ":" + cfg.Port,
		Handler:           shared.Logging(log, shared.CORS(mux)),
		ReadHeaderTimeout: 5 * time.Second,
	}
	log.Info("backend listening", "addr", srv.Addr, "ai_provider", cfg.AIProvider)
	if err := srv.ListenAndServe(); err != nil {
		log.Error("server", "err", err)
		os.Exit(1)
	}
}
