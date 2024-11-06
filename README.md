## Introduction
In this project, we create a loan process that includes generating billing schedules until payments are completed. Our focus is primarily on billing services for the end user, from schedule creation to payment completion.

While we implemented the loan request use case, the loan approval and disbursement processes—which usually involve a third-party payment gateway—are outside the scope of this project. Therefore, they are either skipped or simulated in this example.

There are two applications:

- billing.com: The frontend, built with Next.js.
- rest-api: The backend API, built with GoLang. We use PostgreSQL as the database, and the migration schema can be found in the migrations folder.

## Codebase & Demo
- Github: https://github.com/herryg91/billing
- Demo Website: https://herryg-billing.netlify.app
- Demo API: https://billing-api.coderbased.com

## Demo user
| email | password | Notes |
| --- | ----------- | ----- |
| user1@billing.com | user1@billing.com | Normal Case + 1 Billing Left Use Case |
| user2@billing.com | user2@billing.com | Delinquent Use Case |
| user3@billing.com | user3@billing.com | Loan Done Use Case |
| user4@billing.com | user4@billing.com | Empty |
| user5@billing.com | user5@billing.com | Empty |