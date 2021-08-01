package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"graphgl-demo/graph/generated"
	"graphgl-demo/graph/model"
	"math/rand"
)

func (r *mutationResolver) CreateAccount(ctx context.Context, input model.NewAccount) (*model.Account, error) {
	account := model.Account{
		ID:      fmt.Sprintf("T%d", rand.Int()),
		Title:   input.Title,
		Address: input.Address,
	}

	user := &model.User{Name: input.UserID}

	account.User = user

	r.AccountsDB[account.ID] = &account
	return &account, nil
}

func (r *mutationResolver) CreateLoan(ctx context.Context, input model.NewLoan) (*model.Loan, error) {
	user := &model.User{Name: "mohaimen"}

	card := &model.Card{
		ID:        fmt.Sprintf("T%d", rand.Int()),
		Number:    "19823746591823",
		UserName:  user.Name,
		ValidThru: "08/25",
		Cvv:       653,
	}

	loan := &model.Loan{
		ID:             fmt.Sprintf("T%d", rand.Int()),
		Name:           nil,
		Purpose:        nil,
		Amount:         input.Amount,
		Duration:       input.Duration,
		MonthlyPayment: input.MonthlyPayment,
		Interest:       input.Interest,
		Fees:           input.Fees,
		Payments:       r.Payments,
		Card:           card,
		User:           user,
	}

	r.LoansDB[loan.ID] = loan

	return loan, nil
}

func (r *queryResolver) Accounts(ctx context.Context) ([]*model.Account, error) {
	var accounts []*model.Account

	for _, v := range r.AccountsDB {
		accounts = append(accounts, v)
	}
	return accounts, nil
}

func (r *queryResolver) Loans(ctx context.Context) ([]*model.Loan, error) {
	var loans []*model.Loan

	for _, v := range r.LoansDB {
		loans = append(loans, v)
	}

	return loans, nil
}

func (r *queryResolver) AccountByID(ctx context.Context, input string) (*model.Account, error) {
	return r.AccountsDB[input], nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
