package database

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"saml_sso/internal/models"
)

func Connect() (*gorm.DB, error) {
	dbPath := "./mydatabase.db"
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&models.Tenant{}, &models.Member{})

	fmt.Println(dbPath)

	fmt.Println(err)

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	// sqlDB.SetMaxIdleConns(10)

	// // SetMaxOpenConns sets the maximum number of open connections to the database.
	// sqlDB.SetMaxOpenConns(100)

	// // SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	// sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}
