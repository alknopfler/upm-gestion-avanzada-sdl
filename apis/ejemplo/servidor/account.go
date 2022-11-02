package main

// swagger:model Account
type Account struct {
	// The id of an account
	// example: 6204037c-30e6-408b-8aaa-dd8219860b4b
	ID string `json:"id"`
	// The name of an account
	// example: account1
	Name string `json:"name"`
	// The balance of an account
	// example: 100
	Balance int `json:"balance"`
	// The currency of an account
	// example: EUR
	Currency string `json:"currency"`
	// The owner of an account
	// example: owner1
	Owner string `json:"owner"`
	// The date of creation of an account
	// example: 2019-01-01T00:00:00Z
	CreatedAt string `json:"created_at"`
	// The date of update of an account
	// example: 2019-01-01T00:00:00Z
	UpdatedAt string `json:"updated_at"`
	// The date of deletion of an account
	// example: 2019-01-01T00:00:00Z
	DeletedAt string `json:"deleted_at"`
}

// swagger:model ListAccount
type ListAccount struct {
	// The list of accounts
	List []Account
}

// swagger:parameters create-account
type InputAccount struct {
	// in:body
	Name string `json:"name"`
	// in:body
	Balance int `json:"balance"`
	// in:body
	Currency string `json:"currency"`
	// in:body
	Owner string `json:"owner"`
}

// swagger:model Balance
type Balance struct {
	// The balance of an account
	Balance int `json:"balance"`
}

const schemaDB string = `
  CREATE TABLE IF NOT EXISTS account (
  id TEXT NOT NULL PRIMARY KEY,
  name TEXT NOT NULL,
  balance INTEGER NOT NULL,
  currency TEXT NOT NULL,
  owner TEXT NOT NULL,
  createdat TEXT NOT NULL,
  updatedat TEXT NOT NULL,
  deleteat TEXT NOT NULL
  );`
