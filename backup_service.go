package services

import (
    "fmt"
    "os"
    "time"
)

func BackupUsers() {
    fileName := "users.json"
    f, err := os.Create(fileName)
    if err != nil {
        fmt.Println("Error creando users.json:", err)
        return
    }
    defer f.Close()

    json.NewEncoder(f).Encode(Users)

    backupFile := fmt.Sprintf("backup_users_%d.json", time.Now().Unix())
    input, _ := os.ReadFile(fileName)
    os.WriteFile(backupFile, input, 0644)
    fmt.Println("Backup creado en", backupFile)
}