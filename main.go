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

func (acc account) outputPassword() {
	fmt.Println(acc.password)
}

func main() {
	fmt.Println(generatePassword(12))
	login := promtData("Введите логин:")
	password := promtData("Введите пароль:")
	url := promtData("Введите URL:")

	myAccount := account{
		login,
		password,
		url,
	}
	myAccount.outputPassword()
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLM")

func promtData(promt string) string {
	fmt.Println(promt)
	var res string
	fmt.Scan(&res)
	return res
}

func outputPassword(acc *account) {
	fmt.Println(acc.login, acc.password, acc.url)
}

func generatePassword(n int) string {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(res)
}
