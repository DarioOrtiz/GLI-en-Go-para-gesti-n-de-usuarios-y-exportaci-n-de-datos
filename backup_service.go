package services

import (
	"fmt"
	"os"
	"time"
)

func BackupUsers() {
	src := "./data/users.db"
	if _, err := os.Stat(src); os.IsNotExist(err) {
		fmt.Println("Archivo de DB no existe:", src)
		return
	}
	dst := fmt.Sprintf("./data/backup_users_%d.db", time.Now().Unix())
	input, _ := os.ReadFile(src)
	os.WriteFile(dst, input, 0644)
	fmt.Println("Backup creado en", dst)
}