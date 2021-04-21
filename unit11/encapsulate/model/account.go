package model

import (
	"fmt"
)

type account struct {
	accountNo string
	pwd string
	balance float64
}

func NewAccount(accountNo string, pwd string, balance float64) *account {
	if len(accountNo) < 6 || len(accountNo) > 10 {
		fmt.Println("accountNo len error")
		return nil
	}

	if len(pwd) != 6 {
		fmt.Println("pwd len error")
		return nil
	}

	if balance < 20.0 {
		fmt.Println("balance error")
		return nil
	}

	return &account {
		accountNo : accountNo,
		pwd : pwd,
		balance : balance,
	}
}