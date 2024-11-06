package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ServiceName string `envconfig:"SERVICE_NAME" default:"rest-api"`
	Environment string `envconfig:"ENVIRONMENT" default:"dev"`
	Maintenance bool   `envconfig:"MAINTENANCE" default:"false"`
	RestPort    int    `envconfig:"REST_PORT" default:"38000" required:"true"` // 28003
	GrpcPort    int    `envconfig:"GRPC_PORT" default:"39000" required:"true"` // 29003

	DBHost         string `envconfig:"DB_HOST" default:"localhost"`
	DBPort         int    `envconfig:"DB_PORT" default:"3306"`
	DBUserName     string `envconfig:"DB_USERNAME" default:"root"`
	DBPassword     string `envconfig:"DB_PASSWORD" default:"password"`
	DBDatabaseName string `envconfig:"DB_DBNAME" default:"postgres"`
	DBLogEnable    bool   `envconfig:"DB_LOG_ENABLE" default:"true"`
	DBLogLevel     int    `envconfig:"DB_LOG_LEVEL" default:"3"`
	DBLogThreshold int    `envconfig:"DB_LOG_THRESHOLD" default:"1"`

	// Usertoken Related
	AuthTokenSecret    string `envconfig:"AUTH_TOKEN_SECRET" default:"auth_token_secret"`
	AuthTokenExpiry    int    `envconfig:"AUTH_TOKEN_EXPIRY" default:"3600"`
	RefreshTokenSecret string `envconfig:"REFRESH_TOKEN_SECRET" default:"refresh_token_secret"`
	RefreshTokenExpiry int    `envconfig:"REFRESH_TOKEN_EXPIRY" default:"86400"`

	// Auth & User Related
	PasswordSalt     string `envconfig:"PASSWORD_SALT" default:"password_salt"`
	SuperPassword    string `envconfig:"SUPER_PASSWORD" default:"password1234567890"`
	ValidateTokenURL string `envconfig:"VALIDATE_TOKEN_URL" default:"http://localhost:38000/validate"`

	FlatInterestRate float64 `envconfig:"FLAT_INTEREST_RATE" default:"10"`
	RunMigration     bool    `envconfig:"RUN_MIGRATION" default:"true"`
}

func New() Config {
	_ = godotenv.Overload()
	cfg := Config{}
	envconfig.MustProcess("", &cfg)
	return cfg
}
