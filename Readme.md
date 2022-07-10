# Storefront for Kantin Kejujuran

Kantin kejujuran project in go and nextjs

## Demo

Try the demo here [Demo Link](https://compfest.febrilian.com)

Login as:

- Student ID: 12306
- password: 123123

## Project Setup

For running this project, you need to setup Postgres with uuid-ossp and pgcrypto extension, go and nodejs.

### Postgres Setup

```
  CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
  CREATE EXTENSION IF NOT EXISTS "pgcrypto";
```

The database dump filled with sample data is in /database

### REST API

Run the rest api with `go run cmd/rest/main.go`, by default, the rest server will be run in port 9000

### Nextjs Client App

Run with `pnpm run build && pnpm run start`

## Project Documentation

```
├── Readme.md                       # Project Documentation
├── cmd                             # Contains executable codes and serves as entry point of the app
│   └── rest
├── database                        # SQL dump with example data
│   └── compfest_store_Jul10.sql
├── internal                        # Contracts of the system
│   ├── app                         # Contains dependency injection of the app and other app's related configs
│   ├── delivery                    # Delivery layer of the app
│   ├── entities                    # Data structures
│   ├── helper                      # Helper methods
│   ├── repo                        # Database implementation layer
│   ├── usecases                    # Business logic layer
│   ├── repos.go                    # Interfaces / Contracts of all the repositories
│   └── usecases.go                 # Interfaces / Contracts of all the usecases
```

### Endpoints

"An API is only as good as the clients/libraries for it." - [a hackernews user](https://compfest.febrilian.com)

So I built the client library in client/src/services/api.ts

Frontend developers can use the library right away.

Here is the implementation example:

```
const uploadRes = await storeAPI.products.upload(file);
const req = await storeAPI.balance.add(5000);
const req = await storeAPI.balance.withdraw(5000);
const req = await storeAPI.purchase(product_slug);
const res = await storeAPI.products.create(body);
const token = await storeAPI.user.login(form.student_id, form.password);
const req = await storeAPI.user.register(form.student_id, form.password);
```
