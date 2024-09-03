# Loan API Project

## Features

- **User Authentication & Authorization:**
  - User registration, login, and profile management.
  - Role-based access control (admin/user).
  - Secure JWT-based token management with access and refresh tokens.


## Tech Stack

- **Programming Language:** Go
- **Framework:** Gin
- **Database:** MongoDB
- **Authentication:** JWT (JSON Web Tokens)
- **Documentation:** Postman

### Key Components:

- **Delivery:** Contains the controllers and routers, handling HTTP requests and responses.
- **Domain:** Contains the core business logic and domain models.
- **Infrastructure:** Handles external services and infrastructure-related code such as authentication and database connections.
- **Repositories:** Implements the data access layer, interacting with the MongoDB database.
- **Usecases:** Contains the application-specific business logic.

## Getting Started

### Prerequisites

- Go 1.19 or higher
- MongoDB
- Postman (for API testing)


## Getting Started

### Prerequisites

Before you begin, ensure you have the following installed on your machine:

- [Go](https://golang.org/doc/install) (version 1.19 or higher)
- [MongoDB](https://www.mongodb.com/)

### Setup

1. **Clone the repository:**

    ```sh
    git clone git@github.com:Ararsa-Derese/Loan-Tracker-API.git
    cd Loan-Tracker-API
    ```
2. **Set up environment variables:**

    Copy the `.env.example` file to `.env` and update the environment variables as needed.

    ```sh
    cp cmd/.env.example cmd/.env
    ```

3. **Install dependencies:**

    ```sh
    go mod download
    ```

4. **Run the backend server:**

    ```sh
    go run cmd/main.go
    ```


Here is the postman Documentation
(https://documenter.getpostman.com/view/30253109/2sAXjGduYC)