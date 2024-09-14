package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Определение структуры (модель)
type User struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"size:100"`
	Email string `gorm:"unique"`
	Age   int
}

func main() {
	// Строка подключения к базе данных
	dsn := "host=localhost user=postgres password=admin dbname=postgres port=5432 sslmode=disable"

	// Подключение к базе данных
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Автоматическая миграция (создание таблиц на основе моделей)
	db.AutoMigrate(&User{})

	// Создание новой записи в таблице
	user := User{Name: "John Doe", Email: "john.doe@example.com", Age: 30}
	result := db.Create(&user) // INSERT INTO users (name, email, age) VALUES ("John Doe", "john.doe@example.com", 30)
	if result.Error != nil {
		log.Fatalf("Failed to create user: %v", result.Error)
	}

	fmt.Println("New user ID:", user.ID)

	// Чтение данных
	var readUser User
	db.First(&readUser, user.ID) // SELECT * FROM users WHERE id = user.ID LIMIT 1
	fmt.Printf("User: %+v\n", readUser)

	// Обновление данных
	db.Model(&readUser).Update("Age", 31) // UPDATE users SET age = 31 WHERE id = user.ID

	// Удаление записи
	db.Delete(&readUser) // DELETE FROM users WHERE id = user.ID
}
