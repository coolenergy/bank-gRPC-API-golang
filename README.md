# Bank Service

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Bank is a backend service that provides gRPC APIs to the frontend, facilitating the following functionalities:

1. **Create and Manage Bank Accounts**: Users can create bank accounts with details including owner’s name, balance, and currency.
2. **Record Balance Changes**: The service records every transaction that results in a balance change, creating an account entry record for each such instance.
3. **Money Transfer**: Enables users to perform money transfers between two accounts within a transaction, ensuring that either both accounts’ balances are updated successfully, or none of them are.
4. **User Authentication**: Authenticates users and ensures that they can only access and manage their own accounts.
5. **Role-based Functionality**: The service includes role-based access control, with specific roles such as "banker" and "depositor". There are several ways to create a banker user. You can create the first banker user either via a DB migration, or a script that runs on the production server. Once the first banker user is created, they can access an API that allows them to create other banker users.

### Used Technologies

The project utilizes the following technologies:

- Golang
- gRPC
- PostgreSQL
- Redis
- SQLC
- Asynq

## Setup Local Development

### Install Tools

- **Docker Desktop**: [Installation Guide](https://www.docker.com/products/docker-desktop)
- **TablePlus**: [Installation Guide](https://tableplus.com/)
- **Golang**: [Installation Guide](https://golang.org/)
- **Migrate**:
  ```bash
  sudo install golang-migrate
  ```
- **DB Docs**:
  ```bash
  npm install -g dbdocs
  dbdocs login
  ```
- **DBML CLI**:
  ```bash
  npm install -g @dbml/cli
  dbml2sql --version
  ```
- **Sqlc**:
  ```bash
  brew install sqlc
  ```
- **Gomock**:
  ```bash
  go install github.com/golang/mock/mockgen@v1.6.0
  ```

### Setup Infrastructure

- **Create Bank Network**:
  ```bash
  make network
  ```
- **Start Postgres Container**:
  ```bash
  make postgres
  ```
- **Create Bank Database**:
  ```bash
  make createdb
  ```
- **Run Database Migrations**:
  ```bash
  make migrateup
  make migratedown
  ```

### Documentation

- Generate DB documentation:
  ```bash
  make db_docs
  ```
- Access the DB documentation at [Swagger Hub](https://app.swaggerhub.com/apis-docs/MAKHARADZEGIORGI00/bank-gRPC/1.2) or at `localhost:8080/swagger` when you run the app locally.

### Code Generation

- **Generate Schema SQL File with DBML**:
  ```bash
  make db_schema
  ```
- **Generate SQL CRUD with Sqlc**:
  ```bash
  make sqlc
  ```
- **Generate DB Mock with Gomock**:
  ```bash
  make mock
  ```
- **Create a New DB Migration**:
  ```bash
  make new_migration name=<migration_name>
  ```

### How to Run

- **Run Server**:
  ```bash
  make server
  ```

Alternatively, if you have Docker installed on your local machine, you can simply run:

```bash
docker compose up
```

- **Run Tests**:
  ```bash
  make test
  ```
