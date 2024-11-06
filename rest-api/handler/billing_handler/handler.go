package billing_handler

import (
	"context"
	"errors"
	"net/http"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/herryg91/billing/rest-api/app/usecase/billing_usecase"
	"github.com/herryg91/billing/rest-api/app/usecase/loan_usecase"
	pb "github.com/herryg91/billing/rest-api/handler/grst/billing"
	"github.com/herryg91/billing/rest-api/pkg/interceptor/usertoken_interceptor"
	grst_errors "github.com/herryg91/cdd/grst/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type handler struct {
	pb.BillingApiServer
	billing_uc billing_usecase.UseCase
}

func New(billing_uc billing_usecase.UseCase) pb.BillingApiServer {
	return &handler{billing_uc: billing_uc}
}

func (h *handler) GetBillingByLoanCode(ctx context.Context, req *pb.GetBillingByLoanCodeReq) (*pb.Billings, error) {
	if err := pb.ValidateRequest(req); err != nil {
		return nil, err
	}
	user_ctx, err := usertoken_interceptor.GetUserContext(ctx)
	if err != nil {
		return nil, err
	}
	billings, err := h.billing_uc.GetBillingByLoanCode(user_ctx.UserId, req.LoanCode)
	if err != nil {
		if errors.Is(err, loan_usecase.ErrUnauthorized) {
			return nil, grst_errors.New(http.StatusForbidden, codes.PermissionDenied, 403, err.Error())
		}
		return nil, grst_errors.New(http.StatusInternalServerError, codes.Internal, 500, err.Error())
	}
	resp := &pb.Billings{
		Billings: []*pb.Billing{},
	}
	for _, b := range billings {
		resp.Billings = append(resp.Billings, &pb.Billing{
			Id:                int32(b.Id),
			LoanCode:          b.LoanCode,
			InstallmentNumber: int32(b.InstallmentNumber),
			DueDate:           b.DueDate.String(),
			Principal:         b.Principal,
			InterestAmount:    b.InterestAmount,
			TotalAmount:       b.TotalAmount,
			PaymentStatus:     string(b.PaymentStatus),
		})
	}
	return resp, nil
}
func (h *handler) GetBillingOverDue(ctx context.Context, req *empty.Empty) (*pb.Billings, error) {
	user_ctx, err := usertoken_interceptor.GetUserContext(ctx)
	if err != nil {
		return nil, err
	}
	billings, err := h.billing_uc.GetBillingOverDue(user_ctx.UserId)
	if err != nil {
		return nil, grst_errors.New(http.StatusInternalServerError, codes.Internal, 500, err.Error())
	}
	resp := &pb.Billings{
		Billings: []*pb.Billing{},
	}
	for _, b := range billings {
		resp.Billings = append(resp.Billings, &pb.Billing{
			Id:                int32(b.Id),
			LoanCode:          b.LoanCode,
			InstallmentNumber: int32(b.InstallmentNumber),
			DueDate:           b.DueDate.String(),
			Principal:         b.Principal,
			InterestAmount:    b.InterestAmount,
			TotalAmount:       b.TotalAmount,
			PaymentStatus:     string(b.PaymentStatus),
		})
	}
	return resp, nil
}

func (h *handler) GenerateBillingPayment(ctx context.Context, req *pb.GenerateBillingPaymentReq) (*pb.Billing, error) {
	if err := pb.ValidateRequest(req); err != nil {
		return nil, err
	}
	user_ctx, err := usertoken_interceptor.GetUserContext(ctx)
	if err != nil {
		return nil, err
	}
	b, err := h.billing_uc.GeneratePaymentInfo(user_ctx.UserId, req.LoanCode, int(req.InstallmentNumber))
	if err != nil {
		if errors.Is(err, loan_usecase.ErrNotFound) {
			return nil, grst_errors.New(http.StatusNotFound, codes.NotFound, 404, err.Error())
		} else if errors.Is(err, loan_usecase.ErrUnauthorized) {
			return nil, grst_errors.New(http.StatusForbidden, codes.PermissionDenied, 403, err.Error())
		}
		return nil, grst_errors.New(http.StatusInternalServerError, codes.Internal, 500, err.Error())
	}
	resp := &pb.Billing{
		Id:                int32(b.Id),
		LoanCode:          b.LoanCode,
		InstallmentNumber: int32(b.InstallmentNumber),
		DueDate:           b.DueDate.String(),
		Principal:         b.Principal,
		InterestAmount:    b.InterestAmount,
		TotalAmount:       b.TotalAmount,
		PaymentStatus:     string(b.PaymentStatus),
		PaymentBank:       b.PaymentBank,
		PaymentVA:         b.PaymentVA,
		PaymentExpiredAt:  timestamppb.New(b.PaymentExpiredAt),
	}

	return resp, nil
}

func (h *handler) SettleBillingPayment(ctx context.Context, req *pb.GenerateBillingPaymentReq) (*empty.Empty, error) {
	if err := pb.ValidateRequest(req); err != nil {
		return nil, err
	}
	user_ctx, err := usertoken_interceptor.GetUserContext(ctx)
	if err != nil {
		return nil, err
	}
	err = h.billing_uc.SettlePayment(user_ctx.UserId, req.LoanCode, int(req.InstallmentNumber))
	if err != nil {
		if errors.Is(err, loan_usecase.ErrNotFound) {
			return nil, grst_errors.New(http.StatusNotFound, codes.NotFound, 404, err.Error())
		} else if errors.Is(err, loan_usecase.ErrUnauthorized) {
			return nil, grst_errors.New(http.StatusForbidden, codes.PermissionDenied, 403, err.Error())
		}
		return nil, grst_errors.New(http.StatusInternalServerError, codes.Internal, 500, err.Error())
	}
	return &empty.Empty{}, nil
}
