package config

import (
	"fmt"
	"os"
	"test-case-vhiweb/internal/logger"
	"test-case-vhiweb/internal/models"

	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

func InitDB() *gorm.DB {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	user := os.Getenv("NAME")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")

	defaultConnStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=postgres sslmode=disable", host, port, user, password)
	defaultDB, err := sql.Open("pgx", defaultConnStr)
	if err != nil {
		logger.Log.Fatal(err)
	}
	defer defaultDB.Close()

	var exists bool
	err = defaultDB.QueryRow("SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = $1)", dbname).Scan(&exists)
	if err != nil {
		logger.Log.Fatal(err)
	}
	if !exists {
		_, err = defaultDB.Exec(fmt.Sprintf("CREATE DATABASE %s", dbname))
		if err != nil {
			logger.Log.Fatalf("Failed to create database: %v", err)
		}
		logger.Log.Printf("Database %s created", dbname)
	}

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Info),
	})
	if err != nil {
		logger.Log.Fatal(err)
	}

	err = DropAllTables(db)
	if err != nil {
		logger.Log.Fatal(err)
	}

	err = CreateAllTables(db)
	if err != nil {
		logger.Log.Fatal(err)
	}

	return db
}

func DropAllTables(db *gorm.DB) error {
	return db.Migrator().DropTable(
		&models.Product{},
		&models.Vendor{},
		&models.User{},
	)
}

func CreateAllTables(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.Product{},
		&models.Vendor{},
		&models.User{},
	)
}
