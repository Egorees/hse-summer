package server

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
	"server/internal/repository"
	"server/internal/service"
	"server/proto"
)

type Server struct {
	proto.UnimplementedAccountsServer
	srv *service.Service
}

func BuildServer(srv *service.Service) *Server {
	return &Server{srv: srv}
}

func (s *Server) CreateAccount(ctx context.Context, req *proto.CreateAccountRequest) (*proto.Empty, error) {
	err := s.srv.CreateAccount(repository.CreateAccountRequest{
		Username: req.Username,
		Amount:   int(req.Amount),
	})
	if err != nil {
		if errors.Is(err, service.AccountAlreadyExist) {
			slog.Info("username already taken error with:", req.Username)
			return nil, status.Errorf(codes.AlreadyExists, "username %s is already taken", req.Username)
		}
		return nil, status.Error(codes.Internal, "something went wrong")
	}
	slog.Info("Created:", req.Username)
	return nil, nil
}

func (s *Server) PatchAccount(ctx context.Context, req *proto.PatchAccountRequest) (*proto.Empty, error) {
	err := s.srv.PatchAccount(repository.PatchAccountRequest{
		Username: req.Username,
		Amount:   int(req.Amount),
	})
	if err != nil {
		if errors.Is(err, service.AccountDoesNotExist) {
			slog.Info("account does not exist while patching", req.Username)
			return nil, status.Errorf(codes.Unknown, "username %s does not exist", req.Username)
		}
		return nil, status.Error(codes.Internal, "something went wrong")
	}
	slog.Info("patched account", req.Username)
	return nil, nil
}

func (s *Server) ChangeAccount(ctx context.Context, req *proto.ChangeAccountRequest) (*proto.Empty, error) {
	err := s.srv.ChangeAccount(repository.ChangeAccountRequest{
		LastName: req.LastName,
		NewName:  req.NewName,
	})
	if err != nil {
		if errors.Is(err, service.AccountDoesNotExist) {
			slog.Info("account does not exist while changing", req.LastName)
			return nil, status.Errorf(codes.Unknown, "username %s does not exist", req.LastName)
		}
		if errors.Is(err, service.AccountAlreadyExist) {
			slog.Info("account with new username already exist", req.NewName)
			return nil, status.Errorf(codes.AlreadyExists, "username %s already exist", req.NewName)
		}
		slog.Error("error while changing account", err)
		return nil, status.Error(codes.Internal, "something went wrong")
	}
	slog.Info("changed account", req.LastName, req.NewName)
	return nil, nil
}

func (s *Server) DeleteAccount(ctx context.Context, req *proto.DeleteAccountRequest) (*proto.Empty, error) {
	err := s.srv.DeleteAccount(req.Username)
	if err != nil {
		return nil, status.Error(codes.Internal, "something went wrong")
	}
	slog.Info("deleted account", req.Username)
	return nil, nil
}

func (s *Server) GetAccount(ctx context.Context, req *proto.GetAccountRequest) (*proto.GetAccountResponse, error) {
	resp, err := s.srv.GetAccount(req.Username)
	if err != nil {
		if errors.Is(err, service.AccountDoesNotExist) {
			slog.Info("account does not exist while getting", req.Username)
			return nil, status.Errorf(codes.Unknown, "username %s does not exist", req.Username)
		}
		return nil, status.Error(codes.Internal, "something went wrong")
	}
	slog.Info("gotten account", req.Username)
	return &proto.GetAccountResponse{
		Username: resp.Username,
		Amount:   int32(resp.Amount),
	}, nil
}
