package internalsql

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func Connect(dataSource string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dataSource), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed connect to dabaeses")
		return nil, err
	}

	return db, nil
}
