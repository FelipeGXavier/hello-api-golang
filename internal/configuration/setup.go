package configuration

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type Env struct {
	Db *pgxpool.Pool
	Logger zap.Logger
}

func LoadEnvObject(config Config) *Env {
	result := Env{};
	logger, _ := zap.NewProduction();
	defer logger.Sync();
	conn, err := pgxpool.New(context.Background(), getDatabaseConnectionUrl(config));
	if (err != nil) {
		logger.Error("Error while connection to database", zap.Error(err));
		panic("Error while connection to database");
	}
	result.Db = conn;
	result.Logger = *logger;
	return &result;
}

func getDatabaseConnectionUrl(config Config) string {
	urlFormat := "postgresql://%s/%s?user=%s&password=%s";
	return fmt.Sprintf(urlFormat, config.DBHost, config.DBName, config.DBUser, config.DBPassword);
}