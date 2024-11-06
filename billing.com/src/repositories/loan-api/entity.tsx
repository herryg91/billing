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
