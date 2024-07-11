package handler

import (
	"accounts/dto"
	"accounts/internal/repository"
	"accounts/internal/service"
	"encoding/json"
	"errors"
	"github.com/valyala/fasthttp"
	"net/http"
)

func BuildHandler(src *service.Service) func(ctx *fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/account/create":
			createAccountHandler(ctx, src)
		case "/account":
			getAccountHandler(ctx, src)
		case "/account/delete":
			deleteAccountHandler(ctx, src)
		case "/account/patch":
			patchAccountHandler(ctx, src)
		case "/account/change":
			changeAccountHandler(ctx, src)
		default:
			ctx.Error("Unsupported path", http.StatusNotFound)
		}
	}
}

// Создаёт аккаунт по имени и деньгам
func createAccountHandler(ctx *fasthttp.RequestCtx, srv *service.Service) {
	var request dto.CreateAccountRequest
	err := json.Unmarshal(ctx.PostBody(), &request)
	if err != nil {
		ctx.Error("Wrong create request format", http.StatusBadRequest)
	}

	err = srv.CreateAccount(repository.CreateAccountRequest{
		Username: request.Username,
		Amount:   request.Amount,
	})
	if err != nil {
		if errors.Is(err, service.AccountAlreadyExist) {
			ctx.Error("Account already exist", http.StatusForbidden)
		} else {
			ctx.Error("Error while creating account", http.StatusInternalServerError)
		}
		return
	}
	ctx.SetStatusCode(http.StatusCreated)
}

// Меняет значение денег по имени
func patchAccountHandler(ctx *fasthttp.RequestCtx, srv *service.Service) {
	var request dto.PatchAccountRequest
	err := json.Unmarshal(ctx.PostBody(), &request)
	if err != nil {
		ctx.Error("Wrong create request format", http.StatusBadRequest)
		return
	}

	err = srv.PatchAccount(repository.PatchAccountRequest{
		Username: request.Username,
		Amount:   request.Amount,
	})

	if err != nil {
		if errors.Is(err, service.AccountDoesNotExist) {
			ctx.Error("Account doesn't exist", http.StatusBadRequest)
		} else {
			ctx.Error("Error while patching account", http.StatusInternalServerError)
		}
		return
	}
	ctx.SetStatusCode(http.StatusOK)
}

// Меняет имя по старому имени и новому
func changeAccountHandler(ctx *fasthttp.RequestCtx, srv *service.Service) {
	var request dto.ChangeAccountRequest
	err := json.Unmarshal(ctx.PostBody(), &request)
	if err != nil {
		ctx.Error("Wrong create request format", http.StatusBadRequest)
		return
	}

	err = srv.ChangeAccount(repository.ChangeAccountRequest{
		LastName: request.LastName,
		NewName:  request.NewName,
	})

	if err != nil {
		if errors.Is(err, service.AccountDoesNotExist) {
			ctx.Error("Account doesn't exist", http.StatusBadRequest)
		} else if errors.Is(err, service.AccountAlreadyExist) {
			ctx.Error("Account with new-name already exist", http.StatusForbidden)
		} else {
			ctx.Error("Error while changing account", http.StatusInternalServerError)
		}
		return
	}
	ctx.SetStatusCode(http.StatusOK)
}

// Удаляет аккаунт по имени
func deleteAccountHandler(ctx *fasthttp.RequestCtx, srv *service.Service) {
	var request dto.DeleteAccountRequest
	err := json.Unmarshal(ctx.PostBody(), &request)
	if err != nil {
		ctx.Error("Wrong create request format", http.StatusBadRequest)
		return
	}

	err = srv.DeleteAccount(request.Username)

	if err != nil {
		ctx.Error("Error while deleting account", http.StatusInternalServerError)
		return
	}
	ctx.SetStatusCode(http.StatusOK)
}

// Возвращает аккаунт по имени
func getAccountHandler(ctx *fasthttp.RequestCtx, srv *service.Service) {
	username := string(ctx.QueryArgs().Peek("username"))

	account, err := srv.GetAccount(username)
	if err != nil {
		if errors.Is(err, service.AccountDoesNotExist) {
			ctx.Error("Account doesn't exist", http.StatusBadRequest)
		} else {
			ctx.Error("Error while getting account", http.StatusInternalServerError)
		}
		return
	}

	doJSONWrite(ctx, http.StatusOK, account)
}
