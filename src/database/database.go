package database

import (
	"app/src/config"
	"app/src/database/seeders"
	"app/src/model"
	"app/src/utils"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect(dbHost, dbName string) *gorm.DB {
	// hihihi maap
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		dbHost, config.DBUser, config.DBPassword, dbName, config.DBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Info),
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		TranslateError:         true,
	})
	if err != nil {
		utils.Log.Errorf("Failed to connect to database: %+v", err)
	}

	if err := db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";").Error; err != nil {
		utils.Log.Errorf("Failed to enable uuid-ossp extension: %+v", err)
	}

	// AutoMigrate untuk membuat tabel berdasarkan model
	err = db.AutoMigrate(
		&model.User{},
		&model.Token{},
		&model.Article{},
		&model.ArticleCategory{},
		&model.MealHistory{},
		&model.MealHistoryDetail{},
		&model.ProductToken{},
		&model.Recipe{},
		&model.UsersStar{},
		&model.UsersWeightHeightHistory{},
		&model.UsersWeightHeightTarget{},
	)
	if err != nil {
		utils.Log.Errorf("Failed to auto-migrate database: %+v", err)
	}

	seeders.RunSeeder(db)

	sqlDB, errDB := db.DB()
	if errDB != nil {
		utils.Log.Errorf("Failed to connect to database: %+v", errDB)
	}

	// Config connection pooling
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(60 * time.Minute)

	return db
}
