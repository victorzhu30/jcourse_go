package dal

import (
	"fmt"
	"log"
	"os"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"jcourse_go/util"
)

var dbClient *gorm.DB

func GetDBClient() *gorm.DB {
	return dbClient
}

func initPostgresql() error {
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DBNAME")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", host, user, password, dbname, port)
	var err error
	dbClient, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return err
}

func InitDBClient() {
	err := initPostgresql()
	if err != nil {
		panic(err)
	}
}

func InitMockDBClient() (sqlmock.Sqlmock, error) {
	var err error
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, err
	}
	config := &gorm.Config{}
	if util.IsDebug() {
		newLogger := logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				LogLevel: logger.Silent, // Log level
			},
		)
		config.Logger = newLogger
	}
	dbClient, err = gorm.Open(postgres.New(postgres.Config{Conn: db}), config)
	if err != nil {
		return nil, err
	}
	return mock, nil
}
