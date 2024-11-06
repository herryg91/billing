-- +goose Up
CREATE TABLE loan (
    id SERIAL NOT NULL,
    code text NOT NULL UNIQUE,
    user_id int NOT NULL,

    description text not null,
    installment_cycle text NOT NULL default 'WEEKLY', -- for now we just have weekly
    installment_length int NOT NULL,
    interest_type text NOT NULL default 'FLAT', --FLAT, for now we just have flat
    interest_percent decimal(16,2) NOT NULL, 
    
    principal decimal(16,2) NOT NULL, 
    interest_amount decimal(16,2) NOT NULL, 
    total_amount decimal(16,2) NOT NULL, 
    status text not null default 'DRAFT', -- DRAFT|APPROVED|ACTIVE|DONE

    disbursed_at TIMESTAMPTZ NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ( id )
);
create index index_loan_user_id on loan (user_id);


-- +goose Down
DROP TABLE IF EXISTS loan;
