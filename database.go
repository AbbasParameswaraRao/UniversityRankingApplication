package main

import (
	"encoding/json"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func ConnectDB(config *Config) (*gorm.DB, error) {
	// a connection to the MySQL database using the configuration values from `config`
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		config.Database.User,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.Name))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}

	return db, nil
}

type University struct {
	Rank uint16 `json:"ranking" gorm:"primary_key"`
	Name string `json:"title" gorm:"type:varchar(100);column:name"`
	City string `json:"location" gorm:"type:varchar(100);column:city"`
}

var DB *gorm.DB
var institutions []University

func ReadFile(path string) {

	file, _ := os.ReadFile(path)
	json.Unmarshal([]byte(file), &institutions)
}

func PushToDB(config *Config) {
	db, err := ConnectDB(config)
	if err != nil {
		LogError(err)
		return
	}

	// Creates the `universities` database if it doesn't exist
	err = db.Exec("CREATE DATABASE IF NOT EXISTS universities").Error
	if( err != nil){
		LogError(err)
		return
	}


	// Switches to the `universities` database
	db.Exec("USE universities")
	// Auto migrates the `University` struct to create the necessary table in the database
	db.AutoMigrate(&University{})
	// Inserts or updates each university record in the `institutions` slice into the database
	for _, value := range institutions {
		_ = db.Where(University{Rank: value.Rank}).Assign(University{Rank: value.Rank, Name: value.Name, City: value.City}).FirstOrCreate(&University{}).Error
	}
	DB = db
}

func ReadAndLoad() {
	config, err := LoadConfig()
	if err != nil {
		LogError(err)
		return
	}
	ReadFile("universities_ranking.json")
	PushToDB(config)
}
