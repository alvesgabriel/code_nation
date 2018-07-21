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
	Birth  string
}

var players [][]string

func main() {
	//Todas as perguntas são referentes ao arquivo data.csv
	for i, v := range players[0] {
		if v == "birth_date" {
			fmt.Println(i, v)
		}
	}
	// fmt.Println(players[1][17])
	// fmt.Println(players[42][17])
	// fmt.Println(q1())
	// fmt.Println(q2())
	// fmt.Println(q3())
	// fmt.Println(q4())
	names, _ := q5()
	for _, name := range names {
		fmt.Println(name)
	}
	// fmt.Println(q6())
}

func init() {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(fmt.Errorf("Can't open the file(%+v)\n%+v",
			filename, err))
	}

	r := csv.NewReader(strings.NewReader(string(b)))
	players, err = r.ReadAll()
	if err != nil {
		panic(fmt.Errorf("Can't convert CSV datan\n%+v",
			err))
	}
}

//Quantas nacionalidades (coluna `nationality`) diferentes existem no arquivo?
func q1() (int, error) {
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
	var ages = []Player{}
	for _, player := range players[1:] {
		age, _ := strconv.Atoi(player[6])
		p := Player{
			Name:  player[2],
			Age:   age,
			Birth: player[8],
		}
		ages = append(ages, p)
	}
	sort.Slice(ages, func(i, j int) bool {
		return ages[i].Age > ages[j].Age
	})
	eldestPlayers := ages[:10]
	sort.Slice(eldestPlayers, func(i, j int) bool {
		return eldestPlayers[i].Birth < eldestPlayers[j].Birth
	})
	var names = []string{}
	for _, player := range eldestPlayers {
		names = append(names, player.Name)
	}

	return names, nil
	// return []string{}, fmt.Errorf("Not implemented")
}

//Conte quantos jogadores existem por idade. Para isso, construa um mapa onde as chaves são as idades e os valores a contagem.
func q6() (map[int]int, error) {
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
