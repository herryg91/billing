package loan_handler

import (
	"context"
	"errors"
	"net/http"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/herryg91/billing/rest-api/app/entity"
	"github.com/herryg91/billing/rest-api/app/usecase/loan_usecase"
	pb "github.com/herryg91/billing/rest-api/handler/grst/loan"
	"github.com/herryg91/billing/rest-api/pkg/interceptor/usertoken_interceptor"
	grst_errors "github.com/herryg91/cdd/grst/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type handler struct {
	pb.LoanApiServer
	loan_uc loan_usecase.UseCase
}

func New(loan_uc loan_usecase.UseCase) pb.LoanApiServer {
	return &handler{loan_uc: loan_uc}
}

func (h *handler) SimulateLoan(ctx context.Context, req *pb.SimulateLoanReq) (*pb.LoanSimulation, error) {
	if err := pb.ValidateRequest(req); err != nil {
		return nil, err
	}

	l, simulationBillings := h.loan_uc.SimulateLoan(entity.Loan{
		InstallmentCycle:  entity.InstallmentCycle_Weekly,
		InstallmentLength: int(req.InstallmentLength),
		Principal:         req.Principal,
	})
	resp := &pb.LoanSimulation{
		InstallmentLength: req.InstallmentLength,
		Principal:         l.Principal,
		InterestAmount:    l.InterestAmount,
		TotalAmount:       l.TotalAmount,
		Billings:          []*pb.LoanBillingSimulation{},
	}
	for _, b := range simulationBillings {
		resp.Billings = append(resp.Billings, &pb.LoanBillingSimulation{
			InstallmentNumber: int32(b.InstallmentNumber),
			Principal:         b.Principal,
			InterestAmount:    b.InterestAmount,
			TotalAmount:       b.TotalAmount,
		})
	}
	return resp, nil
}

func (h *handler) CreateLoan(ctx context.Context, req *pb.CreateLoanReq) (*empty.Empty, error) {
	if err := pb.ValidateRequest(req); err != nil {
		return nil, err
	}

	user_ctx, err := usertoken_interceptor.GetUserContext(ctx)
	if err != nil {
		return nil, err
	}

	err = h.loan_uc.CreateLoanRequest(entity.Loan{
		UserId:            user_ctx.UserId,
		Description:       req.Description,
		InstallmentCycle:  entity.InstallmentCycle_Weekly,
		InstallmentLength: int(req.InstallmentLength),
		InterestType:      entity.InterestType_Flat,
		Principal:         req.Principal,
	})
	if err != nil {
		return nil, grst_errors.New(http.StatusInternalServerError, codes.Internal, 500, err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (h *handler) GetLoans(ctx context.Context, req *empty.Empty) (*pb.Loans, error) {
	user_ctx, err := usertoken_interceptor.GetUserContext(ctx)
	if err != nil {
		return nil, err
	}

	loans, err := h.loan_uc.GetLoans(user_ctx.UserId)
	if err != nil {
		return nil, grst_errors.New(http.StatusInternalServerError, codes.Internal, 500, err.Error())
	}

	resp := &pb.Loans{Loans: []*pb.Loan{}}
	for _, l := range loans {
		resp.Loans = append(resp.Loans, &pb.Loan{
			Code:              l.Code,
			Description:       l.Description,
			InstallmentCycle:  string(l.InstallmentCycle),
			InstallmentLength: int32(l.InstallmentLength),
			InterestType:      string(l.InterestType),
			InterestPercent:   l.InterestPercent,
			Principal:         l.Principal,
			InterestAmount:    l.InterestAmount,
			TotalAmount:       l.TotalAmount,
			Outstanding:       l.Outstanding,
			Status:            string(l.Status),
			CreatedAt:         timestamppb.New(l.CreatedAt),
		})
	}

	return resp, nil
}

func (h *handler) GetMySummary(ctx context.Context, req *empty.Empty) (*pb.LoanSummary, error) {
	user_ctx, err := usertoken_interceptor.GetUserContext(ctx)
	if err != nil {
		return nil, err
	}
	loan_count, total_outstanding, is_delinquent, err := h.loan_uc.GetUserSummary(user_ctx.UserId)
	if err != nil {
		return nil, grst_errors.New(http.StatusInternalServerError, codes.Internal, 500, err.Error())
	}
	return &pb.LoanSummary{
		LoanCount:        int32(loan_count),
		TotalOutstanding: total_outstanding,
		IsDelinquent:     is_delinquent,
	}, nil
}

func (h *handler) GetLoanByCode(ctx context.Context, req *pb.GetLoanByCodeReq) (*pb.Loan, error) {
	if err := pb.ValidateRequest(req); err != nil {
		return nil, err
	}
	user_ctx, err := usertoken_interceptor.GetUserContext(ctx)
	if err != nil {
		return nil, err
	}
	loan, err := h.loan_uc.GetLoanByCode(user_ctx.UserId, req.Code)
	if err != nil {
		if errors.Is(err, loan_usecase.ErrNotFound) {
			return nil, grst_errors.New(http.StatusNotFound, codes.NotFound, 404, err.Error())
		} else if errors.Is(err, loan_usecase.ErrUnauthorized) {
			return nil, grst_errors.New(http.StatusForbidden, codes.PermissionDenied, 403, err.Error())
		}
		return nil, grst_errors.New(http.StatusInternalServerError, codes.Internal, 500, err.Error())
	}

	resp := &pb.Loan{
		Code:              loan.Code,
		Description:       loan.Description,
		InstallmentCycle:  string(loan.InstallmentCycle),
		InstallmentLength: int32(loan.InstallmentLength),
		InterestType:      string(loan.InterestType),
		InterestPercent:   loan.InterestPercent,
		Principal:         loan.Principal,
		InterestAmount:    loan.InterestAmount,
		TotalAmount:       loan.TotalAmount,
		Status:            string(loan.Status),
		CreatedAt:         timestamppb.New(loan.CreatedAt),
	}

	return resp, nil
}
