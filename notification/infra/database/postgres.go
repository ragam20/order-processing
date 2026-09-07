package database

import (
	"fmt"

	"ops/notification/infra/database/entity"
	"ops/shared/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgres(env *config.Env) (*gorm.DB, error) {
	sslMode := env.Database.SSLMode
	if sslMode == "" {
		sslMode = "disable"
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		env.Database.Host,
		env.Database.User,
		env.Database.Password,
		env.Database.Name,
		env.Database.Port,
		sslMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("open postgres: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("get postgres connection: %w", err)
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("ping postgres: %w", err)
	}

	return db, nil
}

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(entity.NotificationEntity{}); err != nil {
		return err
	}
	return nil
}
