## Prerequisites

- Go 1.16 or later

## How To Run

1. Clone the repository.
2. Run `go install` to download all the packages needed.
3. Make sure you create the .env file by copying the .env.example and rename it to .env
4. Run `go run main.go` to start the application.

### Project Layout

```tree
├── adapters
│   ├── controllers
│   └── db
├── core
│   ├── models
├── └── services
└── pkg
```

- `adapter` is the code who adapt the business logic code with external. resource like a HTTP request or database.
- `core` this is where we put our business logic related code.
- `core/model` this is for the database model.
- `main.go` this is the main file that run our backend.
- `pkg` this is where we store all the packages that we use thru out the program
