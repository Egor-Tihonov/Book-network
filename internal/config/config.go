package config

type Config struct {
	PostgresDBURL string `env:"POSTGRES_DB_URL"`
	CookieName    string `env:"COOKIE_NAME"`
	CookieMaxAge  int    `env:"COOKIE_MAX_AGE"`
	CookiePath    string `env:"COOKIE_PATH"`
	JWTKey        string `env:"JWT_KEY"`
}
