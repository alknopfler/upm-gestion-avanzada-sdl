package main

// swagger:parameters get-account create-account-balance get-account-balance delete-account
type _ struct {
	// The ID of an account
	// in:path
	UUID string `json:"id"`
}

// swagger:model ErrorResponse
type ErrorResponse struct {
	// The error message
	Error string `json:"error"`
}

// swagger:route GET /api/v1/accounts Account get-accounts
/////////////////////////////////////////////////////////////
// This is the summary for getting all accounts
//
// This is the description for getting all accounts. Which can be longer.
//
// responses:
//   200: []Account
//   416: ErrorResponse
//   500: ErrorResponse
//   405: ErrorResponse

// swagger:route GET /api/v1/account/{id} Account get-account
/////////////////////////////////////////////////////////////
// This is the summary for getting an account by its ID
//
// This is the description for getting an account by its ID. Which can be longer.
//
// responses:
//   200: Account
//   404: ErrorResponse
//   500: ErrorResponse
//   405: ErrorResponse

// swagger:route POST /api/v1/account Account create-account
/////////////////////////////////////////////////////////////
// This is the summary for creating an account
//
// This is the description for creating an account. Which can be longer.
//
//
// responses:
//   201:
//   400: ErrorResponse
//   500: ErrorResponse

// swagger:route POST /api/v1/account/{id}/balance Account create-account-balance
/////////////////////////////////////////////////////////////
// This is the summary for creating an account balance
//
// This is the description for creating an account balance. Which can be longer.
//

// responses:
//   200:
//   400: ErrorResponse
//   500: ErrorResponse

// swagger:route GET /api/v1/account/{id}/balance Account get-account-balance
/////////////////////////////////////////////////////////////
// This is the summary for getting an account balance
//
// This is the description for getting an account balance. Which can be longer.
//
// responses:
// 	200: Balance
// 	404: ErrorResponse
// 	500: ErrorResponse
// 	405: ErrorResponse

// swagger:route DELETE /api/v1/account/{id} Account delete-account
/////////////////////////////////////////////////////////////
// This is the summary for deleting an account
//
// This is the description for deleting an account. Which can be longer.
//
// responses:
//   200:
//   404: ErrorResponse
//   500: ErrorResponse
//   405: ErrorResponse

// generation of the swagger file
// swagger generate spec -o ./swagger.json --scan-models
