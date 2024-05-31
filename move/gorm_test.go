package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	db, err := gorm.Open(postgres.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Failed to open database: %v", err))
	}
	db.AutoMigrate(&UserGorm{})
	return db
}

func TestAddUser(t *testing.T) {
	db := setupTestDB()

	t.Run("successfully add user", func(t *testing.T) {
		err := AddUser(db, "John Doe", "john.doe@example.com", 30)
		assert.NoError(t, err)

		var user UserGorm
		db.First(&user, "email = ?", "john.doe@example.com")
		assert.Equal(t, "John Doe", user.Fullname)
	})

	t.Run("fail to add user with existing email", func(t *testing.T) {
		err := AddUser(db, "Jane Doe", "john.doe@example.com", 28)
		assert.EqualError(t, err, "email already exists")
	})
}
