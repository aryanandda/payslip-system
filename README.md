# payslip-system

## Overview

**payslip-system** is a scalable payslip generation system designed to automate employee payslip creation based on attendance, overtime, and reimbursement rules.

---

## Technology Stack
Detail the technologies used in the project, including:
- Programming Languages: Go
- Frameworks & Libraries: Gin
- Databases: PostgreSQL
- Tools: Docker

---

## Getting Started

## Clone the repository

```bash
git clone https://github.com/aryanandda/payslip-system.git
cd payslip-system
```

## Run Service
- Run service from docker, then run `docker compose up --build -d` to start all docker container.
- To see the logs, run `docker logs -f payslip_app`.

## Run test
Run `go test ./services -v` to run unit test.

## API Documentation

http://localhost:8080/swagger/index.html#/
