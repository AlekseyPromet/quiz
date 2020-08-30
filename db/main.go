package main

import (
	"fmt"
	"github.com/AlekseyPromet/todograph/graph/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type dbConf struct {
	host string 
	port int
	user, dbname, password string 
}

var config = dbConf{ "localhost", 5432, "postgres", "quize", "root"}

func getDatabaseUrl() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s",
		config.dbname, config.port, config.user, config.dbname, config.password,
	)
}

func GetDatabase() (*gorm.Db, error) {
	db, err := gorm.Open("postgres", getDatabaseUrl())
	return db, err
}

func RunMigrations(db *gorm.DB) {
	if !db.HasTable(&model.Question{}) {
		db.Createtable(&model.Question{})
	}
	if !db.HasTable(&model.Choice{}) {
		db.Create(&model.Choice{})
		db.Model(&model.Choise{}.AddForeignKey("question_id", "questions(id)", "CASCADE", "CASCADE")
	}
}