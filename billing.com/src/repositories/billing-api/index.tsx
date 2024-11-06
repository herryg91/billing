import request from "@/pkg/api/request"
import { ApiSuccessResponse } from "@/pkg/api/response"
import { Billing, Billings } from "./entity";

const GetBillingByLoanCode = (loan_code: string) :Promise<ApiSuccessResponse<Billings>> => {
  return request({
    url: `billing/loan/${loan_code}`,
    method: "GET",
    headers: {needauth: true},
  });
};

const GetBillingOverDue = () :Promise<ApiSuccessResponse<Billings>> => {
  return request({
    url: `billing/overdue`,
    method: "GET",
    headers: {needauth: true},
  });
};
const GenerateBillingPayment = (loan_code: string, installment_number: number) :Promise<ApiSuccessResponse<Billing>> => {
  return request({
    url: `billing/payment/${loan_code}/${installment_number}`,
    method: "GET",
    headers: {needauth: true},
  });
};
const SettleBillingPayment = (loan_code: string, installment_number: number) :Promise<ApiSuccessResponse<void>> => {
  return request({
    url: `billing/settle/${loan_code}/${installment_number}`,
    method: "POST",
    headers: {needauth: true},
  });
};

export const BillingApi = {
  GetBillingByLoanCode,
  GetBillingOverDue,
  GenerateBillingPayment,
  SettleBillingPayment
};
