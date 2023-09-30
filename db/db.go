package db

import (
	"fmt"
	"log"
	"os"

	"github.com/zohaibsoomro/go-course-api-postgresql/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func connect(c *models.DbConfig) (*gorm.DB, error) {
	//	dns := "postgres://user:pass@host/dbName?port=portNum&sslmode=verify-full&connect_timeout=20"
	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", c.Host, c.User, c.Password, c.DBName, c.Port, c.SSLEnabled)
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{
		Logger: logger.New(log.New(os.Stdout, fmt.Sprint(log.Lmsgprefix), log.LstdFlags), logger.Config{
			Colorful: true,
		}),
	})
	return db, err
}

func NewDbConnection(c *models.DbConfig) *models.MyDb {
	db, err := connect(c)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	log.Println("Connection established.")
	//if table is already created then no need to do it
	db.AutoMigrate(&models.Course{})
	return &models.MyDb{
		DB: db,
	}
}
