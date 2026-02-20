package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/MauricioGiaconia/go_base_api/pkg/config"
	_ "github.com/go-sql-driver/mysql"
)

// NewConnection inicializa y retorna una conexión a la base de datos MySQL.
// Carga las variables de entorno necesarias y configura el pool de conexiones.
func NewConnection() *sql.DB {
	envToLoad := []string{"DB_USER", "DB_HOST", "DB_PORT", "DB_NAME"}
	config.LoadEnv(envToLoad)

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	// Configurar el pool de conexiones
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("MySQL DB: Successful connection!")

	return db
}
