package database

import (
	"fmt"
	"synergize/backend-test/pkg/facades"
	"synergize/backend-test/pkg/helper"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func bootDatabase() *gorm.DB {
	db, err := gorm.Open(postgres.Open(GetStringConfigDatabase()), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info),
	})

	helper.PanicIfError(err)

	return db
}

func GetStringConfigDatabase() string {
	config := facades.Config.Get("database.pgsql").(map[string]any)

	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		config["host"],
		config["username"],
		config["password"],
		config["database"],
		config["port"],
	)
}
