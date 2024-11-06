import request from "@/pkg/api/request"
import { ApiSuccessResponse } from "@/pkg/api/response"
import { Loan, Loans, LoanSimulation } from "./entity";

const SimulateLoan = (principal:number, installment_length: number) :Promise<ApiSuccessResponse<LoanSimulation>> => {
  return request({
    url: `loan/simulate`,
    method: "POST",
    data:{
      principal: principal,
      installment_length: installment_length,
    }
  });
};

const CreateLoan = (description: string, principal:number, installment_length: number) :Promise<ApiSuccessResponse<void>> => {
  return request({
    url: `loan`,
    method: "POST",
    headers: {needauth: true},
    data:{
      description:description,
      principal: principal,
      installment_length: installment_length,
    }
  });
};


const GetLoans = () :Promise<ApiSuccessResponse<Loans>> => {
  return request({
    url: `loan`,
    method: "GET",
    headers: {needauth: true},
  });
};

const GetLoanByCode = (code: string) :Promise<ApiSuccessResponse<Loan>> => {
  return request({
    url: `loan/${code}/view`,
    method: "GET",
    headers: {needauth: true},
  });
};


const GetMySummary = () :Promise<ApiSuccessResponse<{loan_count: number, total_outstanding: number,is_delinquent: boolean }>> => {
  return request({
    url: `loan/my-summary`,
    method: "GET",
    headers: {needauth: true},
  });
};

export const LoanApi = {
  GetMySummary,
  SimulateLoan,
  CreateLoan,
  GetLoans,
  GetLoanByCode,
};
