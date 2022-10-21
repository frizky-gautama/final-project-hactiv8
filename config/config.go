package config

import (
	"MyGram/model"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GetEnv() (*string, *string, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("ENV ERROR")
		return nil, nil, err
	}
	port_app := os.Getenv("PORT_APP")
	host := os.Getenv("HOST")
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	port := os.Getenv("PORT")
	db := os.Getenv("DB")
	config := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=true&autocommit=true", username, password, host, port, db)
	// dsn := "philo:password@tcp(10.5.99.76:3306)/digital_form?multiStatements=true&autocommit=true"
	return &config, &port_app, nil
}

func GetConnection() (*gorm.DB, error) {
	// GET ENV
	dsn, _, errCon := GetEnv()
	if errCon != nil {
		log.Fatal("Env is Not Connected")
	}

	// Connection DB
	db, err := gorm.Open(mysql.Open(*dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database not Connected")
	}

	fmt.Println("Connect to Database .....")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	dbConfig, _ := db.DB()
	defer cancel()
	err = dbConfig.PingContext(ctx)
	if err != nil {
		return db, err
	}

	// connection pool
	dbConfig.SetMaxOpenConns(20)
	dbConfig.SetMaxIdleConns(10)
	dbConfig.SetConnMaxIdleTime(10 * time.Second)
	dbConfig.SetConnMaxLifetime(20 * time.Second)

	DB = db

	return db, nil
}

func StartDB() {
	db, err := GetConnection()
	if err != nil {
		log.Fatal("Database not Connected", err)
		return
	}
	db.Debug().AutoMigrate(model.User{}, model.Photo{}, model.Comment{}, model.SocialMedia{})
}
