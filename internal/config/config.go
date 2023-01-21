// Package config ...
package config

// Config for env values
type Config struct {
	PostgresDBURL   string `env:"POSTGRES_DB_URL"`
	CookieTokenName string `env:"COOKIE_TOKEN_NAME"`
	CookieUserName  string `env:"COOKIE_USER_NAME"`
	CookieMaxAge    int    `env:"COOKIE_MAX_AGE"`
	CookiePath      string `env:"COOKIE_PATH"`
	JWTKey          string `env:"JWT_KEY"`
}
