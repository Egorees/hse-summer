package service

import (
	"SSO/internal/repository"
	"errors"
	"log/slog"
)

type Service struct {
	repo *repository.Repository
}

func BuildService(repo *repository.Repository) *Service {
	return &Service{repo: repo}
}

// Errors
var AccountAlreadyExist = errors.New("account already exist")
var AccountDoesNotExist = errors.New("account doesn't exist")

func (srv *Service) CreateAccountSrv(req repository.CreateAccountRequest) error {
	exist, err := srv.repo.AccountExist(req.Username)
	if err != nil {
		slog.Error("Error while account exist check: ", err)
		return err
	}

	if exist {
		return AccountAlreadyExist
	}

	err = srv.repo.CreateAccount(req)

	if err != nil {
		slog.Error("Error while creating account: ", err)
	}

	return err
}

func (srv *Service) PatchAccountSrv(req repository.PatchAccountRequest) error {
	exist, err := srv.repo.AccountExist(req.Username)
	if err != nil {
		slog.Error("Error while account exist check:", err)
		return err
	}

	if !exist {
		return AccountDoesNotExist
	}

	err = srv.repo.PatchAccount(req)

	if err != nil {
		slog.Error("Error while patching account", err)
	}
	return err
}

func (srv *Service) ChangeAccountSrv(req repository.ChangeAccountRequest) error {
	exist, err := srv.repo.AccountExist(req.LastName)
	if err != nil {
		slog.Error("Error while account exist check:", err)
		return err
	}

	if !exist {
		return AccountDoesNotExist
	}

	err = srv.repo.ChangeAccount(req)

	if err != nil {
		slog.Error("Error while changing account", err)
	}
	return err
}

func (srv *Service) DeleteAccountSrv(username string) error {
	err := srv.repo.DeleteAccount(username)
	if err != nil {
		slog.Error("Error while deleting account", err)
	}
	return err
}

func (srv *Service) GetAccountSrv(username string) (*repository.GetAccountResponse, error) {
	exist, err := srv.repo.AccountExist(username)
	if err != nil {
		slog.Error("Error while account exist check:", err)
		return &repository.GetAccountResponse{}, err
	}

	if !exist {
		return &repository.GetAccountResponse{}, AccountDoesNotExist
	}

	resp, err := srv.repo.GetAccount(username)
	if err != nil {
		slog.Error("Error while getting account", err)
	}
	return resp, err
}
