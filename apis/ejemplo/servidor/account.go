package main

/*
Account is the struct for the account with the next attributes:

	ID is the id of the account
	Name is the name of the account
	Balance is the balance of the account
	Currency is the currency of the account
	Owner is the owner of the account
	CreatedAt is the date of creation of the account
	UpdatedAt is the date of update of the account
	DeletedAt is the date of deletion of the account
*/
type Account struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Balance   int    `json:"balance"`
	Currency  string `json:"currency"`
	Owner     string `json:"owner"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}
