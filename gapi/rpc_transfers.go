package gapi

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	db "github.com/GiorgiMakharadze/bank-API-golang/db/sqlc"
	"github.com/GiorgiMakharadze/bank-API-golang/pb"
	"github.com/GiorgiMakharadze/bank-API-golang/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateTransfer(ctx context.Context, req *pb.CreateTransferRequest) (*pb.CreateTransferResponse, error) {
	fromAccount, valid, msg := server.validAccount(ctx, req.FromAccountId, req.Currency)
	if !valid {
		return nil, status.Errorf(codes.NotFound, msg)
	}

	authPayload, err := server.authorizeUser(ctx, []string{util.BankerRole, util.DepositorRole})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated: %v", err)
	}

	if fromAccount.Owner != authPayload.Username {
		return nil, status.Errorf(codes.PermissionDenied, "from account doesn't belong to the authenticated user")
	}

	_, valid, msg = server.validAccount(ctx, req.ToAccountId, req.Currency)
	if !valid {
		return nil, status.Errorf(codes.InvalidArgument, msg)
	}

	arg := db.TransferTxParams{
		FromAccountID: req.FromAccountId,
		ToAccountID:   req.ToAccountId,
		Amount:        req.Amount,
	}

	result, err := server.store.TransferTx(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal error: %v", err)
	}

	response := &pb.CreateTransferResponse{
		Id:            result.Transfer.ID,
		FromAccountId: result.FromAccount.ID,
		ToAccountId:   result.ToAccount.ID,
		Amount:        result.Transfer.Amount,
		CreatedAt:     result.Transfer.CreatedAt.Format(time.RFC3339),
	}
	return response, nil
}

func (server *Server) validAccount(ctx context.Context, accountID int64, currency string) (db.Account, bool, string) {
	account, err := server.store.GetAccount(ctx, accountID)
	if err != nil {
		if err == sql.ErrNoRows {
			return account, false, "account not found"
		}
		return account, false, "internal server error"
	}

	if account.Currency != currency {
		return account, false, fmt.Sprintf("account currency mismatch: expected %s, got %s", currency, account.Currency)
	}

	return account, true, ""
}
