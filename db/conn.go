package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib" // _ untuk init
	"github.com/joho/godotenv"
)

var DB *sql.DB

func Connect() {
	// baca file .env
	if err := godotenv.Load(); err != nil {
		panic("gagal baca file .env" + err.Error())
	}

	// ambil URL di .env
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		panic("DB_URL tidak di temukan")
	}

	//membuka koneksi dengan driver pgx
	var err error
	DB, err = sql.Open("pgx", dbURL)
	if err != nil {
		panic("tidak dapat terkoneksi dengan database" + err.Error())
	}

	if err := DB.Ping(); err != nil {
		panic("tidak dapat terkoneksi ke database" + err.Error())
	}

	fmt.Println("database terkoneksi")

}
