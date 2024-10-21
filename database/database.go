package database

import (
    "fmt"
    _ "github.com/lib/pq"
    "github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

// Fungsi untuk inisialisasi koneksi database
func InitDB() {
    var err error
    connStr := "user=postgres password=arnoarno dbname=faqdb sslmode=disable"
    DB, err = sqlx.Connect("postgres", connStr)
    if err != nil {
        fmt.Println("Failed to connect to the database:", err)
        return
    }

    fmt.Println("Database connected successfully!")
    createTable()
}

// Fungsi untuk migrate tables jika belum ada
func createTable() {
    query := `
    CREATE TABLE IF NOT EXISTS faqs (
        id SERIAL PRIMARY KEY,
        question TEXT NOT NULL,
        answer TEXT NOT NULL
    );

    CREATE TABLE settings (
    id SERIAL PRIMARY KEY,
    appname VARCHAR(100),
    description TEXT,
    about TEXT,
    phone VARCHAR(20),
    email VARCHAR(100),
    location TEXT
);

INSERT INTO settings (appname, description, about, phone, email, location) 
VALUES ('MyApp', 'This is my app description', 'About my app', '123456789', 'admin@myapp.com', 'My Location');
    `
    _, err := DB.Exec(query)
    if err != nil {
        fmt.Println("Failed to create table:", err)
    }
}
