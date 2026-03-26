package main

import (
	"flag"
	"fmt"
	"os"

	"go-api-cli/services"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: go-api-cli <comando> [flags]")
		fmt.Println("Comandos: users, metrics, export, backup, python-process")
		return
	}

	cmd := os.Args[1]

	switch cmd {
	case "users":
		services.ListUsers()

	case "metrics":
		services.ShowMetrics()

	case "export":
		exportFlags := flag.NewFlagSet("export", flag.ExitOnError)
		format := exportFlags.String("format", "json", "Formato de exportación: json/csv")
		domain := exportFlags.String("domain", "", "Filtrar usuarios por dominio")
		exportFlags.Parse(os.Args[2:])
		services.ExportDataFiltered(*format, *domain)

	case "backup":
		services.BackupUsers()

	case "python-process":
		services.RunPythonAnalysis()

	default:
		fmt.Println("Comando desconocido:", cmd)
	}
}