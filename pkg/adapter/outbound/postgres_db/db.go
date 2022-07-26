package postgres_db

import (
	"context"
	"github.com/jackc/pgconn"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func Init() (*pgconn.PgConn, error) {
	var err error
	zapLogger, err := zap.NewProduction()
	defer zapLogger.Sync()
	logger := zapLogger.Sugar()
	connConf := &pgconn.Config{
		Host:           viper.GetString("POSTGRES_HOST"),
		Port:           uint16(viper.GetUint("POSTGRES_PORT")),
		Database:       viper.GetString("POSTGRES_DB"),
		User:           viper.GetString("POSTGRES_USER"),
		Password:       viper.GetString("POSTGRES_PASSWORD"),
		ConnectTimeout: 15,
	}
	maxRetries := 10
	numRetry := 0
	var pgConn *pgconn.PgConn
	for numRetry < maxRetries {
		numRetry++
		logger.Infof("Connecting to Postgres DB at host %s", connConf.Host)
		pgConn, err = pgconn.ConnectConfig(context.Background(), connConf)
		if err != nil {
			logger.Errorf("Connected to Postgres DB at host %s failed", connConf.Host)
		}
		logger.Infof("Connected to Postgres DB at host %s successfully", connConf.Host)
		break
	}
	return pgConn, err
}
