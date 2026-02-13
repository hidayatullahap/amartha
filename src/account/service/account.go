package service

import (
	"amartha/generated/sqlc"
	"fmt"
)

type AccountService interface {
	Login()
}

type accountService struct {
}

func NewAccountService(queries *sqlc.Queries) AccountService {
	return accountService{}
}

func (s accountService) Login() {
	fmt.Println("login")
}
