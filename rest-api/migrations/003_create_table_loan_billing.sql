-- +goose Up
CREATE TABLE loan_billing (
    id SERIAL NOT NULL,
    loan_id int NOT NULL,
    installment_number int NOT NULL,
    due_date date NOT NULL,
    principal decimal(16,2) NOT NULL, 
    interest_amount decimal(16,2) NOT NULL, 

    -- payment info 
    payment_bank text not null default 'BCA',
    payment_va text not null default '',
    payment_status text not null default 'UNPAID', -- UNPAID|WAIT_FOR_PAYMENT|SUCCESS
    payment_expired_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    payment_ref text not null default '',

    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ( id ),
    UNIQUE (loan_id, installment_number)
);


-- +goose Down
DROP TABLE IF EXISTS loan_billing;
