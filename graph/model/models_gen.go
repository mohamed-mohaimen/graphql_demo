// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Account struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Address string `json:"address"`
	User    *User  `json:"user"`
}

type Card struct {
	ID        string `json:"id"`
	Number    string `json:"number"`
	UserName  string `json:"userName"`
	ValidThru string `json:"validThru"`
	Cvv       int    `json:"cvv"`
}

type Loan struct {
	ID             string     `json:"id"`
	Name           *string    `json:"name"`
	Purpose        *string    `json:"purpose"`
	Amount         int        `json:"amount"`
	Duration       int        `json:"duration"`
	MonthlyPayment float64    `json:"monthlyPayment"`
	Interest       *float64   `json:"interest"`
	Fees           *float64   `json:"fees"`
	Payments       []*Payment `json:"payments"`
	Card           *Card      `json:"card"`
	User           *User      `json:"user"`
}

type NewAccount struct {
	Address string `json:"address"`
	Title   string `json:"title"`
	UserID  string `json:"userId"`
}

type NewLoan struct {
	Amount         int      `json:"amount"`
	Duration       int      `json:"duration"`
	MonthlyPayment float64  `json:"monthlyPayment"`
	Interest       *float64 `json:"interest"`
	Fees           *float64 `json:"fees"`
	UserID         string   `json:"userId"`
}

type Payment struct {
	ID     string  `json:"id"`
	Amount float64 `json:"amount"`
	Date   string  `json:"date"`
	Status string  `json:"status"`
	UserID string  `json:"userId"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
