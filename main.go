package main

import (
	"fmt"
	"math/rand"
)

type account struct {
	login    string
	password string
	url      string
}

func (acc *account) outputPassword() {
	fmt.Println(acc.login, acc.password, acc.url)
}

func (acc *account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	acc.password = string(res)
}

func main() {
	login := promtData("Введите логин:")
	url := promtData("Введите URL:")

	myAccount := account{
		login: login,
		url:   url,
	}
	myAccount.generatePassword(12)
	myAccount.outputPassword()
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLM")

func promtData(promt string) string {
	fmt.Println(promt)
	var res string
	fmt.Scan(&res)
	return res
}
