package repository

import (
	"github.com/jmoiron/sqlx"
)

func BuildRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

type Repository struct {
	db *sqlx.DB
}

func (repo *Repository) AccountExist(username string) (bool, error) {
	existRequest := `SELECT EXISTS(SELECT 1 FROM users WHERE username=$1) as "exist"`

	response := repo.db.QueryRow(existRequest, username)

	var exist bool
	err := response.Scan(&exist)
	if err != nil {
		return false, err
	}

	return exist, nil
}

type CreateAccountRequest struct {
	Username string
	Amount   int
}

func (repo *Repository) CreateAccount(req CreateAccountRequest) error {
	createRequest := `INSERT INTO users(username, amount) VALUES($1, $2)`

	if _, err := repo.db.Exec(createRequest, req.Username, req.Amount); err != nil {
		return err
	}
	return nil
}

type PatchAccountRequest struct {
	Username string
	Amount   int
}

func (repo *Repository) PatchAccount(req PatchAccountRequest) error {
	patchRequest := `UPDATE users SET amount = $1 WHERE username = $2`

	if _, err := repo.db.Exec(patchRequest, req.Amount, req.Username); err != nil {
		return err
	}
	return nil
}

type ChangeAccountRequest struct {
	LastName string
	NewName  string
}

func (repo *Repository) ChangeAccount(req ChangeAccountRequest) error {
	patchRequest := `UPDATE users SET username = $1 WHERE username = $2`

	if _, err := repo.db.Exec(patchRequest, req.NewName, req.LastName); err != nil {
		return err
	}
	return nil
}

func (repo *Repository) DeleteAccount(username string) error {
	deleteRequest := `DELETE FROM users WHERE username = $1`

	if _, err := repo.db.Exec(deleteRequest, username); err != nil {
		return err
	}
	return nil
}

type GetAccountResponse struct {
	Username string
	Amount   int
}

func (repo *Repository) GetAccount(username string) (*GetAccountResponse, error) {
	getRequest := `SELECT * FROM users WHERE username = $1`

	resp := repo.db.QueryRow(getRequest, username)
	account := new(GetAccountResponse)
	err := resp.Scan(&(account.Username), &(account.Amount))
	if err != nil {
		return &GetAccountResponse{}, err
	}

	return account, nil
}
