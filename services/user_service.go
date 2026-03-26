package services

import (
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// ExportDataFiltered exporta los usuarios, opcionalmente filtrando por dominio
func ExportDataFiltered(format string, domainFilter string) {
	users := getUsersFromDB(domainFilter)

	switch format {
	case "json":
		fileName := fmt.Sprintf("users_%d.json", time.Now().Unix())
		f, _ := os.Create(fileName)
		defer f.Close()
		json.NewEncoder(f).Encode(users)
		fmt.Println("Datos exportados a", fileName)

	case "csv":
		fileName := fmt.Sprintf("users_%d.csv", time.Now().Unix())
		f, _ := os.Create(fileName)
		defer f.Close()
		w := csv.NewWriter(f)
		w.Write([]string{"name", "email"})
		for _, u := range users {
			w.Write([]string{u.Name, u.Email})
		}
		w.Flush()
		fmt.Println("Datos exportados a", fileName)

	default:
		fmt.Println("Formato no soportado:", format)
	}
}

// ListUsers lista usuarios desde DB
func ListUsers() {
	users := getUsersFromDB("")
	fmt.Println("Listado de usuarios:")
	for _, u := range users {
		fmt.Printf("- %s (%s)\n", u.Name, u.Email)
	}
}

// getUsersFromDB lee usuarios desde SQLite
func getUsersFromDB(domain string) []User {
	db, _ := sql.Open("sqlite3", "./data/users.db")
	defer db.Close()

	query := "SELECT name, email FROM users"
	if domain != "" {
		query += " WHERE email LIKE ?"
		rows, _ := db.Query(query, "%"+domain)
		return scanUsers(rows)
	}
	rows, _ := db.Query(query)
	return scanUsers(rows)
}

func scanUsers(rows *sql.Rows) []User {
	var users []User
	for rows.Next() {
		var u User
		rows.Scan(&u.Name, &u.Email)
		users = append(users, u)
	}
	return users
}