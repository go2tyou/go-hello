package main

import (
	"fmt"
)

type Account struct {
	Account string
	Pwd string
	Balance float64
}

func (acc *Account) Deposite(money float64, pwd string) {
	if pwd != acc.Pwd {
		fmt.Println("pwd error")
		return
	}
	if money <= 0 {
		fmt.Println("money error")
		return
	}
	acc.Balance += money
	fmt.Println("deposite success")
}

func (acc *Account) WithDraw(money float64, pwd string) {
	if pwd != acc.Pwd {
		fmt.Println("pwd error")
		return
	}
	if money <= 0 || money > acc.Balance {
		fmt.Println("money error")
		return
	}
	acc.Balance -= money
	fmt.Println("withDraw success")
}

func (acc *Account) Query(pwd string) {
	if pwd != acc.Pwd {
		fmt.Println("pwd error")
		return
	}
	fmt.Println("query success")
	fmt.Printf("%v have %.2f\n", acc.Account, acc.Balance)
}

func main()  {
	acc := &Account {
		Account : "ssr",
		Pwd : "666666",
		Balance : 100.0,
	}
	acc.Query("666666")
	acc.Deposite(200.0, "666666")
	acc.Query("666666")
	acc.WithDraw(150.0, "666666")
	acc.Query("666666")
}