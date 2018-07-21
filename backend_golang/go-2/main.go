package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

const filename = "data.csv"

type Player struct {
	Name   string
	Salary float64
	Age    int
}

func main() {
	//Todas as perguntas são referentes ao arquivo data.csv
}

func getPlayers() ([][]string, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("Can't open the file(%+v)\n%+v",
			filename, err)
	}

	r := csv.NewReader(strings.NewReader(string(b)))
	players, err := r.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("Can't convert CSV datan\n%+v",
			err)
	}
	return players, nil
}

//Quantas nacionalidades (coluna `nationality`) diferentes existem no arquivo?
func q1() (int, error) {
	players, err := getPlayers()
	if err != nil {
		return 0, err
	}

	var nationality = make(map[string]int)
	for _, player := range players[1:] {
		if _, ok := nationality[player[14]]; !ok {
			nationality[player[14]] = 0
		}
		nationality[player[14]]++
	}
	return len(nationality), nil
	// return 0, fmt.Errorf("Not implemented")
}

//Quantos clubes (coluna `club`) diferentes existem no arquivo?
func q2() (int, error) {
	players, err := getPlayers()
	if err != nil {
		return 0, err
	}

	var clubs = make(map[string]int)
	for _, player := range players[1:] {
		if _, ok := clubs[player[3]]; !ok {
			clubs[player[3]] = 0
		}
		clubs[player[3]]++
	}
	return len(clubs), nil
	// return 0, fmt.Errorf("Not implemented")
}

//Liste o primeiro nome dos 20 primeiros jogadores de acordo com a coluna `full_name`.
func q3() ([]string, error) {
	players, err := getPlayers()
	if err != nil {
		return nil, err
	}

	var names = []string{}
	for _, player := range players[1:21] {
		name := strings.Split(player[2], " ")[0]
		names = append(names, name)
	}
	return names, nil
	// return []string{}, fmt.Errorf("Not implemented")
}

//Quem são os top 10 jogadores que ganham mais dinheiro (utilize as colunas `full_name` e `eur_wage`)?
func q4() ([]string, error) {
	players, err := getPlayers()
	if err != nil {
		return nil, err
	}

	var salaries = []Player{}
	for _, player := range players[1:] {
		salary, _ := strconv.ParseFloat(player[17], 64)
		p := Player{
			Name:   player[2],
			Salary: salary,
		}
		salaries = append(salaries, p)
	}
	sort.Slice(salaries, func(i, j int) bool {
		return salaries[i].Salary > salaries[j].Salary
	})

	var names = []string{}
	for _, player := range salaries[:10] {
		names = append(names, player.Name)
	}

	return names, nil
	// return []string{}, fmt.Errorf("Not implemented")
}

//Quem são os 10 jogadores mais velhos?
func q5() ([]string, error) {
	players, err := getPlayers()
	if err != nil {
		return nil, err
	}

	var ages = []Player{}
	for _, player := range players[1:] {
		age, _ := strconv.Atoi(player[6])
		p := Player{
			Name: player[2],
			Age:  age,
		}
		ages = append(ages, p)
	}
	sort.Slice(ages, func(i, j int) bool {
		return ages[i].Age > ages[j].Age
	})

	var names = []string{}
	for _, player := range ages[:10] {
		names = append(names, player.Name)
	}

	return names, nil
	// return []string{}, fmt.Errorf("Not implemented")
}

//Conte quantos jogadores existem por idade. Para isso, construa um mapa onde as chaves são as idades e os valores a contagem.
func q6() (map[int]int, error) {
	players, err := getPlayers()
	if err != nil {
		return nil, err
	}

	idades := make(map[int]int)
	for _, player := range players[1:] {
		age, _ := strconv.Atoi(player[6])
		if _, ok := idades[age]; !ok {
			idades[age] = 0
		}
		idades[age]++
	}

	return idades, nil
	// return idades, fmt.Errorf("Not implemented")
}
