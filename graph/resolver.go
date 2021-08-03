package graph

import (
	"graphgl-demo/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{

	AccountsDB map[string]*model.Account
	LoansDB    map[string]*model.Loan
	Payments   []*model.Payment
}

func NewResolver() *Resolver{

	accountsDB := make(map[string]*model.Account)
	loansDB    := make(map[string]*model.Loan)
	payments   := make([]*model.Payment, 0)

	pay1 := &model.Payment{
		ID: "1",
		Amount: 220.32,
		Date: "25 Jul 2021",
		Status: "Awaiting",
	}

	pay2 := &model.Payment{
		ID: "2",
		Amount: 220.32,
		Date: "25 Aug 2021",
		Status: "Awaiting",
	}

	pay3 := &model.Payment{
		ID: "3",
		Amount: 220.32,
		Date: "25 Sep 2021",
		Status: "Awaiting",
	}

	pay4 := &model.Payment{
		ID: "4",
		Amount: 220.32,
		Date: "25 Oct 2021",
		Status: "Awaiting",
	}

	pay5 := &model.Payment{
		ID: "5",
		Amount: 220.32,
		Date: "25 Nov 2021",
		Status: "Awaiting",
	}

	payments = append(payments, pay1)
	payments = append(payments, pay2)
	payments = append(payments, pay3)
	payments = append(payments, pay4)
	payments = append(payments, pay5)

	return &Resolver{
		AccountsDB: accountsDB,
		LoansDB:    loansDB,
		Payments:   payments,
	}
}

