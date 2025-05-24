# Boilerplate
This is a boilerplate, quite simple but it is a good starting point for your golang rest api application.

## Project Structure

```
boilerplate/
├── cmd/                    # Application entry point
│   └── server/             # Application server
│       └── main.go
├── internal/               # Internal packages
│   ├── platform/           # Platform integration
│   │   ├── config.go       # Configuration
│   │   └── postgres.go     # Database connection
│   └── server/             # Server router
│       ├── response.go     # Response struct
│       └── router.go       # Router configuration
├── pkg/                    # Common packages
│   └── web/                # Web package
│       └── json.go         # JSON encoder and decoder
├── docs/                   # API documentation
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── go.mod                  # Go module
├── go.sum                  # Go module
└── README.md               # README
```

## How to run
1. Setup environment variables
```bash
make setup
```
2. Generate swagger documentation
```bash
make swagger
```
3. Run the application
```bash
make run
```
Take look to Makefile for more commands

## API Documentation

You can access the API documentation at http://localhost:8080/docs

## Created by

- Ahmad Syaifudin
