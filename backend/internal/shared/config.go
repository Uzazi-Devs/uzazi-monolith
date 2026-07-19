package shared

import "os"

// Config is loaded from the environment. Defaults target local `docker compose`.
type Config struct {
	Port        string
	DatabaseURL string
	AuthJWKSURL string
	AIProvider  string
	RedisURL    string
}

func Load() Config {
	return Config{
		Port:        env("PORT", "8080"),
		DatabaseURL: env("DATABASE_URL", "postgres://uzazi:uzazi@localhost:5432/uzazi?sslmode=disable"),
		AuthJWKSURL: env("AUTH_JWKS_URL", "http://localhost:3000/api/auth/jwks"),
		AIProvider:  env("AI_PROVIDER", "gemma"),
		RedisURL:    env("REDIS_URL", "redis://localhost:6379"),
	}
}

func env(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
