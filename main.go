package main

import (
	"errors"
	"fmt"
	"math/rand"
	"net/url"
	"time"
)

type account struct {
	login    string
	password string
	url      string
}

type accountWithTimeStamp struct {
	createdAt time.Time
	updatedAt time.Time
	account
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

func newAccount(login, password, urlString string) (*accountWithTimeStamp, error) {
	if login == "" {
		return nil, errors.New("login required")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("Invalid URL")
	}

	newAcc := &accountWithTimeStamp{
		account:   account{login: login, password: password, url: urlString},
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}

	if password == "" {
		newAcc.generatePassword(12)
	}

	return newAcc, nil
}

func newAccount(login, password, urlString string) (*account, error) {
	if login == "" {
		return nil, errors.New("login required")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("Invalid URL")
	}

	newAcc := &account{
		login,
		password,
		urlString,
	}

	if password == "" {
		newAcc.generatePassword(12)
	}

	return newAcc, nil
}

func main() {
	login := promtData("Введите логин:")
	password := promtData("Введите пароль:")
	url := promtData("Введите URL:")

	myAccount, err := newAccount(login, password, url)
	if err != nil {
		fmt.Println(err)
		return
	}
	myAccount.outputPassword()
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLM")

func promtData(promt string) string {
	fmt.Print(promt)
	var res string
	fmt.Scanln(&res)
	return res
}
