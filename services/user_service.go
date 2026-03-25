package services

import (
    "encoding/csv"
    "encoding/json"
    "fmt"
    "os"
    "time"
)

var Users = []map[string]string{
    {"name": "Alice", "email": "alice@example.com"},
    {"name": "Bob", "email": "bob@biz.com"},
    {"name": "Charlie", "email": "charlie@biz.com"},
    {"name": "David", "email": "david@biz.com"},
    {"name": "Eva", "email": "eva@company.org"},
    {"name": "Frank", "email": "frank@biz.com"},
}

func ListUsers() {
    fmt.Println("Listado de usuarios:")
    for _, u := range Users {
        fmt.Printf("- %s (%s)\n", u["name"], u["email"])
    }
}

func ExportData(format string) {
    switch format {
    case "json":
        fileName := fmt.Sprintf("users_%d.json", time.Now().Unix())
        f, _ := os.Create(fileName)
        defer f.Close()
        json.NewEncoder(f).Encode(Users)
        fmt.Println("Datos exportados a", fileName)
    case "csv":
        fileName := fmt.Sprintf("users_%d.csv", time.Now().Unix())
        f, _ := os.Create(fileName)
        defer f.Close()
        w := csv.NewWriter(f)
        w.Write([]string{"name", "email"})
        for _, u := range Users {
            w.Write([]string{u["name"], u["email"]})
        }
        w.Flush()
        fmt.Println("Datos exportados a", fileName)
    default:
        fmt.Println("Formato no soportado:", format)
    }
}