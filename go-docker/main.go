package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func seedData() {
	var count int64
	db.Model(&User{}).Count(&count)

	if count == 0 {
		users := []User{
			{Name: "Juan Pérez", Email: "juan@example.com"},
			{Name: "María López", Email: "maria@example.com"},
			{Name: "Carlos Díaz", Email: "carlos@example.com"},
		}
		db.Create(&users)
		log.Println("Datos iniciales insertados en la tabla users.")
	} else {
		log.Println("Seed data ya existe, no se insertó nada.")
	}
}

var db *gorm.DB

func main() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		getEnv("DB_USER", "root"),
		getEnv("DB_PASSWORD", "password"),
		getEnv("DB_HOST", "db"),
		getEnv("DB_PORT", "3306"),
		getEnv("DB_NAME", "testdb"),
	)

	var err error

	// Quitar esto si ya se agrego healthcheck
	for i := 0; i < 10; i++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		log.Printf("Intento %d: No se pudo conectar a la DB: %v", i+1, err)
		time.Sleep(3 * time.Second)
	}

	if err != nil {
		log.Fatalf("No se pudo conectar a la base de datos: %v", err)
	}

	db.AutoMigrate(&User{})
	seedData()

	http.HandleFunc("/users", getUsers)
	log.Println("Servidor escuchando en :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	var users []User
	db.Find(&users)
	for _, u := range users {
		fmt.Fprintf(w, "ID: %d, Name: %s, Email: %s\n", u.ID, u.Name, u.Email)
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
