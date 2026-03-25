package cmd

import (
    "flag"
    "fmt"
    "os"

    "go-api-cli/services"
)

func Execute() {
    if len(os.Args) < 2 {
        fmt.Println("Uso: go-api-cli <comando> [flags]")
        fmt.Println("Comandos disponibles: users, metrics, export")
        return
    }

    cmd := os.Args[1]

    switch cmd {
    case "users":
        services.ListUsers()
    case "metrics":
        services.ShowMetrics()
    case "export":
        format := flag.String("format", "json", "Formato de exportación: json/csv")
        flag.CommandLine.Parse(os.Args[2:])
        services.ExportData(*format)
	case "process-biz":
    services.ProcessBizUsers()
    default:
        fmt.Println("Comando desconocido:", cmd)
    }
}