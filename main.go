package main

import "fmt"

type account struct {
	login    string
	password string
	url      string
}

func main() {
	login := promtData("Введите логин:")
	password := promtData("Введите пароль:")
	url := promtData("Введите URL")

	myAccount := account{
		login:    login,
		password: password,
		url:      url,
	}
	outputPassword(&myAccount)
}

func promtData(promt string) string {
	fmt.Println(promt)
	var res string
	fmt.Scan(&res)
	return res
}

func outputPassword(acc *account) {
	fmt.Println(acc.login, acc.password, acc.url)
}
