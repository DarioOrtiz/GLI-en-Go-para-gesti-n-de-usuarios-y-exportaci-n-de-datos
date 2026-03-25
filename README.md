# Go API CLI

CLI escrita en Go para gestionar usuarios y exportar datos en distintos formatos.

## Descripción

Este proyecto es una herramienta de línea de comandos que permite:

- Listar usuarios simulados.
- Mostrar métricas de usuarios (total y distribución por dominio de email).
- Exportar los datos de usuarios a archivos JSON o CSV.
- (Opcional) Crear backups y procesar análisis de datos en Python.

Es un ejemplo práctico de cómo organizar un proyecto Go con múltiples servicios y comandos.

---

## Comandos disponibles

| Comando             | Descripción                                           |
|--------------------|-------------------------------------------------------|
| `users`            | Muestra todos los usuarios.                           |
| `metrics`          | Muestra métricas de los usuarios (total y dominios). |
| `export -format=…` | Exporta los usuarios en JSON o CSV.                  |
| `backup`           | Crea un backup del archivo JSON más reciente.        |
| `python-process`   | Ejecuta un análisis de usuarios con Python.          |

**Ejemplo de uso:**

```bash
# Listar usuarios
go run main.go users

# Mostrar métricas
go run main.go metrics

# Exportar datos en CSV
go run main.go export -format=csv

# Exportar datos en JSON
go run main.go export -format=json

# Crear backup de usuarios
go run main.go backup

# Ejecutar análisis Python
go run main.go python-process