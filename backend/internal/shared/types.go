package shared

// Config is loaded from the environment. Defaults target local `docker compose`.
type Config struct {
	Port        string
	DatabaseURL string
	AuthJWKSURL string
	AIProvider  string
	RedisURL    string
}
