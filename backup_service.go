package services

import (
	"fmt"
	"os"
	"time"
)

func BackupUsers() {
	srcFile := "users.json"
	if _, err := os.Stat(srcFile); os.IsNotExist(err) {
		fmt.Println("No se encontró el archivo", srcFile)
		return
	}

	input, err := os.ReadFile(srcFile)
	if err != nil {
		fmt.Println("Error leyendo archivo:", err)
		return
	}

	dstFile := fmt.Sprintf("backup_users_%d.json", time.Now().Unix())
	err = os.WriteFile(dstFile, input, 0644)
	if err != nil {
		fmt.Println("Error creando backup:", err)
		return
	}

	fmt.Println("Backup creado en", dstFile)
}