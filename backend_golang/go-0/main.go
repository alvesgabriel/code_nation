package main

import (
	"fmt"
)

func main() {
	
}

func os10maioresEstadosDoBrasil() ([]string, error) {
	var data = []string{
		"Amazonas", "Pará", "Mato Grosso", "Minas Gerais", "Bahia",
		"Mato Grosso do Sul", "Goiás", "Maranhão", "Rio Grande do Sul",
		"Tocantins"}
	if len(data) == 10 {
		return data, nil
	}
	return data, fmt.Errorf("Not implemented")
}
