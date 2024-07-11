package dto

type CreateAccountRequest struct {
	Username string `json:"username"`
	Amount   int    `json:"amount"`
}

type PatchAccountRequest struct {
	Username string `json:"username"`
	Amount   int    `json:"amount"`
}

type ChangeAccountRequest struct {
	LastName string `json:"last_name"`
	NewName  string `json:"new_name"`
}

type DeleteAccountRequest struct {
	Username string `json:"username"`
}
