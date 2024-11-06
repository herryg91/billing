export type LoanSimulation = {
	installment_length: number
	principal: number
	interest_amount: number
	total_amount: number
	billings: LoanBillingSimulation[]
}

export type LoanBillingSimulation = {
	installment_number: number
	principal: number
	interest_amount: number
	total_amount: number
}

export type Loans = {
	loans: Loan[]
}
export type Loan = {
	code: string
	description: string
	installment_cycle: "WEEKLY"
	installment_length: number
	interest_type: string
	interest_percent: number
	principal: number
	interest_amount: number
	total_amount: number
	outstanding: number
	status: "PENDING"|"APPROVED"|"ACTIVE"|"DONE"
	created_at: Date
}

export type Billing = {
	id: number
	loan_code: string
	installment_number: number
	due_date: Date
	principal: number
	interest_amount: number
	total_amount: number
	payment_status: "UNPAID"|"WAIT_FOR_PAYMENT"|"PAID"
	payment_bank: string
	payment_va: string
	payment_expired_at: Date
}
export type Billings = {
	billings: Billing[]
}

