package main

// Account is the struct for the account
type Account struct {
	// ID is the id of the account
	ID string `json:"id"`
	// Name is the name of the account
	Name string `json:"name"`
	// Balance is the balance of the account
	Balance int `json:"balance"`
	// Currency is the currency of the account
	Currency string `json:"currency"`
	// Owner is the owner of the account
	Owner string `json:"owner"`
	// CreatedAt is the date of creation of the account
	CreatedAt string `json:"created_at"`
	// UpdatedAt is the date of the last update of the account
	UpdatedAt string `json:"updated_at"`
	// DeletedAt is the date of the deletion of the account
	DeletedAt string `json:"deleted_at"`
}
