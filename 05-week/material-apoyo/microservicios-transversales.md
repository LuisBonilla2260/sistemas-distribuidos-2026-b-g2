# Transversal microservices: support guide

> Week 05 status: this document is study material and a reference for future decisions. The Di-Lucca MVP **does not implement transversal microservices**; it currently operates as a monolith with Angular 20, Spring Boot, and PostgreSQL.

## What are they?

A transversal microservice addresses a capability that can be required by several system domains without becoming the owner of each domain's business logic. Its purpose is to provide a clear, reusable interface while preserving the autonomy of business services.

Common examples include:

- Identity and access: authentication, authorization, and role management.
- Notifications: email, SMS, or in-app messages triggered by requests or events.
- Audit: recording relevant actions for traceability.
- Observability: metrics, logs, and traces used to operate the system.
- Configuration: environment parameters and securely managed secrets.
- Error handling: response, error-code, and technical logging conventions.

Not every reusable piece of code should become a microservice. A local utility or shared library can be enough when it does not require independent deployment, data, scaling, or lifecycle management.

## Relationship with Di-Lucca domains

The system's candidate business domains are Patients, Appointments, Clinical Care, and Billing. Each domain must retain ownership of its rules and data. A transversal capability must not replace that ownership.

```text
Patients ─────┐
Appointments ─┼──► explicit interface or event ──► transversal capability
Clinical Care ┤                                  (for example, notifications)
Billing ──────┘
```

Communication must use explicit contracts — APIs or events — rather than direct access to another domain's database.

## Risks to avoid

- Creating a shared database for every service.
- Placing Patients, Appointments, or Billing business rules inside a transversal service.
- Coupling domains to a concrete email, logging, or authentication implementation.
- Extracting a service before identifying its owner, contract, data, and operating model.

## Criteria for a future extraction

Before extracting a capability from the monolith, the team should answer:

1. What need does it address, and who owns it technically?
2. Which contract does it expose, and which domains consume it?
3. Which data does it own, and which data must it not read directly?
4. How will it be tested, observed, deployed, and recovered after a failure?
5. What value does extraction provide compared with keeping it inside the MVP?

## Current MVP priority

During Week 05, the priority is to strengthen the existing monolith: preserve the `application`, `domain`, and `infrastructure` layers, expand unit and integration testing, and validate API contracts between frontend and backend. Extracting transversal microservices should only be evaluated when a real need and a validated boundary exist.

![Visual reference: transversal capabilities as a future proposal](week-05-mvp-testing-transversal-services.png)
