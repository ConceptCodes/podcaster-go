package postgres

import (
	"fmt"

	"github.com/rs/zerolog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"podcaster/config"
	"podcaster/internal/constants"
)

var db *gorm.DB
var err error

func New(l zerolog.Logger) (*gorm.DB, error) {
	l.Debug().Msg("Connecting to PostgreSQl")

	db, err = gorm.Open(postgres.New(postgres.Config{
		DSN: fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
			config.AppConfig.DbHost,
			config.AppConfig.DbUser,
			config.AppConfig.DbPass,
			config.AppConfig.DbName,
			config.AppConfig.DbPort,
		),
	}), &gorm.Config{
		Logger: logger.New(
			&l,
			logger.Config{
				LogLevel:             logger.Info,
				Colorful:             config.AppConfig.Env == constants.LocalEnv,
				ParameterizedQueries: true,
			},
		),
	})

	return db, err
}
