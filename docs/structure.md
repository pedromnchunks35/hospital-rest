# Structure
```
hospital-rest/
│
├── cmd/                       # Main entry point of the application
│   └── main.go                # Initializes and runs the application
│
├── configs/                    # Configuration files (e.g., database, environment variables)
│   └── config.go              # Loads configuration from environment variables or files
│
├── controllers/                # API route handlers (controllers)
│   └── user_controller.go     # Handles requests related to users (e.g., Create, Read)
│
├── routes/                     # Routing definitions
│   └── router.go              # Defines all routes and applies middleware
│
├── models/                     # Data models (usually for databases)
│   └── user.go                # Defines user model (e.g., for ORM or struct definition)
│
├── services/                   # Business logic layer, interfacing between models and controllers
│   └── user_service.go        # Logic for user-related operations
│
├── middlewares/                # Custom middleware (e.g., authentication, logging)
│   └── auth_middleware.go     # Middleware for user authentication
│
├── utils/                      # Utility functions (e.g., logging, validation, etc.)
│   └── response.go            # Utility for formatting responses
│
│
│── mocks/                      # Where mocks are generated using mockgen from google
│
│── docs/                       # Where we produce some docs/observations regarding the project
│
│── tests/                        # Where tests are created
│   └── <SOME_TEST>_test.go     # FORMATING OF THE TESTS
│
├── .env                        # Environment variables for development (e.g., DB credentials, secrets)
├── go.mod                      # Go modules file for dependencies
├── go.sum                      # Go checksum file for dependency versions
└── README.md                   # Documentation for the project
```
