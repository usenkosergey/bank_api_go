package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
	env "github.com/joho/godotenv"
)

var (
	once     sync.Once
	instance *sql.DB
	initErr  error
)

func DbConnect() {
	psqlInfo := getConnectString()
	// Подключение к базе
	db, err := sql.Open("pgx", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	// Настроим пул (опционально)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(30 * time.Minute)
	// Проверка подключения
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected to PostgreSQL!")
	instance = db
}

// GetDB возвращает синглтон пула *sql.DB
func GetDB() (*sql.DB, error) {
	once.Do(DbConnect)
	return instance, initErr
}

// Получение строки для подключения к БД
func getConnectString() string {
	errEnf := env.Load()
	if errEnf != nil {
		log.Printf("Предупреждение: Не удалось загрузить .env файл")
	}
	// Получение переменных окружения
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "54321")
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "")
	dbname := getEnv("DB_NAME", "testdb")
	sslMode := getEnv("SSL_MODE", "disable")
	// Формирование connection string
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslMode)
	return psqlInfo
}

// Пользовательская функция getEnv
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
