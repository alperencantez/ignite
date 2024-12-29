# Ignite API

Ignite is a minimalistic API architecture template designed to prioritize simplicity and scalability.

## Features

- **Modular Structure**: Organized into distinct components such as handlers, services, and routers to promote clean code separation.
- **Middleware Integration**: Includes middleware support for tasks like logging, authentication, and request validation.
- **Database Ready**: Configured to connect with databases, facilitating data persistence.
- **Docker Support**: Comes with a `Dockerfile` and `docker-compose.yml` for containerized deployments.
- **Environment Configuration**: Utilizes environment variables for flexible configuration across different environments.

## Getting Started

### Prerequisites

- Go 1.16 or higher
- Docker (optional, for containerization)

### Installation

1. **Clone the repository**:

   ```bash
   git clone https://github.com/alperencantez/ignite.git
   ```

2. Navigate to the project directory:

   ```bash
   cd ignite
   ```

3. Install dependencies:

   ```bash
   go mod tidy
   ```

4. Set up environment variables:

Copy the .env.development file and adjust the variables as needed.

5. Run the application:

   ```bash
   go run main.go
   ```

### Using Docker

1. **Build the Docker image**:

   ```bash
   docker build -t ignite .
   ```

2. Navigate to the project directory:

   ```bash
   docker run -p 8080:8080 --env-file .env.development ignite
   ```

- You can also use the preset commands for a better devex.

  ```bash
  make <environment> # "dev" or "prod"
  ```

### Project Folder Structure

```
ignite/
├── .github/
│ └── workflows/
├── database/
├── handler/
├── middleware/
├── model/
├── pkg/
├── router/
├── service/
├── test/
├── util/
├── .air.toml
├── .env.development
├── .gitignore
├── Dockerfile
├── LICENSE.md
├── Makefile
├── README.md
├── docker-compose.yml
├── go.mod
├── go.sum
└── main.go
```

#### `database/:` Database connection and migration files.

#### `handler/:` HTTP request handlers & business logic.

#### `middleware/:` Custom middleware functions.

#### `model/:` Data models and schemas.

#### `pkg/:` Utility packages.

#### `router/:` API route definitions.

#### `service/:` Services and execution for database queries.

#### `test/:` Test cases and test utilities.

#### `util/:` Helper functions and utilities.

### Contributing

Contributions are welcome, please fork the repository and create a pull request with your changes.
