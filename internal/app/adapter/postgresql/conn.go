package postgresql

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

// Connection gets connection of postgresql database
func Connection() (db *gorm.DB) {
	host := viper.Get("DB_HOST")
	user := viper.Get("DB_USER")
	pass := viper.Get("DB_PASSWORD")
	dsn := fmt.Sprintf("host=%v user=%v password=%v", host, user, pass)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db
}
