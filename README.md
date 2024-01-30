# Simple Bank Backend Web Service in Golang

## Overview

This project was the outcome of following the course [Backend Master Class by TECH SCHOOL](https://www.udemy.com/course/backend-master-class-golang-postgresql-kubernetes/?couponCode=A468BE23956E907B8CBC) where we designed, developed, tested and deployed a backend web service for a simple bank.
Three main functions were implemented:
- Create and manage bank accounts.
- Record all balance changes to each of the accounts.
- Perform a money transfer between 2 accounts.

## Backend Web Development Topics
- Database design and database migrations
- RESTful HTTP APIs using Gin
- Build the application with Docker
- Deploying the application to production with Kubernetes + AWS
- Set up GitHub Actions for automatic test, build and deploy
- Develop gRPC APIs with unit testing
- gRPC gateway to serve both gRPC and HTTP requests at the same time
- Swagger API documentation
- Asynchronous processing in Golang using background workers and Redis as message queue
- Create and send verification email

## Build image
```
docker build -t simple-bank:latest .
```

## Setup infrastructure

Create the bank-network

``` bash
make network
```

Start postgres container:

```bash
make postgres
```

Create simple_bank database:

```bash
make createdb
```

Run db migration up all versions:

```bash
make migrateup
```

Run db migration up 1 version:

```bash
make migrateup1
```

Run db migration down all versions:

```bash
make migratedown
```

Run db migration down 1 version:

```bash
make migratedown1
```

## How to generate code

Generate schema SQL file with DBML:

```bash
make db_schema
```

Generate SQL CRUD with sqlc:

```bash
make sqlc
```

Generate DB mock with gomock:

```bash
make mock
```

Create a new db migration:

```bash
migrate create -ext sql -dir db/migration -seq $(name)
```

## How to run

Run server:

```bash
make server
```

Run test:

```bash
make test
```