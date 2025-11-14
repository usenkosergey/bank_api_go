package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
	env "github.com/joho/godotenv"
)

func DbConnect() *sql.DB {
	psqlInfo := getConnectString()
	// Подключение к базе
	db, err := sql.Open("pgx", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	// Проверка подключения
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected to PostgreSQL!")

	return db
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
