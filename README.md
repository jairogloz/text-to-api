# Project text-to-api

One Paragraph of project description goes here

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

## MakeFile

Run build make command with tests
```bash
make all
```

Build the application
```bash
make build
```

Run the application
```bash
make run
```

Clean up binary from the last build:
```bash
make clean
```

## About the OpenAI Model Used
- Model: gpt-3.5-turbo
- Token Limit: 4096
- Approx max size in words: 3000
- Approx max size in bytes: 16384 (16KB)

## About the storage

### MongoDB

This project currently uses MongoDB as primary storage for the data. For the right behaviour of the application, the following is required:

#### Indexes

- Collection: `users`
  - Index: `client_id, user_id` (unique) 

### Postgres

We use Supabase to store user data. A `User` in supabase is a `Client` in our system. The supabase database is configured
in mode `session` to allow for transactional queries. See "session modes" section [here](https://supabase.com/docs/guides/database/connecting-to-postgres#how-connection-pooling-works).
