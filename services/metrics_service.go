package services

import "fmt"

func ShowMetrics() {
	total := len(Users)
	domainCount := make(map[string]int)
	for _, u := range Users {
		email := u["email"]
		at := 0
		for i, c := range email {
			if c == '@' {
				at = i
				break
			}
		}
		domain := email[at+1:]
		domainCount[domain]++
	}

	fmt.Println("Métricas de la aplicación:")
	fmt.Println("Usuarios totales:", total)
	fmt.Println("Distribución por dominio:")
	for d, c := range domainCount {
		fmt.Printf("- %s: %d\n", d, c)
	}
}