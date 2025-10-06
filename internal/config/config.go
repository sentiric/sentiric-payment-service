// sentiric-payment-service/internal/config/config.go
package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type Config struct {
	GRPCPort string
	HttpPort string
	CertPath string
	KeyPath  string
	CaPath   string
	LogLevel string
	Env      string

	// Ödeme servisi bağımlılıkları (Placeholder)
	StripeAPIKey   string
	PaymentAdapter string // Stripe, Iyzico, etc.
}

func Load() (*Config, error) {
	godotenv.Load()

	// Harmonik Mimari Portlar (Yatay Yetenek, 171XX bloğu atandı)
	return &Config{
		GRPCPort: GetEnv("PAYMENT_SERVICE_GRPC_PORT", "17111"),
		HttpPort: GetEnv("PAYMENT_SERVICE_HTTP_PORT", "17110"),

		CertPath: GetEnvOrFail("PAYMENT_SERVICE_CERT_PATH"),
		KeyPath:  GetEnvOrFail("PAYMENT_SERVICE_KEY_PATH"),
		CaPath:   GetEnvOrFail("GRPC_TLS_CA_PATH"),
		LogLevel: GetEnv("LOG_LEVEL", "info"),
		Env:      GetEnv("ENV", "production"),

		StripeAPIKey:   GetEnv("STRIPE_API_KEY", ""),
		PaymentAdapter: GetEnv("PAYMENT_ADAPTER", "stripe"),
	}, nil
}

func GetEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func GetEnvOrFail(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatal().Str("variable", key).Msg("Gerekli ortam değişkeni tanımlı değil")
	}
	return value
}
