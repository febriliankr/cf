# Storefront API

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

##
