package db

import (
	"log"
	"os"
	"time"

	customLogger "backend/pkg/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Database *gorm.DB

func InitDatabase() {
	dsn := os.Getenv("DATABASE_CONNECTION_STRING")

	if dsn == "" {
		customLogger.Error("Please Set The Database Connection String In ENV file")
		os.Exit(1)
	}
	gormLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,       // Don't include params in the SQL log
			Colorful:                  true,        // Disable color
		},
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: gormLogger})
	if err != nil {
		customLogger.Trace(err)
		os.Exit(1)
	}
	db.Logger.LogMode(logger.Info)
	Database = db
	customLogger.Success("DATABASE INIT SUCCESSFULLY")
}
